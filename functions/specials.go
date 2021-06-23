package functions

import "fmt"

func Fibonachi(size int)  {
	var arry [200]int
	arry[0]=0
	arry[1]=1
	i:=2
	for  i<=size {
		arry[i]=arry[i-1]+arry[i-2]
		i++
	}
	i=0
	for  i<=size {
		fmt.Println(arry[i])
		i++
	}
}

func FizBuzz(number int)  {
	i:=0
	for i<=number {
		if i%3==0 && i%5==0{
			fmt.Println("FizzBuzz")
		}else if i%3==0 {
			fmt.Println("Fizz")
		} else if i%5==0 {
			fmt.Println("Buzz")
		}else {
			fmt.Println(i)
		}
		i++
	}

}
