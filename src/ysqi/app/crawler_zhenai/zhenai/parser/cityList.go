package parser

import (
	"regexp"
	"ysqi/app/crawler_zhenai/engine"
)

func ParseCityList(contents []byte) engine.ParseResult {
	compile, e := regexp.Compile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	if e != nil {
		panic(e)
	}

	findAll := compile.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	i := 0
	for _, value := range findAll {
		//fmt.Printf("City :%s , URL: %s \n", value[2], value[1])
		// 城市名字
		//result.Items = append(result.Items, string(value[2]))
		// 城市url
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(value[1]),
			ParserFunc: ParseCity,
		})
		i++
		if i == 1 {
			break
		}
	}

	return result

}
