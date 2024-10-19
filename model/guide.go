package model

import (
	"gorm.io/gorm"
	"web3/constants"
)

type Guide struct {
	Id          int    `json:"id"`
	Owner       string `json:"owner"`
	SalesStatus int    `json:"sales_status"`
	Price       string `json:"price"`
	Content     string `json:"content"`
	LikeCount   int    `json:"like_count"`
	CreateTime  int64  `json:"create_time"`
}

type GuideRec struct {
	Id          int    `json:"id"`
	Owner       string `json:"owner"`
	SalesStatus int    `json:"sales_status"`
	Price       string `json:"price"`
	Content     string `json:"content"`
	LikeCount   int    `json:"like_count"`
	CreateTime  int64  `json:"create_time"`
	FansCount   int    `json:"fans_count"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

type GuideDao interface {
	CreateGuide(guide Guide) (success bool, err error)
	DetailGuide(id string) (guide *Guide, err error)
	ThumbsUp(id string, status string) (success bool, err error)
	ListByOwner(owner string) (guides *[]Guide, err error)
	Rec() (guideRecs *[]GuideRec, err error)
}

func NewGuideDao() GuideDao {
	return &guideDao{}
}

type guideDao struct {
}

func (s *guideDao) operateTable() *gorm.DB {
	return db.Table("guide")
}

func (s *guideDao) CreateGuide(guide Guide) (success bool, err error) {
	if res := s.operateTable().Create(&guide); res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (s *guideDao) DetailGuide(id string) (guide *Guide, err error) {
	if res := s.operateTable().First(&guide, id); res.Error != nil {
		return nil, res.Error
	}
	return guide, nil
}

func (s *guideDao) ListByOwner(owner string) (guides *[]Guide, err error) {
	if res := s.operateTable().Where("owner = ?", owner).Find(&guides); res.Error != nil {
		return guides, res.Error
	}
	return guides, nil
}

func (s *guideDao) ThumbsUp(id string, status string) (bool bool, err error) {
	var guide Guide
	if resF := s.operateTable().First(&guide, id); resF.Error != nil {
		return false, resF.Error
	}
	if constants.LIKE == status {
		guide.LikeCount++
	} else {
		guide.LikeCount--
	}
	if resS := s.operateTable().Save(&guide); resS.Error != nil {
		return false, resS.Error
	}
	return true, err
}

func (s *guideDao) Rec() (guideRecs *[]GuideRec, err error) {
	if res := s.operateTable().Select("guide.id AS Id, guide.owner AS Owner,guide.sales_status AS SalesStatus,guide.price AS Price,guide.content AS Content," +
		"guide.like_count AS LikeCount,guide.create_time AS CreateTime,user.fans_count AS FansCount, user.avatar AS Avatar,user.description AS Description").
		Joins("guide LEFT JOIN user ON guide.owner = user.id").
		Scan(&guideRecs); res.Error != nil {
		return guideRecs, res.Error
	}
	return guideRecs, nil
}
