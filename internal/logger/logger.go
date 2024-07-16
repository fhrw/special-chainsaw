package logger

import (
	"context"
	"fmt"
)

type LogEnv struct {
	LogInput chan string
	Ctx      context.Context
}

func (e *LogEnv) LogMan() {
	for {
		select {
		case str := <-e.LogInput:
			fmt.Println("===== LOG MESSSAGE =====")
			fmt.Println(str)
		case <-e.Ctx.Done():
			fmt.Println("Logger signing off...")
			return
		}
	}
}
