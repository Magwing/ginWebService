package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"time"
)

func Run(servers ...IServer) {

	// defer panic
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("server start err: %+v", err)
		}
	}()

	// start servers
	for _, s := range servers {
		go s.Start()
	}

	// check servers status
	for {
		count := 0
		for _, s := range servers {
			if s.Running() {
				count++
			}
		}
		time.Sleep(1)
		if count == len(servers) {
			log.Println("servers starting ...")
			break
		}
	}

	// deal with signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	for exitSig := range sig {
		log.Println("receive exit signal ", exitSig)
		switch exitSig {
		case os.Interrupt:
			fallthrough
		case syscall.SIGTERM:
			log.Println("exiting ...")
			for _, s := range servers {
				s.Stop()
			}
			log.Println("exit success")
			os.Exit(0)
		}
	}

}
