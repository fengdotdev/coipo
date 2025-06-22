package blueprints

// if some err for wasm go onload
// try update the wasm_exec.js file
// some missings are fix with the latest go version

type IndexBlueprint struct {
	Favicon  string // Path to favicon
	Style    string // Path to CSS style
	WasmExec string // Path to wasm_exec.js
	Wasm     string // Path to main.wasm
	Title    string // Page title
	Body     string // Body content of the page
}
