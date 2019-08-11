package main

import (
	"fmt"

	"github.com/myamyu/image-collector/domain/model"
	"github.com/myamyu/image-collector/domain/service"
	"github.com/myamyu/image-collector/infrastructure/repository"
)

func main() {
	r := repository.NewGoogleWebSearchRepository()
	s := service.NewSearchService(r)

	res, err := s.Search(model.SearchQuery{SiteDomain: "twitter.com", SearchWord: "バール"})
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	println(res)
	fmt.Printf("%+v\n", res)
}
