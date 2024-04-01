package service

import "BitoPro_interview_question/model"

type MatchingService struct {
	males   []*model.SinglePerson
	females []*model.SinglePerson
}

func NewMatchingService() *MatchingService {
	return &MatchingService{
		males:   make([]*model.SinglePerson, 0),
		females: make([]*model.SinglePerson, 0),
	}
}

func (s *MatchingService) AddSinglePersonAndMatch(person *model.SinglePerson) {
	// 實現添加單身人士並進行配對的邏輯
}

func (s *MatchingService) RemoveSinglePerson(name string) {
	// 實現移除單身人士的邏輯
}

func (s *MatchingService) QuerySinglePeople(gender string, limit int) []*model.SinglePerson {
	// 實現查詢單身人士的邏輯
	return nil
}
