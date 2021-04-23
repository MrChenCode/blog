package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	
	Name string `json:"name"`
	CreatedBy string `json:"created_by" gorm:"column:created_by"`
	ModifiedBy string `json:"modified_by" gorm:"column:modified_by"`
	State int `json:"state"`
	ModifiedOn int `json:"modified_on" gorm:"column:modified_on"`
}
type Auth struct {
	Model
}

func ExistTagByName (name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByID( id int) bool {
	var tag Tag
	db.Select("name").Where("id = ?", id).First(&tag)
	if tag.Name != ""{
		return true
	}
	return false
}
func EditTag(id int, d interface{})  bool {
	db.Model(&Tag{}).Where("id = ?",  id) . Update(d)

	return true
}

func AddTag(name, createdBy string, state int, ) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}
func GetTags(pageNum, pagesize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pagesize).Find(&tags)

	return
}

func GetTagTotal(maps interface {}) (count int){
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func (tag *Tag)BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
