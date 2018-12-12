package rmq
 
import (
    "log"
    "github.com/streadway/amqp"
	"fmt"
	"go-ini/ini_utils"

)
 
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


var conn *amqp.Connection
var channel *amqp.Channel



//读取配置文件
func GetMQConfig() (string, string, string, string) {
	ini_parser := ini_utils.IniParser{}
	ini_path := "/etc/msg_service.ini"
	if err := ini_parser.Load(ini_path); err != nil {
        fmt.Printf("try load config file[%s] error[%s]\n", ini_path, err.Error())
        return "Read ini file Error", "", "", ""
	}
	user := ini_parser.GetString("rabbit_mq", "user")
	passwd := ini_parser.GetString("rabbit_mq", "passwd")
	ip := ini_parser.GetString("rabbit_mq", "ip")
	port := ini_parser.GetString("rabbit_mq", "port")
	return user, passwd, ip, port
}


func SetupRMQ() (err error) {
	if channel == nil {
	    mq_user, mq_passwd, mq_ip, mq_port := GetMQConfig()
		rmqAddr := "amqp://" + mq_user + ":" + mq_passwd + "@" + mq_ip + ":" + mq_port + "/"
		fmt.Printf(rmqAddr)
		conn, err := amqp.Dial(rmqAddr) // 建立连接
		failOnError(err, "Failed to connect to RabbitMQ")

		channel, err = conn.Channel()  // 创建channel
		failOnError(err, "Failed to open a channel")
	}
	return nil
}


func PushRMQ(msg string) (err error){
	q, err := channel.QueueDeclare(   // 创建消息队列，queue，并分配默认binding，empty exchange
		"sendmsg", // name 消息队列的名字
		false,   // durable   // 队列持久化
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
 
	body := msg
	err = channel.Publish(     // 发布消息，第一个参数表示路由名称（exchange），""则表示使用默认消息路由
		"",     // exchange
		q.Name, // routing key  "hello"
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message")
		return nil
	}
	log.Printf(" [x] Sent %s", body)
	return nil
}