// +build ignore

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/myamyu/image-collector/domain/model"
	"github.com/myamyu/image-collector/domain/service"
	"github.com/myamyu/image-collector/infrastructure/repository"
)

func main() {
	r := repository.NewGoogleWebSearchRepository()
	service := service.NewSearchService(r)

	http.HandleFunc("/image-collector", func(res http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		site := query.Get("s")
		if site == "" {
			site = "twitter.com"
		}
		word := query.Get("q")
		if word == "" {
			word = "スタバなう"
		}
		limit, err := strconv.Atoi(query.Get("l"))
		if err != nil {
			limit = 0
		}

		images, err := service.Search(model.SearchQuery{SiteDomain: site, SearchWord: word, Limit: limit})
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Printf("%+v", err)
			return
		}

		json, err := json.Marshal(images)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Printf("%+v", err)
			return
		}

		header := res.Header()
		header.Add("Content-Type", "application/json; charset=utf-8")
		res.WriteHeader(http.StatusOK)
		res.Write(json)
	})
	http.Handle("/", http.FileServer(http.Dir("web/static")))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
