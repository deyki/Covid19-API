package service

import (
	"encoding/json"
	"net/http"
)


type Country struct {
	ID             string `json:"id"`
	Country        string `json:"country"`
	CountryCode    string `json:"countryCode"`
	Slug           string `json:"slug"`
	NewConfirmed   int    `json:"newConfirmed"`
	TotalConfirmed int    `json:"totalConfirmed"`
	NewDeaths      int    `json:"newDeaths"`
	TotalDeaths    int    `json:"totalDeaths"`
	NewRecovered   int    `json:"newRecovered"`
	TotalRecovered int    `json:"totalRecovered"`
	Date           string `json:"date"`
	Premium        string `json:"premium"`
}


type Global struct {
	ID             int    `json:"id"`
	NewConfirmed   int    `json:"newConfirmed"`
	TotalConfirmed int    `json:"totalConfirmed"`
	NewDeaths      int    `json:"newDeaths"`
	TotalDeaths    int    `json:"totalDeaths"`
	NewRecovered   int    `json:"newRecovered"`
	TotalRecovered int    `json:"totalRecovered"`
	Date           string `json:"date"`
}


type Statistic struct {
	ID      string    `json:"id"`
	Message string    `json:"message"`
	Global  Global    `json:"global"`
	Country []Country `json:"countries"`
}


type ErrorMessage struct {
	HttpStatus int
	Message    string
}


var countries []Country


func DownloadDataFromUrl() *ErrorMessage {
	response, err := http.Get("https://api.covid19api.com/summary")
	if err != nil {
		return &ErrorMessage{
			HttpStatus: 	http.StatusInternalServerError,
			Message: 	"Failed to get data from url",
		}
	}

	var statistic Statistic

	json.NewDecoder(response.Body).Decode(&statistic)

	countries = append(countries, statistic.Country...)

	return nil
}


func GetCountryByCountryCode(countryCode string) (*Country, *ErrorMessage) {
	for _, country := range countries {
		if country.CountryCode == countryCode {
			return &country, nil
		}
	}

	return nil, &ErrorMessage{HttpStatus: http.StatusBadRequest, Message: "Invalid country code or caching is in progress! Try again later."}
}
