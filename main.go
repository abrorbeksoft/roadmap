package main

import (
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host= "localhost"
	port =5432
	user="user"
	password="password"
	dbname="belt_master"
	)

func main()  {

	fmt.Println("Hello world")

}