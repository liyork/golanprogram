package cg

import "fmt"

type Player struct {
	Name  string "name"
	Level int    "level"
	Exp   int    "exp"
	Room  int    "room"

	mq chan *Message //等待收到的消息
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}

	go func(p *Player) { //新gorutine执行
		for {
			msg := <-p.mq //读取消息
			fmt.Println(p.Name, "received message:", msg.Content)
		}
	}(player)

	return player
}
