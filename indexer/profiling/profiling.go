package profiling

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
)

func StartProfiling() (cpuFile *os.File, memFile *os.File) {
	// Create CPU profiling file
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("Error creating CPU profile file:", err)
	}
	// Start CPU profiling
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal("Error starting CPU profiling:", err)
	}

	// Create Memory profiling file
	memFile, err = os.Create("mem.prof")
	if err != nil {
		log.Fatal("Error creating memory profile file:", err)
	}

	return cpuFile, memFile
}

func StopProfiling(cpuFile *os.File, memFile *os.File) {
	// Stop CPU profiling
	pprof.StopCPUProfile()
	cpuFile.Close()

	// Run memory profiling
	runtime.GC() // Force garbage collection before memory profiling
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		log.Fatal("Error writing memory profile:", err)
	}
	memFile.Close()
}

func StartHTTPProfiler() {
	go func() {
		log.Println("Starting pprof HTTP server on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Println("Error starting pprof HTTP server:", err)
		}
	}()
}
