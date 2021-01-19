package main

import (
	"fmt"
)

//////go:generate stringer -type=Pill
//go:generate stringer -type=Pill -linecomment
type Pill int

const (
	Placebo Pill = iota
	Aspirin //我是第二个的描述啊
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

//go:generate echo GoGoGo!
//go:generate go run main.go
//go:generate echo $GOARCH $GOOS $GOFILE $GOLINE $GOPACKAGE

func main() {
	fmt.Println("go rum main.go!")

	var a Pill = 1
	fmt.Printf("%s\n", a)
	fmt.Printf("%v\n", a)
	fmt.Println(a.String())

}
