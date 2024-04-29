run: build
	@./bin/app

build:
	@go build -o bin/app .

css:
	@npx tailwindcss -i views/css/app.css -o public/styles.css --watch

templ:
	@templ generate --watch --proxy=http://localhost:3000

init:
	@go install github.com/cosmtrek/air@latest
	@go install github.com/a-h/templ/cmd/templ@latest