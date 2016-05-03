package main

import _ "fmt"
import _ "os"
import _ "os/signal"
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	_ "syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
