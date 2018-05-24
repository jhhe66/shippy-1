package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	uuid2 "github.com/satori/go.uuid"
	"github.com/labstack/gommon/log"
)

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid2.NewV4()
	if err != nil {
		log.Fatalf("created uuid error: %v\n", err)
	}
	return scope.SetColumn("Id", uuid.String())
}
