package klm

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Airline struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Airport struct {
	Code string `json:"code"`
}

type LegTime struct {
	Scheduled       time.Time `json:"scheduled"`
	LatestPublished time.Time `json:"latestPublished"`
	Estimated       struct {
		Value time.Time `json:"value"`
	} `json:"estimated"`
}

type LegInformation struct {
	Airport Airport `json:"airport"`
	Times   LegTime `json:"times"`
}

type FlightLeg struct {
	StatusName           string         `json:"statusName"`
	PublishedStatus      string         `json:"publishedStatus"`
	DepartureInformation LegInformation `json:"departureInformation"`
	ArrivalInformation   LegInformation `json:"arrivalInformation"`
}

type FlightStatus struct {
	FlightNumber int         `json:"flightNumber"`
	Route        []string    `json:"route"`
	Airline      Airline     `json:"airline"`
	FlightLegs   []FlightLeg `json:"flightLegs"`
}

func (klm *KLM) FlightStatusList(start time.Time, end time.Time) ([]FlightStatus, error) {

	fs := []FlightStatus{}

	page := 0
	remaining_pages := 1

	for remaining_pages > 0 {
		url := fmt.Sprintf("/flightstatus?startRange=%s&endRange=%s&pageSize=%d&pageNumber=%d",
			start.UTC().Format(time.RFC3339),
			end.UTC().Format(time.RFC3339),
			100,
			page)

		resp, err := klm.sh.Get(url)
		if err != nil {
			return []FlightStatus{}, err
		}

		d := struct {
			FlightStatus []FlightStatus `json:"operationalFlights"`
			Page         struct {
				PageSize   int `json:"pageSize"`
				PageNumber int `json:"pageNumber"`
				FullCount  int `json:"fullCount"`
				PageCount  int `json:"pageCount"`
				TotalPages int `json:"totalPages"`
			} `json:"page"`
		}{}

		if err := json.Unmarshal([]byte(resp), &d); err != nil {
			return []FlightStatus{}, err
		}

		log.Println(d.Page)
		remaining_pages = d.Page.TotalPages - d.Page.PageNumber - 1
		log.Printf("Remaining pages: %d\n", remaining_pages)
		page++

		for _, s := range d.FlightStatus {
			fs = append(fs, s)
		}

	}

	return fs, nil

}
