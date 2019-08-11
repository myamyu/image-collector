package repository

import (
	"encoding/json"
	"fmt"

	"github.com/myamyu/image-collector/domain/model"
	"github.com/myamyu/image-collector/infrastructure/persistence"
)

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

type GoogleWebSearchRepository struct {
	p persistence.GoogleWebProvider
}

func NewGoogleWebSearchRepository() *GoogleWebSearchRepository {
	p := persistence.NewGoogleWebProvider()
	return &GoogleWebSearchRepository{*p}
}

func (rep *GoogleWebSearchRepository) Search(query model.SearchQuery) ([]model.ImageInfo, error) {
	q := fmt.Sprintf("site:%s %s", query.SiteDomain, query.SearchWord)
	res, err := rep.p.SearchImage(q)
	if err != nil {
		return nil, err
	}

	var images []model.ImageInfo

	for _, s := range res {
		img := new(ImageSearchResponse)
		if err := json.Unmarshal(([]byte)(s), img); err != nil {
			fmt.Printf("%s\n", s)
			return nil, err
		}
		images = append(images, model.ImageInfo{
			ImageURL:   img.ImageURL,
			Text:       img.PageTitle,
			WebPageURL: img.PageURL,
		})
	}

	return images, nil
}
