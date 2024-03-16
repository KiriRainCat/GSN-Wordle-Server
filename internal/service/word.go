package service

import (
	"gsn-wordle/internal/dao"
	"gsn-wordle/internal/model"
	"gsn-wordle/internal/pkg/errs"
	"gsn-wordle/internal/pkg/util"
	"math/rand"
)

var Word = &WordService{dao: dao.Word}

type WordService struct {
	dao *dao.WordDao
}

func (s *WordService) GetList() (list []*model.Word, err error) {
	return s.dao.GetList()
}

func (s *WordService) GetById(id int) (word *model.Word, err error) {
	return s.dao.GetById(id)
}

func (s *WordService) GetWordOfTheDay() (word *model.Word, err error) {
	word, err = s.dao.GetWordOfTheDay()

	if err != nil {
		word, err = s.GetRandomWord()
		if err != nil {
			return nil, err
		}
	}

	return
}

func (s *WordService) GetRandomWord() (word *model.Word, err error) {
	list, err := s.GetList()
	if err != nil {
		return nil, err
	}

	return list[rand.Intn(len(list))], nil
}

func (s *WordService) Create(word string, definition string) error {
	if err := s.dao.Create(word, definition); err != nil {
		if util.IsViolatingUniqueConstraint(err) {
			return errs.ErrUniqueConstraint("word")
		}
		return errs.ErrServer
	}

	return nil
}

func (s *WordService) Update(id int, word string, definition string) error {
	return s.dao.Update(id, word, definition)
}

func (s *WordService) Delete(id int) error {
	return s.dao.Delete(id)
}
