package models

import (
	"database/sql"
	"libria/utility"
)

type Answer struct {
	ID          string `json:"id"`
	TopicID     string `json:"topic_id"`
	Text        string `json:"body"`
	Votes       string `json:"votes"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}

type AnswerDB struct {
	ID          string
	TopicID     string
	Text        sql.NullString
	Votes       sql.NullString
	CreatedDate sql.NullString
	UpdatedDate sql.NullString
}

func (dbV *AnswerDB) GetAnswer() (a Answer) {
	a.ID = dbV.ID
	a.TopicID = dbV.TopicID
	a.Text = utility.GetStringValue(dbV.Text)
	a.Votes = utility.GetStringValue(dbV.Votes)
	a.CreatedDate = utility.GetStringValue(dbV.CreatedDate)
	a.UpdatedDate = utility.GetStringValue(dbV.UpdatedDate)
	return a
}
