package main

import (
	"context"
	"fmt"
	"scratch-go/internal/foo"
	"scratch-go/internal/logger"
)

type GlobalEnv struct {
	LogChan chan string
	FooChan chan foo.GetMsg
}

func main() {
	global := GlobalEnv{
		LogChan: make(chan string),
		FooChan: make(chan foo.GetMsg),
	}
	ctx, cancel := context.WithCancel(context.Background())
	env := foo.InitFooEnv(global.FooChan, global.LogChan)
	env.Ctx = ctx

	go env.FooMan()

	logCtx, logCancel := context.WithCancel(context.Background())
	logEnv := logger.LogEnv{LogInput: global.LogChan, Ctx: logCtx}

	go logEnv.LogMan()

	otherEnv := &OtherEnv{
		ReqChan: global.FooChan,
		LogChan: global.LogChan,
	}

	if got, err := otherEnv.Getter(foo.Id(10)); err == nil {
		fmt.Println(got)
	}
	cancel()
	logCancel()

	for {
	}
}

type OtherEnv struct {
	ReqChan chan foo.GetMsg
	LogChan chan string
}

func (e *OtherEnv) Getter(id foo.Id) (foo.Foo, error) {
	retChan := make(chan foo.Foo)
	e.ReqChan <- foo.GetMsg{Id: id, ReturnChan: retChan}

	gotJourn, ok := <-retChan
	if !ok {
		e.LogChan <- "getter didn't get anything back"
		return nil, foo.ErrIdNotFound
	}
	e.LogChan <- "Getter got something"
	return gotJourn, nil
}

type event struct {
	message string
}
