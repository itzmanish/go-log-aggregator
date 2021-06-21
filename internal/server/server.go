package server

import (
	"bufio"
	"math/rand"
	"net"
	"time"

	"github.com/itzmanish/go-loganalyzer/internal/logger"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	port, err := cmd.Flags().GetString("port")
	if err != nil {
		logger.Error(err)
	}
	logger.Info(port)
	l, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		logger.Error(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			logger.Error(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buffer, err := bufio.NewReader(conn).ReadBytes('\n')

		if err != nil {
			return
		}

		logger.Info("Client message:", string(buffer[:len(buffer)-1]))

		conn.Write(buffer)
	}
}
