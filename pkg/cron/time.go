package cron

import (
	"strconv"
	"time"
)

const (
	TimeFormatHHMM   = "15:04" //HH:MM
	TimeFormatSimple = "3:04"  //hh:mm
	Today            = "today"
	Tomorrow         = "tomorrow"
)

func setTime(slicedString string) (int, error) {
	if slicedString != "*" {
		parsedInt, err := strconv.ParseInt(slicedString, 10, 8)
		if err != nil {
			return 0, err
		}

		return int(parsedInt), nil
	}

	return 0, nil
}

func dayType(newHour time.Time, parsedMinutesHoursNow time.Time) string {
	day := Today
	if newHour.Before(parsedMinutesHoursNow) {
		day = Tomorrow
	}

	return day
}

func timeFormat(newHour time.Time) string {
	timeFormat := TimeFormatHHMM
	if newHour.Hour() < 10 {
		timeFormat = TimeFormatSimple
	}

	return timeFormat
}
