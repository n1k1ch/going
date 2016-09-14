package main

import (
	"fmt"

	"github.com/n1k1ch/going/stringutil"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/n1k1ch/going/calc"
	"os"
	"reflect"
)

func main() {

	printSome("", "%q", "Hallo, world!")
	printSome("stringutil.Reverse(\"\\n!aloh\") =", "%s", stringutil.Reverse("\n!aloh"))
	printSome("calc.Add(1,2) =", "%d", calc.Add(1, 2))
	printSome("1", "%q", "1")

	p1 := Person{"Alina", 1}
	printSome("p1", "%q", p1)

	var p2 Person
	printSome("p2", "%q", p2)

	p3 := new(Person)
	printSome("p3", "%q", p3)
	p3.name = "Alevtina"
	printSome("p3 again", "%q", p3)

	p4 := &Person{name: "Serega", age: 28}
	printSome("p4", "%q", p4)
	printSome("*p4", "%q", *p4)

	//see print.go: 985
	arrP := []Person{{"Sub-Zero", 2001}, {"Mileena", 2002}}
	printSome("arrP", "%q", arrP)

	i := 2001
	printSome("i:", "%q", i)

	arrP2 := [...]Person{{"Griezzmann", 24}, {"Toldo", 45}}
	printSome("arrP2", "%q", arrP2)

	m := map[int]string{1: "1", 2: "2"}
	printSome("m", "%q", m)

	m2 := map[string]Person{"good": {"Buffon", 46}, "bad": {"Materazzi", 99}}
	printSome("m2", "%q", m2)

	made := make([]int, 10, 20)
	printSome("make", "%q", made)

	arr3 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSome("array formatted as Runes", "%q", arr3)
	printSome("array formatted as strings", "%s", arr3)
	printSome("array formatted as ints", "%d", arr3)
	printSome("array formatted as base 2 (see fmt/doc.go:55)", "%b", arr3)

	slice := arr3[1:2]
	printSome("slice \"1:\"", "%d", slice)

	printSome("reflect.TypeOf(arr3)", "%q", reflect.TypeOf(arr3))
	printSome("reflect.TypeOf(slice)", "%q", reflect.TypeOf(slice))

	a,b,c := multipleReturn()
	printSome("multipleReturn", "%q", a, b, c)
	printSome("multipleReturn, a", "%d", a)
	printSome("multipleReturn, b", "%s", b)
	printSome("multipleReturn, c", "%s", c)
	printSome("reflect.TypeOf(c)", "%q", reflect.TypeOf(a))
	printSome("reflect.TypeOf(c)", "%q", reflect.TypeOf(b))
	printSome("reflect.TypeOf(c)", "%q", reflect.TypeOf(c))

}

var count int = 0
var printer = color.New(color.BgHiCyan, color.FgBlack).PrintfFunc()

func printSome(pad string, f string, s ...interface{}) {
	if isatty.IsTerminal(os.Stdout.Fd()) && count%2 == 0 {
		printer("%-47s", pad)
		printer(":\t")
		printer(f, s)
	} else {
		fmt.Printf("%-47s", pad)
		fmt.Print(":\t")
		fmt.Printf(f, s)
	}
	count++

	fmt.Println()
}

type Person struct {
	name string
	age  int
}

func multipleReturn()(int, string, float32) {
	return 1, "xx", 6.77
}
