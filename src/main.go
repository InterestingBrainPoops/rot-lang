package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
type variable struct{
	lifetime int
	value int
}
func (v *variable) rmhealth(howmuch int){
	v.lifetime --;
}
func SomeFunc(param string) int {
    result, _ := strconv.Atoi(param)
    return result
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
		line := strings.TrimSpace(scanner.Text())
		if(line[0] == 'v'){
			vars[string(line[2])] = &variable{value: SomeFunc(string(line[6])), lifetime: SomeFunc(string(line[8]))}
		}else if(line[0:5] == "print"){
			
			fmt.Println(vars[string(line[6])].value)
			ptr := vars[string(line[6])]
			ptr.rmhealth(1)
			if(vars[string(line[6])].lifetime == 0){
				vars[string(line[6])] = nil
			}
			if(vars[string(line[6])] == nil){
				panic("you attempted to access variable " + string(line[6]) + " after the lifetime finished.")
			}
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}