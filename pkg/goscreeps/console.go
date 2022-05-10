package goscreeps

import "syscall/js"

type Console struct {
	ref js.Value
}

var console = Console{
	ref: js.Global().Get("console"),
}

func (c Console) Log(args ...interface{}) {
	c.ref.Call("log", args...)
}
