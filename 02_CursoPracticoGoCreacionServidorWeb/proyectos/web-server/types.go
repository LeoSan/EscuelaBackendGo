package main

import "encoding/json"

type User struct {
	Name  string `json:"name"` //Anotación especial indicando que si es json debe ser asi
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) IsValid() bool {
	if u.Email == "" {
		return false
	}
	if u.Name == "" {
		return false
	}
	if u.Phone == "" {
		return false
	}
	return true
}

//metodo que devuelve un slice de byte, se necesita para devolver nuestra respuesta en Json
func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u) //Permite empezar a encodear el struc a un json
}

//Nota usamos astericos para no usar copias que no queremos usar
