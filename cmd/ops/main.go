package main

import (
	"bytes"
	"flag"
	"os"
	"os/exec"
)

var configFile = flag.String("c", "configs/config.yml", "set config file which viper will loading.")

func main() {
	flag.Parse()

	//app, err := CreateApp(*configFile)
	//if err != nil {
	//	panic(err)
	//}
	//
	//if err := app.Start(); err != nil {
	//	panic(err)
	//}
	//
	//app.AwaitSignal()

	result, s, _ := RunWithResult("pwd", "/Users/life/Project/go-template/cmd/ops")
	println(result)
	println(s)
}


func Run(cmd string, dir string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Dir = dir
	c.Env = os.Environ()
	if err := c.Run(); err != nil {
		return err
	}
	return nil
}

func RunWithResult(cmd string, dir string, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c := exec.Command(cmd, args...)
	c.Dir = dir
	c.Stdout = &stdout
	c.Stderr = &stderr
	c.Env = os.Environ()
	err := c.Run()
	return stdout.String(), stderr.String(), err
}

func RunCombine(cmd string, dir string, args ...string) (string, error) {
	var stdout bytes.Buffer
	c := exec.Command(cmd, args...)
	c.Dir = dir
	c.Env = os.Environ()
	c.Stdout = &stdout
	if err := c.Run(); err != nil {
		return stdout.String(), err
	}
	return stdout.String(), nil
}