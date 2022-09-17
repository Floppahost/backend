package database

import (
	"fmt"
	"strconv"

	argonpass "github.com/dwin/goArgonPass"
)

func Login(database any, usuario string, senha string) (uint, error) {
	db := DB

	// declaramos o map result
	result := map[string]any{}

	// fazemos a query, puxando senha e id
	db.Model(database).Select("senha, id").Where("usuario = ?", usuario).Find(&result)
	// Em SQL: SELECT senha, id FROM users WHERE usuario=usuario

	// transformamos a senha encriptada em string
	senhaEncriptada := fmt.Sprintf("%v", result["senha"])

	// transformamos o id em int
	id, _ := strconv.ParseUint(fmt.Sprintf("%v", result["id"]), 10, 64)

	// verificamos a veracidade da senha
	err := argonpass.Verify(senha, senhaEncriptada)

	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
