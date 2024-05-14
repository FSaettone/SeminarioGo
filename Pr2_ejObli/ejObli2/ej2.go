package main

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type Transaccion struct {
	monto      float64
	ID_Envio   string
	ID_Recibio string
	timestamp  time.Time
}

type Bloque struct {
	hash       []byte
	hashPrevio []byte
	data       Transaccion
	timestamp  time.Time
}

type Blockchain []Bloque

type Billetera struct {
	ID     string
	Nombre string
}

func GenerarID() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func CrearBilletera(nombre string) Billetera {
	return Billetera{
		ID:     GenerarID(),
		Nombre: nombre,
	}
}

func main() {
	b1 := CrearBilletera("Facundo Saettone")
	b2 := CrearBilletera("Nicolas Mariani")

}
