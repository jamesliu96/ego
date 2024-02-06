//go:build js && wasm

package ego

// Keep thread alive
func KeepAlive() {
	<-make(chan struct{})
}
