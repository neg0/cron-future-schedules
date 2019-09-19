package cron

import (
	"fmt"
	"time"
)

// Returns the next time the command is scheduled to run
func (j *Job) NextSchedule(currentTime string) ([]byte, error) {
	parsedMinutesHoursNow, err := time.Parse(TimeFormatHHMM, currentTime)
	if err != nil {
		return nil, err
	}

	parsedCronTime, err := time.Parse(TimeFormatHHMM, j.formattedTime())
	if err != nil {
		return nil, err
	}

	if j.Recurring {
		result, err := j.processRecurringJob(parsedMinutesHoursNow, parsedCronTime)
		if err != nil {
			return nil, err
		}

		if result != nil {
			return result, nil
		}
	}

	day := Today
	if parsedCronTime.Before(parsedMinutesHoursNow) {
		day = Tomorrow
	}

	return []byte(fmt.Sprintf(
		"%s %s - %s\n",
		parsedCronTime.Format(timeFormat(parsedCronTime)),
		day,
		j.BinaryPath,
	)), nil
}
