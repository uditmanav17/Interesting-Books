package main

import (
	"fmt"
	"reflect"

	"github.com/headfirstgo/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	fmt.Println("Type of", reflect.TypeOf(player))
	player.Play("Test Track")
	player.Stop()

	// try concrete conversion
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func main() {
	fmt.Println("For Tape Player")
	var player Player = gadget.TapePlayer{}
	TryOut(player)

	fmt.Println("\nFor Tape Recorder")
	player = gadget.TapeRecorder{}
	TryOut(player)

}
