package utils

import (
	"bytes"
	"os"
	"os/exec"
)

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
