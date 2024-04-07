package service

import (
	"sort"

	"BitoPro_interview_question/model"
)

type MatchingService struct {
	males   []*model.SinglePerson
	females []*model.SinglePerson
}

func NewMatchingService() *MatchingService {

	result := &MatchingService{}

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

	var index int

	if person.GenderType == model.Male {
		index = sort.Search(len(candidates), func(i int) bool {
			return person.Height <= candidates[i].Height
		})
	} else if person.GenderType == model.Female {
		index = sort.Search(len(candidates), func(i int) bool {
			return person.Height >= candidates[i].Height
		})
	}

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

func (s *MatchingService) RemoveSinglePerson(name string) bool {
	// 先在 males 中尋找符合名字的人
	for i, m := range s.males {
		if m.Name == name {
			// 找到後,將其從 males 中移除
			s.males = append(s.males[:i], s.males[i+1:]...)
			return true
		}
	}

	// 如果在 males 中沒找到,再到 females 中尋找
	for i, f := range s.females {
		if f.Name == name {
			// 找到後,將其從 females 中移除
			s.females = append(s.females[:i], s.females[i+1:]...)
			return true
		}
	}

	// 如果都沒找到,表示沒有這個人,返回 false
	return false
}

func (s *MatchingService) QuerySinglePeople(gender model.Gender, limit int) []*model.SinglePerson {
	var candidates []*model.SinglePerson

	if gender == model.Male {
		candidates = s.males
	} else {
		candidates = s.females
	}

	if limit <= 0 || limit > len(candidates) {
		return candidates
	}

	return candidates[:limit]
}
