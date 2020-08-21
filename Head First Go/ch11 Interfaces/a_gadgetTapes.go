package main

import (
	"github.com/headfirstgo/gadget"
)

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"A", "B", "C"}
	playList(player, mixtape)

	// can't do this, since gadget type is TapePlayer
	// player2 := gadget.TapeRecorder{}
	// playList(player2, mixtape)

}
