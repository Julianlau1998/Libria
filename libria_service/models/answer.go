package models

import (
	"database/sql"
	"libria/utility"
)

type Answer struct {
	ID          string `json:"id"`
	TopicID     string `json:"topic_id"`
	UserID      string `json:"user_id"`
	Text        string `json:"body"`
	Username    string `json:"username"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	Votes       []Vote `json:"votes"`
}

type AnswerDB struct {
	ID          string
	TopicID     string
	UserID      sql.NullString
	Text        sql.NullString
	Username    sql.NullString
	CreatedDate sql.NullString
	UpdatedDate sql.NullString
}

func (dbV *AnswerDB) GetAnswer() (a Answer) {
	a.ID = dbV.ID
	a.UserID = utility.GetStringValue(dbV.UserID)
	a.TopicID = dbV.TopicID
	a.Text = utility.GetStringValue(dbV.Text)
	a.Username = utility.GetStringValue(dbV.Username)
	a.CreatedDate = utility.GetStringValue(dbV.CreatedDate)
	a.UpdatedDate = utility.GetStringValue(dbV.UpdatedDate)
	return a
}
