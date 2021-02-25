package slice

import (
	"strings"

	"../common"
)

var Funcs = []common.Identifier{common.Identifier{Function:bmake, Name:"b"}, common.Identifier{Function:badd, Name:"append"} ,common.Identifier{Function:bset, Name:"set"}}

func bmake(vars map[string]*common.Variable, line string){
	// fmt.Println("thing")
	parsedline := common.Parse(line)
	// syntax : name, type, size, lifetime
	switch parsedline[2]{
	case "string":
		vars[parsedline[1]] = &common.Variable{Value: make([]string, common.Strtoint(parsedline[3])), Lifetime:common.Strtoint(parsedline[4])}
	case "int":
		vars[parsedline[1]] = &common.Variable{Value: make([]int, common.Strtoint(parsedline[3])), Lifetime:common.Strtoint(parsedline[4])}
	}

}

func badd(vars map[string]*common.Variable, line string){
	parsedline := common.Parse(line)
	thing := vars[parsedline[1]].Value
	switch c := thing.(type){
	case []int:
		vars[parsedline[1]].Value = append(c, common.Strtoint(parsedline[2]))
	case []string:
		
		vars[parsedline[1]].Value = append(c, line[strings.Index(line, "\"")+ 1: strings.LastIndex(line, "\"")])
	}
	vars[parsedline[1]].Lifetime --;
	
}

func bset(vars map[string]*common.Variable, line string){
	parsedline := common.Parse(line)
	thing := vars[parsedline[1]].Value
	// usage : variable, index, item
	switch c := thing.(type){
	case []int:
		c[common.Strtoint(parsedline[2])] = common.Strtoint(parsedline[3])
		vars[parsedline[1]].Value = c
	case []string:
		c[common.Strtoint(parsedline[2])] = line[strings.Index(line, "\"")+ 1: strings.LastIndex(line, "\"")]
		vars[parsedline[1]].Value = c
	}
	vars[parsedline[1]].Lifetime --;
}