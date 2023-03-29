## MDParseWASM

一个将markdown文档解析为html的WASM程序



### 编译

#### 如果您使用go编译
```shell
go env -w GOOS=js GOARCH=wasm
go build -ldflags="-s -w" -o main.wasm main.go
```
您还可以安装[bianryen](https://github.com/WebAssembly/binaryen/releases)进一步减小二进制文件大小
```shell
wasm-opt -Oz -o main_opt.wasm main.wasm
```

#### 如果您使用TinyGo编译
请确保您以安装[bianryen](https://github.com/WebAssembly/binaryen/releases)
```shell
tinygo build -o wasm.wasm -target wasm -opt=1 -no-debug main.go
```