package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var clear *exec.Cmd = exec.Command("tput cuu 1")

func main() {
	binName := os.Args[1]
	args := os.Args[2:]

	path, err := exec.LookPath(binName)
	if err != nil {
		_, err = os.Stat(binName)
		if err != nil {
			panic("file not exists")
		}
		path = "./" + binName
	}

	ctx := context.Background()

	cmd := exec.CommandContext(ctx, path, args...)

	startTime := time.Now()

	go TimeCount(startTime)

	err = cmd.Run()

	if err != nil {
		panic(err)
	}

	fmt.Printf("\033[1A\033[K")
	fmt.Printf("time: %s\n", time.Since(startTime))
}

func TimeCount(counter time.Time) {
	for {
		fmt.Println(time.Since(counter))
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("\033[1A\033[K")
	}
}
