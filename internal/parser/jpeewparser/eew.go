package jpeewparser

import (
	"earthquake-crawler/internal/model"
	"earthquake-crawler/internal/util"
	"fmt"
	"strconv"
	"time"
)

func ParseJapanEEWData(rawData *model.RawJapanEEWData) (*model.JapanEEWData, error) {
	// 没有接收到EEW就直接返回空
	if len(rawData.RegionName) == 0 {
		return nil, nil
	}
	eewData := new(model.JapanEEWData)
	eewData.AlertFlg = *rawData.AlertFlg
	eewData.CalcIntensity = rawData.Calcintensity
	eewData.Depth = rawData.Depth
	var err error
	eewData.IsCancel, err = util.ToBool(rawData.IsCancel)
	if err != nil {
		return nil, fmt.Errorf("解析is_cancel字段时出现错误: %v", err)
	}
	eewData.IsFinal, err = util.ToBool(rawData.IsFinal)
	if err != nil {
		return nil, fmt.Errorf("解析is_final字段时出现错误: %v", err)
	}
	eewData.IsTraining, err = util.ToBool(rawData.IsTraining)
	if err != nil {
		return nil, fmt.Errorf("解析is_training字段时出现错误: %v", err)
	}
	eewData.Latitude, err = strconv.ParseFloat(rawData.Latitude, 64)
	if err != nil {
		return nil, fmt.Errorf("解析latitude字段时出现错误: %v", err)
	}
	eewData.Longitude, err = strconv.ParseFloat(rawData.Longitude, 64)
	if err != nil {
		return nil, fmt.Errorf("解析longitude字段时出现错误: %v", err)
	}
	tmpMagn, err := strconv.ParseFloat(rawData.Magunitude, 64)
	if err != nil {
		eewData.Magnitude = nil
	} else {
		eewData.Magnitude = &tmpMagn
	}
	eewData.RegionName = rawData.RegionName
	eewData.ReportNum, err = strconv.Atoi(rawData.ReportNum)
	if err != nil {
		return nil, fmt.Errorf("解析report_num字段时出现错误: %v", err)
	}
	layout := "2006/01/02 15:04:05"
	simpleLayout := "20060102150405"
	t, err := time.Parse(layout, rawData.ReportTime)
	if err != nil {
		return nil, fmt.Errorf("解析report_time字段时出现错误: %v", err)
	}
	eewData.ReportTime = t.Format(simpleLayout)
	eewData.ReportID = rawData.ReportID
	tokyoTimeLoc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, fmt.Errorf("在读取时区过程中出现错误: %v", err)
	}
	t, err = time.ParseInLocation(simpleLayout, rawData.OriginTime, tokyoTimeLoc)
	if err != nil {
		return nil, fmt.Errorf("解析origin_time字段时出现错误: %v", err)
	}
	eewData.OriginTime = t.Format("2006-01-02T15:04:05-0700")
	return eewData, nil
}
