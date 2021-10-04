package main

import (
	"fmt"
	"net"
	"os"

	myLogger "github.com/naka-gawa/pingo/pkg/logger"
	"go.uber.org/zap"
)

type icmpEchoPacket struct {
	Type     byte
	Code     byte
	Checksum uint16
	Identify uint16
	Sequence uint16
	Data     []byte
}

func checkArgs(args []string) string {
	if len(args) != 2 {
		zap.S().Errorf("args error: %v", args)
	}
	ipAddr := args[1]
	zap.S().Infof("args: %v", ipAddr)

	_, err := net.LookupIP(ipAddr)
	if err != nil {
		zap.S().Errorf("invalid address: %v", ipAddr)
		os.Exit(1)
	}

	zap.S().Infof("args addr: %v", ipAddr)
	return ipAddr
}

func pinger(id uint16) {
	fmt.Println(id)
}

func main() {
	myLogger.InitialConfig()
	checkArgs(os.Args)

	id := uint16(os.Getpid())
	pinger(id)
}
