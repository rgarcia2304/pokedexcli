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


type Location struct {
	Name string `json:"name"`
}

type PokemonEncounter struct {
	Pokemon        Location                        `json:"pokemon"`        
}

type Pokemon struct{
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct{
		BaseStat int `json:"base_stat"`
		Stat struct{
			Name string `json:"name"`
		}`json:"stat"`
	}`json:"stats"`
	Types []struct{
		Type struct{
			Name string `json:"name"`
		}`json:"type"`
	}`json:"types"`
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

func (c *Client) CatchPokemon(url string) (Pokemon, error){
	pokeResp := Pokemon{}
	cacheResp, found := c.cache.Get(url)

		if found{
		//will get the raw bytes from the response
		if err := json.Unmarshal(cacheResp, &pokeResp); err != nil{
			return Pokemon{}, err
		}

		return pokeResp, nil
	}

	client := c.httpClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return Pokemon{}, err
	}

	res, err := client.Do(req)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	
	if err := json.Unmarshal(body, &pokeResp); err != nil{
		return Pokemon{}, err
	}

	//cache the response
	c.cache.Add(url, body)
	
	return pokeResp, nil


}


