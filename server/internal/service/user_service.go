package service

import (
    "errors"
    "sellcard/server/internal/model"
    "sellcard/server/internal/repository"

    "golang.org/x/crypto/bcrypt"
)

func Authenticate(username, password string) (*model.User, error) {
    u, err := repository.GetUserByUsername(username)
    if err != nil {
        return nil, err
    }
    if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
        return nil, errors.New("invalid credentials")
    }
    return u, nil
}

func CreateUser(username, password, role string) (*model.User, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    u := &model.User{Username: username, PasswordHash: string(hash), Role: role}
    if err := repository.CreateUser(u); err != nil {
        return nil, err
    }
    return u, nil
}
