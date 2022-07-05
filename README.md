## 接入文档

#### 新增消息主题(Topic)

```shell
curl --location --request POST 'http://localhost:8900/alarm-api/alarm/topic/addTopic' \
--header 'Content-Type: application/json' \
--data-raw '{"remark":"xxx"}'
```



#### 新增消息监听(Listener)

##### 钉钉
```shell
curl --location --request POST 'http://localhost:8900/alarm-api/alarm/listener/addListener' \
--header 'Content-Type: application/json' \
--data-raw '{"accessToken":"xxx","sn":"x1","noticeType":"DT"}'
```
##### TG
```shell
curl --location --request POST 'http://localhost:8900/alarm-api/alarm/listener/addListener' \
--header 'Content-Type: application/json' \
--data-raw '{"accessToken":"xx","sn":"x12","chatId":"-1254","noticeType":"TG"}'
```


#### 配置主题和监听者(Topic>Listener)

##### 钉钉 或者 TG 不区分会话小组
```shell
curl --location --request POST 'http://localhost:8900/alarm-api/alarm/listener/addTopicListener' \
--header 'Content-Type: application/json' \
--data-raw '{"topic":"ebcb0302dcbc0a23218ba69d554087ac","accessToken":"xxx"}'
```
##### TG 单个会话小组监听
```shell
curl --location --request POST 'http://localhost:8900/alarm-api/alarm/listener/addTopicListener' \
--header 'Content-Type: application/json' \
--data-raw '{"topic":"ebcb0302dcbc0a23218ba69d554087ac","accessToken":"xx","chatId":"-1254"}'
```


#### 配置业务SQL监控定时任务

```shell
curl --location --request POST 'http://localhost:8900/alarm-api/alarm/monitor/add' \
--header 'Content-Type: application/json' \
--data-raw '{
	"cron": "0 0/1 * * * ?",
	"sql": "SELECT * FROM al_test.test1;",
	"status": 1,
	"dbId": 1,
	"note": "test：",
	"sn": "a967bb0b4324fee168fa803e3a4d6dcc",
	"showSql": "1"
}'
```



#### 开启定时任务

```shell
curl --location --request POST 'http://localhost:8900/alarm-api/alarm/monitor/start/task' \
--header 'Content-Type: application/json' \
--data-raw '{"id":5}'
```



#### 关闭定时任务

```shell
curl --location --request POST 'http://localhost:8900/alarm-api/alarm/monitor/stop/task' \
--header 'Content-Type: application/json' \
--data-raw '{"id":5}'
```



#### 获取定时任务列表

```shell
curl --location --request POST 'http://localhost:8900/alarm-api/monitor/alarm/list/task' \
--header 'Content-Type: application/json' \
--data-raw '{}'
```



#### 发送普通消息

```shell
curl --location --request POST 'http://localhost:8900/alarm-api/alarm/msg/save' \
--header 'Content-Type: application/json' \
--data-raw '{"msg":"test","sn":"a967bb0b4324fee168fa803e3a4d6dcc","requestNo":"x1"}'
```

