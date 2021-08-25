package client

import (
	"github.com/secr3t/otcommerce-client/model"
	"testing"
)

const apiKey = "xxxxx"

func TestSearchClient_SearchItems(t *testing.T) {
	c := NewSearchClient(apiKey)

	uri := `https://s.taobao.com/search?spm=a2141.241046-cn.lead_cate3.d1_Clothes.41ca6f11kHviFH&q=T%E6%81%A4%E7%94%B7&imgfile=&js=1&stats_click=search_radio_all%3A1&initiative_id=staobaoz_20210402&ie=utf8&cps=yes&ppath=20000%3A719106774&cat=50354021&sort=default&filter=reserve_price%5B20%2C60.8%5D`
	sp := SearchParamFromUri(0, uri)

	t.Log(model.ToJson(sp))

	search, err := c.SearchItems(sp)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(model.ToJson(search))
}
