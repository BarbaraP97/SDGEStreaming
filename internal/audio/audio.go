package audio

import (
    "SDGEStreaming/internal/categories"
    "SDGEStreaming/internal/contentclass"
    "SDGEStreaming/internal/errors"
    "SDGEStreaming/internal/genres"
    "SDGEStreaming/internal/ratings"
)

// Estructura para contenido de audio (música, podcasts, audiolibros)
type AudioContent struct {
    ID            int
    Title         string
    Type          string // "Música", "Podcast", "Audiolibro"
    Genre         string
    Duration      int    // en minutos
    AgeRating     string // "Infantil", "Adolescente", "Adulto"
    Artist        string
    Album         string
    TrackNumber   int
    AverageRating float64
    IsAvailable   bool
}

// Variables globales para almacenamiento en memoria
var (
    contents []AudioContent
    nextID   = 1
)

// Inicializo contenido de audio de ejemplo
func init() {
    AddContent("Sinfonía del Amanecer", "Música", "Clásica", 15, "Adolescente", "Orquesta Sinfónica", "Clásicos Eternos", 1)
    AddContent("Tecnología Hoy", "Podcast", "Tecnología", 45, "Adolescente", "Podcaster Tech", "Episodios Tech", 5)
    AddContent("Cuentos de la Noche", "Audiolibro", "Infantil", 30, "Infantil", "Narrador Infantil", "Colección Noche", 1)
}

// Agrego nuevo contenido de audio
func AddContent(title, contentType, genre string, duration int, ageRating string, artist string, album string, trackNumber int) error {
    // Valido el tipo de contenido
    validTypes := map[string]bool{"Música": true, "Podcast": true, "Audiolibro": true}
    if !validTypes[contentType] {
        return errors.NewAppError("CONTENT_007", "Tipo de contenido de audio inválido", contentType)
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
    newContent := AudioContent{
        ID:            nextID,
        Title:         title,
        Type:          contentType,
        Genre:         genre,
        Duration:      duration,
        AgeRating:     ageRating,
        Artist:        artist,
        Album:         album,
        TrackNumber:   trackNumber,
        AverageRating: 0.0,
        IsAvailable:   true,
    }
    
    contents = append(contents, newContent)
    nextID++
    return nil
}

// Listo todo el contenido de audio disponible
func ListAll() []AudioContent {
    var availableContents []AudioContent
    for _, c := range contents {
        if c.IsAvailable {
            availableContents = append(availableContents, c)
        }
    }
    return availableContents
}

// Obtengo contenido por ID
func GetByID(id int) (*AudioContent, error) {
    for i, c := range contents {
        if c.ID == id {
            return &contents[i], nil
        }
    }
    return nil, errors.ErrContentNotFound
}

// Califico un contenido de audio
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
func FilterByType(contentType string) []AudioContent {
    var filtered []AudioContent
    for _, c := range contents {
        if c.Type == contentType && c.IsAvailable {
            filtered = append(filtered, c)
        }
    }
    return filtered
}

// Filtro contenido por género
func FilterByGenre(genre string) []AudioContent {
    var filtered []AudioContent
    for _, c := range contents {
        if c.Genre == genre && c.IsAvailable {
            filtered = append(filtered, c)
        }
    }
    return filtered
}

// Filtro contenido por clasificación de edad
func FilterByAgeRating(ageRating string) []AudioContent {
    var filtered []AudioContent
    for _, c := range contents {
        if c.AgeRating == ageRating && c.IsAvailable {
            filtered = append(filtered, c)
        }
    }
    return filtered
}