package banner

import (
	"Zscan/core/httpx"
	_ "embed"
	"encoding/json"
	"github.com/Knetic/govaluate"
)

type Banner struct {
	RuleId         string `json:"rule_id"`
	Level          string `json:"level"`
	SoftHard       string `json:"softhard"`
	Product        string `json:"product"`
	Company        string `json:"company"`
	Category       string `json:"category"`
	ParentCategory string `json:"parent_category"`
	Condition      string `json:"Condition"`
}
type BannerPrints []Banner

//go:embed data/banner.json
var BannerJson []byte

func InitBanner() (BannerPrints, error) {
	// 序列化指纹json数据
	var Banners BannerPrints
	err := json.Unmarshal(BannerJson, &Banners)
	if err != nil {
		return nil, err
	}
	return Banners, err
}

func (f *Banner) Matcher(response *httpx.Response) (bool, error) {
	expString := f.Condition
	expression, err := govaluate.NewEvaluableExpressionWithFunctions(expString, HelperFunctions(response))
	if err != nil {
		return false, err
	}
	paramters := make(map[string]interface{})
	paramters["title"] = response.Title
	paramters["server"] = response.GetHeader("server")
	paramters["protocol"] = "http"

	result, err := expression.Evaluate(paramters)
	if err != nil {
		return false, err
	}
	t := result.(bool)
	return t, err
}
func (f *BannerPrints) Matcher(response *httpx.Response) ([]string, error) {
	ret := make([]string, 0)
	for _, item := range *f {
		v, err := item.Matcher(response)
		if err != nil {
			return nil, err
		}
		if v {
			n := item.Product
			ret = append(ret, n)
		}
	}
	return ret, nil
}
