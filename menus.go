package main

import (
	"fmt"
)

// MainMenu array initialized here
var MainMenu = []string{
	"Continue",
	"Interactive",
	"Tunnels",
	"File Transfer",
	"Local Shell",
	"Exit",
}

// TunnelMenu Initialized here.
var TunnelMenu = []string{
	"Forward",
	"Reverse",
	"Cancel",
	"Main Menu",
}

// FileMenu Initialized here.
var FileMenu = []string{
	"SCP Upload",
	"SCP Download",
	"SFTP Upload",
	"SFTP Download",
	"Cat Binary Upload",
	"Cat Binary Download",
	"Main Menu",
}

// MenuDisplay function here for above menus.
func MenuDisplay(menu []string) int {
	var choice int = 0
	for true {
		if choice <= 0 || choice > len(menu) {
			for num, option := range menu {
				fmt.Println(num+1, ": ", option)
			}
			fmt.Println("Please Enter a Selection Number: ")
			fmt.Scanf("%d", &choice)
		} else if choice >= 1 || choice < len(menu) {
			break
		} else {
			fmt.Println("Incorrect Value Entered, Try again!")
			choice = 0
		}
	}
	return choice
}
