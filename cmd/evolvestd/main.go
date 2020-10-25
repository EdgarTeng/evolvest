package main

import (
	"github.com/EdgarTeng/evolvest/pkg/embed"
	"github.com/EdgarTeng/evolvest/pkg/kit"
)

func main() {
	port := ":8762"
	go embed.StartServer(port)
	kit.WaitSignal()
}
