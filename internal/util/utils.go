package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func GetCurrentJapanTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	tokyoTime := time.Now().In(loc)
	return tokyoTime
}

func GetTokyoLocation() *time.Location {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("无法使用时区信息: %v", err)
	}
	return loc
}

func ToBool(v interface{}) (bool, error) {
	switch val := v.(type) {
	case bool:
		return val, nil
	case string:
		s := strings.ToLower(strings.TrimSpace(val))
		if s == "true" || s == "1" {
			return true, nil
		} else if s == "false" || s == "0" || s == "" {
			return false, nil
		} else {
			return false, fmt.Errorf("cannot convert string %q to bool", val)
		}
	case float64: // json.Unmarshal 把数字默认转成 float64
		return val != 0, nil
	case nil:
		return false, nil
	default:
		return false, fmt.Errorf("unsupported type: %T", val)
	}
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(out, in)
	return err
}
