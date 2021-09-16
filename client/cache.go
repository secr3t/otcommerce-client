package client

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	resetCache = "ResetInstanceCaches"
)

type ResetCacheClient struct {
	ApiKey string
}

func NewResetCacheClient(apiKey string) *ResetCacheClient {
	return &ResetCacheClient{
		ApiKey: apiKey,
	}
}

func (c *ResetCacheClient) ResetCaches() bool {
	query := url.Values{}
	query.Set("instanceKey", c.ApiKey)
	uri := GetUri(host, resetCache, query.Encode())

	req, _ := http.NewRequest("GET", uri, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return gjson.ParseBytes(body).Get("ErrorCode").String() == "Ok"
}
