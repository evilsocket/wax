all: wax

install: wax
	@cp wax /usr/local/bin/

deps: 
	@dep ensure

wax: deps
	@go build -o wax . 

clean:
	@rm -rf wax
