package main

import (
	"syscall/js"

	"github.com/go-wasm/src/messages"
)

func main() {
	document := js.Global().Get("document")

	p := document.Call("createElement", "p")
	p.Set("innerHTML", messages.Message)
	p.Set("className", "lead")

	document.Call("getElementById", "content").Call("appendChild", p)
}

// RegisterFunction registers a JS function
func RegisterFunction(funcName string, myfunc func(js.Value, []js.Value) interface{}) {
	js.Global().Set(funcName, js.FuncOf(myfunc))
}
