package model

import (
	"database/sql"
)

type Model struct {
	db                *sql.DB
	languageModelPath string
	probWords         []ProbWord
}

func NewModel(db *sql.DB, languageModelPath string) *Model {
	return &Model{db: db, languageModelPath: languageModelPath}
}
