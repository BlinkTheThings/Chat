package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Chat")

	localAddr, err := net.ResolveUDPAddr("udp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	udpConn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Fatal(err)
	}

	dgram := make([]byte, 65535)

	for {
		cnt, remoteAddr, err := udpConn.ReadFromUDP(dgram)
		if err != nil {
			log.Print(err)
		}

		fmt.Printf("Read %d bytes from %s\n", cnt, remoteAddr.String())
	}

}
