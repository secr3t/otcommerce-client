package client

import "testing"

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