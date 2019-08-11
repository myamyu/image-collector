package service

import (
	"github.com/myamyu/image-collector/domain/model"
	"github.com/myamyu/image-collector/domain/repository"
)

type SearchService struct {
	rep repository.SearchRepository
}

func NewSearchService(r repository.SearchRepository) *SearchService {
	return &SearchService{r}
}

func (s *SearchService) Search(query model.SearchQuery) ([]model.ImageInfo, error) {
	return s.rep.Search(query)
}
