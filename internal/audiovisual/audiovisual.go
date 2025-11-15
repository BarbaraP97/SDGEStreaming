package audiovisual

import (
    "SDGEStreaming/internal/categories"
    "SDGEStreaming/internal/contentclass"
    "SDGEStreaming/internal/errors"
    "SDGEStreaming/internal/genres"
    "SDGEStreaming/internal/ratings"
)

// Estructura para contenido audiovisual (películas, series, documentales)
type AudiovisualContent struct {
    ID            int
    Title         string
    Type          string // "Película", "Serie", "Documental"
    Genre         string
    Duration      int    // en minutos
    AgeRating     string // "Infantil", "Adolescente", "Adulto"
    Synopsis      string
    ReleaseYear   int
    Director      string
    AverageRating float64
    IsAvailable   bool
}

// Variables globales para almacenamiento en memoria
var (
    contents []AudiovisualContent
    nextID   = 1
)

// Inicializo contenido audiovisual de ejemplo
func init() {
    AddContent("El Viaje Infinito", "Película", "Ciencia Ficción", 120, "Adolescente", "Una aventura épica por el espacio", 2024, "Director X")
    AddContent("Misterios del Océano", "Documental", "Documental", 90, "Infantil", "Descubre los secretos del mar", 2023, "Documentalista Y")
    AddContent("Risas en la Ciudad", "Serie", "Comedia", 45, "Adolescente", "Comedia sobre la vida urbana", 2024, "Creador Z")
}

// Agrego nuevo contenido audiovisual
func AddContent(title, contentType, genre string, duration int, ageRating string, synopsis string, releaseYear int, director string) error {
    // Valido el tipo de contenido
    validTypes := map[string]bool{"Película": true, "Serie": true, "Documental": true}
    if !validTypes[contentType] {
        return errors.NewAppError("CONTENT_006", "Tipo de contenido audiovisual inválido", contentType)
    }
    
    // Valido género
    if !genres.IsSupportedGenre(genre) {
        return errors.NewAppError("CONTENT_005", "Género inválido", genre)
    }
    
    // Valido duración
    if duration <= 0 {
        return errors.ErrInvalidDuration
    }
    
    // Valido clasificación por edad
    if _, err := contentclass.GetRatingByName(ageRating); err != nil {
        return err
    }
    
    // Creo el nuevo contenido
    newContent := AudiovisualContent{
        ID:            nextID,
        Title:         title,
        Type:          contentType,
        Genre:         genre,
        Duration:      duration,
        AgeRating:     ageRating,
        Synopsis:      synopsis,
        ReleaseYear:   releaseYear,
        Director:      director,
        AverageRating: 0.0,
        IsAvailable:   true,
    }
    
    contents = append(contents, newContent)
    nextID++
    return nil
}

// Listo todo el contenido audiovisual disponible
func ListAll() []AudiovisualContent {
    var availableContents []AudiovisualContent
    for _, c := range contents {
        if c.IsAvailable {
            availableContents = append(availableContents, c)
        }
    }
    return availableContents
}

// Obtengo contenido por ID
func GetByID(id int) (*AudiovisualContent, error) {
    for i, c := range contents {
        if c.ID == id {
            return &contents[i], nil
        }
    }
    return nil, errors.ErrContentNotFound
}

// Califico un contenido audiovisual
func RateContent(contentID, userID int, rating float64) (string, error) {
    content, err := GetByID(contentID)
    if err != nil {
        return "", err
    }
    
    if rating < 1.0 || rating > 10.0 {
        return "", errors.ErrInvalidRating
    }
    
    // Uso el módulo de ratings para manejar la calificación
    message, err := ratings.RateContent(contentID, userID, rating)
    if err != nil {
        return "", err
    }
    
    // Recalculo el promedio
    avg, _ := ratings.GetAverage(contentID)
    content.AverageRating = avg
    
    return message, nil
}

// Obtengo las calificaciones individuales de un contenido
func GetIndividualRatings(contentID int) ([]categories.UserRating, error) {
    return ratings.GetRatings(contentID)
}

// Filtro contenido por tipo
func FilterByType(contentType string) []AudiovisualContent {
    var filtered []AudiovisualContent
    for _, c := range contents {
        if c.Type == contentType && c.IsAvailable {
            filtered = append(filtered, c)
        }
    }
    return filtered
}

// Filtro contenido por género
func FilterByGenre(genre string) []AudiovisualContent {
    var filtered []AudiovisualContent
    for _, c := range contents {
        if c.Genre == genre && c.IsAvailable {
            filtered = append(filtered, c)
        }
    }
    return filtered
}

// Filtro contenido por clasificación de edad
func FilterByAgeRating(ageRating string) []AudiovisualContent {
    var filtered []AudiovisualContent
    for _, c := range contents {
        if c.AgeRating == ageRating && c.IsAvailable {
            filtered = append(filtered, c)
        }
    }
    return filtered
}