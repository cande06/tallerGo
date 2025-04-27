package bases

import "errors"

// ## ESTRUCTURA
type Dispositivo struct {
	Nombre string
	Estado bool
}

// ## INTERFAZ
type Controlable interface {
	Encender() error
	Apagar() error
	EstadoActual() string
}

// ## FUNCIONES

// Devolver estado del dispositivo
func (d Dispositivo) EstadoActual() string {
	if !d.Estado {
		return "Apagado"
	} else {
		return "Encendido"
	}
}

// Definir errores
func (d *Dispositivo) Encender() error {
	if d.Estado {
		return errors.New("el dispositivo ya se encuentra encendido")
	}
	d.Estado = true
	return nil
}

func (d *Dispositivo) Apagar() error {
	if !d.Estado {
		return errors.New("el dispositivo ya se encuentra apagado")
	}
	d.Estado = false
	return nil
}
