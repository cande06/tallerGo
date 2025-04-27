package user

//capa de persistencia

type LocalStorage struct {
	m map[string]*User
}

//pasar puntero es mas rapido que pasar la estructura completa
//es mas optimo acceder directamente
func (l *LocalStorage) Set(user *User) error {
	l.m[user.ID] = user
	return nil
}

// func (l *LocalStorage) Get(id string) (*User, error) {

// }

// func (l *LocalStorage) Delete(id string) error {

// }
