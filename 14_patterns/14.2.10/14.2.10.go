package main

import (
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

// メッセージの構造体
type hello struct{ Who string }

// アクターの構造体
type helloActor struct{}

// アクターのメールボックス受信時に呼ばれるメソッド
func (state *helloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

func main() {
	props := actor.FromProducer(func() actor.Actor { return &helloActor{} })
	pid := actor.Spawn(props)
	pid.Tell(&hello{Who: "Roger"})
	console.ReadLine()
}
