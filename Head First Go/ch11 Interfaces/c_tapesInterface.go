package main

import (
	"fmt"

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

func main() {
	var player Player // important
	player = gadget.TapePlayer{}
	fmt.Println("For Tape Player")
	mixtape := []string{"A", "B", "C"}
	playList(player, mixtape)

	// Now, we can do this
	fmt.Println("\nFor Tape Recorder")
	player = gadget.TapeRecorder{}
	playList(player, mixtape)

}
