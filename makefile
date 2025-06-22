PHONY: front back

main_frontend = frontend/cmd/main/main.go
main_backend = backend/cmd/main/main.go

web_assets = builders/forweb/assets
web_scripts = builders/forweb//scripts
web_templates = builders/forweb/templates

web_output = output/web
wasi_output = output/wasi


sand:


run:
	 $(web_output)/main


front: front_build front_assets
	@echo "Frontend build complete. Output at $(web_output)"

front_build:
	GOOS=js GOARCH=wasm go build -o $(web_output)/main.wasm $(main_frontend)

front_assets:
	mkdir -p $(web_output)/static/assets
	mkdir -p $(web_output)/static/scripts
	mkdir -p $(web_output)/static/templates
	mkdir -p $(web_output)/static/css
	mkdir -p $(web_output)/static/fonts
	mkdir -p $(web_output)/static/images
	cp $(web_assets)/* $(web_output)/static/assets/
	cp $(web_scripts)/* $(web_output)/static/scripts/
	cp $(web_templates)/* $(web_output)/static/templates/


wasi:
	GOOS=wasip1 GOARCH=wasm go build -o $(wasi_output)/main.wasm $(main_frontend)

runwasi:
	wasmtime $(wasi_output)/main.wasm

back:
	go run $(main_backend)