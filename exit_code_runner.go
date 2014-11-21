package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

// Runs a binary a number of times, and records the number of succesful
// runs it had
func main() {
	iterations, err := strconv.ParseInt(os.Args[1], 10, 0)
	if err != nil {
		log.Fatal("not a number", os.Args[1])
	}

	cmd_to_exec := os.Args[2]
	params := os.Args[3:]

	succesful_runs := 0

	for i := 0; i < int(iterations); i++ {
		fmt.Println("running", cmd_to_exec, "with args", params)
		out, err := exec.Command(cmd_to_exec, params...).CombinedOutput()
		if err == nil {
			succesful_runs += 1
		}
		fmt.Println("run", i, "output:")
		fmt.Println(string(out))
	}

	fmt.Println("succesful runs", succesful_runs, "total runs", iterations)

}
