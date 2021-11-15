package models

import "time"

type SearchRequest struct {
	Cacheable            bool
	AirlineCode          string
	DepartureAirportCode string
	ArrivalAirportCode   string
	DepartureDateTime    time.Time
	ArrivalDateTime      time.Time
	RoundTrip            bool
	BookingTime          time.Time
}

type FlightCacheSearchQuery struct {
	DepartureDateTimeInUtc string `json:"departure_date_time_in_utc"`
	AirlineCode            string `json:"airline_code"`
	BookingTimeInUtc       string `json:"booking_time_in_utc"`
	Origin                 string `json:"origin"`
	Destination            string `json:"destination"`
	JourneyType            string `json:"journeyType"`
}

type SearchResponse struct {
	Cacheable   bool   `json:"cacheable"`
	AirlineCode string `json:"airlineCode"`
}

type KnowledgeBaseForCacheRule struct {
	Name    string
	Version string
}
