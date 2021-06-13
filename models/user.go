package models

import (
	"context"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type SysUser struct {
	Id       int64  `json:"id"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Enabled  int8   `json:"enabled"`
	Password string `json:"password"`
	Username string `json:"username" valid:"Required;"`
	//DeptId int32
	Phone string `json:"phone"`
	//JobId int32
	NickName    string     `json:"nickName"`
	Sex         string     `json:"sex"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(SysUser))
}

//根据用户名返回
func GetUserByUsername(name string) (v *SysUser, err error) {
	o := orm.NewOrm()
	user := &SysUser{}
	err = o.QueryTable(new(SysUser)).Filter("username", name).RelatedSel().One(user)
	if err == nil {
		return user, nil
	}
	return nil, err
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int64) (v *SysUser, err error) {
	//var userlist []User
	o := orm.NewOrm()
	v = &SysUser{Id: id}

	err = o.QueryTable(new(SysUser)).Filter("Id", id).RelatedSel().One(v)
	if err == nil {
		return v, nil
	}

	return nil, err
}

// get all
func GetAllUser(tableName string,page int,size int, condition string) (int, []SysUser) {
	users := []SysUser{}
	total, _, rs := GetPagesInfo(tableName, page, size, condition)
	rs.QueryRows(&users)

	return total, users
}

func UpdateCurrentUser(m *SysUser) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(m)
	return
}

func AddUser(m *SysUser) (err error) {
	o := orm.NewOrm()
	//transaction
	err = o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) (error) {
		// data
		_, e := txOrm.Insert(m)
		return e
	})
	return err
}

func UpdateByUser(m *SysUser) (err error) {
	o := orm.NewOrm()
	//transaction
	err = o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		// data
		_, e := txOrm.Update(m)
		if e != nil {
			return e
		}
		return e
	})
	logs.Error(err)
	return
}

func DelByUser(id int64) (err error) {
	o := orm.NewOrm()
	err = o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		_, e := txOrm.Raw("UPDATE sys_user SET is_del = ? WHERE id =", 1, id).Exec()
		return e
	})
	return
}
