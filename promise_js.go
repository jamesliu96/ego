//go:build js && wasm

package ego

import (
	"fmt"
	"sync"
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
				x := fn(this, args)
				switch x := x.(type) {
				case []byte:
					resolve.Invoke(BytesToUint8Array(x))
				default:
					resolve.Invoke(x)
				}
			}()
			return nil
		}))
	})
}

// Awaits a promise to be resolved and returned synchronously or rejected as an error.
func Await(v js.Value) (ret js.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%+v", r)
		}
	}()
	if !v.InstanceOf(js.Global().Get("Promise")) {
		ret = v
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	v.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
		ret = args[0]
		wg.Done()
		return nil
	}), js.FuncOf(func(this js.Value, args []js.Value) any {
		err = fmt.Errorf("%+v", args[0])
		wg.Done()
		return nil
	}))
	wg.Wait()
	return
}
