package main

import (
	"bufio"

	"log"
	"os"

	"strings"
	"./tokens"
	"./std"
)

// splits a function up into a list of usable directions/parts. 
// "print(x)" becomes
// ["print","x"]
// "v a = 3, 4" becomes 
// ["v", "a", "=", "3", "4"]

func Parse(line string) []string{
	return strings.FieldsFunc(line, isDelimiter)
}
func isDelimiter(r rune) bool{
	for _, x := range tokens.Delimiters{
		if(x == r){
			return true
		}
	}
	return false
}
func main() {
	
	vars := make(map[string]*std.Variable)
    file, err := os.Open("../main.rot")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	for _, x := range std.Functions{
		std.Stdlib[x.Name] = x.Function
	}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		parsedline := Parse(strings.TrimSpace(scanner.Text())) // parses the line
		std.Stdlib[parsedline[0]](vars, parsedline)
		// essentially deallocates any vars that have run out of life
		for x := range vars{
			if(vars[x].Lifetime <= 0){
				delete(vars, x)
			}
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}