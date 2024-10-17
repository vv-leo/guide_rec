package model

import (
	"gorm.io/gorm"
	"time"
)

type Guide struct {
	Id          int       `json:"id"`
	Owner       string    `json:"owner"`
	SalesStatus int       `json:"sales_status"`
	Price       int       `json:"price"`
	Content     string    `json:"content"`
	LikeCount   int       `json:"like_count"`
	CreateTime  time.Time `json:"create_time"`
}

type GuideDao interface {
	SaveGuide(guide Guide) (success bool, err error)
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

func (s *guideDao) DetailGuide(id string) (guide Guide, err error) {
	if res := s.operateTable().First(&Guide{}, id); res.Error != nil {
		return , res.Error
	}
	//if res := s.operateTable().Create(&guide); res.Error != nil {
	//	return false, res.Error
	//}
	return true, nil
}
