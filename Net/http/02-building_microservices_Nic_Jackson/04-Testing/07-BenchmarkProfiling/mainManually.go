package main

import (
	"04-Testing/handlers"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to file")

// go run mainManually.go -cpuprofile ./handlers/cpu.prof -memprofile ./hendlers/heap.prof
// go tool pprof ./mainManually ./handlers/cpu.prof

// brew install graphviz
// apt-get install graphviz
func main() {
	flag.Parse()

	//we are checking whether we have specified an output file for CPU profiling, and if so,
	//we are creating the file and then starting the profiler with pprof.StartCPUProfile(f),
	//and passing it a reference to the file:
	//func StartCPUProfile(w io.Writer) error
	if *cpuprofile != "" {
		fmt.Println("Running with CPU profile")
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("Finished")
		if *memprofile != "" {
			f, err := os.Create(*memprofile)
			if err != nil {
				log.Fatal(err)
			}
			runtime.GC()
			pprof.WriteHeapProfile(f)
			defer f.Close()
		}
		if *cpuprofile != "" {
			pprof.StopCPUProfile()
		}

		os.Exit(0)
	}()

	handler := handlers.Search{}
	err := http.ListenAndServe(":8323", &handler)
	if err != nil {
		log.Fatal(err)
	}
}
