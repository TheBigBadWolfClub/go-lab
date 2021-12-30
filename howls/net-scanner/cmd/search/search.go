package main

import (
	"fmt"

	"github.com/TheBigBadWolfClub/go-lab/howls/net-scanner/internal"
)

func main() {
	maxPort := 5000
	ports := make([]uint16, 0, maxPort)
	for i := maxPort; i > 0; i-- {
		ports = append(ports, uint16(i))
	}
	scan := internal.Scan(ports...)
	fmt.Println(scan)
}
