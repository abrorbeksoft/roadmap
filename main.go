package main

import (
	"awesomeProject/database"
	"awesomeProject/functions"
	"awesomeProject/models"
	"bufio"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"strings"
)

func main()  {
	var myUser models.User
	var users []models.User

	db:=database.Connect()

	rows,err:=db.Query("select * from users")

	if err != nil {
		panic(err)
	}

	for rows.Next(){
		p := models.User{}
		err := rows.Scan(&p.Id, &p.Name, &p.Surname)
		if err != nil{
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}

	for _,u:=range users{
		fmt.Println(u.Id,u.Name,u.Surname)
	}

	reader := bufio.NewReader(os.Stdin)

	text:=""
	myUser=models.User{
		Id: 0,
		Name: "",
		Surname: "",
	}


	for  {
		if strings.Compare("",text)==0 {
			fmt.Print("Enter:\n 0 to login -> \n 1 to registration -> \n 2 to add new contacts-> \n 3 to see all your contacts ->\n 4 for fibonachi numbers ->\n 5 for FizzBuz numbers ->\n-> ")
			text,_=reader.ReadString('\n')
			text=strings.Replace(text,"\n","",-1)

		}

		// login bolish uchun
			//name va surname
		if strings.Compare("0",text)==0 {
			name,surname:=functions.ReadUser(reader)
			myUser=functions.CheckAuth(name,surname,users)

			if myUser.Id!=0 {
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

			name,surname:=functions.ReadUser(reader)

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
			if myUser.Id==0 {
				name,surname:=functions.ReadUser(reader)
				myUser=functions.CheckAuth(name,surname,users)
			} else {
				name,surname,phone:=functions.ReadContact(reader)

				result,err:=db.Exec("insert into contats (name,surname,phone,user_id) values ($1,$2,$3,$4)",name,surname,phone,myUser.Id)
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
			if myUser.Id==0 {
				name,surname:=functions.ReadUser(reader)
				surname=strings.Replace(surname,"\n","",-1)

				myUser=functions.CheckAuth(name,surname,users)
			} else {
				row,err:=db.Query("select * from contats where user_id=$1",myUser.Id)
				if err != nil {
					panic(err)
				}

				contacts := []models.Contact{}
				for row.Next(){
					p := models.Contact{}
					err := row.Scan(&p.Id, &p.Name, &p.Surname,&p.Phone,&p.User_id)
					if err != nil{
						fmt.Println(err)
						continue
					}
					contacts = append(contacts, p)
				}
				fmt.Println(contacts)

				for _,u:=range contacts{
					fmt.Println(u.Id,u.Name,u.Surname,u.Phone)
				}
			}

			text=""
		}

		if strings.Compare("4",text)==0 {
			fmt.Print("Enter number ->")

			number,_:=reader.ReadString('\n')
			number=strings.Replace(number,"\n","",-1)

			if temp,err:=strconv.Atoi(number); err == nil {
				functions.Fibonachi(temp)
			}
			text=""
		}

		if strings.Compare("5",text)==0 {
			fmt.Print("Enter number ->")

			number,_:=reader.ReadString('\n')
			number=strings.Replace(number,"\n","",-1)

			if temp,err:=strconv.Atoi(number); err == nil {
				functions.FizBuzz(temp)
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
