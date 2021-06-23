package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type user struct {
	id int
	name string
	surname string

}

func main()  {

	conectStr:="user=root password=password dbname=roadmap sslmode=disable"
	db,err:=sql.Open("postgres",conectStr)

	if err!=nil{
		panic(err)
	}


	rows,err:=db.Query("select * from users")

	if err != nil {
		panic(err)
	}
	users := []user{}
	for rows.Next(){
		p := user{}
		err := rows.Scan(&p.id, &p.name, &p.surname)
		if err != nil{
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}

	for _,u:=range users{
		fmt.Println(u.id,u.name,u.surname)
	}

	reader := bufio.NewReader(os.Stdin)
	for  {
		fmt.Println("->")
		text,_:=reader.ReadString('\n')

		fmt.Println(text)
	}


}