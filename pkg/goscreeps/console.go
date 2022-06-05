package goscreeps

import (
	"github.com/mjmar01/go-screeps/pkg/goscreeps/resources"
	"syscall/js"

	_ "unsafe"
)

type Console struct {
	ref js.Value
}

var console = Console{
	ref: js.Global().Get("console"),
}

func (c Console) Log(args ...interface{}) {
	jsArgs := make([]interface{}, len(args))
	for i, arg := range args {
		switch arg.(type) {
		case resources.IReference:
			jsArgs[i] = resources.Ref(arg.(resources.IReference))
		default:
			jsArgs[i] = js.ValueOf(arg)
		}
	}
	c.ref.Call("log", jsArgs...)
}

func (c Console) CheckErr(err error) {
	if err != nil {
		c.ref.Call("log", err.Error())
	}
}
