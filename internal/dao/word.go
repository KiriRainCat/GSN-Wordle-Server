package dao

import (
	"context"
	"gsn-wordle/internal/model"
)

var Word = &WordDao{}

type WordDao struct{}

func (*WordDao) GetList() (list []*model.Word, err error) {
	err = DB.Find(&list).Error
	return
}

func (*WordDao) GetById(id int) (word *model.Word, err error) {
	err = DB.Find(&word, id).Error
	return
}

func (*WordDao) Create(subject string, word string, definition string) (int, error) {
	data := model.Word{Subject: subject, Value: word, Definition: definition, Length: len(word)}
	result := DB.Create(&data)
	return data.Id, result.Error
}

func (*WordDao) Update(id int, subject string, word string, definition string) error {
	return DB.Model(&model.Word{Id: id}).Updates(map[string]any{"subject": subject, "value": word, "length": len(word)}).Error
}

func (*WordDao) Delete(id int) error {
	return DB.Delete(&model.Word{}, id).Error
}

func (*WordDao) SetActiveState(id int, active bool) error {
	return DB.Model(&model.Word{Id: id}).Select("active").UpdateColumn("active", active).Error
}

func (dao *WordDao) GetWordOfTheDay() (word *model.Word, err error) {
	result := Redis.Get(context.Background(), "wordOfTheDay")
	if result.Err() != nil {
		return nil, result.Err()
	}

	id, err := result.Int()
	if err != nil {
		return nil, err
	}

	return dao.GetById(id)
}

func (*WordDao) SetWordOfTheDay(id int) error {
	return Redis.Set(context.Background(), "wordOfTheDay", id, 0).Err()
}
