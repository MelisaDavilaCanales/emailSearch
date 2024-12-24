package profiling

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
)

func StartProfiling() (cpuFile *os.File, memFile *os.File) {
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

	return cpuFile, memFile
}

func StopProfiling(cpuFile *os.File, memFile *os.File) {
	pprof.StopCPUProfile()
	cpuFile.Close()

	// Run memory profiling
	runtime.GC() // Force garbage collection before memory profiling
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		fmt.Println("Error writing memory profile:", err)
	}
	memFile.Close()
}

func StartHTTPProfiler() {
	go func() {
		fmt.Println("Starting pprof HTTP server on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			fmt.Println("Error starting pprof HTTP server:", err)
		}
	}()
}
