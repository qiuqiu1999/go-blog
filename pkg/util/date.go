package util

import (
	"fmt"
	"time"
)

func GetDate() string {
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	return fmt.Sprintf("%s-%s-%s", year, month, day)
}
