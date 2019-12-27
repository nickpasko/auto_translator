package model

import "math/rand"

type WcStatusData struct {
	Color string
	Text string
}

func ReadWcStatusFromDevice() (WcStatusData, error) {
	busy := rand.Intn(2) == 1
	if busy {
		return WcStatusData{
			Color:"red",
			Text:"Занят",
		}, nil
	}

	return WcStatusData{
		Color:"green",
		Text:"Свободен",
	}, nil
}