package mysql

import (
	"UserApi/gen/user"
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"type:varchar(100);unique_index;unique" json:"email"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Surname  string `gorm:"type:varchar(100)" json:"surname"`
	Phone    string `gorm:"type:varchar(100)" json:"phone"`
	Address  string `gorm:"type:varchar(200)" json:"address"`
}

func CreateUser(ctx context.Context, db *gorm.DB, user *User) (*gorm.DB, error) {
	result := db.WithContext(ctx).Create(user)
	return result, result.Error
}

func GetUserByID(ctx context.Context, db *gorm.DB, id uint) (*User, error) {
	var user User
	result := db.WithContext(ctx).First(&user, id)
	user.Password = "sensitive"
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UpdateUser(ctx context.Context, db *gorm.DB, p *user.UpdatePayload) (userStr string, err error) {
	var user User
	userID, ok := ctx.Value("userID").(float64)
	if !ok {
		return "", errors.New("user ID not found in the context")
	}
	result := db.First(&user, uint(userID))
	user.Password = "sensitive"
	if result.Error != nil {
		return "", result.Error
	}
	bsonPayload, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(bsonPayload, &user); err != nil {
		return "", err
	}
	result = db.Save(&user)
	if result.Error != nil {
		return "", result.Error
	}
	resBson, _ := json.Marshal(&user)
	userStr = string(resBson)

	return userStr, nil
}

func DeleteUser(ctx context.Context, db *gorm.DB, id *uint) error {
	result := db.WithContext(ctx).Delete(&User{}, id)
	return result.Error
}

func GetUserByEmail(ctx context.Context, db *gorm.DB, email string) (*User, error) {
	var user User
	result := db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func MakeNewUserObj(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}
