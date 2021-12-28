package pots

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const PROTOCOL = "tcp"

func TcpServer(port uint16) error {
	lstPort := ":" + strconv.Itoa(int(port))

	listen, err := net.Listen(PROTOCOL, lstPort)
	if err != nil {
		return fmt.Errorf("fail to listen %w", err)
	}
	defer listen.Close()
	fmt.Println("listening " + PROTOCOL + lstPort)

	for {
		conn, err := listen.Accept()
		if err != nil {
			return fmt.Errorf("fail to conn %w", err)
		}
		fmt.Println("Accept " + PROTOCOL + lstPort)

		go handleConn(conn)

	}

	return nil
}

func handleConn(conn net.Conn) error {
	clientID := conn.RemoteAddr().String()
	fmt.Println("handling..." + clientID)
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			return fmt.Errorf("fail to read %w", err)
		}

		msg := string(bytes[:])
		fmt.Printf("msg %s: %s", clientID, msg)
		if strings.TrimSpace(msg) == "exit" {
			fmt.Println("exit for: " + clientID)
			return nil
		}

		// response
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		conn.Write([]byte(myTime))
	}
}

func TcpClient(port int) error {
	time.Sleep(10 * time.Second)
	lstPort := ":" + strconv.Itoa(int(port))

	dial, err := net.Dial(PROTOCOL, lstPort)
	if err != nil {
		return fmt.Errorf("fail to dial %w", err)
	}
	defer dial.Close()
	fmt.Println("connected to " + PROTOCOL + lstPort)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(dial, text+"\n")

		message, _ := bufio.NewReader(dial).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "Bye" {
			fmt.Println("TCP client exiting...")
			return nil
		}
	}

	return nil
}
