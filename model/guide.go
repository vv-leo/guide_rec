package model

import (
	"gorm.io/gorm"
	"time"
)

type Guide struct {
	Id          int       `json:"id"`
	Owner       string    `json:"owner"`
	SalesStatus int       `json:"sales_status"`
	Price       float64   `json:"price"`
	Content     string    `json:"content"`
	LikeCount   int       `json:"like_count"`
	CreateTime  time.Time `json:"create_time"`
}

func (Guide) TableName() string {
	return "guide"
}

type GuideDao interface {
	SaveGuide(guide Guide) (success bool, err error)
	DetailGuide(id string) (guide *Guide, err error)
}

func NewGuideDao() GuideDao {
	return &guideDao{}
}

type guideDao struct {
}

func (dao *guideDao) operateTable() *gorm.DB {
	return db.Table("guide")
}

func (s *guideDao) SaveGuide(guide Guide) (success bool, err error) {
	if res := s.operateTable().Create(&guide); res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (s *guideDao) DetailGuide(id string) (guide *Guide, err error) {
	if res := db.First(&guide, id); res.Error != nil {
		return nil, res.Error
	}
	return guide, nil
}
