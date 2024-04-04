package service

import (
	"BitoPro_interview_question/model"

	"sort"
)

type MatchingService struct {
	males   []*model.SinglePerson
	females []*model.SinglePerson
}

func createMockData() ([]*model.SinglePerson, []*model.SinglePerson) {
	males := []*model.SinglePerson{
		{Name: "John", Height: 180, GenderType: model.Male, RemainingDates: 3},
		{Name: "Mike", Height: 165, GenderType: model.Male, RemainingDates: 2},
		{Name: "Tom", Height: 168, GenderType: model.Male, RemainingDates: 1},
		{Name: "David", Height: 178, GenderType: model.Male, RemainingDates: 4},
		{Name: "Chris", Height: 182, GenderType: model.Male, RemainingDates: 2},
	}

	females := []*model.SinglePerson{
		{Name: "Emily", Height: 165, GenderType: model.Female, RemainingDates: 2},
		{Name: "Emma", Height: 170, GenderType: model.Female, RemainingDates: 3},
		{Name: "Olivia", Height: 168, GenderType: model.Female, RemainingDates: 1},
		{Name: "Sophia", Height: 154, GenderType: model.Female, RemainingDates: 1},
		{Name: "Ava", Height: 167, GenderType: model.Female, RemainingDates: 3},
	}

	return males, females
}

func NewMatchingService() *MatchingService {
	males, females := createMockData()

	result := &MatchingService{
		males:   males,
		females: females,
	}

	sort.Slice(result.males, func(i, j int) bool {
		return result.males[i].Height < result.males[j].Height
	})

	sort.Slice(result.females, func(i, j int) bool {
		return result.females[i].Height < result.females[j].Height
	})

	return result
}

func (s *MatchingService) AddSinglePersonAndMatch(person *model.SinglePerson) *model.SinglePerson {
	if s.isPersonExist(person) {
		return nil
	}

	if person.GenderType == model.Male {
		s.males = append(s.males, person)
		sort.Slice(s.males, func(i, j int) bool {
			return s.males[i].Height < s.males[j].Height
		})
		return s.matchPerson(person, s.females)
	} else if person.GenderType == model.Female {
		s.females = append(s.females, person)
		sort.Slice(s.females, func(i, j int) bool {
			return s.females[i].Height < s.females[j].Height
		})
		return s.matchPerson(person, s.males)
	}
	return nil
}

func (s *MatchingService) isPersonExist(person *model.SinglePerson) bool {
	if person.GenderType == model.Male {
		for _, p := range s.males {
			if p.Name == person.Name {
				return true
			}
		}
	} else if person.GenderType == model.Female {
		for _, p := range s.females {
			if p.Name == person.Name {
				return true
			}
		}
	}
	return false
}

func (s *MatchingService) matchPerson(person *model.SinglePerson, candidates []*model.SinglePerson) *model.SinglePerson {
	index := sort.Search(len(candidates), func(i int) bool {
		return person.Height <= candidates[i].Height
	})

	for i := index - 1; i >= 0; i-- {
		candidate := candidates[i]
		if candidate.RemainingDates > 0 {
			candidate.RemainingDates--
			person.RemainingDates--
			if candidate.RemainingDates == 0 {
				s.removePerson(candidate)
			}
			if person.RemainingDates == 0 {
				s.removePerson(person)
			}
			return candidate
		}
	}
	return nil
}

func (s *MatchingService) removePerson(person *model.SinglePerson) {
	if person.GenderType == model.Male {
		s.removeMale(person)
	} else if person.GenderType == model.Female {
		s.removeFemale(person)
	}
}

func (s *MatchingService) removeMale(male *model.SinglePerson) {
	for i, m := range s.males {
		if m == male {
			s.males = append(s.males[:i], s.males[i+1:]...)
			break
		}
	}
}

func (s *MatchingService) removeFemale(female *model.SinglePerson) {
	for i, f := range s.females {
		if f == female {
			s.females = append(s.females[:i], s.females[i+1:]...)
			break
		}
	}
}

func (s *MatchingService) RemoveSinglePerson(name string) {
	// 實現移除單身人士的邏輯
}

func (s *MatchingService) QuerySinglePeople(gender string, limit int) []*model.SinglePerson {
	// 實現查詢單身人士的邏輯
	return nil
}
