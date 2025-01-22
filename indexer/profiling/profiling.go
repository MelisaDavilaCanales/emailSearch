package profiling

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
)

func StartProfiling() (cpuFile, memFile, traceFile *os.File) {
	// Crear archivo para perfil de CPU
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Error creating CPU profile file:", err)
	}
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		fmt.Println("Error starting CPU profiling:", err)
	}

	// Crear archivo para perfil de memoria
	memFile, err = os.Create("mem.prof")
	if err != nil {
		fmt.Println("Error creating memory profile file:", err)
	}

	// Crear archivo para perfil de rastreo
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
	// Detener perfilado de CPU
	pprof.StopCPUProfile()
	if cpuFile != nil {
		cpuFile.Close()
	}

	// Forzar recolección de basura antes del perfil de memoria
	runtime.GC()
	if memFile != nil {
		if err := pprof.WriteHeapProfile(memFile); err != nil {
			fmt.Println("Error writing memory profile:", err)
		}
		memFile.Close()
	}

	// Detener rastreo
	trace.Stop()
	if traceFile != nil {
		traceFile.Close()
	}
}

func StartHTTPProfiler() {
	go func() {
		fmt.Println("Starting pprof HTTP server on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			fmt.Println("Error starting pprof HTTP server:", err)
		}
	}()
}
