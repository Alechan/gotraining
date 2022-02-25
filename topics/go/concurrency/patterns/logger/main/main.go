// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program demonstrates how the logger package works.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

// device allows us to mock a device we write logs to.
type device struct {
	problem bool
}

// Write implements the io.Writer interface.
func (d *device) Write(p []byte) (n int, err error) {
	for d.problem {

		// Simulate disk problems.
		time.Sleep(time.Second)
	}

	fmt.Print(string(p))
	return len(p), nil
}

func main() {

	// Number of goroutines that will be writing logs.
	const grs = 10

	// Create a logger value with a buffer of capacity
	// for each goroutine that will be logging.
	var d device
	// Versión de logging “original” (usa el log oficial)
	l := log.New(&d, "prefix", 0)
	// Versión de logging “arreglada” (el logger no es el oficial de Go)
	//l := logger.New(&d, grs)

	// Generate goroutines, each writing to disk

	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				// Versión original que printea a STDOUT
				l.Println(fmt.Sprintf("%d: log data", id))
				// Versión arreglada que usa el logger ad-hoc
				//l.Write(fmt.Sprintf("%d: log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// We want to control the simulated disk blocking. Capture
	// interrupt signals to toggle device issues. Use <ctrl> z
	// to kill the program.

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan

		// I appreciate we have a data race here with the Write
		// method. Let's keep things simple to show the mechanics.
		d.problem = !d.problem
	}
}
