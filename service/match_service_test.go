package service

import (
	"testing"

	"BitoPro_interview_question/model"
)

func TestAddSinglePersonAndMatch(t *testing.T) {
	service := NewMatchingService()

	// 測試添加一個新的男性用戶,並成功配對到一個女性用戶
	male := &model.SinglePerson{Name: "John", Height: 180, GenderType: model.Male, RemainingDates: 1}
	service.females = []*model.SinglePerson{{Name: "Emma", Height: 165, GenderType: model.Female, RemainingDates: 1}}
	matchedFemale := service.AddSinglePersonAndMatch(male)
	if matchedFemale == nil {
		t.Error("Expected a matched female, but got nil")
	}

	// 測試添加一個新的女性用戶,並成功配對到一個男性用戶
	female := &model.SinglePerson{Name: "Olivia", Height: 170, GenderType: model.Female, RemainingDates: 1}
	service.males = []*model.SinglePerson{{Name: "David", Height: 175, GenderType: model.Male, RemainingDates: 1}}
	matchedMale := service.AddSinglePersonAndMatch(female)
	if matchedMale == nil {
		t.Error("Expected a matched male, but got nil")
	}

	// 測試添加一個新用戶,但沒有找到合適的配對對象
	newMale := &model.SinglePerson{Name: "Mike", Height: 190, GenderType: model.Male, RemainingDates: 1}
	service.females = []*model.SinglePerson{}
	matchedPerson := service.AddSinglePersonAndMatch(newMale)
	if matchedPerson != nil {
		t.Error("Expected a matched male, but got nil")
	}

	// 測試添加一個已存在的用戶（根據姓名判斷）,驗證是否返回 nil
	existingMale := &model.SinglePerson{Name: "John", Height: 180, GenderType: model.Male, RemainingDates: 1}
	matchedPerson = service.AddSinglePersonAndMatch(existingMale)
	if matchedPerson != nil {
		t.Error("Expected nil for an existing person, but got a matched person")
	}
}

func TestRemoveSinglePerson(t *testing.T) {
	service := NewMatchingService()
	service.males = []*model.SinglePerson{
		{Name: "John", Height: 180, GenderType: model.Male, RemainingDates: 1},
		{Name: "David", Height: 175, GenderType: model.Male, RemainingDates: 1},
	}
	service.females = []*model.SinglePerson{
		{Name: "Emma", Height: 165, GenderType: model.Female, RemainingDates: 1},
		{Name: "Olivia", Height: 170, GenderType: model.Female, RemainingDates: 1},
	}

	// 測試成功移除一個男性用戶
	if !service.RemoveSinglePerson("John") {
		t.Error("Expected successful removal of a male person, but got false")
	}

	// 測試成功移除一個女性用戶
	if !service.RemoveSinglePerson("Emma") {
		t.Error("Expected successful removal of a female person, but got false")
	}

	// 測試移除一個不存在的用戶,驗證是否返回 false
	if service.RemoveSinglePerson("NonExistentPerson") {
		t.Error("Expected false for removing a non-existent person, but got true")
	}
}

func TestQuerySinglePeople(t *testing.T) {
	service := NewMatchingService()
	service.males = []*model.SinglePerson{
		{Name: "John", Height: 180, GenderType: model.Male, RemainingDates: 1},
		{Name: "David", Height: 175, GenderType: model.Male, RemainingDates: 1},
		{Name: "Mike", Height: 190, GenderType: model.Male, RemainingDates: 1},
	}
	service.females = []*model.SinglePerson{
		{Name: "Emma", Height: 165, GenderType: model.Female, RemainingDates: 1},
		{Name: "Olivia", Height: 170, GenderType: model.Female, RemainingDates: 1},
	}

	// 測試查詢男性用戶列表,並驗證返回的結果是否正確
	maleList := service.QuerySinglePeople(model.Male, 10)
	if len(maleList) != 3 {
		t.Errorf("Expected 3 males, but got %d", len(maleList))
	}

	// 測試查詢女性用戶列表,並驗證返回的結果是否正確
	femaleList := service.QuerySinglePeople(model.Female, 10)
	if len(femaleList) != 2 {
		t.Errorf("Expected 2 females, but got %d", len(femaleList))
	}

	// 測試查詢用戶列表時,限制返回的數量,驗證返回的結果是否符合限制
	limitedList := service.QuerySinglePeople(model.Male, 2)
	if len(limitedList) != 2 {
		t.Errorf("Expected 2 persons, but got %d", len(limitedList))
	}
}
