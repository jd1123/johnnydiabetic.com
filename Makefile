all:
	make clean
	make build
	make cookiegen
	make createuser

clean:
	rm -f johnnydiabetic.com
	rm -f createuser
	rm -f cookiegen

build:
	go build

createuser:
	go build helpers/createuser/createuser.go

cookiegen:
	go build helpers/cookiegen/cookiegen.go  
