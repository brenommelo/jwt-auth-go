build:
	@go build -o bin/jwtauth.exe

run: build
	@CompileDaemon --build="go build -o ./bin/jwtauth.exe" --command="./bin/jwtauth.exe"

br: build
	@./bin/jwtauth.exe

test:
	@go test -v ./...