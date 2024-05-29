package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Call struct {
	From      string
	To        string
	StartTime time.Time
	EndTime   time.Time
}

func main() {
	file, err := os.Open("class017/call_time.csv")
	if err != nil {
		log.Printf("Error opening class017/call_time.csv: %v", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	var calls []Call

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Error reading csv: %v", err)
			return
		}

		startTime, _ := time.Parse(time.RFC3339, record[2])
		endTime, _ := time.Parse(time.RFC3339, record[3])

		call := Call{
			From:      record[0],
			To:        record[1],
			StartTime: startTime,
			EndTime:   endTime,
		}

		calls = append(calls, call)
	}

	name := "beth"

	// John == john == JOHN == jOHN

	var totalIncoming time.Duration
	var totalOutgoing time.Duration

	for _, call := range calls {
		if strings.EqualFold(call.From, name) {
			totalOutgoing = totalOutgoing + call.EndTime.Sub(call.StartTime)
		} else if strings.EqualFold(call.To, name) {
			totalIncoming += call.EndTime.Sub(call.StartTime)
		}
	}

	fmt.Printf("totalIncoming: %v\n", totalIncoming)
	fmt.Printf("totalOutgoing: %v\n", totalOutgoing)

}

// Take name of the person to query the call time from std IO
// If the name isn't available in the list, print 404
