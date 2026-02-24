package repository

import (
    "sellcard/server/internal/model"
)

func GetUserByUsername(username string) (*model.User, error) {
    var u model.User
    if err := db.Where("username = ?", username).First(&u).Error; err != nil {
        return nil, err
    }
    return &u, nil
}

func CreateUser(u *model.User) error {
    return db.Create(u).Error
}
