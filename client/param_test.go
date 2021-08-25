package client

import (
	"testing"
)

func TestSearchItemsParameters_Ppath(t *testing.T) {
	p := &SearchItemsParameters{}
	p.Ppath("22200:29382")

	t.Log(p.ToXml())
}

func TestNewSearchParam(t *testing.T) {
	p := NewParams().
		Q("shoes")

	t.Log(p.ToXml())
}

func TestSearchParamFromUri(t *testing.T) {
	uri := `https://s.taobao.com/search?spm=a2141.241046-cn.lead_cate3.d1_Clothes.41ca6f11kHviFH&q=T%E6%81%A4%E7%94%B7&imgfile=&js=1&stats_click=search_radio_all%3A1&initiative_id=staobaoz_20210402&ie=utf8&cps=yes&ppath=20000%3A719106774&cat=50354021&sort=default&filter=reserve_price%5B20%2C%5D`
	sp := SearchParamFromUri(0, uri)

	t.Log(sp)
}

func TestSearchParam_ToQuery(t *testing.T) {
	uri := `https://s.taobao.com/search?spm=a2141.241046-cn.lead_cate3.d1_Clothes.41ca6f11kHviFH&q=T%E6%81%A4%E7%94%B7&imgfile=&js=1&stats_click=search_radio_all%3A1&initiative_id=staobaoz_20210402&ie=utf8&cps=yes&ppath=20000%3A719106774&cat=50354021&sort=default&filter=reserve_price%5B20%2C60.8%5D`
	sp := SearchParamFromUri(0, uri)

	t.Log(sp.ToQuery("xxxxx-xxxx-xxx-xxxx"))
}
