.PHONY: run
run:
	@go run cmd/main.go

# Create a tailwind.config.js file
.PHONY: tw-init
tw-init:
	./tailwindcss init

# Start a watcher
.PHONY: tw-watch
tw-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch


# Compile and minify your CSS for production
.PHONY: tw-build
tw-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
