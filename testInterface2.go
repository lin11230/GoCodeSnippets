package main

import "fmt"

function main(){
	names := []string{"stanely", "david", "oscar"}
	vals := make([]interface{}, len(names))
	for i,v := range names{
		vals[i] = v
	}

    PrintAll(vals)	
}

func PrintAll(vals []interface{}){
	for _, val := range vals{
		fmt.Println(val)
	}
}
