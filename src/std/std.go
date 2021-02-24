package std

import (
	"fmt"
	"strconv"
	"strings"

	"../tokens"
)
// variable struct
type Variable struct {
	Lifetime int
	value    interface{}
}
// identifier, used for functions
type Identifier struct{
	Function func(map[string]*Variable, string)
	Name string
}
// removes 1 health given a variable
func (v *Variable) rmhealth(howmuch int) {
	v.Lifetime--
}
// string to int
func strtoint(param string) int {
	result, _ := strconv.Atoi(param)
	return result
}
// defines a variable
func v(vars map[string]*Variable, line string){
	parsedline := Parse(line)
	variable := strings.TrimSpace(line[strings.Index(line, "=")+1:strings.Index(line, ",")])
	// fmt.Println(variable)
	if(variable[0] == '"'){
		vars[parsedline[1]] = &Variable{value:variable[1:len(variable)-1], Lifetime: strtoint(parsedline[len(parsedline)-1])}	
	}else{
		vars[parsedline[1]] = &Variable{value:strtoint(variable), Lifetime: strtoint(parsedline[len(parsedline)-1])}
	}
	
}
// checks if a rune (char) is a delimiter
func isDelimiter(r rune) bool{
	for _, x := range tokens.Delimiters{
		if(x == r){
			return true
		}
	}
	return false
}
// parses a line
func Parse(line string) []string{
	return strings.FieldsFunc(line, isDelimiter)
}
// prints the given variable
func print(vars map[string]*Variable, line string){
	parsedline := Parse(line)
	// fmt.Println(vars[parsedline[1]])
	fmt.Println(vars[parsedline[1]].value)
	vars[parsedline[1]].rmhealth(1)
}
// deallocates the given vairable
func free(vars map[string]*Variable, line string){
	parsedline := Parse(line)
	delete(vars, parsedline[1])
}
// extends the lifetime of a given variable
func extend(vars map[string]*Variable, line string){
	parsedline := Parse(line)
	vars[parsedline[1]].Lifetime += strtoint(parsedline[2])
}
func IMPORT(vars map[string]*Variable, line string){
	parsedline := Parse(line)
	fmt.Println(parsedline)
	for _, x := range packages[parsedline[1]]{
		// UsableFunctions[x.Name] = x.Function
	}
}
type module struct{
	name string
	reference []Identifier
}
var stdlib = []Identifier{Identifier{Function: v, Name:"v"}, Identifier{Function: print, Name:"print"}, Identifier{Function: free, Name:"free"}, Identifier{Function: extend, Name:"extend"}}
var packages = map[string][]Identifier{"stdlib":stdlib,}

var Functions = [...]Identifier{Identifier{Function: IMPORT, Name:"import"}}

var UsableFunctions = make(map[string](func(map[string]*Variable, string)))