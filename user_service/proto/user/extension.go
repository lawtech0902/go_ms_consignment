package user

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid_ := uuid.NewV4()
	// if err != nil {
	// 	return err
	// }
	return scope.SetColumn("Id", uuid_.String())
}
