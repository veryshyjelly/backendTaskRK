package helpers

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/veryshyvelly/task2-backend/database"
	"github.com/veryshyvelly/task2-backend/models"
)

type SignedDetails struct {
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	Name     string `json:"name"`
	RollNo   string `json:"roll_no"`
	jwt.RegisteredClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(name string, email string, userType string, userId string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:    email,
		UserType: userType,
		Name:     name,
		RollNo:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 1)),
		},
	}
	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24)),
		},
	}

	signedToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
	}

	signedRefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
	}

	return
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userID uint, userType string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if userType == "student" {
		var student models.Student
		database.DB.First(&student, userID)
		student.Token = &signedToken
		student.RefreshToken = &signedRefreshToken
		database.DB.WithContext(ctx).Save(&student)
	} else if userType == "admin" {
		var admin models.Admin
		database.DB.First(&admin, userID)
		admin.Token = &signedToken
		admin.RefreshToken = &signedRefreshToken
		database.DB.WithContext(ctx).Save(&admin)
	}
}

func ValidateToken(signedToken string) (claims *SignedDetails, err error) {
	claims = &SignedDetails{}
	token, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		err = errors.New("invalid token")
		return
	}

	if !token.Valid {
		err = jwt.ErrSignatureInvalid
		return
	}

	if claims.ExpiresAt.Before(time.Now()) {
		err = errors.New("token expired")
		return
	}

	return
}
