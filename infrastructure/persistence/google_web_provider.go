package persistence

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ImageSearchResponse Google Image Searchの検索結果JSON
type ImageSearchResponse struct {
	ImageURL         string `json:"ou"`
	ImageType        string `json:"ity"`
	ImageWidth       int    `json:"ow"`
	ImageHeight      int    `json:"oh"`
	ImageThumbURL    string `json:"tu"`
	ImageThumbWidth  int    `json:"tw"`
	ImageThumbHeight int    `json:"th"`
	PageTitle        string `json:"pt"`
	PageURL          string `json:"ru"`
	SiteTitle        string `json:"st"`
}

type GoogleWebProvider struct {
}

func NewGoogleWebProvider() *GoogleWebProvider {
	return &GoogleWebProvider{}
}

func (p *GoogleWebProvider) SearchImage(q string) ([]ImageSearchResponse, error) {
	tbs := fmt.Sprintf("tbs=isz:%s,islt:%s,itp:%s,qdr:%s",
		"lt",    // 最小サイズ
		"2mp",   // 200万画素
		"photo", // 写真
		"m",     // 1か月以内
	)

	reqURL := fmt.Sprintf("https://www.google.co.jp/search?q=%s&tbm=isch&%s", url.QueryEscape(q), tbs)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// ".rg_meta.notranslate" にはJSONが詰まっている
	selection := doc.Find(".rg_meta.notranslate")
	var (
		dataArray    []ImageSearchResponse
		errTextArray []string
	)
	selection.Each(func(index int, s *goquery.Selection) {
		innerText := s.Text()

		img := new(ImageSearchResponse)
		if err := json.Unmarshal(([]byte)(innerText), img); err != nil {
			errTextArray = append(errTextArray, innerText)
			return
		}

		dataArray = append(dataArray, *img)
	})

	var parseError error
	if len(errTextArray) == 0 {
		parseError = nil
	} else {
		parseError = fmt.Errorf("JSON parse error[%s]", strings.Join(errTextArray, ", "))
	}

	return dataArray, parseError
}
