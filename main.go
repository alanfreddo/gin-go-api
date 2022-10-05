package main

import (
	"github.com/alanfreddo/gin-go-api/models"
	"github.com/alanfreddo/gin-go-api/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Alan", CPF: "000000000", RG: "9876347283723"},
		{Nome: "Joya", CPF: "000000001", RG: "0000000000000"},
	}
	routes.HandleRequests()

}
