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

func (*WordDao) Create(word string, definition string) error {
	return DB.Create(&model.Word{Value: word, Definition: definition, Length: len(word)}).Error
}

func (*WordDao) Update(id int, word string, definition string) error {
	return DB.Save(&model.Word{Id: id, Value: word, Definition: definition, Length: len(word)}).Error
}

func (*WordDao) Delete(id int) error {
	return DB.Delete(&model.Word{}, id).Error
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
