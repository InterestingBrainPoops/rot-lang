package packagemanager

import ("../std"
 "../mutate"
"../common")


func IMPORT(vars map[string]*common.Variable, line string){
	parsedline := common.Parse(line)
	//fmt.Println(parsedline)
	for _, x := range packages[parsedline[1]]{
		UsableFunctions[x.Name] = x.Function
	}
}
var packages = map[string][]common.Identifier{"stdlib":std.Funcs,"mutate":mutate.Funcs,}

var Functions = [...]common.Identifier{common.Identifier{Function: IMPORT, Name:"import"}}

var UsableFunctions = make(map[string](func(map[string]*common.Variable, string)))