package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"./tokens"
)
type variable struct{
	lifetime int
	value interface{}
}
func (v *variable) rmhealth(howmuch int){
	v.lifetime --;
}
func strtoint(param string) int {
    result, _ := strconv.Atoi(param)
    return result
}
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
	
	vars := make(map[string]*variable)
    file, err := os.Open("../main.rot")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		parsedline := Parse(strings.TrimSpace(scanner.Text())) // parses the line
		switch  parsedline[0]{
		case "v": // variable declaration
			vars[parsedline[1]] = &variable{value:parsedline[3], lifetime: strtoint(parsedline[4])}
		case "print": // print statement
			fmt.Println(vars[parsedline[1]].value)
			vars[parsedline[1]].rmhealth(1)
		}
		// essentially deallocates any vars that have run out of life
		for x := range vars{
			if(vars[x].lifetime <= 0){
				delete(vars, x)
			}
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}