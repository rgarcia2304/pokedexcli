package pokeapi


import(
	"net/http"
	"time"
	"encoding/json"
	"io"
	"github.com/rgarcia2304/pokedexcli/internal/pokecache"
)

type LocationAreaResponse struct{
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}`json:"results"`
}


type ExploreAreaResponse struct {
	               
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`    
}

type EncounterMethodRate struct {
	EncounterMethod Location                           `json:"encounter_method"`
	VersionDetails  []EncounterMethodRateVersionDetail `json:"version_details"` 
}

type Location struct {
	Name string `json:"name"`
}

type EncounterMethodRateVersionDetail struct {
	Rate    int64    `json:"rate"`   
	Version Location `json:"version"`
}

type Name struct {
	Language Location `json:"language"`
	Name     string   `json:"name"`    
}

type PokemonEncounter struct {
	Pokemon        Location                        `json:"pokemon"`        
}

type PokemonEncounterVersionDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int64             `json:"max_chance"`       
	Version          Location          `json:"version"`          
}

type EncounterDetail struct {
	Chance          int64         `json:"chance"`          
	ConditionValues []interface{} `json:"condition_values"`
	MaxLevel        int64         `json:"max_level"`       
	Method          Location      `json:"method"`          
	MinLevel        int64         `json:"min_level"`       
}




type Client struct{
	httpClient http.Client
	cache *pokecache.Cache
}

func NewClient(timeout time.Duration, cache *pokecache.Cache) Client{
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}



func (c *Client) ListLocations(url string)(LocationAreaResponse, error){
	

	//Before making request check if entry is in cache 
	locationresp := LocationAreaResponse{}
	cacheResponse, found := c.cache.Get(url)
	if found{
		//will get the raw bytes from the response
		if err := json.Unmarshal(cacheResponse, &locationresp); err != nil{
			return LocationAreaResponse{}, err
		}

		return locationresp, nil
	}

	client := c.httpClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return LocationAreaResponse{}, err
	}

	res, err := client.Do(req)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	
	if err := json.Unmarshal(body, &locationresp); err != nil{
		return LocationAreaResponse{}, err
	}

	//cache the response
	c.cache.Add(url, body)
	
	return locationresp, nil

}

func (c *Client) ExploreArea(url string) (ExploreAreaResponse, error){
	areaResp := ExploreAreaResponse{}
	cacheResp, found := c.cache.Get(url)

		if found{
		//will get the raw bytes from the response
		if err := json.Unmarshal(cacheResp, &areaResp); err != nil{
			return ExploreAreaResponse{}, err
		}

		return areaResp, nil
	}

	client := c.httpClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return ExploreAreaResponse{}, err
	}

	res, err := client.Do(req)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	
	if err := json.Unmarshal(body, &areaResp); err != nil{
		return ExploreAreaResponse{}, err
	}

	//cache the response
	c.cache.Add(url, body)
	
	return areaResp, nil


}
