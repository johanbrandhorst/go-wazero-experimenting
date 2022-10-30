wasm/protoc-gen-go.wasm:
	GOOS=js GOARCH=wasm go build -o wasm/protoc-gen-go.wasm google.golang.org/protobuf/cmd/protoc-gen-go

wasm/protoc-gen-go-grpc.wasm:
	GOOS=js GOARCH=wasm go build -o wasm/protoc-gen-go-grpc.wasm google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc-gen-go: main.go
	go build -o protoc-gen-go ./main.go

protoc-gen-go-grpc: main.go
	go build -o protoc-gen-go-grpc ./main.go

all: wasm/protoc-gen-go.wasm wasm/protoc-gen-go-grpc.wasm protoc-gen-go protoc-gen-go-grpc

generate: all
	go run github.com/bufbuild/buf/cmd/buf@v1.9.0 generate

clean:
	rm -r wasm
	rm protoc-gen-go protoc-gen-go-grpc
