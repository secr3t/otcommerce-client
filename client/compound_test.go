package client

import (
	"github.com/secr3t/otcommerce-client/model"
	"testing"
	"time"
)

func TestCompoundClient_SearchAndGetDetails(t *testing.T) {
	c := NewCompoundClient("c1ec36ce-0572-43f2-b536-7467d2f65190", 80)

	p := SearchParamFromUri(0, "https://s.taobao.com/search?q=%E7%BA%B1%E5%B7%BE%E5%9B%B4%E5%B7%BE%E6%98%A5%E7%A7%8B%E5%AD%A3%E8%96%84%E5%A6%88%E5%A6%88&imgfile=&js=1&stats_click=search_radio_all%3A1&initiative_id=staobaoz_20210820&ie=utf8&bcoffset=4&ntoffset=4&p4ppushleft=2%2C48&s=0")

	start := time.Now()
	detailChan, err := c.SearchAndGetDetailsMultiRequestOneTime(&p)

	if err != nil {
		t.Fatal(err)
	}


	var details []model.DetailItem

	for detail := range detailChan {
		details = append(details, detail)
	}

	elapsed := time.Since(start)
	t.Log(elapsed)
	t.Log(len(details))
}

func TestCompoundClient_SearchAndGetDetail(t *testing.T) {
	c := NewCompoundClient("c1ec36ce-0572-43f2-b536-7467d2f65190", 80)
	p := SearchParamFromUri(0, "https://s.taobao.com/search?q=T%E6%81%A4%E7%94%B7&imgfile=&js=1&stats_click=search_radio_all%3A1&initiative_id=staobaoz_20210825&ie=utf8&cps=yes&ppath=20000%3A661710970")

	start := time.Now()

	details, err := c.SearchAndGetDetail(&p)

	if err != nil {
		t.Fatal(err)
	}

	elapsed := time.Since(start)
	t.Log(elapsed)

	t.Log(details)
}

func TestNewCompoundClient(t *testing.T) {
	//c := NewCompoundClient("", "",600)

	//t.Log(c.SearchLimit)
}
