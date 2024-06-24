
package models

import (
    "time"

    "gorm.io/gorm"
)

type Staff struct {
    ID        uint           `gorm:"primaryKey"`
    Username  string         `gorm:"unique"`
    Password  string
    Role      string         // "admin" or "staff"
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
    Contracts []Contract
}
