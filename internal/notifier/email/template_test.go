package email

import (
	"earthquake-crawler/internal/model"
	"testing"
)

// 用于测试地震情报邮件是否能正常渲染的函数
func TestRenderJapanEarthquakeEmailTemplate(t *testing.T) {
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
	resultHtml, err := RenderJapanEarthquakeEmailTemplate(&data)
	if err != nil {
		t.Errorf("[测试]在渲染JPQuake邮件的过程中发送错误: %v", err)
	}
	t.Logf("[测试]生成邮件结果: %v", *resultHtml)
}
