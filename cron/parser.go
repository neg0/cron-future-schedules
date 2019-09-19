package cron

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Job struct {
	Hours      int
	Minutes    int
	BinaryPath string
	Recurring  bool
}

func NewJob() *Job {
	return &Job{}
}

func (j *Job) ParseSingle(cron string) (*Job, error) {
	slicedCronStr := strings.Fields(cron)
	if len(slicedCronStr) != 3 {
		return nil, errors.New("must be three length long. e.g. 30 * /bin/run_me_every_thirty_minutes")
	}

	minutes, err := setTime(slicedCronStr[0])
	if err != nil {
		return nil, err
	}
	j.Minutes = minutes

	hours, err := setTime(slicedCronStr[1])
	if err != nil {
		return nil, err
	}
	j.Hours = hours

	j.Recurring = false
	if j.Hours == 0 || j.Minutes == 0 {
		j.Recurring = true
	}
	j.BinaryPath = slicedCronStr[2]

	return j, nil
}

func (j *Job) formattedTime() string {
	parsedCronTimeStr := fmt.Sprintf("%d:%d", j.Hours, j.Minutes)
	if j.Hours < 10 {
		parsedCronTimeStr = fmt.Sprintf("0%d:%d", j.Hours, j.Minutes)
	}

	if j.Minutes < 10 {
		parsedCronTimeStr = fmt.Sprintf("%d:0%d", j.Hours, j.Minutes)
	}

	if j.Minutes < 10 && j.Hours < 10 {
		parsedCronTimeStr = fmt.Sprintf("0%d:0%d", j.Hours, j.Minutes)
	}

	return parsedCronTimeStr
}

func (j *Job) processNoneRecurringJob(parsedMinutesHoursNow time.Time, parsedCronTime time.Time) ([]byte, error) {
	if j.Hours == 0 && j.Minutes != 0 {
		newHour, err := time.Parse(TimeFormatHHMM, fmt.Sprintf("%d:%d", parsedMinutesHoursNow.Hour(), j.Minutes))
		if err != nil {
			return nil, err
		}

		return []byte(fmt.Sprintf("%s %s - %s\n",
			newHour.Format(TimeFormat(newHour)),
			DayType(newHour, parsedMinutesHoursNow),
			j.BinaryPath,
		)), nil
	}

	if j.Minutes == 0 && j.Hours != 0 {
		newHour, err := time.Parse(TimeFormatHHMM, fmt.Sprintf("%d:00", j.Hours))
		if err != nil {
			return nil, err
		}

		return []byte(fmt.Sprintf("%s %s - %s\n",
			newHour.Format(TimeFormat(newHour)),
			DayType(newHour, parsedMinutesHoursNow),
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
		result, err := j.processNoneRecurringJob(parsedMinutesHoursNow, parsedCronTime)
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
		parsedCronTime.Format(TimeFormat(parsedCronTime)),
		day,
		j.BinaryPath,
	)), nil
}