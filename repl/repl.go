package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ECecillo/hexautis/format"
	"github.com/ECecillo/hexautis/parser"
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

		// fmt.Printf("String Value : %s, Hexadecimal Content : %# x \n", line, line)
		fmt.Println(format.Hex(line))

		fmt.Println("")
		result, err := parser.HexToString(line)
		if err != nil {
			fmt.Println("Error : ", err)
		}
		fmt.Println("String value : ", string(result))
	}
}
