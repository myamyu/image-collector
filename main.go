// +build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/myamyu/image-collector/domain/model"
	"github.com/myamyu/image-collector/domain/service"
	"github.com/myamyu/image-collector/infrastructure/repository"
)

func main() {
	r := repository.NewGoogleWebSearchRepository()
	s := service.NewSearchService(r)

	stdin := bufio.NewScanner(os.Stdin)

	fmt.Print("サイト：")
	stdin.Scan()
	site := stdin.Text()

	fmt.Print("ワード：")
	stdin.Scan()
	word := stdin.Text()

	fmt.Print("件数：")
	stdin.Scan()
	limit, err := strconv.Atoi(stdin.Text())
	if err != nil {
		limit = 0
	}

	res, err := s.Search(model.SearchQuery{SiteDomain: site, SearchWord: word, Limit: limit})
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Printf("[%d]件 %+v\n", len(res), res)
}
