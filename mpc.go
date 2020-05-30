package main

import (
	"log"
	"os/exec"
)

type mpcCommand = int

const (
	volumeUp   mpcCommand = 0
	volumeDown mpcCommand = 1
)

func translateCommand(msg mpcCommand) []string {
	switch msg {
	case volumeDown:
		return []string{"volume", "-10"}
	case volumeUp:
		return []string{"volume", "+10"}
	default:
		return []string{}
	}
}

func sendCommand(msg mpcCommand) {
	args := []string{"-h", "localhost", "-p", "6600"}
	for _, sb := range translateCommand(msg) {
		args = append(args, sb)
	}
	log.Println(args)
	output, err := exec.Command("mpc", args...).CombinedOutput()
	if err != nil {
		log.Print(err.Error())
	} else {
		log.Print(string(output))
	}
}
