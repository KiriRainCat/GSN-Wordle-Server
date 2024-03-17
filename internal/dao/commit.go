package dao

import "gsn-wordle/internal/model"

var Commit = &CommitDao{}

type CommitDao struct{}

func (*CommitDao) Create(wid int, subject string, word string, definition string) error {
	return DB.Create(&model.Commit{WordId: wid, Subject: subject, Word: word, Definition: definition}).Error
}

func (*CommitDao) GetList() (commits []*model.Commit, err error) {
	err = DB.Find(&commits).Error
	return
}

func (*CommitDao) GetByID(id int) (commit *model.Commit, err error) {
	err = DB.Find(&commit, id).Error
	return
}

func (*CommitDao) GetListByWordId(wid int) (commits []*model.Commit, err error) {
	err = DB.Find(&commits, "wid = ?", wid).Error
	return
}

func (*CommitDao) Delete(id int) error {
	return DB.Delete(&model.Commit{}, id).Error
}
