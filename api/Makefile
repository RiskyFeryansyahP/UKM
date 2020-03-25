dep:
    GO111MODULE=on go mod download
tidy:
	go mod tidy
	gofmt -s -w ./**/*.go
	goimports -w ./**/*.go

build: dep
	GO111MODULE=on CGO_ENABLED=0 go build -ldflags="-s -w" \
	-o build/ukm .

run: build
	./build/ukm

migrate: build
	./build/ukm --migrate=true

generate:
	entc generate ./ent/schema