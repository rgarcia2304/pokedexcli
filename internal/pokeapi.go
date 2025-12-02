package pokeapi


import(
	"net/http"
	"time"
	"encoding/json"
	"io"
)

type LocationAreaResponse struct{
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}`json:"results"`
}

type Client struct{
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client{
	return Client{
		httpClient: http.Client{
			Timeout: timeout, 
		},
	}
}



func (c *Client) ListLocations(url string, cfg *Config)(LocationAreaResponse, error){
	

	//Before making request check if entry is in cache 
	locationresp := LocationAreaResponse{}
	cacheResponse, found := cfg.Cache.Get(url)
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
	pokecache.Add(url, body)
	
	return locationresp, nil

}
