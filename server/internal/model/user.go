package model

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    Username     string    `gorm:"uniqueIndex;size:64" json:"username"`
    PasswordHash string    `gorm:"size:255" json:"-"`
    Role         string    `gorm:"size:32;default:'user'" json:"role"`
    CreatedAt    time.Time `json:"created_at"`
}
