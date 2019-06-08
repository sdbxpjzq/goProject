珍爱网 爬虫项目



# 前端- frontend目录
启动 start.go
访问
`localhost:9527/`

# 并发版本 - main.go
运行`main.go`

# 分布式rpc - rpcRun 目录
先运行`rpcRun/server/main.go`
在运行`rpcRun/main.go`

# 使用到的中间件
ElasticSearch

```shell

docker pull docker.elastic.co/elasticsearch/elasticsearch:6.3.2
docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.3.2

```
ElasticSearch的默认端口是9200，我们把宿主环境9200端口映射到Docker容器中的9200端口，就可以访问到Docker容器中的ElasticSearch服务了，同时我们把这个容器命名为es。

https://gopkg.in/olivere/elastic.v6

参考:
参考:
https://chaindesk.cn/witbook/22/495

