package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string
	Password string // This will store the hashed password.
	Address  string
	Job      string
	Reason   string
}

var Users = make(map[string]*User)

func EncryptPassword(inputtedPassword []byte) string {
	hash, err := bcrypt.GenerateFromPassword(inputtedPassword, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func PasswordCheck(storedPassword, inputtedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputtedPassword))
	return err == nil
}

func AddUser(email, password, address, job, reason string) {
	encryptedPassword := EncryptPassword([]byte(password))
	Users[email] = &User{
		Email:    email,
		Password: encryptedPassword,
		Address:  address,
		Job:      job,
		Reason:   reason,
	}
}

func GetUsersEmail() []string {
	emailList := make([]string, 0, len(Users))
	for email := range Users {
		emailList = append(emailList, email)
	}

	return emailList
}

func GetUsersData() []*User {
	userDataList := make([]*User, 0, len(Users))

	for _, user := range Users {
		userWithoutPassword := &User{
			Email:   user.Email,
			Address: user.Address,
			Job:     user.Job,
			Reason:  user.Reason,
		}
		userDataList = append(userDataList, userWithoutPassword)
	}

	return userDataList
}

func GetUserDetails(email string) *User {
	if user, exist := Users[email]; exist {
		return &User{
			Email:   user.Email,
			Address: user.Address,
			Job:     user.Job,
			Reason:  user.Reason,
		}
	}

	return nil
}
