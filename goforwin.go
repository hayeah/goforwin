package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"time"
)

var retryCount int
var retryDelay int
var showHelp bool

func init() {
	flag.IntVar(&retryCount, "n", 2, "number of times to retry")
	flag.IntVar(&retryDelay, "d", 5, "seconds to wait before retry")
	flag.BoolVar(&showHelp, "h", false, "show help")
}

func main() {
	var err error
	flag.Parse()

	if showHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 1 {
		log.Fatalln("Please specify a command to run.")
	}

	count := retryCount
	for {
		log.Printf("exec: %v\n", args)
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err == nil {
			os.Exit(0)
		} else {
			log.Println(err)
			if count > 0 {
				count--
				log.Printf("retry in %d seconds...", retryDelay)
				time.Sleep(time.Duration(retryDelay) * time.Second)
			} else {
				log.Printf("retried %d times. give up.", retryCount)
				os.Exit(1)
			}
		}
	}
}
