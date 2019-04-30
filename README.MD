# 配置文件
>client端需要配置/etc/msg_service.ini
```
<<<<<<< HEAD
[server]
ip = 10.0.0.1
port = 50052
=======
[server]  
ip = 10.0.0.1 
port = 50052 
>>>>>>> 15a92fe05d52b97e466a1446518abccd94502987
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


# 编译：
mac系统需要交叉编译才能在linux上运行：
```GOOS=linux GOARCH=amd64 go build server.go```

# 部署：
部署到主机
