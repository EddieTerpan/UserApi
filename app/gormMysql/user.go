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
	Email    string `gorm:"type:varchar(100);unique_index;unique"`
	Password string `gorm:"type:varchar(100)" json:"-"`
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
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UpdateUser(ctx context.Context, db *gorm.DB, p *user.UpdatePayload) (user *User, err error) {
	userID, ok := ctx.Value("userID").(float64)
	if !ok {
		return nil, errors.New("user ID not found in the context")
	}
	result := db.First(&user, uint(userID))
	if result.Error != nil {
		return nil, result.Error
	}
	bsonPayload, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bsonPayload, &user); err != nil {
		return nil, err
	}
	result = db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
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
