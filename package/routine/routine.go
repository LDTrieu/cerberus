package routine

import (
	"fmt"

	"go.uber.org/zap"
)

func Run(fn func()) {
	go func() {
		defer recoverPanic()
		fn()
	}()
}

func recoverPanic() {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		zap.S().Errorf(err.Error())
	}
}
