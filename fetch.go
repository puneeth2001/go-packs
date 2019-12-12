package packs

import (
	"fmt"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"log"
	"strconv"
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

type Details struct {
	BackdropPath string `json:"backdrop_path"`
	CreatedBy    []struct {
		ID          int    `json:"id"`
		CreditID    string `json:"credit_id"`
		Name        string `json:"name"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"created_by"`
	EpisodeRunTime []int  `json:"episode_run_time"`
	FirstAirDate   string `json:"first_air_date"`
	Genres         []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage         string   `json:"homepage"`
	ID               int      `json:"id"`
	InProduction     bool     `json:"in_production"`
	Languages        []string `json:"languages"`
	LastAirDate      string   `json:"last_air_date"`
	LastEpisodeToAir struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		ID             int     `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int     `json:"show_id"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float64 `json:"vote_average"`
		VoteCount      int     `json:"vote_count"`
	} `json:"last_episode_to_air"`
	Name             string      `json:"name"`
	NextEpisodeToAir interface{} `json:"next_episode_to_air"`
	Networks         []struct {
		Name          string `json:"name"`
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"networks"`
	NumberOfEpisodes    int      `json:"number_of_episodes"`
	NumberOfSeasons     int      `json:"number_of_seasons"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalName        string   `json:"original_name"`
	Overview            string   `json:"overview"`
	Popularity          float64  `json:"popularity"`
	PosterPath          string   `json:"poster_path"`
	ProductionCompanies []struct {
		ID            int    `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	Seasons []struct {
		AirDate      string `json:"air_date"`
		EpisodeCount int    `json:"episode_count"`
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Overview     string `json:"overview"`
		PosterPath   string `json:"poster_path"`
		SeasonNumber int    `json:"season_number"`
	} `json:"seasons"`
	Status      string  `json:"status"`
	Type        string  `json:"type"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
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
	
}
func FetchTotalSeasons(name string) []byte{
	a16 := strconv.FormatInt(GetID(name),10)
	fmt.Println(string(a16))
	u, err := url.Parse("https://api.themoviedb.org/3/tv/"+string(a16)+"?language=en-US&api_key=jhyjh")
	if err != nil {
		log.Fatal(err)			
		}
	q:= u.Query()
	q.Set("api_key","85024bf9f2db24e284e8959926cd3226")
	q.Set("language","en-US")
	
	u.RawQuery = q.Encode()
	stringURL := u.String()
	url := stringURL
	
	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return (body)
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
func GetSeasons(name string) int64 {
	details := &Details{}
	var byteData = FetchTotalSeasons(name)
	err := json.Unmarshal(byteData, details)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(details.NumberOfSeasons)
	return (int64(details.NumberOfSeasons))
}

