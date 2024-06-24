
package models

import (
    "time"

    "gorm.io/gorm"
)

type Contract struct {
    ID        uint           `gorm:"primaryKey"`
    StaffID   uint           `gorm:"index"`
    Title     string
    Content   string
    Status    string
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
