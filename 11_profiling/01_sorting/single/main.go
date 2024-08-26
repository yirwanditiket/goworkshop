package main

import (
	"context"
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	pprof.Do(context.Background(), pprof.Labels("customlabel", "calculate fibbonachi of 43"), func(ctx context.Context) {
		log.Println(wrap(43))
	})
}

func wrap(x int) int {
	return fib(x)
}

func fib(x int) int {
	if x <= 1 {
		return x
	}

	return fib(x-1) + fib(x-2)
}
