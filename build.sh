env GOOS=js GOARCH=wasm go build -o docs/heat.wasm app/cmd/run
cp $(go env GOROOT)/misc/wasm/wasm_exec.js docs/
npx http-server docs/
