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
	// player.Record() // Wont work

	// type conversion wont work either
	// recorder := gadget.TapeRecorder(player)
	// recorder.Record()
}

func main() {
	var player Player = gadget.TapePlayer{}
	TryOut(player)

	player = gadget.TapeRecorder{}
	TryOut(player)
	// record still wont work, for tape recorder
	// player.Record()

	// need to recover concrete type using type assertion
	tapeRec := player.(gadget.TapeRecorder)
	tapeRec.Record()

}
