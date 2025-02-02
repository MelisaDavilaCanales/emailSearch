package profiling

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
)

func StartProfiling() (cpuFile, memFile, traceFile *os.File) {
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Error creating CPU profile file:", err)
	}

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		fmt.Println("Error starting CPU profiling:", err)
	}

	memFile, err = os.Create("mem.prof")
	if err != nil {
		fmt.Println("Error creating memory profile file:", err)
	}

	traceFile, err = os.Create("trace.out")
	if err != nil {
		fmt.Println("Error creating trace file:", err)
	}

	if err := trace.Start(traceFile); err != nil {
		fmt.Println("Error starting trace:", err)
	}

	return cpuFile, memFile, traceFile
}

func StopProfiling(cpuFile, memFile, traceFile *os.File) {
	pprof.StopCPUProfile()

	if cpuFile != nil {
		cpuFile.Close() // nolint: errcheck
	}

	runtime.GC()

	if memFile != nil {
		if err := pprof.WriteHeapProfile(memFile); err != nil {
			fmt.Println("Error writing memory profile:", err)
		}

		memFile.Close() // nolint: errcheck
	}

	trace.Stop()

	if traceFile != nil {
		traceFile.Close() //nolint :errcheck
	}
}
