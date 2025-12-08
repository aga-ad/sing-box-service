package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sing-box-service/pkg/service"
	"sing-box-service/pkg/singbox"
	"syscall"
)

func runTerminal() {
	ctx, cancel := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	err := singbox.Run(ctx)
	if err != nil {
		fmt.Println(err)
	}
}

func runService() {
	service.Run("sing-box-service", func(ctx context.Context) { singbox.Run(ctx) })
}

func main() {
	isService, err := service.IsWindowsService()
	if err != nil {
		log.Printf("Can't check service mode: %v\n", err)
		return
	}

	if !isService {
		runTerminal()
	} else {
		runService()
	}
}
