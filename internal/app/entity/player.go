package entity

import (
	"crypto/md5"
	"encoding/hex"
)

type Player struct {
	id   string
	Name string
}

func NewPlayer(name string) *Player {
	hash := md5.Sum([]byte(name))
	return &Player{
		id:   hex.EncodeToString(hash[:]),
		Name: name,
	}
}
