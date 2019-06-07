package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestParseProfile(t *testing.T) {
	//bytes, e := fetcher.Fetcher("http://album.zhenai.com/u/1928969584")
	//if e != nil {
	//	panic(e)
	//}
	//ParseProfile(bytes)
	url := "http://www.zhenai.com/zhenghun/aba"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	all, err := ioutil.ReadAll(resp.Body)
	ParseCity(all)

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s\n", all)

}
