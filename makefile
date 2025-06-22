PHONY: front

main_frontend = frontend/cmd/main/main.go


front:
	go run GOOS=js GOARCH=wasm go build -o main.wasm $(main_frontend)