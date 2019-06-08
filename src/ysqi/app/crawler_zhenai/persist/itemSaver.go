package persist

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"log"
	"ysqi/app/crawler_zhenai/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	//关闭内网的sniff
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://47.244.160.71:9200"),
	)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("Item Saver: Got %d  item : %v", itemCount, item)
			err := save(client, index, item) //保存item
			if err != nil {
				panic(err)
			}

		}
	}()
	return out, nil

}

func save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type ..")
	}

	indexService := client.Index(). //存储数据，可以添加或者修改，要看id是否存在
					Index(index).
					Type(item.Type).
					BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
