package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	fname_reader := bufio.NewReader(os.Stdin)
	fmt.Print("First Name: ")
	fname, _ := fname_reader.ReadString('\n')
	fname = strings.TrimSpace(fname)

	lname_reader := bufio.NewReader(os.Stdin)
	fmt.Print("Last Name: ")
	lname, _ := lname_reader.ReadString('\n')
	lname = strings.TrimSpace(lname)

	fmt.Println(fname, lname)

}
