package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSearchResponse_AddDays(t *testing.T) {
	request := SearchRequest{
		Cacheable:            false,
		AirlineCode:          "",
		DepartureAirportCode: "",
		ArrivalAirportCode:   "",
		DepartureDateTime:    time.Now(),
		ArrivalDateTime:      time.Time{},
		RoundTrip:            false,
		BookingTime:          time.Time{},
	}

	response := SearchResponse{
		Cacheable:   false,
		AirlineCode: "AF",
	}
	assert.Equal(t, time.Now().Add(time.Hour*72).Day(), response.AddDays(request.DepartureDateTime, 3).Day(), "Add days "+
		"function is not working as expected")
}

func TestSearchResponse_IsPastDate(t *testing.T) {
	request := SearchRequest{
		Cacheable:            false,
		AirlineCode:          "AF",
		DepartureAirportCode: "",
		ArrivalAirportCode:   "",
		DepartureDateTime:    time.Now(),
		ArrivalDateTime:      time.Time{},
		RoundTrip:            false,
		BookingTime:          time.Time{},
	}

	response := SearchResponse{
		Cacheable:   false,
		AirlineCode: request.AirlineCode,
	}
	assert.Equal(t, true, response.IsPastDate(request.DepartureDateTime.Add(-1)), "Add days "+
		"function is not working as expected")
	assert.Equal(t, false, response.IsPastDate(request.DepartureDateTime.Add(time.Hour*24)), "Add days "+
		"function is not working as expected")
	assert.Equal(t, "AF", response.AirlineCode, "Add days "+
		"function is not working as expected")

}
