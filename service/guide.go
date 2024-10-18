package service

import (
	"time"
	"web3/model"
)

var guideDao = model.NewGuideDao()

type GuideService interface {
	Create(guide model.Guide) (success bool, err error)
	Detail(id string) (guide *model.Guide, err error)
	ThumbsUp(id string, account string, status string) (success bool, err error)
}

func NewGuideService() GuideService {
	return &guideService{}
}

type guideService struct {
}

func (s *guideService) Create(guide model.Guide) (success bool, err error) {

	guide.LikeCount = 0
	guide.SalesStatus = 0
	guide.Price = 0
	guide.CreateTime = time.Now()
	saveGuide, err := guideDao.SaveGuide(guide)
	return saveGuide, err

}

func (s *guideService) Detail(id string) (guide *model.Guide, err error) {
	guideDetail, err := guideDao.DetailGuide(id)
	return guideDetail, err
}

func (s *guideService) ThumbsUp(id string, account string, status string) (success bool, err error) {

	model.GetRedisValue(account)
	return true, nil
}
