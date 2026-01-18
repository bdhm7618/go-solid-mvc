package main

import "fmt"

func main() {
	fmt.Print("hello MVC Framework") ; 


	var data map[string]any = map[string]any{
		"name" : "Abdo Mostafa" , 
		"email":"abdo@gmail.com" , 
		"phone" : "010953285567",
		"address" : [] any{
			"address one" , 
			"address two" , 
			12321 , 
		} ,
	}

	rules := map[string]any {
        "name": "required|min:3" , 
		"email": "required" , 
	}
	
	validate(data ,  rules) ; 
	fmt.Print(data , rules) ;
}