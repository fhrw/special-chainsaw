package foo

import (
	"context"
	"fmt"
)

type EnvFooer struct {
	InputChan chan GetMsg
	LogChan   chan string
	Ctx       context.Context
	Cache     FooCache
}

type GetMsg struct {
	Id         Id
	ReturnChan chan Foo
}

func InitFooEnv(input chan GetMsg, log chan string) *EnvFooer {
	return &EnvFooer{
		InputChan: input,
		LogChan:   log,
		Cache:     InitFooCache(),
	}
}

func (e *EnvFooer) FooMan() {
	for {
		select {
		case getMsg := <-e.InputChan:
			got, err := e.Cache.GetFoo(getMsg.Id)
			if err != nil {
				close(getMsg.ReturnChan)
				e.LogChan <- fmt.Sprintf("didn't get %v", getMsg.Id)
				continue
			}
			e.LogChan <- fmt.Sprintf("got %v", getMsg.Id)
			getMsg.ReturnChan <- got
		case <-e.Ctx.Done():
			e.LogChan <- "FooMan stopping"
			return
		}
	}
}
