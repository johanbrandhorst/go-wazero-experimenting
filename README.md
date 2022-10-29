# go-wazero-experimenting

Playing around with Wazero and Go wasm binaries

# Usage

Build the plugins with Go's WebAssembly and the local runner:

```shell
$ make all
GOOS=js GOARCH=wasm go build -o wasm/protoc-gen-go.wasm google.golang.org/protobuf/cmd/protoc-gen-go
GOOS=js GOARCH=wasm go build -o wasm/protoc-gen-go-grpc.wasm google.golang.org/grpc/cmd/protoc-gen-go-grpc
go build -o protoc-gen-go ./main.go
go build -o protoc-gen-go-grpc ./main.go
```

Run the generator:

```
$ make generate
```

This will run the webassembly compiled plugins using wazero's VM. Note that the first time will be slower as it has to compile the wasm binary to a local binary. The second invocation will be significantly faster:

```
$ time make generate
go run github.com/bufbuild/buf/cmd/buf@v1.9.0 generate

real    0m9.699s
user    0m30.246s
sys     0m1.143s
$ time make generate
go run github.com/bufbuild/buf/cmd/buf@v1.9.0 generate

real    0m1.672s
user    0m2.444s
sys     0m0.736s
```
