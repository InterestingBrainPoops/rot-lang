package common

import ("strings"
 "../tokens"
"strconv")
// variable struct
type Variable struct {
	Lifetime int
	Value    interface{}
}
// identifier, used for functions
type Identifier struct{
	Function func(map[string]*Variable, string)
	Name string
}
// removes 1 health given a variable
func (v *Variable) Rmhealth(howmuch int) {
	v.Lifetime--
}
// string to int
func Strtoint(param string) int {
	result, _ := strconv.Atoi(param)
	return result
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