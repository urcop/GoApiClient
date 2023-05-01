package GoAPIClient

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Client struct {
	client *http.Client
	apiKey string
}

func NewClient(apiKey string) (*Client, error) {
	return &Client{
		client: &http.Client{},
		apiKey: apiKey,
	}, nil
}

func (c *Client) GetBrawlers() ([]Brawler, error) {
	bearer := "Bearer " + os.Getenv("API_KEY")
	url := "https://api.brawlstars.com/v1/brawlers"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", bearer)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var b BrawlersResponse
	if err := json.Unmarshal(body, &b); err != nil {
		return nil, err
	}

	return b.Brawlers, nil
}

func (c *Client) GetBrawler(BrawlerId string) (Brawler, error) {
	bearer := "Bearer " + os.Getenv("API_KEY")
	url := fmt.Sprintf("https://api.brawlstars.com/v1/brawlers/%s", BrawlerId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", bearer)

	resp, err := c.client.Do(req)
	if err != nil {
		return Brawler{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Brawler{}, err
	}

	var b Brawler
	if err := json.Unmarshal(body, &b); err != nil {
		return Brawler{}, err
	}

	return b, nil
}
