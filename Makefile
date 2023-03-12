
all: build	

fmt:
	go fmt .

tidy:   fmt
	go mod tidy

build:  tidy
	go build .

clean:
	rm -f zoom-wh
 
