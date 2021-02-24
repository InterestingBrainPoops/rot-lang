package main

import (
	"bufio"

	"log"
	"os"
	"./packagemanager"
	"./common"
	"strings"
)

// splits a function up into a list of usable directions/parts.
// "print(x)" becomes
// ["print","x"]
// "v a = 3, 4" becomes
// ["v", "a", "=", "3", "4"]

func main() {

	vars := make(map[string]*common.Variable)
	file, err := os.Open("../main.rot")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, x := range packagemanager.Functions {
		packagemanager.UsableFunctions[x.Name] = x.Function
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsedline := common.Parse(strings.TrimSpace(scanner.Text())) // parses the line
		packagemanager.UsableFunctions[parsedline[0]](vars, scanner.Text())
		// essentially deallocates any vars that have run out of life
		for x := range vars {
			if vars[x].Lifetime <= 0 {
				delete(vars, x)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
