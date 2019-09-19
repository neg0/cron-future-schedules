package cron

import (
	"strings"
	"testing"
)

var testCases = map[string]string{
	"30 1 /bin/run_me_daily":        "1:30 tomorrow - /bin/run_me_daily",
	"45 * /bin/run_me_hourly":       "16:45 today - /bin/run_me_hourly",
	"* * /bin/run_me_every_minute":  "16:10 today - /bin/run_me_every_minute",
	"* 14 /bin/run_me_every_minute": "14:00 tomorrow - /bin/run_me_every_minute",
	"* 19 /bin/run_me_sixty_times":  "19:00 today - /bin/run_me_sixty_times",
	"* 22 /bin/run_me_once_at_22":   "22:00 today - /bin/run_me_once_at_22",
	"10 10 /bin/run_me_daily_10_10": "10:10 tomorrow - /bin/run_me_daily_10_10",
	"* 9 /bin/run_me_daily_10_10":   "9:00 tomorrow - /bin/run_me_daily_10_10",
}

func TestParser(t *testing.T) {
	for mockString, expected := range testCases {
		actual, err := NewJob().ParseSingle(mockString)
		if err != nil {
			t.Error(err)
		}

		newActual, err := actual.NextSchedule("16:10")
		if err != nil {
			t.Error(err)
		}

		if !strings.Contains(string(newActual), expected) {
			t.Log("Actual " + string(newActual))
			t.Log("Expected " + expected)
			t.Fail()
		}
	}
}