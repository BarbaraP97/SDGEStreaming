package contentclass

import (
    "SDGEStreaming/internal/categories"
    "SDGEStreaming/internal/errors"
)

// Clasificaciones por edad soportadas
var AgeRatings = []string{"Infantil", "Adolescente", "Adulto"}

// Variables globales para clasificaciones
var (
    ratings   = make(map[string]categories.ContentRating)
    nextID    = 1
)

// Inicializo las clasificaciones por defecto
func init() {
    AddRating("Infantil", "Contenido adecuado para niños menores de 13 años", 0)
    AddRating("Adolescente", "Contenido adecuado para adolescentes (13+)", 13)
    AddRating("Adulto", "Contenido para adultos (18+)", 18)
}

// Agrego una nueva clasificación por edad
func AddRating(name, description string, minAge int) *categories.ContentRating {
    newRating := categories.ContentRating{
        ID:          nextID,
        Name:        name,
        Description: description,
        MinAge:      minAge,
    }
    
    ratings[name] = newRating
    nextID++
    return &newRating
}

// Obtengo una clasificación por nombre
func GetRatingByName(name string) (*categories.ContentRating, error) {
    rating, exists := ratings[name]
    if !exists {
        return nil, errors.NewAppError("CONTENT_004", "Clasificación por edad inválida", name)
    }
    return &rating, nil
}

// Obtengo todas las clasificaciones disponibles
func GetAllRatings() []categories.ContentRating {
    var allRatings []categories.ContentRating
    for _, rating := range ratings {
        allRatings = append(allRatings, rating)
    }
    return allRatings
}

// Valido si un usuario puede acceder a contenido basado en su edad
func CanAccessContent(userAge int, contentRating string) bool {
    rating, err := GetRatingByName(contentRating)
    if err != nil {
        return false
    }
    
    return userAge >= rating.MinAge
}