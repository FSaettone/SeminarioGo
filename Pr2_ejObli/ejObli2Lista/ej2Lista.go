package main

import (
	"container/list"
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

type Blockchain struct {
	Bloques *list.List
}

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
	newBloque := CrearBloque(t, nil)
	bc.Bloques.PushBack(newBloque)

	return nil
}

func (bc *Blockchain) AgregarSaldo(destino string, monto float64) {
	t := Transaccion{
		monto:      monto,
		ID_origen:  "Ingreso",
		ID_destino: destino,
		timestamp:  time.Now(),
	}
	newBloque := CrearBloque(t, nil)
	bc.Bloques.PushBack(newBloque)

}

func (bc *Blockchain) ObtenerSaldo(id string) float64 {
	saldo := 0.0
	for i := bc.Bloques.Front(); i != nil; i = i.Next() {
		bloque := i.Value.(Bloque)
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
	if bc.Bloques.Len() == 0 {
		return true
	}
	hashAnt := bc.Bloques.Front().Value.(Bloque).hash
	for i := bc.Bloques.Front().Next(); i != nil; i = i.Next() {
		bloqueAct := i.Value.(Bloque)
		if !iguales(bloqueAct.hash, hashAnt) {
			return false
		}
		hashAnt = bloqueAct.hash
	}
	return true
}

func main() {
	fmt.Println("")
	bc := Blockchain{
		Bloques: list.New(),
	}

	b1 := CrearBilletera("Piti Fernandez")
	b2 := CrearBilletera("Ciro Martinez")
	b3 := CrearBilletera("Ricardo iorio")

	bc.AgregarSaldo(b1.ID, 10000.0)
	bc.AgregarSaldo(b2.ID, 12000.0)

	fmt.Printf("Saldo de %s: %.2f\n", b1.Nombre, bc.ObtenerSaldo(b1.ID))
	fmt.Printf("Saldo de %s: %.2f\n", b2.Nombre, bc.ObtenerSaldo(b2.ID))
	fmt.Printf("Saldo de %s: %.2f\n", b3.Nombre, bc.ObtenerSaldo(b3.ID))
	err := bc.EnviarTransaccion(b1.ID, b2.ID, 1100.0)
	if err != nil {
		fmt.Println(err)
	}

	err = bc.EnviarTransaccion(b2.ID, b3.ID, 1200.0)
	if err != nil {
		fmt.Println(err)
	}

	err = bc.EnviarTransaccion(b3.ID, b2.ID, 1300.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("")
	fmt.Println("Luego de las transacciones")
	fmt.Println("")
	fmt.Printf("Saldo de %s: %.2f\n", b1.Nombre, bc.ObtenerSaldo(b1.ID))
	fmt.Printf("Saldo de %s: %.2f\n", b2.Nombre, bc.ObtenerSaldo(b2.ID))
	fmt.Printf("Saldo de %s: %.2f\n", b3.Nombre, bc.ObtenerSaldo(b3.ID))
}
