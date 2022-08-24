package main

import "fmt"

func main() {
	//var num1  = 280
	//var num2 = 263
	//var num3 = 344
	//fmt.Println("Hello World", num3, num2, num1)
	//
	//scanner := bufio.NewScanner(os.Stdin)
	//fmt.Println("Enter the first integer:  ")
	//scanner.Scan()
	//input1 := scanner.Text()
	//fmt.Println("Enter the second integer:  ")
	//scanner.Scan()
	//input2 := scanner.Text()
	//fmt.Println("Enter the third integer:  ")
	//scanner.Scan()
	//input3 := scanner.Text()
	//
	//fmt.Println(input1, " ", input2, " ", input3)
	//fmt.Println(len("Hello"))
	//var (
	//	x = "hello"
	//	y = "world"
	//)
	//fmt.Println(x == y)

	//fmt.Print("Enter a number: ")
	//var input int
	//fmt.Scanf("%d", &input)
	//output := input * 2
	//fmt.Println(output)

	var C float64
	var F float64
	fmt.Println("Enter the Fahrenheit")
	fmt.Scanf("%f", &F)
	C = (F - 32) * 5 / 9
	fmt.Printf("The resulting Celsius is %.2f", C)

}
