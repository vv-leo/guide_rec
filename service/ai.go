package service

import "web3/utils"

type AiService interface {
	GenerateGuide(cueword string) (res string, err error)
}

func NewAiService() AiService {
	return &aiService{}
}

type aiService struct {
}

func (s aiService) GenerateGuide(cueword string) (res string, err error) {
	role := "你是豆包，是一名出游规划专家"
	res, err = utils.GetByDoubao(role, cueword)
	return res, err
}
