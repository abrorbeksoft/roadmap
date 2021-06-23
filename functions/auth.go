package functions

import "awesomeProject/models"

func CheckAuth(name string,surname string,users []models.User) models.User {
	var myUser models.User

	for _,value:=range users{
		if value.Name==name && value.Surname==surname {
			myUser.Id=value.Id
			myUser.Name=value.Name
			myUser.Surname=value.Surname
		}
	}

	return myUser
}
