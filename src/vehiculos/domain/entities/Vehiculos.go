package entities

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Vehiculo struct {
	ID       int    `json:"id"`
	Marca    string `json:"marca"`
	Modelo   string `json:"modelo"`
	Año      int    `json:"año"`
	Color    string `json:"color"`
	Placa    string `json:"placa"`
	Password string `json:"password,omitempty"` 
	Hash     string `json:"-"`                  
}


func NewVehiculo(id int, marca, modelo string, año int, color, placa, password string) (*Vehiculo, error) {
	v := &Vehiculo{
		ID:     id,
		Marca:  marca,
		Modelo: modelo,
		Año:    año,
		Color:  color,
		Placa:  placa,
	}

	if password != "" {
		if err := v.SetPassword(password); err != nil {
			return nil, err
		}
	}

	return v, nil
}

// Establece el password generando el hash
func (v *Vehiculo) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("no se pudo generar el hash de la contraseña")
	}
	
	v.Hash = string(hash)
	v.Password = "" // Limpiamos el password plano
	return nil
}

// Verifica si el password coincide con el hash
func (v *Vehiculo) VerifyPassword(password string) error {
	if v.Hash == "" {
		return errors.New("no hay hash almacenado para comparar")
	}
	return bcrypt.CompareHashAndPassword([]byte(v.Hash), []byte(password))
}

// Getters
func (v *Vehiculo) GetID() int { return v.ID }
func (v *Vehiculo) GetMarca() string { return v.Marca }
func (v *Vehiculo) GetModelo() string { return v.Modelo }
func (v *Vehiculo) GetAño() int { return v.Año }
func (v *Vehiculo) GetColor() string { return v.Color }
func (v *Vehiculo) GetPlaca() string { return v.Placa }
func (v *Vehiculo) GetHash() string { return v.Hash }

// Setters
func (v *Vehiculo) SetID(id int) { v.ID = id }
func (v *Vehiculo) SetMarca(marca string) { v.Marca = marca }
func (v *Vehiculo) SetModelo(modelo string) { v.Modelo = modelo }
func (v *Vehiculo) SetAño(año int) { v.Año = año }
func (v *Vehiculo) SetColor(color string) { v.Color = color }
func (v *Vehiculo) SetPlaca(placa string) { v.Placa = placa }