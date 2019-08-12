package repository

import (
	"fmt"

	"github.com/myamyu/image-collector/domain/model"
	"github.com/myamyu/image-collector/infrastructure/persistence"
)

// GoogleWebSearchRepository Google Web検索から情報を取得するリポジトリ
type GoogleWebSearchRepository struct {
	p persistence.GoogleWebProvider
}

// NewGoogleWebSearchRepository GoogleWebSearchRepositoryを作成する
func NewGoogleWebSearchRepository() *GoogleWebSearchRepository {
	p := persistence.NewGoogleWebProvider()
	return &GoogleWebSearchRepository{*p}
}

// Search 検索を実行する
func (rep *GoogleWebSearchRepository) Search(query model.SearchQuery) ([]model.ImageInfo, error) {
	q := fmt.Sprintf("site:%s %s", query.SiteDomain, query.SearchWord)
	res, err := rep.p.SearchImage(q)
	if err != nil {
		return nil, err
	}

	var images []model.ImageInfo

	for _, img := range res {
		images = append(images, model.ImageInfo{
			ImageURL:      img.ImageURL,
			ImageType:     img.ImageType,
			Text:          img.PageTitle,
			WebPageURL:    img.PageURL,
			ImageThumbURL: img.ImageThumbURL,
		})

		if query.Limit > 0 && len(images) >= query.Limit {
			break
		}
	}

	return images, nil
}
