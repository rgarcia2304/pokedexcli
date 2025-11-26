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



func (c *Client) ListLocations(url string)(LocationAreaResponse, error){
	
	client := c.httpClient
	req, err := http.NewRequest("GET", url, nil)
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
