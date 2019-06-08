珍爱网 爬虫项目

![](http://ww4.sinaimg.cn/large/006tNc79ly1g3sl5u8ipmj32180roagf.jpg)


```shell

docker pull docker.elastic.co/elasticsearch/elasticsearch:6.3.2
docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.3.2


```
ElasticSearch的默认端口是9200，我们把宿主环境9200端口映射到Docker容器中的9200端口，就可以访问到Docker容器中的ElasticSearch服务了，同时我们把这个容器命名为es。


https://gopkg.in/olivere/elastic.v6