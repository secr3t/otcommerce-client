package client

import (
	"encoding/json"
	"github.com/secr3t/otcommerce-client/model"
	"io/ioutil"
	"net/http"
)

const (
	host        = "otapi.net/service-json/"
	searchItems = "item_search"
)

type SearchClient struct {
	ApiKey string
}

func NewSearchClient(apiKey string) *SearchClient {
	return &SearchClient{
		ApiKey: apiKey,
	}
}

func (c *SearchClient) SearchItems(param SearchParam) model.Search {
	query := param.ToQuery()

	uri := GetUri(host, searchItems, query)

	req, _ := http.NewRequest("GET", uri, nil)


	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var search model.Search

	json.Unmarshal(body, &search)

	rateLimit := model.FromHeader(res.Header)
	search.RateLimit = rateLimit

	return search
}
