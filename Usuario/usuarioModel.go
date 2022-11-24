package Usuario

import "gorm.io/gorm"

func main() {
	type Usuario struct {
		gorm.Model
		Nombre   string `json:"nombre"`
		Apellido string `json:"apellido"`
		Email    string `json:"email"`
	}
}
