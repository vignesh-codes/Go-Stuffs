package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//docker run cmd args

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Run what?")
	}
}
func run() {

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("running %v\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}