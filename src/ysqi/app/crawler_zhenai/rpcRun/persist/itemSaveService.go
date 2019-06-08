package persist

import (
	"github.com/olivere/elastic"
	"ysqi/app/crawler_zhenai/engine"
	"ysqi/app/crawler_zhenai/persist"
)

type ItemSaveService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaveService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	}
	return err
}
