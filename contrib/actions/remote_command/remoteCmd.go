package main

import (
	"bytes"
	"flag"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func runRemoteCommand(user, password, address string, port int, cmd string) (string, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", address, port), config)
	if err != nil {
		return "", err
	}

	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(cmd)
	return b.String(), err
}

var userVar string
var passwordVar string
var hostVar string
var portVar int
var cmdVar string

func main() {
	flag.StringVar(&userVar, "u", "root", "SSH user")
	flag.StringVar(&passwordVar, "p", "Password", "SSH password")
	flag.StringVar(&hostVar, "h", "127.0.0.1", "Host address")
	flag.IntVar(&portVar, "P", 22, "SSH port")
	flag.StringVar(&cmdVar, "c", "ls", "Remote command")
	flag.Parse()

	fmt.Printf("Execute remote SSH command \"%s\" on %s\n", cmdVar, hostVar)
	result, err := runRemoteCommand(userVar, passwordVar, hostVar, portVar, cmdVar)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("result: %s", result)
	}
}
