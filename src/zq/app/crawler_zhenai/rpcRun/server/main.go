package main

import (
	"github.com/olivere/elastic"
	"log"
	"zq/app/crawler_zhenai/config"
	"zq/app/crawler_zhenai/rpcRun/persist"
	"zq/app/crawler_zhenai/rpcRun/rpcsupport"
)

func main() {
	//如果发生错误，Fatal()会强制退出。。
	log.Fatal(serveRpc(":1234", config.EsType))

}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(config.EsSetUrl),
	)

	if err != nil {
		return err
	}
	rpcsupport.ServeRpc(host, &persist.ItemSaveService{
		Client: client,
		Index:  index,
	})
	return nil
}
