package main


import (
	"main/com/codeconveyor/MorseDecoder/controller"
	"time"
	"math/big"
	"fmt"
	"log"
	"github.com/wblakecaldwell/profiler"

	"github.com/pkg/profile"
	"net/http"
)

func main() {
	profiler.AddMemoryProfilingHandlers()
	go http.ListenAndServe(":6060", nil)
	defer profile.Start(profile.CPUProfile).Stop()
	fmt.Scanln()
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	controller.Decode()

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	fmt.Scanln()
}