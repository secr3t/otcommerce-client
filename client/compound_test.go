package client

import (
	"github.com/secr3t/otcommerce-client/model"
	"testing"
	"time"
)

func TestCompoundClient_SearchAndGetDetails(t *testing.T) {
	c := NewCompoundClient("","", 100)

	p := NewSearchParam(0, 100)

	start := time.Now()
	detailChan := c.SearchAndGetDetailsMultiRequestOneTime(&p)

	elapsed := time.Since(start)
	t.Log(elapsed)

	var details []model.Detail

	for detail := range detailChan {
		details = append(details, detail)
	}

	t.Log(len(details))
}

func TestNewCompoundClient(t *testing.T) {
	c := NewCompoundClient("", "",600)

	t.Log(c.SearchLimit)
}
