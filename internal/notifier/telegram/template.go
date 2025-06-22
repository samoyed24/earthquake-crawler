package telegram

import (
	"bytes"
	"earthquake-crawler/internal/model"
	"fmt"
)

func RenderJapanEarthquakeTGTemplate(data *model.JapanEarthquakeDetail) (string, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("[日本地震情报更新]%v发生地震\n", data.Center))
	buf.WriteString(fmt.Sprintf("地震发生时间: %v\n", data.OccurTime))
	buf.WriteString(fmt.Sprintf("震中: %v(%v/%v)\n", data.Center, data.Latitude, data.Longitude))
	buf.WriteString("规模: ")
	if data.Magnitude != nil {
		buf.WriteString(fmt.Sprintf("%v\n", *data.Magnitude))
	} else {
		buf.WriteString("未知\n")
	}
	buf.WriteString("最大震度: ")
	if data.MaxIntensity != nil {
		buf.WriteString(fmt.Sprintf("%v\n", *data.MaxIntensity))
	} else {
		buf.WriteString("未知\n")
	}
	buf.WriteString("震源深度: ")
	if data.Depth != nil {
		buf.WriteString(fmt.Sprintf("%v\n", *data.Depth))
	} else {
		buf.WriteString("未知\n")
	}
	buf.WriteString(fmt.Sprintf("地震相关信息: %v\n", data.Info))
	for _, locationReport := range data.LocationReports {
		buf.WriteString(fmt.Sprintf("%v: \n", locationReport.Intensity))
		for _, location := range locationReport.Locations {
			buf.WriteString(fmt.Sprintf("%v ", location.Prefecture))
			for _, subarea := range location.Subareas {
				buf.WriteString(fmt.Sprintf("%v ", subarea))
			}
			buf.WriteString(fmt.Sprintf("\n"))
		}
		buf.WriteString(fmt.Sprintf("\n"))
	}
	return buf.String(), nil
}

func RenderJapanEEWTGTemplate(data *model.JapanEEWData) (string, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("[日本EEW(%v)]%v发生%v级地震, 预估最大震度%v\n", data.AlertFlg, data.RegionName, *data.Magnitude, data.CalcIntensity))
	buf.WriteString(fmt.Sprintf("紧急地震速报(%v) - 第%v报\n", data.AlertFlg, data.ReportNum))
	if data.IsFinal {
		buf.WriteString("**本报为最终报**\n")
	}
	if data.IsTraining {
		buf.WriteString("**训练用**\n")
	}
	if data.IsCancel {
		buf.WriteString("**紧急地震速报已取消**\n")
	}
	buf.WriteString(fmt.Sprintf("推测震中: %v(%v/%v)\n", data.RegionName, data.Latitude, data.Longitude))
	buf.WriteString(fmt.Sprintf("推测最大震度: %v\n", data.CalcIntensity))
	buf.WriteString(fmt.Sprintf("推测规模: %v\n", *data.Magnitude))
	buf.WriteString(fmt.Sprintf("推测震源深度: %v\n", data.Depth))
	buf.WriteString(fmt.Sprintf("地震发生时间: %v\n", data.OriginTime))
	return buf.String(), nil
}
