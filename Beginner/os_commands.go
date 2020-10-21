package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		executeLs()
		executePwd()
		executeLsWithParams("ltr")
	}
}

func executeLsWithParams(p string) {
	out, err := exec.Command("ls", "-" + p).Output()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println("LS Command Successfully Executed with params")
	output := string(out[:])
	fmt.Println(output)
}

func executeLs() {
	out, err := exec.Command("ls").Output()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println("LS Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}



func executePwd()  {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Println("PWD Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}
