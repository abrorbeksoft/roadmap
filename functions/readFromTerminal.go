package functions

import (
	"bufio"
	"fmt"
	"strings"
)

func ReadUser(reader *bufio.Reader) (Name ,Surname string) {
	fmt.Print("Enter name ->")

	name,_:=reader.ReadString('\n')
	name=strings.Replace(name,"\n","",-1)

	fmt.Print("Enter surname ->")

	surname,_:=reader.ReadString('\n')
	surname=strings.Replace(surname,"\n","",-1)

	return name,surname
}

func ReadContact(reader *bufio.Reader) (Name ,Phone,Surname string) {
	fmt.Print("Enter contact name ->")

	name,_:=reader.ReadString('\n')
	name=strings.Replace(name,"\n","",-1)

	fmt.Print("Enter contact surname ->")

	surname,_:=reader.ReadString('\n')
	surname=strings.Replace(surname,"\n","",-1)

	fmt.Print("Enter contact phone ->")

	phone,_:=reader.ReadString('\n')
	phone=strings.Replace(surname,"\n","",-1)

	return name,surname,phone
}