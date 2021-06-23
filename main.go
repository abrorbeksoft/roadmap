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

type Contact struct {
	id int
	name string
	surname string
	phone string
	user_id int
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

	text:=""
	myUser:=user{
		id: 0,
		name: "",
		surname: "",
	}


	for  {
		if strings.Compare("",text)==0 {
			fmt.Print("Enter:\n 0 to login -> \n 1 to registration -> \n 2 to add new contacts-> \n 3 to see all your contacts ->\n-> ")
			text,_=reader.ReadString('\n')
			text=strings.Replace(text,"\n","",-1)

		}

		// login bolish uchun
			//name va surname
		if strings.Compare("0",text)==0 {
			fmt.Print("Enter name ->")

			name,_:=reader.ReadString('\n')
			name=strings.Replace(name,"\n","",-1)

			fmt.Print("Enter surname ->")

			surname,_:=reader.ReadString('\n')
			surname=strings.Replace(surname,"\n","",-1)

			for _,value:=range users{
				if value.name==name && value.surname==surname {
					myUser.id=value.id
					myUser.name=value.name
					myUser.surname=value.surname
				}
			}
			if myUser.id!=0 {
				fmt.Println("Login success!")
				text=""
			} else {
				fmt.Println("User not found")
				text=""
			}

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
		if strings.Compare("2",text)==0{
			if myUser.id==0 {
				fmt.Print("Enter name ->")

				name,_:=reader.ReadString('\n')
				name=strings.Replace(name,"\n","",-1)

				fmt.Print("Enter surname ->")

				surname,_:=reader.ReadString('\n')
				surname=strings.Replace(surname,"\n","",-1)

				for _,value:=range users{
					if value.name==name && value.surname==surname {
						myUser.id=value.id
						myUser.name=value.name
						myUser.surname=value.surname
					}
				}
			}

			if myUser.id!=0{
				fmt.Print("Enter contact name ->")

				name,_:=reader.ReadString('\n')
				name=strings.Replace(name,"\n","",-1)

				fmt.Print("Enter contact surname ->")

				surname,_:=reader.ReadString('\n')
				surname=strings.Replace(surname,"\n","",-1)

				fmt.Print("Enter contact phone ->")

				phone,_:=reader.ReadString('\n')
				phone=strings.Replace(surname,"\n","",-1)

				result,err:=db.Exec("insert into contats (name,surname,phone,user_id) values ($1,$2,$3,$4)",name,surname,phone,myUser.id)
				if err!=nil {
					panic(err)
				}

				fmt.Println(result.LastInsertId())
				text=""
			}
		}


		// 3 contactlarni korish
			// name va surname ni kiritishni soraymiz
		if strings.Compare("3",text)==0{
			if myUser.id==0 {
				fmt.Print("Enter name ->")

				name,_:=reader.ReadString('\n')
				name=strings.Replace(name,"\n","",-1)

				fmt.Print("Enter surname ->")

				surname,_:=reader.ReadString('\n')
				surname=strings.Replace(surname,"\n","",-1)

				for _,value:=range users{
					if value.name==name && value.surname==surname {
						myUser.id=value.id
						myUser.name=value.name
						myUser.surname=value.surname
					}
				}
			}
			if myUser.id !=0 {

				row,err:=db.Query("select * from contats where user_id=$1",myUser.id)
				if err != nil {
					panic(err)
				}

				contacts := []Contact{}
				for row.Next(){
					p := Contact{}
					err := row.Scan(&p.id, &p.name, &p.surname,&p.phone,&p.user_id)
					if err != nil{
						fmt.Println(err)
						continue
					}
					contacts = append(contacts, p)
				}
				fmt.Println(contacts)

				for _,u:=range contacts{
					fmt.Println(u.id,u.name,u.surname,u.phone)
				}
			}

			text=""
		}


		if strings.Compare("hi",text)==0 {
			fmt.Println("Hello to you too!")
		}

		if strings.Compare("q",text)==0{
			break
		}
	}


}
