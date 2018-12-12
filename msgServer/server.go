//server/server.go
package main

import (
    pb "grpcMsg/proto"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "log"
    "net"
    "fmt"
    "grpcMsg/MqHandler"
    "go-ini/ini_utils"
    "strings"
    "google.golang.org/grpc/peer"
)

// 获取client ip
func getClietIP(ctx context.Context) (string, error) {
    pr, ok := peer.FromContext(ctx)
    if !ok {
        return "", fmt.Errorf("[getClinetIP] invoke FromContext() failed")
    }
    if pr.Addr == net.Addr(nil) {
        return "", fmt.Errorf("[getClientIP] peer.Addr is nil")
    }
    addSlice := strings.Split(pr.Addr.String(), ":")
    return addSlice[0], nil
}

//读取配置文件
func ms_server_config() (string, string) {
	ini_parser := ini_utils.IniParser{}
	ini_path := "/etc/msg_service.ini"
	if err := ini_parser.Load(ini_path); err != nil {
        fmt.Printf("try load config file[%s] error[%s]\n", ini_path, err.Error())
        return "Read ini file Error", ""
	}
	ip := "0.0.0.0"//ini_parser.GetString("server", "ip")
	port := ini_parser.GetString("server", "port")
	return ip, port
}

//定义SendMsgService并实现约定的接口
type SendMsgService struct{}


func (h SendMsgService) SendMsg(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {
    resp := new(pb.MsgReply)
    client_ip, err := getClietIP(ctx)
    if err != nil {
        log.Fatalf("get client ip error :%v", err)
    }
    msg := fmt.Sprintf(`{"ip": "%s", "msg": "%s"}`, client_ip, in.Name)
    fmt.Println(msg)
    err = rmq.PushRMQ(msg)
    if err != nil {
        log.Fatalf("send msg rabbit mq error :%v", err)
    }
    resp.Message = "Send Message: " + in.Name + "."
    return resp, nil
}



var _SendMsgService = SendMsgService{}


func main() {
    ip, port := ms_server_config()
    Address := ip + ":" + port
    
    err := rmq.SetupRMQ()
    if err != nil {
        log.Fatalf("Connect rabbit mq error :%v", err)
    }

    listen, err := net.Listen("tcp", Address)
    if err != nil {
        log.Fatalf("failed to listen:%v", err)
    }

    s := grpc.NewServer()                   //实例化grpc Server
    pb.RegisterMsgServer(s, _SendMsgService) //注册_SendMsgService
    
    log.Println("Listen on " + Address)
    s.Serve(listen)
}
