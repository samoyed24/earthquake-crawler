package app

import (
	"fmt"
	"japan-earthquake-webspider/internal/parser"
	"japan-earthquake-webspider/internal/spider"
	"japan-earthquake-webspider/internal/storage"

	log "github.com/sirupsen/logrus"
)

func RunEarthquakeListSpider() error {
	eqListDoc, err := spider.GetEarthquakeListDoc()
	if err != nil {
		return fmt.Errorf("在获取日本地震信息列表的过程中失败: %v", err)
	}

	eqList, err := parser.ParseEarthquakeListDoc(eqListDoc)
	if err != nil {
		return fmt.Errorf("在解析日本地震列表HTML的过程中失败: %v", err)
	}

	eqNotExist, err := storage.GetEarthquakeNotInDB(eqList)
	if err != nil {
		return fmt.Errorf("在查询数据库选择需要获取详情的日本地震列表的过程中失败：%v", err)
	}

	if len(eqNotExist) != 0 {
		log.Infof("解析到%d条未加入的日本地震信息，即将开始获取详情", len(eqNotExist))
	}

	for _, eqTime := range eqNotExist {
		doc, err := spider.GetEarthquakeDetailDoc(eqTime)
		if err != nil {
			log.Errorf("在尝试获取%v发生的日本地震的过程中出现错误: %v", eqTime, err)
			continue
		}
		detail, err := parser.ParseEarthquakeDetailDoc(eqTime, doc)
		if err != nil {
			log.Errorf("在尝试解析%v发生的日本地震的过程中出现错误：%v", eqTime, err)
			continue
		}
		err = storage.AddNewEarthquake(detail)
		if err != nil {
			return err
		}
		log.Infof("新增于%v发生于%v的%v级地震, 最大震度为%v", detail.OccurTime, detail.Center, detail.Magnitude, detail.MaxIntensity)
	}
	return nil
}
