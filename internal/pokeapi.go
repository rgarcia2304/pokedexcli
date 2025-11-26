package pokeapi


import(
	"net/http"
	"fmt"
	"time"
	"encoding/json"
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
			Timeout: timout, 
		},
	}
}



func (c *Client) ListLocations(url string)(LocationAreaResponse, error){
	
	client := &c
	req, err := http.NewRequest("GET", url)
	if err != nil{
		return LocationAreaResponse{}, err
	}

	res, err := client.Do(req)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	locationresp := LocationAreaResponse{}
	if err := json.Unmarshal(body, &locationresp); err != nil{
		return LocationAreaResponse{}, err
	}
	
	return locationresp, nil

}
