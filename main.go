package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Errors
var (
	errNoBin = errors.New("there is no bin found")
	errUsage = errors.New("usage timex <command/bin> <args>")
)

func main() {
	err := Mesure()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func Mesure() error {
	if len(os.Args) < 2 {
		fmt.Println(errUsage)
		return errNoBin
	}
	binName := os.Args[1]
	args := os.Args[2:]

	path, err := exec.LookPath(binName)
	if err != nil {
		_, err = os.Stat(binName)
		if err != nil {
			return err
		}
		path = "./" + binName
	}

	ctx := context.Background()
	cmd := exec.CommandContext(ctx, path, args...)
	startTime := time.Now()
	go TimeCount(startTime)

	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf("\033[1A\033[K")
	fmt.Printf("time: %s\n", time.Since(startTime))
	return nil
}

func TimeCount(counter time.Time) {
	for {
		fmt.Println(time.Since(counter))
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("\033[1A\033[K")
	}
}
