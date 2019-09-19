package cron

import (
	"fmt"
	"time"
)

// Processing None Recurring Cron Jobs
func (j *Job) processRecurringJob(parsedMinutesHoursNow time.Time, parsedCronTime time.Time) ([]byte, error) {
	if j.Hours == 0 && j.Minutes != 0 {
		newHour, err := time.Parse(TimeFormatHHMM, fmt.Sprintf("%d:%d", parsedMinutesHoursNow.Hour(), j.Minutes))
		if err != nil {
			return nil, err
		}

		return []byte(fmt.Sprintf("%s %s - %s\n",
			newHour.Format(timeFormat(newHour)),
			dayType(newHour, parsedMinutesHoursNow),
			j.BinaryPath,
		)), nil
	}

	if j.Minutes == 0 && j.Hours != 0 {
		newHour, err := time.Parse(TimeFormatHHMM, fmt.Sprintf("%d:00", j.Hours))
		if err != nil {
			return nil, err
		}

		return []byte(fmt.Sprintf("%s %s - %s\n",
			newHour.Format(timeFormat(newHour)),
			dayType(newHour, parsedMinutesHoursNow),
			j.BinaryPath,
		)), nil
	}

	day := Tomorrow
	if parsedCronTime.Before(parsedMinutesHoursNow) {
		day = Today
	}

	return []byte(fmt.Sprintf("%s %s - %s\n",
		parsedMinutesHoursNow.Format(TimeFormatHHMM),
		day,
		j.BinaryPath,
	)), nil
}
