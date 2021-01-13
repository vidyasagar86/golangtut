package main

import "fmt"

func main() {
	variable_dec_syntax()
	if_else_syntax(23, 45)
	switch_syntax(5)
	array_syntax()
	pointer_syntax()
	pointer_toarray_syntax()
	array_of_pointers()
	passing_pointer_function()
	return_pointer_from_func()

}
func variable_dec_syntax() {
	//method1
	var result1 int
	result1 = 23
	fmt.Println(result1)
	//method2 dynamic declaration
	result2 := 24
	fmt.Println(result2)
	//method3
	var result3 int = 50
	fmt.Println(result3)
	//method4 mixed declaration
	var a, b, c, d = 1, 2.4, "sagar", true
	fmt.Println(a, b, c, d)
	//const declaration
	const length = 20
	const width = 50
	fmt.Println(length * width)
}
func if_else_syntax(x, y int) {
	if x > y {
		fmt.Println(x)
	} else if x < y {
		fmt.Println(y)
	} else {
		fmt.Println("comparison fialed")
	}
}
func switch_syntax(cond int) {
	switch {
	case cond == 1:
		fmt.Println("1")
	case cond == 2:
		fmt.Println("2")
	default:
		fmt.Println("not a number")
	}
}

func array_syntax() {
	var n [10]int /* n is an array of 10 integers */
	var i int

	/* initialize elements of array n to 0 */
	for i = 0; i < 5; i++ {
		n[i] = i + 100 /* set element at location i to i + 100 */
		fmt.Printf("Element[%d] = %d\n", i, n[i])
	}
	/* manual initislization of array
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}*/
}

func pointer_syntax() {
	var a int = 10
	var p *int
	p = &a
	fmt.Println(*p)
}

func pointer_toarray_syntax() {
	var i int
	var n = [2]int{1, 2}
	var p *[2]int //pointer to an array declaration
	p = &n
	for i = 0; i < 2; i++ {
		fmt.Println((*p)[i])
	}
}
func array_of_pointers() {
	var x int
	var n [2]*int //array of pointers declaration
	for x = 0; x < 2; x++ {
		y := x + 100
		n[x] = &y
	}
	for x = 0; x < 2; x++ {
		fmt.Println(*n[x])
	}
}

/*passing_pointer_function*/
func sample_funcfor_passing_pointer_function(x *int, y *int) {
	fmt.Println(*x + *y)
}
func passing_pointer_function() {
	a := 10
	b := 23
	sample_funcfor_passing_pointer_function(&a, &b)
}

/*sample_return_pointer_from_func*/
func sample_return_pointer_from_func() *int { /* return value of function signature */
	x := 10
	return &x
}
func return_pointer_from_func() {
	n := sample_return_pointer_from_func()
	fmt.Println(*n)
}
