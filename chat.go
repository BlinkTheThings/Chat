package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: chat [listen addr] [speak addr]\n")
		os.Exit(0)
	}

	fmt.Println("Chat")
	fmt.Println("----")

	listenAddr := os.Args[1]
	speakAddr := os.Args[2]

	go listen(listenAddr)

	speak(speakAddr)
}

func listen(listenAddr string) {
	localAddr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	udpConn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Fatal(err)
	}

	dgram := make([]byte, 65535)

	for {
		cnt, _, err := udpConn.ReadFromUDP(dgram)
		if err != nil {
			log.Print(err)
		}

		fmt.Printf("< %s", string(dgram[:cnt]))
	}
}

func speak(speakAddr string) {
	remoteAddr, err := net.ResolveUDPAddr("udp", speakAddr)
	if err != nil {
		log.Fatal(err)
	}

	udpConn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, _ := reader.ReadString('\n')
		fmt.Printf("> %s", line)

		_, err = udpConn.Write([]byte(line))
		if err != nil {
			log.Print(err)
		}
	}
}
