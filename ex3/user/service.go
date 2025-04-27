package user

type Service struct {
	persistence *LocalStorage
}

func (s *Service) Create(user *User) {
	// la capa de servicio es como una capa de ingreso donde puedo obtener datos extra y validar
	//validaciones
	user.ID = "123"
}

// get = leer
// post = crear, cargar(?)
// upd = put o patch
// PUT user: paso el user completo (aunque updatee una sola cosa)
// PATCH modifica una sola cosa, especifico
