package main

import (
	"test/gadget"
)

type Player interface { // Реализация интерфейса
	Play(string)
	Stop()
}

func playList(device Player, song []string) {
	for _, song := range song {
		device.Play(song)

	}

	device.Stop()
}

func TryOut(player Player) {
	player.Play("Track")
	player.Stop()
	recorder := player.(gadget.TapeRecorder)
	recorder.Recoder()
}

func main() {

	var player Player = gadget.TapePlayer{}
	mixtape := []string{"A", "B", "C"} //A,B,С Песни которые будут играть
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	playList(player, mixtape)

	TryOut(gadget.TapeRecorder{})
}
