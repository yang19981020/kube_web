package jwt

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
	models "kube_web/models"
	"strconv"
	"time"
)

type userStdClaims struct {
	models.SysUser
	//*models.User
	jwt.StandardClaims
}

var (
	SignedKey,_ = beego.AppConfig.String("jwt_token") // 签名的key
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
	ErrExpired = "token expired" // 令牌过期
	ErrOther   = "other error"   // 其他错误
)

func GenerateToken(m *models.SysUser,d time.Duration) (string,error) {
	m.Password = ""
	expireTime := time.Now().Add(d)  // 过期时间
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Id: strconv.FormatInt(m.Id,10),
		Issuer:    "kube_web",
	}

	var user = models.SysUser{
		Id:        m.Id,
		Avatar:    m.Avatar,
		Email:     m.Email,
		Username:  m.Username,
		Phone: m.Phone,
		NickName:  m.NickName,
		Sex:       m.Sex,
	}

	uClaims := userStdClaims{
		StandardClaims: stdClaims,
		SysUser: user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,uClaims)
	tokenString,err := token.SignedString([]byte(SignedKey))
	if err != nil {
		logs.Info("token 生成失败", err)
	}
	return tokenString,err
}

func ValidateToken(tokenStr string) (*models.SysUser,error)  {
	if tokenStr == "" {
		return nil, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(SignedKey), nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SignedKey), nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.SysUser, err
}

func GetAdminUserId(c *context.BeegoInput) (int64, error) {
	// todo
	return 1, errors.New("")
}