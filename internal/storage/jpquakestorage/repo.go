package jpquakestorage

import (
	"earthquake-crawler/internal/model"
	"earthquake-crawler/internal/storage"
	"encoding/json"
	"fmt"
)

func GetJapanEarthquakeNotInDB(earthquakeList []string) ([]string, error) {
	if storage.DB == nil {
		return nil, fmt.Errorf("数据库未初始化")
	}
	querySQL := "SELECT earthquake_time FROM japan_earthquake_record WHERE earthquake_time IN (?"
	for i := 1; i < len(earthquakeList); i++ {
		querySQL += ",?"
	}
	querySQL += ")"
	args := make([]any, len(earthquakeList))
	for i, v := range earthquakeList {
		args[i] = v
	}
	rows, err := storage.DB.Query(querySQL, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var resultQuery []string
	for rows.Next() {
		var earthquake_time string
		if err := rows.Scan(&earthquake_time); err != nil {
			return nil, err
		}
		resultQuery = append(resultQuery, earthquake_time)
	}
	var resultList []string
	dbExistSet := make(map[string]struct{})
	for _, val := range resultQuery {
		dbExistSet[val] = struct{}{}
	}

	for _, val := range earthquakeList {
		if _, exists := dbExistSet[val]; !exists {
			resultList = append(resultList, val)
		}
	}
	return resultList, nil
}

func AddNewJapanEarthquake(earthquakeDetail *model.JapanEarthquakeDetail) error {
	if storage.DB == nil {
		return fmt.Errorf("数据库未初始化")
	}
	insertSQL := "INSERT INTO japan_earthquake_record (earthquake_time, earthquake_detail) VALUES (?, ?)"
	jsonBytes, err := json.Marshal(earthquakeDetail)
	if err != nil {
		return err
	}

	_, err = storage.DB.Exec(insertSQL, earthquakeDetail.EarthquakeTime, string(jsonBytes))
	if err != nil {
		return err
	}
	return nil
}
