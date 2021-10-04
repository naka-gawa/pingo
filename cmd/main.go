package main

import (
	"net"
	"os"

	myLogger "github.com/naka-gawa/pingo/pkg/logger"
	"go.uber.org/zap"
)

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

func main() {
	myLogger.InitialConfig()
	checkArgs(os.Args)
}
