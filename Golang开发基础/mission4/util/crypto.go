package util

import (
  "crypto/sha3"
  "encoding/hex"
  "errors"
  "myBlog/model"
  "time"

  "github.com/golang-jwt/jwt/v5"
)

const DEFAULT_SALT = "游龙当归海 海不迎我自来也"

/*
*
使用sha256加密字符串，返回加密后的十六进制字符串
*/
func CryptoStrWithSalt(text string, salt string) string {
  sum := sha3.Sum256([]byte(text + salt))
  return hex.EncodeToString(sum[:])
}

func CryptoStrWithDefaulSalt(text string) string {
  return CryptoStrWithSalt(text, DEFAULT_SALT)
}

const myJwtSigningKey = "my_secret_key"

type UserClaims struct {
  ID       uint
  Username string
  jwt.RegisteredClaims
}

// 生成jwt token
func GenJwtToken(user *model.User) (string, error) {
  userClaim := UserClaims{
    ID:       user.ID,
    Username: user.Username,
    RegisteredClaims: jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
      Issuer:    "zhp",
    },
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
  return token.SignedString([]byte(myJwtSigningKey))
}

// 解析JWT token
func ParseJwtToken(tokenString string) (*UserClaims, error) {
  claims := &UserClaims{}
  token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
    return []byte(myJwtSigningKey), nil
  })

  if err != nil {
    return nil, err
  }
  if token.Valid {
    return claims, nil
  }

  return nil, errors.New("token无效")
}
