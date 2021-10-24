package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	apikey := "faf7e5bb"
	pagination := "2"
	searchword := "batman"
	callList(apikey, pagination, searchword)
	fmt.Println(" ")
	fmt.Println(" ")
	imdbID := "tt4853102"
	callSingle(apikey, imdbID)
}

func callList(apikey string, pagination string, searchword string) {
	link := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%s", apikey, searchword, pagination)
	message := fmt.Sprintf("Calling %s . . .", link)
	fmt.Println(message)
	fmt.Println(" ")
	response := getAPI(link)
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result ResponseList
	json.Unmarshal(responseData, &result)
	fmt.Println(result)
}

func callSingle(apikey string, imdbID string) {
	link := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s", apikey, imdbID)
	message := fmt.Sprintf("Calling %s . . .", link)
	fmt.Println(message)
	fmt.Println(" ")
	response := getAPI(link)
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var result ResponseSingle
	json.Unmarshal(responseData, &result)
	fmt.Println(result)
}

func getAPI(link string) http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	return *resp
}

type DataList struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type ResponseList struct {
	Search      []DataList `json:"Search"`
	TotalResult string     `json:"TotalResult"`
	Response    string     `json:"Response"`
}

type ResponseSingle struct {
	Title      string                  `json:"Title"`
	Year       string                  `json:"Year"`
	ImdbID     string                  `json:"imdbID"`
	Type       string                  `json:"Type"`
	Poster     string                  `json:"Poster"`
	Ratings    []ResponseRatingsSingle `json:"Ratings"`
	Released   string                  `json:"Released"`
	Runtime    string                  `json:"Runtime"`
	Genre      string                  `json:"Genre"`
	Director   string                  `json:"Director"`
	Writer     string                  `json:"Writer"`
	Actors     string                  `json:"Actors"`
	Plot       string                  `json:"Plot"`
	Language   string                  `json:"Language"`
	Country    string                  `json:"Country"`
	Awards     string                  `json:"Awards"`
	Metascore  string                  `json:"Metascore"`
	imdbRating string                  `json:"imdbRating"`
	imdbVotes  string                  `json:"imdbVotes"`
	DVD        string                  `json:"DVD"`
	BoxOffice  string                  `json:"BoxOffice"`
	Production string                  `json:"Production"`
	Website    string                  `json:"Website"`
	Response   string                  `json:"Response"`
}

type ResponseRatingsSingle struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
