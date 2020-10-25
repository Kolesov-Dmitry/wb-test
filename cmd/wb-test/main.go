package main

import (
	"log"
	"os"

	"dima.org/wb-test/internals/app/net"
	"dima.org/wb-test/internals/app/tools"
)

func main() {
	const maxWorkersAmount = 2

	pool := tools.NewURLWorkersPool(
		maxWorkersAmount,
		net.NewCountWorker(net.RequesterFunc(net.RequestURL)),
		os.Stdin,
	)

	pool.Start()

	log.Println("Total:", pool.Result())
}
