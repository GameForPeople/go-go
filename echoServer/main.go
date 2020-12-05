package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

const (
	PORT_NUMBER int = 7777
	LOCAL_HOST string = "127.0.0.1:"
)

func main() {
	portString := strconv.Itoa(PORT_NUMBER)

	// 서버를 Listen 상태로 변경합니다.
	listener, error := net.Listen("tcp", LOCAL_HOST+portString)
	if listener == nil {
		panic("Failed to listen:" + error.Error())
	}

	fmt.Println("Listening Ip:", LOCAL_HOST)
	fmt.Println("Listening Port:", portString)

	clients := AcceptClient(listener)
	for {
		go HandleClient(<-clients)
	}
}

// Accept합니다.
func AcceptClient(listner net.Listener) chan net.Conn {
	ch := make(chan net.Conn)

	go func() {
		for {
			client, error := listner.Accept()
			if client == nil {
				fmt.Println("failed to accept:", error.Error())
			}

			fmt.Printf("new connection: %v <-> %v\n", client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()

	return ch
}

// 
func HandleClient(client net.Conn) {
	defer fmt.Println("close client")
	defer client.Close()

	fmt.Println("Connected from:", client.RemoteAddr().String())

	reader := bufio.NewReader(client)
	for {
		line, error := reader.ReadBytes('\n')
		if error != nil {
			fmt.Println("failed to read from client.", error.Error())
			break
		}

		fmt.Println(string(line))

		_, error = client.Write(line)
		if error != nil {
			fmt.Println("write error:", error)
			return
		}
	}
}
