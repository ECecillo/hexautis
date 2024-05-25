package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print("hexautis>")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		if line == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		fmt.Printf("So you said : %s \n", line)
	}
}
