# 配置文件
>client端需要配置/tzagent/tzmsg/conf/server.ini  
```
[server]  
ip = 50052  
port = 10.0.0.1
```

>服务端需要配置mq: /tzagent/tzmsg/conf/mq.ini    #消息推送至队列
```
[rabbit_mq]  
user = username
passwd = pass
ip = 127.0.0.1  #mq队列的服务地址
port = 5672  
```