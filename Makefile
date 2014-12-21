all:
	make clean
	make build

clean:
	rm johnnydiabetic.com

build:
	go build
