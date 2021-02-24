package std

import (
	"fmt"
	"strconv"
	"strings"

	"../tokens"
)

type Variable struct {
	Lifetime int
	value    interface{}
}
type Identifier struct{
	Function func(map[string]*Variable, string)
	Name string
}
func (v *Variable) rmhealth(howmuch int) {
	v.Lifetime--
}

func strtoint(param string) int {
	result, _ := strconv.Atoi(param)
	return result
}
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
func isDelimiter(r rune) bool{
	for _, x := range tokens.Delimiters{
		if(x == r){
			return true
		}
	}
	return false
}
func Parse(line string) []string{
	return strings.FieldsFunc(line, isDelimiter)
}
func print(vars map[string]*Variable, line string){
	parsedline := Parse(line)
	// fmt.Println(vars[parsedline[1]])
	fmt.Println(vars[parsedline[1]].value)
	vars[parsedline[1]].rmhealth(1)
}

func free(vars map[string]*Variable, line string){
	parsedline := Parse(line)
	delete(vars, parsedline[1])
}
func extend(vars map[string]*Variable, line string){
	parsedline := Parse(line)
	vars[parsedline[1]].Lifetime += strtoint(parsedline[2])
}
var Functions = [...]Identifier{Identifier{Function: v, Name:"v"}, Identifier{Function: print, Name:"print"}, Identifier{Function: free, Name:"free"}, Identifier{Function: extend, Name:"extend"}}
var Stdlib = make(map[string](func(map[string]*Variable, string)))