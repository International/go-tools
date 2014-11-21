package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

	path, err := exec.LookPath(cmd_to_exec)
	if err != nil {
		log.Fatal(err)
	}

	args_to_cmd := strings.Join(params, " ")

	succesful_runs := 0

	for i := 0; i < int(iterations); i++ {
		out, err := exec.Command(path, args_to_cmd).Output()
		if err == nil {
			succesful_runs += 1
		}
		fmt.Println("run", i, "output:")
		fmt.Println(string(out))
	}

	fmt.Println("succesful runs", succesful_runs, "total runs", iterations)

}
