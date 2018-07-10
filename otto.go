package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"syscall/js"
)

func main() {
	document := js.Global().Get("document")
	body := document.Get("body")
	textarea := document.Call("createElement", "textarea")
	textarea.Call("setAttribute", "style", "width:100%;height:100px")
	textarea.Set("value", "function fib(num) {\n  if (num <= 1) return 1;\n  return fib(num - 1) + fib(num - 2);\n}\nfib(20);")
	run := document.Call("createElement", "input")
	run.Call("setAttribute", "type", "submit")
	run.Call("setAttribute", "value", "Run")
	run.Call("setAttribute", "onclick", "runOtto()")

	result := document.Call("createElement", "pre")
	js.Global().Set("runOtto", js.NewCallback(func(args []js.Value) {
		vm := otto.New()
		code := textarea.Get("value").String()
		value, err := vm.Run(code)
		if err != nil {
			result.Set("innerText", fmt.Sprintf("error running javascript\n%v", err))
		}
		result.Set("innerText", fmt.Sprintf("returned\n%v", value))
	}))

	body.Call("appendChild", textarea)
	body.Call("appendChild", document.Call("createElement", "br"))
	body.Call("appendChild", run)
	body.Call("appendChild", document.Call("createElement", "br"))
	body.Call("appendChild", result)
	select {}
}
