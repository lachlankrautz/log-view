package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	args := getArgs()

	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	fmt.Printf("output: \n%s\n", string(out))
	if err != nil {
		log.Fatal("process failed: ", err)
	}
}

func getArgs() []string {
	// get logs from rex
	args := []string{"rex", "logs"}

	// append service names to filter logs
	args = append(args, os.Args[1:]...)

	// if windows run everything in msys shell
	if runtime.GOOS == "windows" {
		args = []string{
			"C:\\msys64\\usr\\bin\\env.exe",
			"MSYSTEM=MINGW64",
			"C:\\msys64\\usr\\bin\\bash.exe",
			"-l",
			"-c",
			strings.Join(args, " ")}
	}

	return args
}
