package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

//用户表
type User struct {
	Id            int64
	Username      string `json:"name",orm:"unique"` //这个拼音的简写
	Nickname      string //中文名，注意这里，很多都要查询中文名才行`orm:"unique;size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(2)"`
	Password      string
	Repassword    string `orm:"-" form:"Repassword" valid:"Required" form:"-"`
	Email         string `orm:"size(32)" form:"Email" valid:"Email"`
	Department    string //分院
	Secoffice     string //科室,这里应该用科室id，才能保证即时重名也不怕。否则，查看科室必须要上溯到分院才能避免科室名称重复问题
	Remark        string `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
	Ip            string //ip地址
	Port          string
	Status        int       `orm:"default(1)";form:"Status";valid:"Range(1,2)"`
	Lastlogintime time.Time `orm:"type(datetime);auto_now_add" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
	Updated       time.Time `orm:"type(datetime);auto_now_add" `
	Role          string    `json:"role";orm:"default('4')"` //这个不是角色，这个无意义
	// Roles         []*Role   `orm:"rel(m2m)"`
}

//用户和openid对应表,一个用户对应多个openid
type UserOpenID struct {
	Id     int64
	Uid    int64
	OpenID string
}

//用户和AvatorUrl对应表,一个用户对应多个AvatorUrl
type UserAvatar struct {
	Id         int64
	Uid        int64
	AvatarUrl  string
	Createtime time.Time `orm:"type(datetime);auto_now_add" `
}

//用户和AppreciationUrl对应表,一个用户对应多个AppreciationUrl
type UserAppreciation struct {
	Id              int64
	Uid             int64
	AppreciationUrl string
	Createtime      time.Time `orm:"type(datetime);auto_now_add" `
}

func init() { //
	orm.RegisterModel(new(User), new(UserOpenID), new(UserAvatar), new(UserAppreciation))
}

//这个是使用的，下面那个adduser不知干啥的
func SaveUser(user User) (uid int64, err error) {
	o := orm.NewOrm()
	var user1 User
	//判断是否有重名
	err = o.QueryTable("user").Filter("username", user.Username).One(&user1, "Id")
	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		uid, err = o.Insert(&user)
		if err != nil {
			return uid, err
		}
	}
	return uid, err
}
