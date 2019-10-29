package entity

import(
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

type User struct {
    Username string
    Password string
    Email string
    Phone string
}

func CheckRegister(username, password, email, phone string) int {
	fileData := ReadFromFile()

	flagE := 0
	for i:=0; i<len(fileData); i++{
		if fileData[i].Username == username {
			flagE = 1//用户名已存在
		}
	}

	if username=="" {
		return 3 //未输入用户名
	} else if flagE == 1{
		return 2 //用户名已存在
	} else {
		user := User{Username: username, Password: password, Email: email, Phone: phone}
		fileData = append(fileData, user)
		WriteToFile(fileData)
		return 1
	}
} 

func CheckLogin(username, password string) int {
	fileData := ReadFromFile()

	flagA := 0
	for i:=0; i<len(fileData); i++{
		if fileData[i].Username == username {
			flagA = 1//用户名已存在
		}
	}

	flagB := 0
	for j:=0; j<len(fileData); j++{
		if fileData[j].Username == username {
			if fileData[j].Password == password {
				flagB = 1//密码正确
			}
		}
	}


	if flagA == 0 {
		return 2 //用户名不存在
	} else if flagB == 0 {
		return 3 //密码错误
	} else {
		return 1
	}
}

func ReadFromFile() []User {
    jsondata, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		fmt.Println(err)
	}
    var user []User
	json.Unmarshal(jsondata, &user)
    return user
}

func WriteToFile(user []User) {
    os.Truncate("./data/data.json", 0)
    jsondata, err := json.Marshal(user)
    if err != nil {
		fmt.Println(err)
	}
    ioutil.WriteFile("./data/data.json", jsondata, os.ModeAppend)
    os.Chmod("./data/data.json", 0777)
}