package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func Server() {

	listener, err := net.Listen("tcp", "127.0.0.1:8081")

	if err != nil {

		log.Println(err)

	}

	conn, err := listener.Accept()

	if err != nil {

		log.Println(err)

	}

	defer conn.Close()

	conn.Write([]byte("message"))

	time.Sleep(10)

	conn.Write([]byte("MesSaGe"))

	time.Sleep(10)

	conn.Write([]byte("MESSAGE"))
}

func main() {
	go Server()

	conn, err := net.Dial("tcp", "127.0.0.1:8081")

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	time.Sleep(time.Second * 3)
	message := make([]byte, 1024) // создадим буфер
	n, err := conn.Read(message)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(strings.ToUpper(string(message[:n]))) // напечатаем сообщение
	//fmt.Println(string(message[:]))  // напечатаем сообщение

}
