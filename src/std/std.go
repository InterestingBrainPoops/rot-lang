package std

import (
	"fmt"
	"strings"
	"../common"
)

// defines a variable
func v(vars map[string]*common.Variable, line string){
	parsedline := common.Parse(line)
	variable := strings.TrimSpace(line[strings.Index(line, "=")+1:strings.Index(line, ",")])
	// fmt.Println(variable)
	if(variable[0] == '"'){
		vars[parsedline[1]] = &common.Variable{Value:variable[1:len(variable)-1], Lifetime: common.Strtoint(parsedline[len(parsedline)-1])}	
	}else{
		vars[parsedline[1]] = &common.Variable{Value:common.Strtoint(variable), Lifetime: common.Strtoint(parsedline[len(parsedline)-1])}
	}
	
}
// prints the given variable
func print(vars map[string]*common.Variable, line string){
	parsedline := common.Parse(line)
	// fmt.Println(vars[parsedline[1]])
	fmt.Println(vars[parsedline[1]].Value)
	vars[parsedline[1]].Rmhealth(1)
}
// deallocates the given vairable
func free(vars map[string]*common.Variable, line string){
	parsedline := common.Parse(line)
	delete(vars, parsedline[1])
}
// extends the lifetime of a given variable
func extend(vars map[string]*common.Variable, line string){
	parsedline := common.Parse(line)
	vars[parsedline[1]].Lifetime += common.Strtoint(parsedline[2])
}

var Funcs = []common.Identifier{common.Identifier{Function: v, Name:"v"}, common.Identifier{Function: print, Name:"print"}, common.Identifier{Function: free, Name:"free"}, common.Identifier{Function: extend, Name:"extend"}}
