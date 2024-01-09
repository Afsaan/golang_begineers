package main

import "fmt"

var Constant = "Constant data" // this constant is Public because the first letter is caps

func main() {
	var username string = "afsan"
	fmt.Println(username)
	fmt.Printf("Varibales is of type: %T \n", username)

	var is_logged_in bool = true
	fmt.Println(is_logged_in)
	fmt.Printf("Varibales is of type: %T \n", is_logged_in)

	var complex_example complex64 = 11 + 10i
	fmt.Println(complex_example)
	fmt.Printf("Varibales is of type: %T \n", complex_example)

	var small_value uint8 = 255
	fmt.Println(small_value)
	fmt.Printf("Varibales is of type: %T \n", small_value)

	var small_float float32 = 255.789
	fmt.Println(small_float)
	fmt.Printf("Varibales is of type: %T \n", small_float)

	// implict type
	var another_varibale = 100
	fmt.Println(another_varibale)
	fmt.Printf("Varibales is of type: %T \n", another_varibale)

	// without var - please note this technique cannot be used as a global varibale
	user_data := "this is user data"
	fmt.Println(user_data)
	fmt.Printf("Varibales is of type: %T \n", user_data)

	fmt.Println(Constant)
	fmt.Printf("Varibales is of type: %T \n", Constant)
}
