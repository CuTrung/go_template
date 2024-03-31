package utils

import (
	"fmt"
	"time"
)

func GetCurrentDate() string {
	t := time.Now()
	return fmt.Sprintf("%d_%02d_%02d",
		t.Year(), t.Month(), t.Day(),
	)
}

func GetCurrentDateHasHour() string {
	t := time.Now()
	return fmt.Sprintf("%d_%02d_%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
