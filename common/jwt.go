// 返回token的文件，不用管，要看也行。模板写法来着，至于为什么，我也不是很懂
package common

import (
	"sse_market_admin/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// 采用jwt方式生成token
var jwtKey []byte

type Claims struct {
	UserID int
	jwt.StandardClaims
}

type Claims_admin struct {
	AdminID int
	jwt.StandardClaims
}

func InitJWTkey() {
	jwtKey = []byte(viper.GetString("crypto.jwtKey"))
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserID: user.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "loginTest",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ReleaseToken_admin(admin model.Admin) (string, error) {
	expirationTime := time.Now().Add(1 * 24 * time.Hour)
	claims := &Claims_admin{
		AdminID: admin.AdminID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "loginTest",
			Subject:   "admin token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}

func ParseToken_admin(tokenString string) (*jwt.Token, *Claims_admin, error) {
	claims_admin := &Claims_admin{}
	token, err := jwt.ParseWithClaims(tokenString, claims_admin, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims_admin, err
}
