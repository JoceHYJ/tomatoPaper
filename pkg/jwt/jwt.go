package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"tomatoPaper/api/entity"
)

const (
	ContextKeyUserObj    = "authedUserObj"
	ContextKeyStudentObj = "authedStudentObj"
	ContextKeyTeacherObj = "authedTeacherObj"
	ContextKeyAdminObj   = "authedAdminObj"
)

// studentStdClaims 自定义 JWT 载荷
type studentStdClaims struct {
	entity.JwtStudentDto
	jwt.StandardClaims
}

// userStdClaims 自定义 JWT 载荷
type userStdClaims struct {
	entity.JwtUser
	jwt.StandardClaims
}

// TokenExpireDuration 设置 Token 的过期时间
const TokenExpireDuration = time.Hour * 24

// Secret token 密钥
var Secret = []byte("tomato-paper")
var (
	ErrAbsent  = "token absent"
	ErrInvalid = "token invalid"
)

// GenerateTokenUser 根据用户信息生成 token
func GenerateTokenUser(user entity.Users) (string, error) {
	var jwtUser = entity.JwtUser{
		UserID: user.UserID,
		//Username: user.Username,
		Password: user.Password,
	}
	c := userStdClaims{
		jwtUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "tomatoPaper",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// GenerateTokenStudent 根据学生信息生成 token
func GenerateTokenStudent(student entity.Students) (string, error) {
	var jwtStudentDto = entity.JwtStudentDto{
		StudentID: student.StudentID,
		//Username: user.Username,
		Password: student.Password,
	}
	c := studentStdClaims{
		jwtStudentDto,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "tomatoPaper",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// ValidateToken 解析 JWT 验证 token 是否有效
func ValidateToken(tokenString string) (*entity.JwtUser, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return Secret, nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.JwtUser, nil
}

// ValidateTokenStudent 解析 JWT 验证 token 是否有效
func ValidateTokenStudent(tokenString string) (*entity.JwtStudentDto, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return Secret, nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	claims := studentStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.JwtStudentDto, nil
}
