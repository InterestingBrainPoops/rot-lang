package std

import (
	"fmt"
	"strconv"
)

type Variable struct {
	Lifetime int
	value    interface{}
}
type Identifier struct{
	Function func(map[string]*Variable, []string)
	Name string
}
func (v *Variable) rmhealth(howmuch int) {
	v.Lifetime--
}

func strtoint(param string) int {
	result, _ := strconv.Atoi(param)
	return result
}
func v(vars map[string]*Variable, parsedline[]string){
	vars[parsedline[1]] = &Variable{value:parsedline[3], Lifetime: strtoint(parsedline[4])}
}

func print(vars map[string]*Variable, parsedline[]string){
	fmt.Println(vars[parsedline[1]].value)
	vars[parsedline[1]].rmhealth(1)
}
var Functions = [...]Identifier{Identifier{Function: v, Name:"v"}, Identifier{Function: print, Name:"print"}}
var Stdlib = make(map[string](func(map[string]*Variable, []string)))