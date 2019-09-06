package main

import (
	"fmt"
)

//func main(){
//	//fmt.Println("My favorite number is",rand.Intn(10))
//	fmt.Println("Now you have %g problems. \n",math.Sqrt(7))
//	fmt.Println(math.Pi)
//}

//const pi =3.14
//
//func swap(x,y string)(string,string)  {
//		return y,x
//}
//
//func main(){
//	a, b := swap("hello", "world")
//	fmt.Println(a,b)
//}

func main()  {
	//sum :=0;
	//for i:=0;i<10 ;i++  {
	//	sum += i
	//}
	//fmt.Println(sum)

	//fmt.Println(math.Sqrt(2))
	//
	//fmt.Println("Go runs on")
	//switch os:= runtime.GOOS;
	//os{
	//case "darwin":fmt.Println("os x")
	//case "linux":
	//	fmt.Println("Linux")
	//default:
	//	fmt.Println("%s.\n",os)
	//}
	//
	//type a struct {
	//	x,y int
	//}
	//fmt.Println(a{1,2}.x)
	//
	//var (
	//	v1 =a{1,2}
	//	v2 =a{x:1}
	//	v3 = a{}
	//	p =&a{1,2}
	//)
	//fmt.Println(v1,v2,v3,p)
	//fmt.Println((*p).x)

	ints := make([]int, 0, 5)
	fmt.Println(ints,len(ints), cap(ints))

	//pow:= []int{1, 2, 4, 8, 16, 32, 64, 128}
	//for i,v:=range pow {
	//	fmt.Println("2**%d = %d\n",i,v)
	//}

	pow:= make([]int, 10)
	for i :=range pow {
		pow[i] =1<<uint(i)
	}
	for _,value:=range pow {
		fmt.Println("%d\n",value)
	}
}