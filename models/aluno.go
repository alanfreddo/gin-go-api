package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model        //insere ID, hora de criação, update e delete.
	Nome       string `json:"nome"`
	CPF        string `json:"cpf"`
	RG         string `json:"rg"`
}
