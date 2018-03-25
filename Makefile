GORUN=go run
GOTEST=go test
GOBUILD=go build
GOCLEAN=go clean
BIN=gradient

run:
	$(GORUN) cmd/main.go -h=100 -out=./gen/gradient.png '#FF0000,#00FF00,#0000FF'

build:
	$(GOBUILD) -o $(BIN) -v ./cmd/main.go

test:
	clear
	$(GOTEST) -v ./pkg/...

bench:
	clear
	$(GOTEST) -v -bench=. ./pkg/...

clean: 
	rm ./$(BIN)
	rm ./gen/*
	$(GOCLEAN)

examples: ./$(BIN)
	./$(BIN) -h=100 -out=./gen/gradient-lin.png '#FF0000,#FF00FF,#0000FF'
	./$(BIN) -h=100 -out=./gen/gradient-quad.png -type=quadratic '#FF0000,#FF00FF,#0000FF'
	./$(BIN) -h=100 -out=./gen/gradient-cub.png -type=cubic '#FF0000,#FF00FF,#FF00FF,#0000FF'