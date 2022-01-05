package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	ListenAndServe(":8000")
}

func ListenAndServe(address string) {
	// 绑定监听端口
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("Listener err: %v", err))
	}
	defer listener.Close()
	log.Println(fmt.Sprintf("bind:%s, start listening", address))
	// 通过Accept()阻塞等待链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(fmt.Sprintf("accept err: %v", err))
		}
		// 开启gorountine处理该链接
		go Handle(conn)
	}

}

func Handle(conn net.Conn) {
	// 使用bufio提供缓冲区功能
	reader := bufio.NewReader(conn)
	for {
		// ReadString 会一直阻塞直到遇到分隔符 '\n'
		msg, err := reader.ReadString('\n')
		if err != nil {
			// EOF 说明链接被中断
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
			return
		}
		b := []byte(msg)
		// 暂时将收到的消息发送给客户端
		conn.Write(b)
	}

}
