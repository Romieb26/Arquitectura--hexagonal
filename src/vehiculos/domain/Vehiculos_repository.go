package domain

import "github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain/entities"

type IVehiculo interface {
	Save(vehiculo entities.Vehiculo) error
	Update(id int, vehiculo entities.Vehiculo) error
	Delete(id int) error
	FindByID(id int) (entities.Vehiculo, error)
	FindByPlaca(placa string) (*entities.Vehiculo, error) 
	GetAll() ([]entities.Vehiculo, error)

}
