package internal

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
)

const (
	TCP = "tcp"
	UDP = "udp"
)

func Scan(ports ...uint16) []uint16 {

	descriptors := maxDescriptors() / 2
	addr := net.TCPAddr{
		IP:   nil,
		Port: 0,
		Zone: "",
	}

	maxChan := make(chan uint16, descriptors) // required chan to not have to many descriptors open
	open := make([]uint16, 0)
	for _, v := range ports {
		go func(port uint16) {
			fmt.Println("in routine: ", port)
			addr.Port = int(port)
			err := IsTcpListening(addr)
			if err != nil && !strings.Contains(err.Error(), "connection refused") {
				fmt.Println(err)
			}

			if err != nil {
				maxChan <- uint16(0)
				return
			}

			maxChan <- port
		}(v)

		if result := <-maxChan; result > 0 {
			open = append(open, result)
		}
		fmt.Println("end routine: ")

	}
	return open
}

func IsTcpListening(addr net.TCPAddr) error {

	network := addr.String()
	tcpDstAddr, err := net.ResolveTCPAddr(TCP, network)
	if err != nil {
		return err
	}
	conn, err := net.DialTCP(TCP, nil, tcpDstAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func IsUdpListening(addr net.UDPAddr) error {

	dstAddr, err := net.ResolveUDPAddr(UDP, addr.String())
	if err != nil {
		return err
	}
	conn, err := net.DialUDP(UDP, nil, dstAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	dd, err := conn.WriteToUDP([]byte("ping"), &addr)
	if err != nil {
		return err
	}
	fmt.Println(dd)
	return nil
}

// maxDescriptors
// get OS max file descriptors,
// opening too much tcp connections, can cause, OS error
func maxDescriptors() int {
	output, _ := exec.Command("ulimit", "-n").Output()
	//output, err := search.Command("launchctl", "limit", "maxfiles").Output()

	str := string(output)
	str = strings.TrimSuffix(str, "\n")
	byteToInt, err := strconv.Atoi(str)
	if err != nil {
		return 100
	}
	return byteToInt
}
