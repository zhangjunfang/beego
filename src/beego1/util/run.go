package util

import (
	"fmt"
	"os"
	"os/exec"
)

func Pwd() string {
	pwd, _ := os.Getwd()
	return pwd
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func Run(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	return cmd.CombinedOutput()
}

func WriteFile(path string, content []byte) error {
	if IsExist(path) {
		os.Remove(path)
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	n, err := f.Write(content)
	fmt.Println("num", n)
	return err
}
