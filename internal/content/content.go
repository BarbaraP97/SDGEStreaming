/* @titulo: modulo de gestion de contenido
   @descripción: Define la estructura de datos para contenido streaming y proporciona funciones para agregar, listar, calificar y ver calificaciones individuales de contenido en memoria. */
package content // 'content' dentro de 'internal' para que sea no sea importado por otros paquetes del proyecto

import (
		"fmt" // Imrpime en consola
		"strconv" // Convierte cadenas a números enteros
		"strings" // Manipula cadenas de texto
)
// Defino la estructura 'UserRating' para almacenar las calificaciónes de los usuarios hacia el contenido
type UserRating struct {
		UserID int		`json:"user_id"` // Campo UserID para identificar al usuario que calificó
		Rating float64	`json:"rating"` // Campo rating para almacenar la puntuación
	}
// funcion para formatear el float de la puntuación maxima a número entero
func mostrarRating(rating float64) string {
	if r == 10.0 {
		return "10" 
	}
}
// Defino la estructura 'Content' para representar cada item de contenido
type Content struct {
	ID       	  int         `json:"id"` // Campo ID para identificar al contenido de forma única
	Title    	  string      `json:"title"` // Campo title para almacenar el título del contenido
	Type     	  string      `json:"type"` // Campo type para almacenar el tipo (Adiovisual/Audio)
	Duration 	  int         `json:"duration"` // Campo duration para almacenar la duración en minutos
	AverageRating float64     `json:"average_rating"` // Almacena la calificación promedio del contenidoq
	Ratings       []UserRating `json:"ratings"` // Almacena todas las calificaciones
}
// Funcion que Calcula cada vez que se agrega un nuevo rating
