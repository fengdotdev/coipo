PHONY: front back

main_frontend = frontend/cmd/main/main.go
main_backend = backend/cmd/main/main.go
web_output = output/web


front:
	GOOS=js GOARCH=wasm go build -o $(web_output)/main.wasm $(main_frontend)

back:
	go run $(main_backend)