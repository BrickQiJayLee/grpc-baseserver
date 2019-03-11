package main

import (
    pb "grpcMsg/proto"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
	"log"
	"os"
	"go-ini/ini_utils"
	"fmt"
)

type MsgStruct struct {

}

//读取配置文件
func ms_server_config() (string, string) {
	ini_parser := ini_utils.IniParser{}
	ini_path := "/etc/msg_service.ini"
	if err := ini_parser.Load(ini_path); err != nil {
        fmt.Printf("try load config file[%s] error[%s]\n", ini_path, err.Error())
        return "Read ini file Error", ""
	}
	ip := ini_parser.GetString("server", "ip")
	port := ini_parser.GetString("server", "port")
	return ip, port
}

func main() {
    ip, port := ms_server_config()
    //client_ip := ms_client_conifg()
    Address := ip + ":" + port
    fmt.Println(Address)
    conn, err := grpc.Dial(Address, grpc.WithInsecure())
    if err != nil {
        log.Fatalln(err)
	os.Exit(-1)
    }
	defer conn.Close()
	
	msg := os.Args[1]  //命令行输入需要发送的消息
	if err != nil {
		log.Fatalln(err)
	}
	
    c := pb.NewMsgClient(conn)
    reqBody := new(pb.MsgRequest)
    reqBody.Name = msg
    r, err := c.SendMsg(context.Background(), reqBody)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println(r.Message)
}
