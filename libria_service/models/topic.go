package models

import (
	"database/sql"
	"libria/utility"
)

type Topic struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Votes       string `json:"votes"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}

type TopicDB struct {
	ID          string
	Title       sql.NullString
	Body        sql.NullString
	Votes       sql.NullString
	CreatedDate sql.NullString
	UpdatedDate sql.NullString
}

func (dbV *TopicDB) GetTopic() (t Topic) {
	t.ID = dbV.ID
	t.Title = utility.GetStringValue(dbV.Title)
	t.Body = utility.GetStringValue(dbV.Body)
	t.Votes = utility.GetStringValue(dbV.Votes)
	t.CreatedDate = utility.GetStringValue(dbV.CreatedDate)
	t.UpdatedDate = utility.GetStringValue(dbV.UpdatedDate)
	return t
}
