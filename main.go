package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strings"
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

	//for _,u:=range users{
	//	fmt.Println(u.id,u.name,u.surname)
	//}

	reader := bufio.NewReader(os.Stdin)

	bolim:=""
	text:=""

	for  {
		if strings.Compare("",text)==0 {
			fmt.Print("Enter:\n 1 to registration -> \n 2 to add new contacts-> \n 3 to see all your contacts ->\n-> ")
			text,_=reader.ReadString('\n')
			text=strings.Replace(text,"\n","",-1)

		}

		// 1 royxatdan o'tish
			// name va surnameini kiritishni soraymiz
		if strings.Compare("1",text)==0{
			fmt.Print("Enter name ->")

			name,_:=reader.ReadString('\n')
			name=strings.Replace(name,"\n","",-1)

			fmt.Print("Enter surname ->")

			surname,_:=reader.ReadString('\n')
			surname=strings.Replace(surname,"\n","",-1)

			result,err:=db.Exec("insert into users (name,surname) values ($1,$2)",name,surname)
			if err!=nil {
				panic(err)
			}

			fmt.Println(result.LastInsertId())
			text=""
		}


		// 2 contact qoshish
			// royxatdan otgan bolishi kerak
			// id sidan foydalanamiz
		// 3 contactlarni korish
			// name va surname ni kiritishni soraymiz
		if strings.Compare("3",bolim)==0{
			fmt.Print("Enter name ->")

			name,_:=reader.ReadString('\n')
			name=strings.Replace(name,"\n","",-1)

			fmt.Print("Enter surname ->")

			surname,_:=reader.ReadString('\n')
			surname=strings.Replace(surname,"\n","",-1)

			result,err:=db.Exec("insert into users set name=$name,surname=$surname",name,surname)
			if err!=nil {
				panic(err)
			}

			fmt.Println(result.RowsAffected())
			bolim=""
		}


		if strings.Compare("hi",text)==0 {
			fmt.Println("Hello to you too!")
		}

		if strings.Compare("q",text)==0{
			break
		}
	}


}