package task

import (
	"earthquake-crawler/internal/crawler/jpeewcrawler"
	japanSpider "earthquake-crawler/internal/crawler/jpquakecrawler"
	japanParser "earthquake-crawler/internal/parser/jpquakeparser"
	japanStorage "earthquake-crawler/internal/storage/jpquakestorage"
	"earthquake-crawler/internal/util"

	"github.com/sirupsen/logrus"
)

func JapanEarthquakeCrawlTask() {
	eqListDoc, err := japanSpider.GetJapanEarthquakeListDoc()
	if err != nil {
		logrus.Errorf("[日本气象厅地震信息]在获取地震信息列表的过程中失败: %v", err)
	}

	eqList, err := japanParser.ParseJapanEarthquakeListDoc(eqListDoc)
	if err != nil {
		logrus.Errorf("[日本气象厅地震信息]在解析地震列表HTML的过程中失败: %v", err)
	}

	eqNotExist, err := japanStorage.GetJapanEarthquakeNotInDB(eqList)
	if err != nil {
		logrus.Errorf("[日本气象厅地震信息]在查询数据库选择需要获取详情的地震列表的过程中失败：%v", err)
	}

	if len(eqNotExist) != 0 {
		logrus.Infof("[日本气象厅地震信息]解析到%d条未加入的地震信息，即将开始获取详情", len(eqNotExist))
	}

	for _, eqTime := range eqNotExist {
		doc, err := japanSpider.GetJapanEarthquakeDetailDoc(eqTime)
		if err != nil {
			logrus.Errorf("[日本气象厅地震信息]在尝试获取%v发生的地震的过程中出现错误: %v", eqTime, err)
			continue
		}
		detail, err := japanParser.ParseJapanEarthquakeDetailDoc(eqTime, doc)
		if err != nil {
			logrus.Errorf("[日本气象厅地震信息]在尝试解析%v发生的地震的过程中出现错误: %v", eqTime, err)
			continue
		}
		err = japanStorage.AddNewJapanEarthquake(detail)
		if err != nil {
			logrus.Errorf("[日本气象厅地震信息]在尝试添加%v发生的地震的过程中出现错误: %v", eqTime, err)
		}
		logrus.Infof("[日本气象厅地震信息]新增于%v发生于%v的%v级地震, 最大震度为%v", detail.OccurTime, detail.Center, detail.Magnitude, detail.MaxIntensity)
	}
}

func JapanEEWCrawlTask() {
	layout := "20060102150405"
	queryTime := util.GetCurrentJapanTime().Format(layout)
	eewData, err := jpeewcrawler.GetJapanEEW(queryTime)
	if err != nil {
		return
	}
	logrus.Info(eewData)
}
