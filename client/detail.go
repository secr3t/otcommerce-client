package client

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/secr3t/otcommerce-client/model"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strings"
)

const (
	getItemDescription = `GetItemDescription`

	InstanceKeyParam  = `instanceKey`
	ItemIdParam       = `itemId`
	DescEncodedString = `\"`
	DescDecodedString = `"`

	DescriptionPath = `OtapiItemDescription.ItemDescription`

	getItemFullInfoWithPromotions = `GetItemFullInfoWithPromotions`

	ItemParam      = `itemParameters`
	ItemParamValue = `<Parameters AllowIncomplete="false" AllowDeleted="false" WaitingTime="500"/>`

	FullInfoPath    = `OtapiItemFullInfo`
	AttributesPath  = `Attributes.#(IsConfigurator==true)#`
	CombinationPath = `ConfiguredItems.#(Quantity!=0)#`
	PromotionsPath  = `Promotions.#.ConfiguredItems`
	OptionName      = `OriginalPropertyName`
	OptionValue     = `OriginalValue`
	OptionImage     = `ImageUrl`
	Configurators   = `Configurators`
)

var (
	DescFail   = errors.New("get desc failed")
	DetailFail = errors.New("get detail failed")
)

type DetailClient struct {
	ApiKey string
}

func NewDetailClient(apiKey string) *DetailClient {
	return &DetailClient{
		ApiKey: apiKey,
	}
}

func (c *DetailClient) GetDescImgs(id string) ([]string, error) {
	q := url.Values{}

	q.Add(ItemIdParam, id)
	q.Add(InstanceKeyParam, c.ApiKey)

	uri := GetUri(host, getItemDescription, q.Encode())

	req, _ := http.NewRequest("GET", uri, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return descResultToImgs(body)
}

func (c *DetailClient) GetDetail(item model.Item) (*model.DetailItem, error) {
	id := item.Id
	descImgs, err := c.GetDescImgs(id)

	if err != nil {
		return nil, err
	}

	q := url.Values{}

	q.Add(ItemIdParam, id)
	q.Add(InstanceKeyParam, c.ApiKey)
	q.Add(ItemParam, ItemParamValue)

	uri := GetUri(host, getItemFullInfoWithPromotions, q.Encode())

	req, _ := http.NewRequest("GET", uri, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return resultToDetailItem(item, body, descImgs)
}

func resultToDetailItem(item model.Item, body []byte, descImgs []string) (*model.DetailItem, error) {
	r := gjson.ParseBytes(body)

	var detailItem model.DetailItem

	if r.Get("ErrorCode").String() != "Ok" {
		return &detailItem, DetailFail
	}

	detailItem.Title = item.Title
	detailItem.Images = item.Imgs
	detailItem.DescImgs = descImgs
	detailItem.NumIid = item.Id
	detailItem.DetailUrl = item.ProductUrl
	detailItem.Options = getOptions(r)

	if detailItem.Options == nil {
		return nil, DetailFail
	}

	return &detailItem, nil
}

func descResultToImgs(json []byte) ([]string, error) {
	//jsonStr := strings.Replace(string(json), DescEncodedString, DescDecodedString, -1)
	r := gjson.ParseBytes(json)

	if r.Get("ErrorCode").String() != "Ok" {
		return nil, DescFail
	}

	desc := strings.Replace(r.Get(DescriptionPath).String(), DescEncodedString, DescDecodedString, -1)
	desc = strings.Replace(desc, `src="//`, `src="http://`, -1)

	imgs := make([]string, 0)

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(desc))

	doc.Find("p > img").Each(func(_ int, s *goquery.Selection) {
			_, exists := s.Attr("usemap")
			src, _ := s.Attr("src")

			if !exists && !strings.HasSuffix(src, "gif") {
				imgs = append(imgs, src)
			}
		})
	doc.Find("div > img").Each(func(_ int, s *goquery.Selection) {
		_, exists := s.Attr("usemap")
		src, _ := s.Attr("src")

		if !exists && !strings.HasSuffix(src, "gif") {
			imgs = append(imgs, src)
		}
	})

	return imgs, nil
}

func getOptions(r gjson.Result) []model.Option {
	f := r.Get(FullInfoPath)

	optionMap := make(map[string]*model.Option)
	combinedMap := make(map[string][]string)
	priceMap := make(map[string]float64)

	f.Get(AttributesPath).ForEach(func(key, value gjson.Result) bool {
		option := &model.Option{
			Name:  value.Get(OptionName).String(),
			Value: value.Get(OptionValue).String(),
			Img:   value.Get(OptionImage).String(),
		}
		optionMap[getOptionPath(value)] = option
		return true
	})

	if len(optionMap) == 0 {
		return nil
	}

	f.Get(CombinationPath).ForEach(func(key, value gjson.Result) bool {
		combinationId := value.Get(IdPath).String()
		optionPaths := make([]string, 0)
		value.Get(Configurators).ForEach(func(key, configurator gjson.Result) bool {
			optionPaths = append(optionPaths, getOptionPath(configurator))
			return true
		})
		combinedMap[combinationId] = optionPaths
		priceMap[combinationId] = value.Get(PricePath).Float()

		return true
	})

	for _, v := range f.Get(PromotionsPath).Array() {
		v.ForEach(func(key, value gjson.Result) bool {
			combinationId := value.Get(IdPath).String()
			price := value.Get(PricePath).Float()
			if val, ok := priceMap[combinationId]; ok && val > price {
				priceMap[combinationId] = price
			}
			return true
		})
	}

	for combinationId, price := range priceMap {
		props := combinedMap[combinationId]
		for _, optionPath := range props {
			bPrice := optionMap[optionPath].Price
			if bPrice == 0 {
				optionMap[optionPath].Price = price
			} else {
				optionMap[optionPath].Price = math.Min(bPrice, price)
			}
		}
	}

	options := make([]model.Option, 0)

	for _, v := range optionMap {
		if v.Price != 0 {
			options = append(options, *v)
		}
	}

	return options
}

func getOptionPath(value gjson.Result) string {
	return value.Get("Pid").String() + ":" + value.Get("Vid").String()
}
