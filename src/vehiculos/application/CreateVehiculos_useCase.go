package application

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain/entities"
	"golang.org/x/crypto/bcrypt"

)

type CreateVehiculoUseCase struct {
	db domain.IVehiculo
}

func NewCreateVehiculoUseCase(db domain.IVehiculo) *CreateVehiculoUseCase {
	return &CreateVehiculoUseCase{
		db: db,
	}
}

func (uc *CreateVehiculoUseCase) Run(vehiculo *entities.Vehiculo) (*entities.Vehiculo, error) {
	err := uc.db.Save(*vehiculo)
	return vehiculo, err

	hash, err := bcrypt.GenerateFromPassword([]byte(vehiculo.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	
	vehiculo.Hash = string(hash)
	vehiculo.Password = "" 
	
	err = uc.db.Save(*vehiculo)
	return vehiculo, err
}
