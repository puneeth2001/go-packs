package packs

import (
	"fmt"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"log"
	"encoding/json"
)
type Series struct {
	Page         int `json:"page"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Results      []struct {
		OriginalName     string   `json:"original_name"`
		GenreIds         []int    `json:"genre_ids"`
		Name             string   `json:"name"`
		Popularity       float64  `json:"popularity"`
		OriginCountry    []string `json:"origin_country"`
		VoteCount        int      `json:"vote_count"`
		FirstAirDate     string   `json:"first_air_date"`
		BackdropPath     string   `json:"backdrop_path"`
		OriginalLanguage string   `json:"original_language"`
		ID               int      `json:"id"`
		VoteAverage      float64  `json:"vote_average"`
		Overview         string   `json:"overview"`
		PosterPath       string   `json:"poster_path"`
	} `json:"results"`
}

func FetchSeriesID(name string) []byte{
	u, err := url.Parse("https://api.themoviedb.org/3/search/tv?page=1")
	if err != nil {
		log.Fatal(err)
	}
	q:= u.Query()
	q.Set("api_key","85024bf9f2db24e284e8959926cd3226")
	q.Set("language","en-US")
	q.Set("query",name)
	
	u.RawQuery = q.Encode()
	stringURL := u.String()
	url := stringURL
	
	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return (body)
	
	
	// Output: https://google.com/search?q=golang
}
func GetID(name string) int64{
	series := &Series{}
	var JSONData = FetchSeriesID(name)
	err := json.Unmarshal(JSONData, series)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(series.Results[0].ID)
	return (int64(series.Results[0].ID))
}
