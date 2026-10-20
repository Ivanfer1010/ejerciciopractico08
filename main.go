// Ivan Gualotuña
package main
import( 
	"fmt"
)
//Capa UIDE 
//Reto 1: Inscribir estadisticas y jugadores
type Estadisticas struct {
	Kills int
	Muertes int
	Asistencias int
}
type Jugador struct{
	Nick string
	Equipo string
	Stats Estadisticas
	Activo bool
}
//Ejericio 2
func registrar (j *jugador, kills, muertes, asistencias int){
	j.Stats.Kills = Kills
	j.Stats.Muertes = muertes
	j.Stats.Asistencias = asistencias
}
func kda(e Estadisticas)float64 {
	muertes := e.Muertes
	if muertes == 0 {
		muertes = 1
	}
	return float64 (e.kills.Asistencias) / float64 (muertes)
}
//Ejercicio 3 (Aqui se devuelve una copia)
func mvp (roster []Jugador) Jugador {
	var mejor Jugador
	for_, j := range roster {
		if j.Activo && kda(j.Stats) > kda (mejor.Stats){
			mejor = j
		}
	}
	return mejor
}