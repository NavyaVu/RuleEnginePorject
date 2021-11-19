package models

import (
	"log"
	"time"
)

func (p *SearchResponse) AddDays(inputTime time.Time, days int64) time.Time {
	log.Println("adding days ", days)
	return inputTime.AddDate(0, 0, int(days))
}
