package parser

import (
	"regexp"
	"ysqi/app/crawler_zhenai/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const cityUrlRe = `<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`

//解析信息
func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, c := range all {
		//result.Items = append(result.Items, "User:"+string(c[2])) //用户名字
		name := string(c[2])
		url := string(c[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, url, name)
			},
		})
	}
	// 下一页
	urlRe := regexp.MustCompile(cityUrlRe)
	all2 := urlRe.FindAllSubmatch(contents, -1)
	for _, c := range all2 {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(c[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
