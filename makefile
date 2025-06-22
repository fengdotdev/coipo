PHONY: front back

main_frontend = frontend/cmd/main/main.go
main_backend = backend/cmd/main/main.go
web_output = output/web
wasi_output = output/wasi


front:
	GOOS=js GOARCH=wasm go build -o $(web_output)/main.wasm $(main_frontend)

wasi:
	GOOS=wasip1 GOARCH=wasm go build -o $(wasi_output)/main.wasm $(main_frontend)

runwasi:
	wasmtime $(wasi_output)/main.wasm

back:
	go run $(main_backend)