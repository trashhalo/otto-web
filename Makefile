PWD = $(shell pwd)

build:
	docker run -v "$(PWD)":/go/src/app nlepage/golang_wasm bash -c "go get -d -v ./... && go build -o app.wasm app && cp /go/app.wasm /go/src/app"

serve:
	docker build -t nginx .
	docker run -p 3000:80 -v $(PWD):/usr/share/nginx/html nginx