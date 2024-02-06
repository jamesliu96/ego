//go:build js && wasm

package ego

import "syscall/js"

func BytesToUint8Array(bytes []byte) js.Value {
	array := js.Global().Get("Uint8Array").New(len(bytes))
	js.CopyBytesToJS(array, bytes)
	return array
}

func Uint8ArrayToBytes(array js.Value) []byte {
	bytes := make([]byte, array.Length())
	js.CopyBytesToGo(bytes, array)
	return bytes
}
