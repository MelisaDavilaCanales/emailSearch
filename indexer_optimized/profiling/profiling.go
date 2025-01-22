package profiling

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
)

func CreateCPUProfiling() {
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Error creating CPU profile file:", err)
	}

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		fmt.Println("Error starting CPU profiling:", err)
	}
	defer pprof.StopCPUProfile()
}

func CreateMemoryProfiling() {
	memFile, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("Error creating memory profile file:", err)
	}
	defer memFile.Close()

	runtime.GC()
}

func CreateTraceProfilin() {
	traceFile, err := os.Create("trace.prof")
	if err != nil {
		fmt.Println("Error creating trace profile file:", err)
	}
	defer traceFile.Close()

	err = trace.Start(traceFile)
	if err != nil {
		fmt.Println("Error starting trace profiling:", err)
	}
	defer trace.Stop()
}

func StartHTTPProfiler() {
	go func() {
		fmt.Println("Starting pprof HTTP server on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			fmt.Println("Error starting pprof HTTP server:", err)
		}
	}()
}
