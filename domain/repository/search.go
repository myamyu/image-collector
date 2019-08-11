package repository

import (
	"github.com/myamyu/image-collector/domain/model"
)

type SearchRepository interface {
	Search(query model.SearchQuery) ([]model.ImageInfo, error)
}
