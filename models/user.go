package models

import (
	"errors"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/adapter/orm"
	"kube_web/utils"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Username string `json:"username"`
	UserID int `json:"user_id"`
	Token string `json:"token"`
}

type User struct {
	Id int
	Username string
	Salt string
	Password string
}

func DoLogin(lr *LoginRequest) (*LoginResponse, int, error){
	// get username and password
	username := lr.Username
	password := lr.Password

	// validate user name and password is they are empty
	if len(username) == 0 || len(password) == 0 {
		return nil, http.StatusBadRequest,errors.New("error: username or password is empty")
	}

	o := orm.NewOrm()

	// check if the username exists
	user := &User{Username: username}
	err := o.Read(user,"Username")
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("error: username doesn't exist")
	}

	// generate the password hash
	hash, err := utils.GeneratePassHash(password,user.Salt)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	if hash != user.Password {
		return nil, http.StatusBadRequest,errors.New("error: password is error")
	}

	// generate token
	tokenString, err := utils.GenerateToken(lr, user.Id, 0)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return &LoginResponse{
		Username: user.Username,
		UserID: user.Id,
		Token: tokenString,
	},http.StatusOK,nil
}

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type CreateResponse struct {
	UserID  int `json:"user_id"`
	Username string `json:"username"`
}

func DoCreateUser(cr *CreateRequest)(*CreateResponse,int,error){
	o := orm.NewOrm()

	// check if username exists
	userNameCheck := User{Username: cr.Username}
	err := o.Read(&userNameCheck,"Username")
	if err == nil {
		return nil, http.StatusBadRequest, errors.New("username has already existed")
	}

	//generate salt
	saltKey, err := utils.GenerateSalt()
	if err != nil {
		logs.Info(err.Error())
		return nil, http.StatusBadRequest, err
	}

	// generate password hash
	hash, err := utils.GeneratePassHash(cr.Password,saltKey)
	if err != nil {
		logs.Info(err.Error())
		return nil, http.StatusBadRequest,err
	}

	// create user
	user := User{}
	user.Username = cr.Username
	user.Password = hash
	user.Salt = saltKey

	_, err = o.Insert(&user)
	if err != nil {
		logs.Info(err.Error())
		return nil, http.StatusBadRequest,err
	}

	return &CreateResponse{
		UserID:user.Id,
		Username: user.Username,
	}, http.StatusOK,nil
}