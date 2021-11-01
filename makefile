win:
	go run .\cmd\web\ .

unix:
	go run ./cmd/web/ *.go
	
tidy:
	go mod tidy