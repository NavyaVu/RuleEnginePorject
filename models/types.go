package models

import "time"

type SearchRequest struct {
	Cacheable             bool
	AirlineCode           string
	DepartureAirportCode  string
	ArrivalAirportCode    string
	DepartureDateTime     time.Time
	ArrivalDateTime       time.Time
	RoundTrip             bool
	BookingTime           time.Time
	RequestType           string
	RuleGroup             string
	IsProcessingCompleted bool
}

type FlightCacheSearchQuery struct {
	DepartureDateTimeInUtc string `json:"departure_date_time_in_utc"`
	AirlineCode            string `json:"airline_code"`
	BookingTimeInUtc       string `json:"booking_time_in_utc"`
	Origin                 string `json:"origin"`
	Destination            string `json:"destination"`
	JourneyType            string `json:"journeyType"`
	RequestType            string `json:"requestType"`
	RequestGroup           string `json:"requestGroup"`
}

type SearchResponse struct {
	Cacheable               bool   `json:"cacheable"`
	IsPeakworkData          bool   `json:"isPeakworkData"`
	AirlineCode             string `json:"airlineCode"`
	Destinations            string `json:"destinations"`
	Origins                 string `json:"origins"`
	PeakworkEarliestDepDate string `json:"peakworkEarliestDepDate"`
	PeakworkLatestDepDate   string `json:"peakworkLatestDepDate"`
	PeakworkFltDurations    string `json:"peakworkFltDurations"`
	IsProcessingCompleted   bool   `json:"isProcessingCompleted"`
}

type KnowledgeBaseForCacheRule struct {
	Name    string
	Version string
}
