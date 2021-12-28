package pots

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const PROTOCOL_U = "udp"

func UdpServer(port uint16) error {
	lstPort := ":" + strconv.Itoa(int(port))

	udpAddr, err := net.ResolveUDPAddr(PROTOCOL_U, lstPort)
	if err != nil {
		return fmt.Errorf("fail to set udp %w", err)
	}

	listen, err := net.ListenUDP(PROTOCOL_U, udpAddr)
	if err != nil {
		return fmt.Errorf("fail to listen %w", err)
	}

	defer listen.Close()
	fmt.Println("listening " + PROTOCOL_U + lstPort)

	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := listen.ReadFromUDP(buffer)

		msg := string(buffer[0 : n-1])
		clientID := addr.String()
		fmt.Printf("msg %s: %s \n", clientID, msg)

		if strings.TrimSpace(string(buffer[0:n])) == "exit" {
			fmt.Println("Exiting UDP server!")
			return nil
		}

		// response
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		_, err = listen.WriteToUDP([]byte(myTime), addr)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func UdpClient(port uint16) error {
	lstPort := ":" + strconv.Itoa(int(port))

	addrUdp, err := net.ResolveUDPAddr(PROTOCOL_U, lstPort)
	if err != nil {
		return fmt.Errorf("fail to set udp %w", err)
	}

	dialUDP, err := net.DialUDP(PROTOCOL_U, nil, addrUdp)
	if err != nil {
		return fmt.Errorf("fail to connect %w", err)
	}
	defer dialUDP.Close()
	fmt.Println("connected to " + PROTOCOL_U + lstPort)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err = dialUDP.Write(data)
		if strings.TrimSpace(string(data)) == "exit" {
			fmt.Println("Exiting UDP client!")
			return nil
		}

		if err != nil {
			fmt.Println(err)
			return nil
		}

		buffer := make([]byte, 1024)
		n, _, err := dialUDP.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
	return nil
}
