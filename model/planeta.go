package model

import (
	"math"
)

/*
Planeta Estructura de un planeta
	nombre: Nombre del planeta
	velocidad: su velocidad en grados por día (velocidad indica movimiento horario,
	negativa indica movimiento anti-horario
	posicion: Coordenadas polares de la posición actual del planeta en cuestión
*/
type Planeta struct {
	nombre    string
	velocidad float64
	posicion  CoordsPolares
}

// CoordsPolares Estructura de las coordenadas polares de un planeta
type CoordsPolares struct {
	radio  float64
	angulo float64
}

// CrearPlaneta Crea un nuevo planeta con su nombre, velocidad y radio ingresados por parámetro
func CrearPlaneta(nombre string, velocidad float64, radio float64) Planeta {
	p := Planeta{
		nombre:    nombre,
		velocidad: velocidad,
		posicion:  CoordsPolares{radio, 0},
	}
	agregarPlaneta(p)
	return p
}

// CalcularDiasPorAnio Calcula los días que contiene un año para el planeta receiver
func (p Planeta) CalcularDiasPorAnio() int {
	return int(math.Abs(360 / p.velocidad))
}
