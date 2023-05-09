package httpserver

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type User struct {
	Age string //年龄
	Sex string //性别
}

type Client struct {
	Conn   net.Conn // 连接信息
	Name   string   // 别名
	IsQuit bool     // 是否退出
	User
}

type Message struct {
	Name string // 用户名
	Op   int    // 操作服务
	Msg  string // 信息内容
}

const (
	Read = iota + 1
	Quit
	NtyLogin
	UpdUser
)

var ConnMap = make(map[string]Client)

func StartListen() {
	// 使用 net 包的 Listen 函数监听 127.0.0.1:8000 上的 tcp 连接
	listen, err := net.Listen("tcp", "127.0.0.1:9016")
	if err != nil {
		fmt.Printf("server listen error:%v", err)
		return
	}
	fmt.Println("服务启动成功。监听端口9016")
	// 使用 defer 在运行结束后优雅的关闭
	defer listen.Close()
	for {
		// 当接收到连接请求时
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn fail ...")
			continue
		}
		// conn.RemoteAddr() 连接的客户端地址
		fmt.Println(conn.RemoteAddr(), "connect successed")

		// handle 为每一个客户端开单独的协程进行业务操作
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	for {
		// 通过 Read 获取数据到 data中
		// ml 即为数据长度
		data := make([]byte, 255)
		ml, err := conn.Read(data)
		if ml == 0 || err != nil {
			// 收到的参数错误忽略、
			continue
		}

		// 解析协议
		//  Name | Op | Msg | ...Other Operation
		fmt.Println(string(data[0:ml]), "listen data")
		msgStr := strings.Split(string(data[0:ml]), "|")
		fmt.Println(msgStr)

		var cMsg Message
		cMsg.Name = msgStr[0]
		cMsg.Op, _ = strconv.Atoi(msgStr[1])
		cMsg.Msg = msgStr[2]

		name := msgStr[0]
		// 每个人的连接信息
		ConnMap[name] = Client{
			Conn: conn,
			Name: name,
		}

		// switch cMsg.Op {
		// case Read:
		// 	cMsg.Read()
		// case Quit:
		// 	cMsg.Quit()
		// case NtyLogin:
		// 	cMsg.ntyLogin()
		// case UpdUser:
		// 	cMsg.UpdUser()

		// default:
		// 	fmt.Println("无效OP")
		// }

	}
}
