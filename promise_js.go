//go:build js && wasm

package ego

import (
	"fmt"
	"syscall/js"
)

// PromiseOf returns an async function to be used by JavaScript.
func PromiseOf(fn func(this js.Value, args []js.Value) any) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return js.Global().Get("Promise").New(js.FuncOf(func(_ js.Value, _args []js.Value) any {
			resolve, reject := _args[0], _args[1]
			go func() {
				defer func() {
					if r := recover(); r != nil {
						reject.Invoke(js.Global().Get("Error").New(fmt.Sprintf("%+v", r)))
					}
				}()
				resolve.Invoke(fn(this, args))
			}()
			return nil
		}))
	})
}
