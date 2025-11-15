package ratings

import (
    "fmt"
    "math"
    "SDGEStreaming/internal/categories"
    "SDGEStreaming/internal/errors"
    "SDGEStreaming/internal/utils"
)

// Variables para almacenar calificaciones
var (
    contentRatings = make(map[int][]categories.UserRating) // contentID -> []ratings
)

// Califico contenido
func RateContent(contentID, userID int, rating float64) (string, error) {
    // Valido rating
    if rating < 1.0 || rating > 10.0 {
        return "", errors.ErrInvalidRating
    }
    
    // Agrego o actualizo la calificación
    ratings := contentRatings[contentID]
    for i, r := range ratings {
        if r.UserID == userID {
            oldRating := r.Rating
            contentRatings[contentID][i] = categories.UserRating{UserID: userID, Rating: rating}
            
            oldStr := utils.FormatRating(oldRating)
            newStr := utils.FormatRating(rating)
            return fmt.Sprintf("Cambiaste la calificación de %s a %s", oldStr, newStr), nil
        }
    }
    
    // Nueva calificación
    contentRatings[contentID] = append(contentRatings[contentID], categories.UserRating{UserID: userID, Rating: rating})
    
    return "Contenido calificado exitosamente", nil
}

// Obtengo calificaciones para un contenido
func GetRatings(contentID int) ([]categories.UserRating, error) {
    ratings, exists := contentRatings[contentID]
    if !exists {
        return nil, errors.ErrContentNotFound
    }
    return ratings, nil
}

// Obtengo el promedio de calificaciones
func GetAverage(contentID int) (float64, error) {
    ratings, err := GetRatings(contentID)
    if err != nil {
        return 0, err
    }
    
    if len(ratings) == 0 {
        return 0.0, nil
    }
    
    var sum float64
    for _, r := range ratings {
        sum += r.Rating
    }
    
    avg := math.Round(sum/float64(len(ratings))*10) / 10
    return avg, nil
}