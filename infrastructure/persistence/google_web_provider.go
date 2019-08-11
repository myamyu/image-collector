package persistence

import (
	"fmt"
	"net/url"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

type GoogleWebProvider struct {
}

func NewGoogleWebProvider() *GoogleWebProvider {
	return &GoogleWebProvider{}
}

func (p *GoogleWebProvider) SearchImage(q string) ([]string, error) {
	reqURL := fmt.Sprintf("https://www.google.co.jp/search?q=%s&tbm=isch", url.QueryEscape(q))

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	selection := doc.Find(".rg_meta.notranslate")
	var dataArray []string

	selection.Each(func(index int, s *goquery.Selection) {
		dataArray = append(dataArray, s.Text())
	})

	return dataArray, nil
}
