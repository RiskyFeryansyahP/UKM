dep:
	GO111MODULE=on go mod download

build: dep
	GO111MODULE=on CGO_ENABLED=0 go build -ldflags="-s -w" \
	-o build/ukm .

run: build
	./build/ukm

migrate: build
	./build/ukm -migrate=true