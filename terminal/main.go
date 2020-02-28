package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"gitlab.com/vishwavijay.com/smd/raja/rio"
)

func main() {
	input, output, err := bash()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input, output, err)
	cli := rio.NewIo()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case line := <-output:
				fmt.Println(line)
			case sig := <-sigs:
				fmt.Println("-------------->", sig)
			default:
			}
		}
	}()
	for {
		input <- cli.ReadLine()
	}
}

func bash() (input chan string, output chan string, err error) {
	cmd := exec.Command("bash")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println(err)
		return
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Println(err)
		return
	}
	if err = cmd.Start(); err != nil {
		log.Println(err)
		return
	}

	input = make(chan string)
	output = make(chan string)

	go func() {
		defer log.Println("Closing go routine for output and error response!")
		log.Println("Writting erro and output to channel stdout !")
		s := bufio.NewScanner(io.MultiReader(stdout, stderr))
		for s.Scan() {
			//fmt.Println("waiting for output channel")
			line := string(s.Bytes())
			//fmt.Println("output : ", line)
			output <- line
		}
	}()

	go func() {
		defer log.Println("Closing go routine for input request!")
		log.Println("Writting channel input to bash stdin!")
		for {
			//fmt.Println("waiting for input channel")
			command := <-input
			if "exit" == command {
				cmd.Process.Signal(syscall.SIGINT)
			} else {
				stdin.Write([]byte(command + "\n"))
			}

		}
	}()
	return
}
