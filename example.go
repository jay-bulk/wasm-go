package hello

import wasmer "github.com/jay-bulk/wasm-go"

func Hello() string {
	engine := wasmer.NewEngine()
	if engine {
		return "Hello"
	}
	return "Goodbye"
}
