package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/ssh"
)

// Init initialization of args from command
func Init() {
	var port int
	var host string
	var user string
	var password string

	flag.IntVar(&port, "port", 22, "port to connect to.")
	flag.StringVar(&host, "host", "", "host to connect to.")
	flag.StringVar(&user, "user", "", "user to connect with.")
	flag.StringVar(&password, "password", "", "password to use.")

	flag.Parse()

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}

	fmt.Println(port)
}
