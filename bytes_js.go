//go:build js && wasm

package ego

import "syscall/js"

// Converts bytes to Uint8Array
func BytesToUint8Array(bytes []byte) js.Value {
	array := js.Global().Get("Uint8Array").New(len(bytes))
	js.CopyBytesToJS(array, bytes)
	return array
}

// Converts Uint8Array to bytes
func Uint8ArrayToBytes(array js.Value) []byte {
	bytes := make([]byte, array.Length())
	js.CopyBytesToGo(bytes, array)
	return bytes
}
