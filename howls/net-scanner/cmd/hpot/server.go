package main

import (
	"fmt"
	"github.com/TheBigBadWolfClub/go-lab/howls/net-scanner/internal/pots"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	port := uint16(4445)

	// TCP
	go func() {
		err := pots.TcpServer(port)
		if err != nil {
			fmt.Println("fail server TCP")
			os.Exit(1)
		}
		wg.Done()

	}()
	wg.Add(1)

	go func() {
		err := pots.UdpServer(port)
		if err != nil {
			fmt.Println("fail server UDP")
			os.Exit(1)
		}
		wg.Done()
	}()
	wg.Add(1)

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")

}
