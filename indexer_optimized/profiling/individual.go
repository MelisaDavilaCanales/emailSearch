package profiling

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
)

// profiling.CreateCPUProfiling()
// profiling.CreateMemoryProfiling()
// profiling.CreateTraceProfilin()

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

	err = pprof.WriteHeapProfile(memFile)
	if err != nil {
		fmt.Println("Error writing memory profile:", err)
	}
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
