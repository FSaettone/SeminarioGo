package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Transaccion struct {
	monto      float64
	ID_origen  string
	ID_destino string
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

func calcularHash(bloque Bloque) []byte {
	montoAStr := strconv.FormatFloat(bloque.data.monto, 'f', -1, 64)
	datos := []byte(string(bloque.hashPrevio) + montoAStr + bloque.data.ID_origen + bloque.data.ID_destino + bloque.data.timestamp.String())
	hash := sha256.Sum256(datos)
	return hash[:]
}

func CrearBloque(t Transaccion, hashPrevio []byte) Bloque {
	newBloque := Bloque{
		data:       t,
		hashPrevio: hashPrevio,
		timestamp:  time.Now(),
	}
	hash := calcularHash(newBloque)
	newBloque.hash = hash
	return newBloque
}

func (bc *Blockchain) EnviarTransaccion(origen, destino string, monto float64) error {
	saldo_id_origen := bc.ObtenerSaldo(origen)
	if saldo_id_origen < monto {
		return fmt.Errorf("Saldo insuficiente")
	}
	t := Transaccion{
		monto:      monto,
		ID_origen:  origen,
		ID_destino: destino,
		timestamp:  time.Now(),
	}
	newBloque := CrearBloque(t, (*bc)[len(*bc)-1].hash)
	*bc = append(*bc, newBloque)
}

func (bc *Blockchain) ObtenerSaldo(id string) float64 {
	saldo := 0.0
	for _, bloque := range *bc {
		if bloque.data.ID_origen == id {
			saldo -= bloque.data.monto
		} else if bloque.data.ID_destino == id {
			saldo += bloque.data.monto
		}
	}
	return saldo
}

func iguales(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func ValidarCadena(bc Blockchain) bool {
	for i := 1; i < len(bc); i++ {
		bloqueAct := bc[i]
		bloqueAnt := bc[i-1]
		if !iguales(bloqueAct.hashPrevio, calcularHash(bloqueAnt)) {
			return false
		}
	}
	return true
}

func main() {
	b1 := CrearBilletera("")
	b2 := CrearBilletera("")

}
