package util

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
)

func GetRandomPort(minPort, maxPort int) (int, error) {
	for i := 0; i < 100; i++ {
		port := rand.Intn(maxPort-minPort) + minPort
		_, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err == nil {
			return port, nil
		}
	}
	return -1, fmt.Errorf("no available port")
}
