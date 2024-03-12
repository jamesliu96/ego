//go:build js && wasm

package ego

import "syscall/js"

func IsObject(v js.Value) bool {
	vType := v.Type()
	return vType == js.TypeObject || vType == js.TypeFunction
}

func IsFunction(v js.Value) bool {
	vType := v.Type()
	return vType == js.TypeFunction
}
