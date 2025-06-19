package task

import (
	"earthquake-crawler/internal/config"
	"earthquake-crawler/internal/crawler/crawler/jpeewcrawler"
	"earthquake-crawler/internal/crawler/crawler/jpquakecrawler"
	"earthquake-crawler/internal/crawler/parser/jpeewparser"
	"earthquake-crawler/internal/crawler/parser/jpquakeparser"
	"earthquake-crawler/internal/model"
	"earthquake-crawler/internal/repo"
	"earthquake-crawler/internal/util"
	"fmt"

	"github.com/sirupsen/logrus"
)

var lastEEWData = new(model.JapanEEWData)

func JapanEarthquakeCrawlTask() {
	eqListDoc, err := jpquakecrawler.GetJapanEarthquakeListDoc()
	if err != nil {
		logrus.Errorf("[日本地震信息]在获取地震信息列表的过程中失败: %v", err)
		return
	}

	eqList, err := jpquakeparser.ParseJapanEarthquakeListDoc(eqListDoc)
	if err != nil {
		logrus.Errorf("[日本地震信息]在解析地震列表HTML的过程中失败: %v", err)
		return
	}

	eqNotExist, err := repo.GetJapanEarthquakeNotInDB(eqList)
	if err != nil {
		logrus.Errorf("[日本地震信息]在查询数据库选择需要获取详情的地震列表的过程中失败：%v", err)
		return
	}

	if len(eqNotExist) != 0 {
		logrus.Infof("[日本地震信息]解析到%d条未加入的地震信息，即将开始获取详情", len(eqNotExist))
	}

	for _, eqTime := range eqNotExist {
		doc, err := jpquakecrawler.GetJapanEarthquakeDetailDoc(eqTime)
		if err != nil {
			logrus.Errorf("[日本地震信息]在尝试获取%v发生的地震的过程中出现错误: %v", eqTime, err)
			continue
		}
		detail, err := jpquakeparser.ParseJapanEarthquakeDetailDoc(eqTime, doc)
		if err != nil {
			logrus.Errorf("[日本地震信息]在尝试解析%v发生的地震的过程中出现错误: %v", eqTime, err)
			continue
		}
		err = repo.AddNewJapanEarthquake(detail)
		if err != nil {
			logrus.Errorf("[日本地震信息]在尝试添加%v发生的地震的过程中出现错误: %v", eqTime, err)
			continue
		}
		var magInfo string
		var intensityInfo string
		if detail.Magnitude == nil {
			magInfo = "暂无地震规模情报, "
		} else {
			magInfo = fmt.Sprintf("地震规模%v级, ", *detail.Magnitude)
		}
		if detail.MaxIntensity == nil {
			intensityInfo = "暂无最大震度情报"
		} else {
			intensityInfo = fmt.Sprintf("最大震度为%v", *detail.MaxIntensity)
		}
		logrus.Infof("[日本地震信息]新增一条于%v发生在%v的地震, %v%v", detail.OccurTime, detail.Center, magInfo, intensityInfo)
	}
}

func JapanEEWCrawlTask() {
	layout := "20060102150405"
	queryTime := util.GetCurrentJapanTime().Format(layout)
	rawData, err := jpeewcrawler.GetJapanEEW(queryTime)
	if err != nil {
		logrus.Errorf("[日本EEW]在获取EEW数据的过程中失败: %v", err)
		return
	}
	eewData, err := jpeewparser.ParseJapanEEWData(rawData)
	if err != nil {
		logrus.Errorf("[日本EEW]在解析EEW数据的过程中失败: %v", err)
		return
	}
	if eewData == nil {
		return
	}
	if eewData.ReportTime != lastEEWData.ReportTime || eewData.IsCancel != lastEEWData.IsCancel {
		logrus.Infof("[日本EEW]------------------------------------")
		logrus.Infof("[日本EEW]紧急地震速报(%v) - 第%v报", eewData.AlertFlg, eewData.ReportNum)
		if eewData.IsFinal {
			logrus.Infof("[日本EEW]**此报为最终报**")
		}
		if eewData.IsCancel {
			logrus.Info("[日本EEW]**本紧急地震速报已取消**")
		}
		if eewData.IsTraining {
			logrus.Info("[日本EEW]**本紧急地震速报为训练**")
		}
		logrus.Infof("[日本EEW]震源地: %v (%v, %v)", eewData.RegionName, eewData.Latitude, eewData.Longitude)
		logrus.Infof("[日本EEW]发震时间: %v", eewData.OriginTime)
		logrus.Infof("[日本EEW]地震规模：%v", *eewData.Magnitude)
		logrus.Infof("[日本EEW]震源深度：%v", eewData.Depth)
		logrus.Infof("[日本EEW]预计最大震度：%v", eewData.CalcIntensity)
		logrus.Infof("[日本EEW]报告时间: %v", util.GetCurrentJapanTime().Format("2006-01-02T15:04:05-0700"))
		logrus.Infof("[日本EEW]------------------------------------")
		if err := repo.AddJapanEEWRecord(eewData); err != nil {
			logrus.Errorf("[日本EEW]在向SQLite数据库写入EEW信息的过程中发生错误: %v", err)
		}
		if config.Cfg.Redis.Enable {
			if err := repo.RPushJapanEEWRecord(eewData); err != nil {
				logrus.Errorf("[日本EEW]在向Redis写入EEW信息的过程中发生错误: %v", err)
			}
		}
	}
	lastEEWData = eewData
}
