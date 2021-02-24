package mutate

import (
	"reflect"
)
import("../common")
var Funcs = []common.Identifier{common.Identifier{Function:vmod, Name:"vmod"}}
func add(int1 int, int2 int) int{
	return int1 + int2
}
func getUnderlyingAsValue(data interface{}) reflect.Value {
	return reflect.ValueOf(data)
}

func vmod(vars map[string]*common.Variable, line string){
	parsedline := common.Parse(line)
	thing := vars[parsedline[1]].Value
	switch c := thing.(type){
	case int:
		out := c + common.Strtoint(parsedline[2])
		vars[parsedline[1]].Value = out
	case string:
		out := c + (parsedline[2])
		vars[parsedline[1]].Value = out
	}
	// vars[parsedline[1]].Value = vars[parsedline[1]].Value + common.Strtoint(parsedline[2])
}