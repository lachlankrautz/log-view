package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	args := []string{"rex logs"}
	if runtime.GOOS == "windows" {
		args = append([]string{
			"C:\\msys64\\usr\\bin\\env.exe",
			"MSYSTEM=MINGW64",
			"C:\\msys64\\usr\\bin\\bash.exe",
			"-l",
			"-c"},
			args...)
	}

	fmt.Print(args)

	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	fmt.Printf("output: \n%s\n", string(out))
	if err != nil {
		log.Fatal("process failed: ", err)
	}
}
