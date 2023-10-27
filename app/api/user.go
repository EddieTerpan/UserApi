package api

import (
	gormMysqlUser "UserApi/app/gormMysql"
	"UserApi/gen/user"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"
)

type UserService struct {
	logger *log.Logger
	gormDB *gorm.DB
}

// NewUser returns the user service implementation.
func NewUser(logger *log.Logger, gormDB *gorm.DB) user.Service {
	return &UserService{logger, gormDB}
}

// JWTAuth implements the authorization logic for service "user" for the "jwt"
func (s *UserService) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	if token == "" {
		return ctx, goa.PermanentError("unauthorized", "invalid token")
	}
	err := scheme.Validate(scheme.Scopes)
	if err != nil {
		return ctx, goa.PermanentError("unauthorized", err.Error())
	}
	// Parse and validate the JWT token
	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return ctx, goa.PermanentError("unauthorized", "invalid token")
		}
		return []byte(os.Getenv("GOA_JWT_SECRET")), nil
	})
	if err != nil {
		return ctx, goa.PermanentError("token parse error: ", err.Error())
	}
	if !tokenParsed.Valid {
		return ctx, goa.PermanentError("unauthorized", "invalid token")
	}
	// The token is valid, so you can extract claims from the token if needed.
	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	_, okExpClaim := claims["exp"]
	_, okUserIdClaim := claims["userId"]

	if !(ok && okExpClaim && okUserIdClaim) {
		return ctx, goa.PermanentError("unauthorized", "empty token claims")
	}
	// Store the user ID in the context
	ctx = context.WithValue(ctx, "userID", claims["userId"])
	timeExp, expParsed := time.Parse(time.RFC3339, claims["exp"].(string))
	if expParsed == nil && timeExp.Before(time.Now()) {
		return ctx, goa.PermanentError("unauthorized", "expired token")
	}
	if int(claims["userId"].(float64)) <= 0 {
		return ctx, goa.PermanentError("unauthorized", "invalid user")
	}
	if tokenParsed.Valid {
		return ctx, nil
	}

	return ctx, goa.PermanentError("unauthorized", "invalid token")
}

// Create implements create.
func (s *UserService) Create(ctx context.Context, p *user.CreatePayload) (res string, err error) {
	s.logger.Print("user.create")
	hash, err := hashPassword(*p.Password)
	if err != nil {
		return "", err
	}
	result, err := gormMysqlUser.CreateUser(ctx, s.gormDB, gormMysqlUser.MakeNewUserObj(*p.Email, hash))
	if err != nil {
		return "", err
	}

	res = "RowsAffected: " + strconv.FormatInt(result.RowsAffected, 10)
	return
}

// Read implements read.
func (s *UserService) Read(ctx context.Context, p *user.ReadPayload) (res string, err error) {
	s.logger.Print("user.read")
	// Retrieve the user ID from the context
	userID, ok := ctx.Value("userID").(float64)
	if !ok {
		return "", errors.New("user ID not found in the context")
	}
	foundedUser, err := gormMysqlUser.GetUserByID(ctx, s.gormDB, uint(userID))
	if err != nil {
		return "", err
	}
	resBson, _ := json.Marshal(foundedUser)
	res = string(resBson)

	return
}

// Update implements update.
func (s *UserService) Update(ctx context.Context, p *user.UpdatePayload) (res string, err error) {
	s.logger.Print("user.update")
	res, err = gormMysqlUser.UpdateUser(ctx, s.gormDB, p)
	if err != nil {
		return "", err
	}

	return
}

// Delete implements delete.
func (s *UserService) Delete(ctx context.Context, p *user.DeletePayload) (res string, err error) {
	s.logger.Print("user.delete")
	err = gormMysqlUser.DeleteUser(ctx, s.gormDB, p.ID)

	return
}

// Token implements token.
// https://self-issued.info/docs/draft-ietf-oauth-json-web-token.html
func (s *UserService) Token(ctx context.Context, p *user.TokenPayload) (res string, err error) {
	s.logger.Print("user.token")
	verifyPwdErr := verifyPassword(*p.Password)
	if verifyPwdErr != nil {
		return "", verifyPwdErr
	}
	userCheck, err := gormMysqlUser.GetUserByEmail(ctx, s.gormDB, *p.Email)
	if err != nil {
		return "", errors.New("email not found")
	}
	if checkPasswordHash(*p.Password, userCheck.Password) {
		return "", errors.New("invalid password")
	}
	// Create the JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userCheck.ID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7)
	// Sign the token
	tokenString, err := token.SignedString([]byte(os.Getenv("GOA_JWT_SECRET"))) // Replace with your secret key
	if err != nil {
		return "", errors.New("signedString error")
	}

	return "Authorization: Bearer " + tokenString, nil
}

// standard property of language better than slow regexp
func verifyPassword(s string) error {
	letters := 0
	sevenOrMore, number, upper, special := false, false, false, false

	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c):
			letters++
		}
	}

	sevenOrMore = letters >= 7
	if !(sevenOrMore && number && upper && special) {
		return fmt.Errorf(
			"invalid password format, sevenOrMore: %t, number: %t, upper: %t, special: %t",
			sevenOrMore,
			number,
			upper,
			special,
		)
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}
