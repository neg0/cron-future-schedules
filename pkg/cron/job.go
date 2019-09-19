package cron

import (
	"errors"
	"strings"
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

// Factory method to create Job from string
func (j *Job) FromSingleLine(cron string) (*Job, error) {
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
