package model

import "gorm.io/gorm"

type BuyRecord struct {
	Id      int    `json:"id"`
	GuideId string `json:"guide_id"`
	Buyer   string `json:"buyer"`
	Seller  string `json:"seller"`
	Price   string `json:"price"`
	Time    int64  `json:"time"`
}

type BuyRecordDao interface {
	CreateBuyRecord(buyRecord BuyRecord) (success bool, err error)
	FindByGuideId(guideId string) (buyRecords []BuyRecord, err error)
}

func NewBuyRecord() BuyRecordDao {
	return &buyRecordDao{}
}

type buyRecordDao struct {
}

func (s *buyRecordDao) operateTable() *gorm.DB {
	return db.Table("buy_record")
}

func (s *buyRecordDao) CreateBuyRecord(buyRecord BuyRecord) (success bool, err error) {
	if res := s.operateTable().Create(&buyRecord); res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (s *buyRecordDao) FindByGuideId(guideId string) (buyRecords []BuyRecord, err error) {
	if res := s.operateTable().Where("guide_id = ?", guideId).Find(&buyRecords); res.Error != nil {
		return nil, res.Error
	}
	return buyRecords, nil
}
