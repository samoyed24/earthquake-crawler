package storage

func InitJapanEarthquakeDB() error {
	initSQL := `
		CREATE TABLE IF NOT EXISTS japan_earthquake_record (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			earthquake_time TEXT NOT NULL UNIQUE,
			earthquake_detail TEXT NOT NULL
		)
	`
	_, err := DB.Exec(initSQL)
	if err != nil {
		return err
	}
	return nil
}

func InitJapanEEWDB() error {
	initSQL := `
		CREATE TABLE IF NOT EXISTS japan_eew_record (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			report_time TEXT NOT NULL UNIQUE,
			eew_detail TEXT NOT NULL
		)
	`
	_, err := DB.Exec(initSQL)
	if err != nil {
		return err
	}
	return nil
}
