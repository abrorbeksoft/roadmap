package main

func fibonachi(number int) *[]int {
	i:=0
	var responce []int

	for i<=number {
		if i==0 {
			responce[i]=0
		}else if i==1 {
			responce[i]=1
		}else {
			responce[i]=responce[i-1]+responce[i-2]
		}
		i++
	}

	number1:=&responce

	return number1
}