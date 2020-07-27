package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {
	// obserbableを作成
	emitter := make(chan interface{})
	source := observable.Observable(emitter)

	// イベントを受け取るobserverを作成
	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			line := item.(string)
			if strings.HasPrefix(line, "func") {
				fmt.Println(line)
			}
		},
		ErrHandler: func(err error) {
			fmt.Println("Done!")
		},
	}

	// observableとobserverを接続（購読）
	sub := source.Subscribe(watcher)

	// observableに値を投入
	go func() {
		content, err := ioutil.ReadFile("14.2.9.go")
		if err != nil {
			emitter <- err
		} else {
			for _, line := range strings.Split(string(content), "\n") {
				emitter <- line
			}
		}
		close(emitter)
	}()

	// 終了待ち
	<-sub
}
