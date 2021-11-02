win:
	go run .\cmd\web\ .

unix:
	go run ./cmd/web/ *.go
	
tidy:
	go mod tidy

test-main-win:
	go test -v .\cmd\web\ 

test-main-unix:
	go test -v ./cmd/web/ 
