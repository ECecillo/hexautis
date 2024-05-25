package main

import (
	"fmt"
	"os"

	"github.com/ECecillo/hexautis/repl"
)

func main() {
	fmt.Println("======= HEXAUTIS REPL / ECecillo ======")
	fmt.Println("HEX to String <-> String to Hex")
	fmt.Println("Type 'exit' to quit the REPL")
	repl.Start(os.Stdin, os.Stdout)
}
