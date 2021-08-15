package client

import (
	"encoding/xml"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type SearchParam struct {
	Q             string
	Page          int
	Sort          string
	PageSize      int
	XmlParameters SearchItemsParameters
}

type SearchItemsParameters struct {
	Provider            string  `xml:"Provider"`
	SearchMethod        string  `xml:"SearchMethod"`
	CategoryId          *string `xml:"CategoryId,omitempty"`
	VendorName          *string `xml:"VendorName,omitempty"`
	VendorId            *string `xml:"VendorId,omitempty"`
	VendorAreaId        *string `xml:"VendorAreaId,omitempty"`
	ItemTitle           *string `xml:"ItemTitle,omitempty"`                  // q
	MinPrice            *int64  `xml:"MinPrice,omitempty"`                   // start price
	MaxPrice            *int64  `xml:"MaxPrice,omitempty"`                   // end price
	UseOptimalFrameSize bool    `xml:"UseOptimalFrameSize"`                  // use optimal frame size instead of given
	Configurators       *Ppath  `xml:"Configurators>Configurator,omitempty"` // ppath for search by brand
}

func NewParams() *SearchItemsParameters {
	return &SearchItemsParameters{
		Provider:            "taobao",
		SearchMethod:        "Official",
		UseOptimalFrameSize: true,
	}
}

func (p *SearchItemsParameters) CatId(categoryId string) *SearchItemsParameters {
	p.CategoryId = &categoryId
	return p
}

func (p *SearchItemsParameters) VName(vendorName string) *SearchItemsParameters {
	p.VendorName = &vendorName
	return p
}

func (p *SearchItemsParameters) VId(vendorId string) *SearchItemsParameters {
	p.VendorId = &vendorId
	return p
}

func (p *SearchItemsParameters) VAreaId(vendorAreaId string) *SearchItemsParameters {
	p.VendorAreaId = &vendorAreaId
	return p
}

func (p *SearchItemsParameters) StartPrice(startPrice int64) *SearchItemsParameters {
	p.MinPrice = &startPrice
	return p
}

func (p *SearchItemsParameters) EndPrice(endPrice int64) *SearchItemsParameters {
	p.MaxPrice = &endPrice
	return p
}

func (p *SearchItemsParameters) Q(query string) *SearchItemsParameters {
	p.ItemTitle = &query
	return p
}

func (p *SearchItemsParameters) Ppath(ppath string) *SearchItemsParameters {
	p.Configurators = NewPpath(ppath)
	return p
}

func (p *SearchItemsParameters) ToXml() string {
	bytes, err := xml.Marshal(p)
	if err != nil {
		log.Print(err)
		return ""
	}

	return string(bytes)
}

type Configurator struct {
	Ppath Ppath
}

type Ppath struct {
	Pid string `xml:"Pid,attr"`
	Vid string `xml:"Vid,attr"`
}

func NewPpath(ppath string) *Ppath {
	paths := strings.Split(ppath, ":")
	return &Ppath{
		Pid: paths[0],
		Vid: paths[1],
	}
}

func NewSearchParam(page, pageSize int) SearchParam {
	return SearchParam{
		Page:       page,
		PageSize:   pageSize,
	}
}

func (p SearchParam) ToQuery() string {
	query := url.Values{}

	query.Add("frameSize", strconv.Itoa(p.PageSize))
	query.Add("framePosition", strconv.Itoa(p.Page))
	query.Add("xmlParameters", p.XmlParameters.ToXml())

	return query.Encode()
}
