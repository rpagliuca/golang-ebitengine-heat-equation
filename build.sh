env GOOS=js GOARCH=wasm go build -o web/heat.wasm app/cmd/run
cp $(go env GOROOT)/misc/wasm/wasm_exec.js web/
npx http-server web/
