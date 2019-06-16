package engine

import (
	"log"
	"zq/app/crawler_zhenai/fetcher"
)

func worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)

	body, err := fetcher.Fetcher(r.Url)

	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
