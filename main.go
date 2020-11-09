package main 

import "fmt"

func main() { 

	var idTonameMap = make(map[float64]string) 
	
	idTonameMap[1] = "Rohit"
	idTonameMap[2] = "Sumit"

	for id, name := range idTonameMap { 
  
        fmt.Println(id, name) 
    } 
} 

