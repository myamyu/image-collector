// +build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/myamyu/image-collector/domain/model"
	"github.com/myamyu/image-collector/domain/service"
	"github.com/myamyu/image-collector/infrastructure/repository"
)

func download(url string, distFile string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(distFile)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, resp.Body)
	return nil
}

func main() {
	r := repository.NewGoogleWebSearchRepository()
	s := service.NewSearchService(r)

	query := model.SearchQuery{
		SiteDomain: "twitter.com",
		SearchWord: "スタバわず",
		Limit:      50,
	}

	log.Printf("画像を探します query[ %+v ]\n", query)

	res, err := s.Search(query)
	if err != nil {
		log.Fatalf("エラー！！！\n%+v", err)
		return
	}

	log.Printf("%d 件見つかりました。\n", len(res))

	absPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatalf("エラー！！！\n%+v", err)
		return
	}
	re := regexp.MustCompile("[/\\\\<>\\?%#]")
	distDir := filepath.Join(
		absPath,
		"download",
		query.SiteDomain,
		re.ReplaceAllString(query.SearchWord, "_"),
	)
	if err := os.MkdirAll(distDir, 0777); err != nil {
		log.Fatalf("エラー！！！\n%+v", err)
		return
	}

	for i, img := range res {
		log.Printf("[%d] %s ...", i+1, img.Text)
		distPath := filepath.Join(distDir, fmt.Sprintf("%d", i+1)+".jpg")
		if err := download(img.ImageThumbURL, distPath); err != nil {
			log.Fatalf("エラー！！！\n%+v", err)
			return
		}
		log.Println("done.")
	}
}
