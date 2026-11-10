package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		log.Fatal("Couldn't resolve address:", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("Listen failed:", err)

	}

	defer conn.Close()
	fmt.Println("UDP server listening on", conn.LocalAddr())
	buffer := make([]byte, 1024)
	for {

		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("read error: %v", err)
			continue
		}

		fmt.Printf("Got message from %s: %s\n", clientAddr, string(buffer[:n]))

		_, err = conn.WriteToUDP(buffer[:n], clientAddr)
		if err != nil {
			log.Printf("write error: %v", err)
		}
	}
}
