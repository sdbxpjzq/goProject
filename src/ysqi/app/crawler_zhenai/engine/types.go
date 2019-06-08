package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string //URL
	Type    string //存储到ElasticSearch时的type
	Id      string //用户Id
	Payload interface{}
}
