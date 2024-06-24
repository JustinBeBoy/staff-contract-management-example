package tasks

import (
	"log"
	"staff-contract-management/models"

	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

func InitNATS() *nats.Conn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	return nc
}

func SubscribeTasks(nc *nats.Conn, db *gorm.DB) {
	nc.QueueSubscribe("delete.contract", "worker", func(msg *nats.Msg) {
		contractID := string(msg.Data)
		DeleteContractBackground(db, contractID)
	})

	nc.QueueSubscribe("delete.staff", "worker", func(msg *nats.Msg) {
		staffID := string(msg.Data)
		DeleteStaffBackground(db, staffID)
	})
}

func DeleteContractBackground(db *gorm.DB, contractID string) {
	var contract models.Contract
	if err := db.Where("id = ?", contractID).First(&contract).Error; err == nil {
		contract.Status = "ARCHIVED"
		db.Save(&contract)
	}
}

func DeleteStaffBackground(db *gorm.DB, staffID string) {
	var staff models.Staff
	if err := db.Where("id = ?", staffID).First(&staff).Error; err == nil {
		staff.Role = "IN-ACTIVE"
		db.Save(&staff)
	}
}
