package main

import (
	"bufio"
	"flag"
	"fmt"
	"lucid/cron"
	"os"
	"strings"
	"time"
)

func main() {
	currentTime := flag.String(
		"current-time",
		time.Now().Format(cron.TimeFormatHHMM),
		"current timestamp in string e.g. 16:10",
	)
	flag.Parse()

	endSTDIN := false
	texts := ""
	counter := 0
	fmt.Println("+================================================+")
	fmt.Println("| To exit press return [Enter] without any entry |")
	fmt.Println("+================================================+")

	for !endSTDIN {
		counter++
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Please enter a single cron job (%d): \n", counter)
		text, _ := reader.ReadString('\n')

		texts += text
		if text == "\n" {
			endSTDIN = true
		}
	}

	collection := strings.Split(texts, "\n")
	for _, cronJob := range collection {
		if cronJob == "" {
			continue
		}

		job, err := cron.NewJob().ParseSingle(cronJob)
		if err != nil {
			panic(err)
		}

		nextCron, err := job.NextSchedule(*currentTime)
		if err != nil {
			panic(err)
		}

		_, _ = os.Stderr.WriteString(string(nextCron) + "\n")
	}
}