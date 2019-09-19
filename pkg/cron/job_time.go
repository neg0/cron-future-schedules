package cron

import "fmt"

// Prepares cron entries  times to be compatible with defined time format HH:MM 00:00
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
