package demo1

import "fmt"

func MyAdd(a, b int) int {
	if b == 100 {
		return a+b
	}
	return a+b
}


func MyAddTwo(a, b int) int {
	type Calc struct {
		A,B,Result int
	}
	_=fmt.Sprintf("%d,%d\n",a,b)
	calc:=&Calc{
		A:      a,
		B:      b,
		Result: 0,
	}
	if b == 100 {

		calc.Result = calc.A+calc.B
		return calc.Result
	}
	return calc.Result
}
