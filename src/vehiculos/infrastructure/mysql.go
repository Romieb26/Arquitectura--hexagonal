package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Romieb26/Arquitectura--hexagonal/src/core"
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain/entities"
)

type MysqlVehiculo struct {
	conn *sql.DB
}

func NewMysqlVehiculoRepository() domain.IVehiculo {
	conn := core.GetDB()
	return &MysqlVehiculo{conn: conn}
}

// Save implementa el método de la interfaz IVehiculos.
func (mysql *MysqlVehiculo) Save(vehiculo entities.Vehiculo) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO vehiculos (marca, modelo, año, color, placa) VALUES (?, ?, ?, ?, ?)",
		vehiculo.Marca,
		vehiculo.Modelo,
		vehiculo.Año,
		vehiculo.Color,
		vehiculo.Placa,
	)
	if err != nil {
		log.Println("Error al guardar el vehículo:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	vehiculo.SetID(int(idInserted))
	return nil
}

// Update implementa el método de la interfaz IVehiculos.
func (mysql *MysqlVehiculo) Update(id int, vehiculo entities.Vehiculo) error {
	result, err := mysql.conn.Exec(
		"UPDATE vehiculos SET marca = ?, modelo = ?, año = ?, color = ?, placa = ? WHERE id = ?",
		vehiculo.Marca,
		vehiculo.Modelo,
		vehiculo.Año,
		vehiculo.Color,
		vehiculo.Placa,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar el vehículo:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró el vehículo con ID:", id)
		return fmt.Errorf("vehículo con ID %d no encontrado", id)
	}

	return nil
}

// Delete implementa el método de la interfaz IVehiculos.
func (mysql *MysqlVehiculo) Delete(id int) error {
	_, err := mysql.conn.Exec("DELETE FROM vehiculos WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar el vehículo:", err)
		return err
	}
	return nil
}

// FindByID implementa el método de la interfaz IVehiculos.
func (mysql *MysqlVehiculo) FindByID(id int) (entities.Vehiculo, error) {
	var vehiculo entities.Vehiculo
	row := mysql.conn.QueryRow("SELECT id, marca, modelo, año, color, placa FROM vehiculos WHERE id = ?", id)

	err := row.Scan(
		&vehiculo.ID,
		&vehiculo.Marca,
		&vehiculo.Modelo,
		&vehiculo.Año,
		&vehiculo.Color,
		&vehiculo.Placa,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Vehículo no encontrado:", err)
			return entities.Vehiculo{}, fmt.Errorf("vehículo con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el vehículo por ID:", err)
		return entities.Vehiculo{}, err
	}

	return vehiculo, nil
}

// GetAll es un método adicional para obtener todos los vehículos.
func (mysql *MysqlVehiculo) GetAll() ([]entities.Vehiculo, error) {
	var vehiculos []entities.Vehiculo

	rows, err := mysql.conn.Query("SELECT id, marca, modelo, año, color, placa FROM vehiculos")
	if err != nil {
		log.Println("Error al obtener todos los vehículos:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var vehiculo entities.Vehiculo
		err := rows.Scan(
			&vehiculo.ID,
			&vehiculo.Marca,
			&vehiculo.Modelo,
			&vehiculo.Año,
			&vehiculo.Color,
			&vehiculo.Placa,
		)
		if err != nil {
			log.Println("Error al escanear vehículo:", err)
			return nil, err
		}
		vehiculos = append(vehiculos, vehiculo)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return vehiculos, nil
}
