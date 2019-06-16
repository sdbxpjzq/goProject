package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//"https://www.zhenai.com/zhenghun"

var rateLimiter = time.Tick(time.Millisecond * 100)

func Fetcher(url string) ([]byte, error) {
	<-rateLimiter
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", " www.zhenai.com")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Cookie", "sid=d7327b05-fec8-4f6f-92cc-05490cb80fbe; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1559172778; __channelId=905821%2C0; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1559779784")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	if resp == nil {
		fmt.Println("resp:", resp)
		return nil, err
	}

	allBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return allBytes, err
}
