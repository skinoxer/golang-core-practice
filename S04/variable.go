package main

import (
	"fmt"
	"strings"
)

// import (
// 	"fmt"
// 	"strings"
// )

// import "fmt"
//var cz = "global"

func main() {
	//*********** Vriable **************
	//method one
	//var name string ='hello'
	//var i int = 5
	//method two
	//var s1 int //default value is 0 -for string ='' boolean->false
	//var s2 string
	//method three
	//var s5=3 //without data type
	// var s1="iran"
	//method foure
	// var i1:=12
	// var i12:="tehran"
	//method five
	// var f1, f2, s5 = 99, 22, 11
	// var(i7=12,f5=22,s8='reza')
	//*********** Constants **************
	// const pi float64 = 3.14
	// const (
	// 	name   = "iran"
	// 	number = 22
	// 	city   = "ahwaz"
	// )
	// fmt.Printf("%f\n", pi)
	// fmt.Printf("name :%s number:%d city:%s", name, number, city)
	// %t type
	// const baseUrlGoogle = "https://www.google.com"
	// const googleMapUrl = "/maps"
	// println(baseUrlGoogle, googleMapUrl)
	//***********Scope Constants And Variable **************
	//var x = 12
	//println(x)  //12
	//println(ax) //ERROR
	//{
	//	println(x) //12

	//}
	//{
	//	x := 33 //33
	//	ax := 3322
	//	println(x)
	//	println(cz)
	//	}
	//	println(cz)
	//*********** String **************
	// contains - count - cut - join and split - repeat - hasprifix and hassuffix
	// replace and replaceall - compare and equalfold - index-tolowercase toupper and title - trim
	// name := "alireza"
	// age := 27
	// id := 199999
	// score := 7.5
	// print("my name is ", name, "my age is", age, id, score)
	// println()
	// fmt.Printf()//%T(type) and %b (binary)
	//mystring := "this is golang language go"

	//fmt.Println(strings.Contains(mystring, "go"))      //true

	//fmt.Println(strings.ContainsAny(mystring, "1te1")) //true this is bad for security i think

	//fmt.Println(strings.Count(mystring, "go"))         //2 golang =>go

	//fmt.Println(strings.Cut(mystring, "go"))           //this is  lang language go true

	//fmt.Println(strings.Cut(mystring, "go1"))          //this is golang language go  false

	//count number of words in sentens with join and split

	//mystringArray := strings.Split(mystring, " ")      //[this is golang language go]

	//fmt.Println(mystringArray)                         //[this is golang language go]

	//fmt.Println(len(mystringArray))                    //5

	//mystringArray2 := strings.Join(mystringArray, " ") //from arr to string with joda konnde

	//fmt.Println(mystringArray2)                        //this is golang language go

	//println(strings.Repeat("iran ", 10))               //iran iran iran iran iran iran iran iran iran iran

	//* tavajo avalin go ke did ro taghir medi hata golan avalsh go dare

	//fmt.Println(strings.Replace(mystring, "go", "golang", 1)) //this is golanglang language golang !
	// println(strings.Compare("golang", "golang")) //0 true

	// println(strings.Compare("Golang", "golang")) //-1 false return int
	// println(strings.Compare("Golang", "GOlang")) //1
	//println("golang" > "Golang")//true fasle from compare acski code
	//println(strings.EqualFold("golang", "Golang")) //t   not case sensetive
	//println(strings.EqualFold("golang", "Go"))     // f
	// prifix yani ba in prifix tarif shode shoro meshe ya na
	//println(strings.HasPrefix("Iran", "Ir")) //t
	//println(strings.HasPrefix("Iran", "IR")) //f
	// suffix az akhare ba aval miad
	//println(strings.HasSuffix("Iran", "n")) //t
	//println(strings.HasSuffix("Iran", "n")) //t
	// println(strings.Index("iran", "r")) //find index or the r in word iran
	// println(strings.ToLower("IRAN"))
	// println(strings.Trim("IRAN          ", " "))
	//println(strings.Trim("IRANzzzz!!", "!"))
	// strings.ToUpper(mystring)
	// *** practice S04 ***
	name := "  HeLLO GOLang, Golang IS gReat!  "
	fmt.Println(strings.Contains(name, "e"))
	fmt.Println(strings.ContainsAny(name, "1Gl"))
	fmt.Println(strings.ContainsAny(name, "1Gl"))

}
