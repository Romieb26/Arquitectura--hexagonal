package entities

type Titular struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	DNI       string `json:"dni"`
	Telefono  int  `json:"telefono"`
	Direccion string `json:"direccion"`
}

func NewTitular(id int, nombre, apellido, dni string, telefono int, direccion string) *Titular {
	return &Titular{
		ID:        id,
		Nombre:    nombre,
		Apellido:  apellido,
		DNI:       dni,
		Telefono:  telefono,
		Direccion: direccion,
	}
}

func (t *Titular) GetID() int {
	return t.ID
}

func (t *Titular) SetID(id int) {
	t.ID = id
}

func (t *Titular) GetNombre() string {
	return t.Nombre
}

func (t *Titular) SetNombre(nombre string) {
	t.Nombre = nombre
}

func (t *Titular) GetApellido() string {
	return t.Apellido
}

func (t *Titular) SetApellido(apellido string) {
	t.Apellido = apellido
}

func (t *Titular) GetDNI() string {
	return t.DNI
}

func (t *Titular) SetDNI(dni string) {
	t.DNI = dni
}

func (t *Titular) GetTelefono() int {
	return t.Telefono
}

func (t *Titular) SetTelefono(telefono int) {
	t.Telefono = telefono
}

func (t *Titular) GetDireccion() string {
	return t.Direccion
}

func (t *Titular) SetDireccion(direccion string) {
	t.Direccion = direccion
}
