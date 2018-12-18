# 配置文件
>client端需要配置/etc/msg_service.ini
```
[server]
ip = 10.0.0.1
port = 50052
```

>服务端需要配置mq: /etc/msg_service.ini    #消息推送至队列
```
[server]
ip = 10.0.0.1
port = 50052

[rabbit_mq]
user = username
passwd = pass
ip = ip
port = 5672
```
