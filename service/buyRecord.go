package service

import "web3/model"

var buyRecordDao = model.NewBuyRecord()

type BuyRecordService interface {
	GetByGuideId(guideId string) (records []model.BuyRecord, err error)
}

func NewBuyRecordService() BuyRecordService {
	return &buyRecordService{}
}

type buyRecordService struct {
}

func (s *buyRecordService) GetByGuideId(guideId string) (records []model.BuyRecord, err error) {
	records, err = buyRecordDao.FindByGuideId(guideId)
	return
}
