package service

import (
	"errors"
	"time"
	"web3/constants"
	"web3/model"
)

var guideDao = model.NewGuideDao()

type GuideService interface {
	Create(guide model.Guide) (success bool, err error)
	Detail(id string) (guide *model.Guide, err error)
	ThumbsUp(id string, from string, status string) (success bool, err error)
	ListByOwner(owner string) (guides *[]model.Guide, err error)
	Rec() (guides *[]model.GuideRec, err error)
}

func NewGuideService() GuideService {
	return &guideService{}
}

type guideService struct {
}

func (s *guideService) Create(guide model.Guide) (success bool, err error) {
	guide.LikeCount = 0
	guide.SalesStatus = 0
	guide.CreateTime = time.Now().Unix()
	saveGuide, err := guideDao.CreateGuide(guide)
	return saveGuide, err
}

func (s *guideService) Detail(id string) (guide *model.Guide, err error) {
	guide, err = guideDao.DetailGuide(id)
	return guide, err
}

func (s *guideService) ListByOwner(owner string) (guides *[]model.Guide, err error) {
	guides, err = guideDao.ListByOwner(owner)
	return
}

func (s *guideService) ThumbsUp(id string, from string, status string) (success bool, err error) {

	k := "thumbsUp:" + from + ":" + id
	v := model.GetRedisValue(k)

	if constants.LIKE == status {
		if len(v) != 0 {
			success, err = false, errors.New("already liked")
			return
		}
		model.SetRedisValue(k, "done", 0)
	} else {
		if len(v) == 0 {
			success, err = false, errors.New("already disliked")
			return
		}
		model.DelRedisKey(k)
	}
	success, err = guideDao.ThumbsUp(id, status)
	return
}

func (s *guideService) Rec() (guides *[]model.GuideRec, err error) {
	guides, err = guideDao.Rec()
	return
}
