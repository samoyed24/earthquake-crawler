package jpquakeparser

import (
	"earthquake-crawler/internal/model"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func splitLatLon(lagLon string) (string, string, error) {
	parts := strings.Split(lagLon, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("经纬度的格式不正确")
	}
	return parts[0], parts[1], nil
}

func splitLocations(locationsString string) []string {
	trimmedString := strings.TrimSpace(locationsString)
	return strings.Fields(trimmedString)
}

func getDetailFromDetailTable(tableSelection *goquery.Selection, detailStruct *model.JapanEarthquakeDetail) error {
	trs := tableSelection.Find("tbody tr")
	timeLayout := "2006年1月2日 15時04分ごろ"
	occurTimeTr := trs.Eq(0).Find("td").Eq(1).Find("small").First().Text()
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t, err := time.ParseInLocation(timeLayout, occurTimeTr, loc)
	if err != nil {
		return err
	}
	tTokyo := t.In(loc)
	detailStruct.OccurTime = tTokyo.Format("2006-01-02T15:04:05-07:00")
	detailStruct.Center = strings.TrimSpace(trs.Eq(1).Find("td").Eq(1).Find("small").First().Text())
	maxInte := strings.TrimSpace(trs.Eq(2).Find("td").Eq(1).Find("small").First().Text())
	if maxInte == "---" {
		detailStruct.MaxIntensity = nil
	} else {
		detailStruct.MaxIntensity = &maxInte
	}
	magnitude, err := strconv.ParseFloat(strings.TrimSpace(trs.Eq(3).Find("td").Eq(1).Find("small").First().Text()), 64)
	var mag *float64
	if err != nil { // 处理震级是"---"的情况，一般是火山喷发等特殊情况
		mag = nil
	} else {
		mag = &magnitude
	}
	detailStruct.Magnitude = mag
	depth := strings.TrimSpace(trs.Eq(4).Find("td").Eq(1).Find("small").First().Text())
	if depth == "---" {
		detailStruct.Depth = nil
	} else {
		detailStruct.Depth = &depth
	}

	latitude, longitude, err := splitLatLon(trs.Eq(5).Find("td").Eq(1).Find("small").First().Text())

	if err != nil {
		return err
	}

	detailStruct.Latitude = strings.TrimSpace(latitude)
	detailStruct.Longitude = strings.TrimSpace(longitude)
	detailStruct.Info = strings.TrimSpace(trs.Eq(6).Find("td").Eq(1).Find("small").First().Text())

	// fmt.Println(detailStruct)
	return nil
}

func getLocationFromLocationTable(tableSelection *goquery.Selection, detailStruct *model.JapanEarthquakeDetail) error {
	if tableSelection.Length() == 0 { // 有一些地震，并没有影响到日本国内，但是仍然会报，这时表格长度为0
		return nil
	}
	trs := tableSelection.ChildrenFiltered("tbody").ChildrenFiltered("tr")
	trs.Each(func(i int, rowSelection *goquery.Selection) {
		locationReport := new(model.LocationReport)
		locationReport.Intensity = strings.TrimSpace(rowSelection.Find("td").Eq(0).Find("small").First().Text())
		rowSelection.Find("td").Eq(1).Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
			location := new(model.Location)
			location.Prefecture = strings.TrimSpace(s.Find("td").Eq(0).Find("small").First().Text())
			location.Subareas = splitLocations(s.Find("td").Eq(1).Find("small").First().Text())
			if len(location.Subareas) > 0 {
				locationReport.Locations = append(locationReport.Locations, *location)
			}
		})
		detailStruct.LocationReports = append(detailStruct.LocationReports, *locationReport)
	})
	return nil
}

func ParseJapanEarthquakeDetailDoc(eqTime string, doc *goquery.Document) (*model.JapanEarthquakeDetail, error) {
	earthquakeDetail := new(model.JapanEarthquakeDetail)
	earthquakeDetail.EarthquakeTime = eqTime
	tables := doc.Find("#eqinfdtl table")
	detailTable := tables.Eq(0)
	locationTable := tables.Eq(1)
	err := getDetailFromDetailTable(detailTable, earthquakeDetail)
	if err != nil {
		return nil, err
	}
	err = getLocationFromLocationTable(locationTable, earthquakeDetail)
	if err != nil {
		return nil, err
	}
	return earthquakeDetail, nil
}
