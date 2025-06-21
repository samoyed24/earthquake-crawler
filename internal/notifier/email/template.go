package email

import (
	"bytes"
	"earthquake-crawler/internal/model"
	"embed"
	"html/template"
)

//go:embed templates/*
var templatesFS embed.FS

// TODO 这里以后重构可以改成通用函数

func RenderJapanEarthquakeEmailTemplate(data *model.JapanEarthquakeDetail) (*string, error) {
	tmpl, err := template.ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, "jpquake.html", data)
	if err != nil {
		return nil, err
	}
	resultStr := buf.String()
	return &resultStr, nil
}

func RenderJapanEEWEmailTemplate(data *model.JapanEEWData) (*string, error) {
	tmpl, err := template.ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, "jpeew.html", data)
	if err != nil {
		return nil, err
	}
	resultStr := buf.String()
	return &resultStr, nil
}
