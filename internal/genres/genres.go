package genres

import (
    "strings"
    "SDGEStreaming/internal/categories"
    "SDGEStreaming/internal/errors"
)

// Géneros soportados para contenido
var SupportedGenres = []string{
    "Acción", "Comedia", "Drama", "Ciencia Ficción", 
    "Romance", "Terror", "Documental", "Música", 
    "Educación", "Infantil", "Deportes", "Noticias",
}

// Variables globales para géneros
var (
    genres    = make(map[string]categories.Genre)
    nextID    = 1
)

// Inicializo los géneros predeterminados
func init() {
    for _, genreName := range SupportedGenres {
        AddGenre(genreName)
    }
}

// Agrego un nuevo género
func AddGenre(name string) *categories.Genre {
    name = strings.Title(strings.ToLower(name))
    
    newGenre := categories.Genre{
        ID:   nextID,
        Name: name,
    }
    
    genres[name] = newGenre
    nextID++
    return &newGenre
}

// Obtengo un género por nombre
func GetGenreByName(name string) (*categories.Genre, error) {
    name = strings.Title(strings.ToLower(name))
    genre, exists := genres[name]
    if !exists {
        return nil, errors.NewAppError("CONTENT_005", "Género inválido", name)
    }
    return &genre, nil
}

// Obtengo todos los géneros disponibles
func GetAllGenres() []categories.Genre {
    var allGenres []categories.Genre
    for _, genre := range genres {
        allGenres = append(allGenres, genre)
    }
    return allGenres
}

// Valido si un género es soportado
func IsSupportedGenre(genre string) bool {
    _, err := GetGenreByName(genre)
    return err == nil
}

// Filtro géneros por tipo de contenido
func FilterByType(contentType string) []categories.Genre {
    var filtered []categories.Genre
    for _, genre := range GetAllGenres() {
        // En una implementación real, esto dependería del tipo de contenido
        // Por ahora, devolvemos todos los géneros
        filtered = append(filtered, genre)
    }
    return filtered
}