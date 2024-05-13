package main

import "fmt"

func main() {

	//INT LITERALS
	// lvalue appears on both side but rvalues appears only on right side
	/* any number starts with "0X" or "0x" it is hexademical.
	with a "0" or "0o" or "0O" it is octal .
	with "0b" or "0B" it is binary.
	without a "0" it is decimal.const
	*/
	//fmt.Println(15 == 017)
	// var a int = 10
	// a++
	// println(a)
	// pre-decrement and pre increment is not present is go.
	// a := 10
	// // fmt.Printf("Binary of %d is %b\n", a, a)
	// // fmt.Printf("Octal of %d is %0o\n", a, a)

	// fmt.Printf("%d %T", a, 15 == 0xf)

	//float32 and float64
	// datatype mention kiya matlab static if nhi kiya ye kiya := to inferred
	//i, j := "hello", "world"
	// fmt.Print(i, "\n")
	// fmt.Print(j, "\n")

	// fmt.Printf("%s\n", i)
	// fmt.Printf("%s\n", j)

	// fmt.Printf("%s %s", i, j)
	//fmt.Printf("'")
	// var i = 15
	// fmt.Printf("%b\n", i)
	// fmt.Printf("%d\n", i)
	// fmt.Printf("%+d\n", i)
	// fmt.Printf("%o\n", i)
	// fmt.Printf("%O\n", i)
	// fmt.Printf("%x\n", i)
	// fmt.Printf("%X\n", i)
	// fmt.Printf("%#x\n", i)
	// fmt.Printf("%4d\n", i)
	// fmt.Printf("%-4d\n", i)
	// fmt.Printf("%04d\n", i)
	var txt = "Hello"
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

}
