win:
	go run .\cmd\web\ .

unix:
	go run ./cmd/web/ *.go
	
tidy:
	go mod tidy

test-win:
	go test -v .\cmd\web\ 

test-unix:
	go test -v ./cmd/web/ 