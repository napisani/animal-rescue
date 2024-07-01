all: build 

build:
	go build -o animal-rescue .

test:
	go test -v ./...

clean:
	rm -f animal-rescue 
