package telegram

import (
	"earthquake-crawler/internal/model"
	"testing"
)

func TestRenderJapanEarthquakeTGTemplate(t *testing.T) {
	MaxIntensity := "7"
	Magnitude := 7.6
	Depth := "10km"

	data := model.JapanEarthquakeDetail{
		EarthquakeTime: "20240101161010",
		OccurTime:      "2024-01-01T16:10:10+0900",
		Center:         "Ishikawa Noto",
		MaxIntensity:   &MaxIntensity,
		Magnitude:      &Magnitude,
		Depth:          &Depth,
		Latitude:       "36.96",
		Longitude:      "137.11",
		Info:           "Otsunami keihou, nigete",
		LocationReports: []model.LocationReport{
			{
				Intensity: "震度7",
				Locations: []model.Location{
					{
						Prefecture: "Ishikawa",
						Subareas: []string{
							"test1",
							"test2",
						},
					},
					{
						Prefecture: "Niigata",
						Subareas: []string{
							"test3",
							"test4",
						},
					},
				},
			},
			{
				Intensity: "震度6强",
				Locations: []model.Location{
					{
						Prefecture: "Ishikawa",
						Subareas: []string{
							"test5",
							"test6",
						},
					},
					{
						Prefecture: "Niigata",
						Subareas: []string{
							"test7",
							"test8",
						},
					},
				},
			},
		},
	}
	resultMsg, err := RenderJapanEarthquakeTGTemplate(&data)
	if err != nil {
		t.Errorf("[测试]在渲染JPQuake信息的过程中发送错误: %v", err)
	}
	t.Logf("[测试]生成信息结果: %v", resultMsg)
}

func TestRenderJapanEEWTGTemplate(t *testing.T) {
	magn := 7.6
	data := model.JapanEEWData{
		ReportTime:    "2024-01-01T16:10:10+0900",
		RegionName:    "Ishikawa",
		OriginTime:    "20240101161010",
		ReportID:      "20240101161010",
		IsCancel:      true,
		IsTraining:    true,
		IsFinal:       true,
		Depth:         "10km",
		Longitude:     1,
		Latitude:      2,
		CalcIntensity: "7",
		Magnitude:     &magn,
		ReportNum:     10,
		AlertFlg:      "Keihou",
	}
	resultMsg, err := RenderJapanEEWTGTemplate(&data)
	if err != nil {
		t.Errorf("[测试]在渲染JPEEW信息的过程中发送错误: %v", err)
	}
	t.Logf("[测试]生成信息结果: %v", resultMsg)
}
