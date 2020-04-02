compile:
	GOOS=js GOARCH=wasm go build -o main.wasm

bench:
	go test -bench=Tick -cpuprofile=cpu.out github.com/d-baranowski/webassembly-game-of-life/src/wasm/life

seeBench:
	go tool pprof life.test cpu.out