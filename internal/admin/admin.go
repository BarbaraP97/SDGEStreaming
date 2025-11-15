package admin

import (
    "SDGEStreaming/internal/audio"
    "SDGEStreaming/internal/audiovisual"
    "SDGEStreaming/internal/categories"
    "SDGEStreaming/internal/errors"
    "SDGEStreaming/internal/profiles"
)

// Verifico si un usuario tiene permisos de administrador
func IsAdmin(userID int) bool {
    user, err := profiles.FindByID(userID)
    if err != nil {
        return false
    }
    return user.IsAdmin
}

// Obtengo todos los usuarios (solo administradores)
func GetAllUsers(adminUserID int) ([]categories.User, error) {
    if !IsAdmin(adminUserID) {
        return nil, errors.ErrPermissionDenied
    }
    return profiles.GetAllUsers(), nil
}

// Obtengo todo el contenido audiovisual (solo administradores)
func GetAllAudiovisualContent(adminUserID int) ([]audiovisual.AudiovisualContent, error) {
    if !IsAdmin(adminUserID) {
        return nil, errors.ErrPermissionDenied
    }
    return audiovisual.ListAll(), nil
}

// Obtengo todo el contenido de audio (solo administradores)
func GetAllAudioContent(adminUserID int) ([]audio.AudioContent, error) {
    if !IsAdmin(adminUserID) {
        return nil, errors.ErrPermissionDenied
    }
    return audio.ListAll(), nil
}

// Agrego contenido audiovisual (solo administradores)
func AddAudiovisualContent(adminUserID int, title, contentType, genre string, duration int, ageRating string, synopsis string, releaseYear int, director string) error {
    if !IsAdmin(adminUserID) {
        return errors.ErrPermissionDenied
    }
    return audiovisual.AddContent(title, contentType, genre, duration, ageRating, synopsis, releaseYear, director)
}

// Agrego contenido de audio (solo administradores)
func AddAudioContent(adminUserID int, title, contentType, genre string, duration int, ageRating string, artist string, album string, trackNumber int) error {
    if !IsAdmin(adminUserID) {
        return errors.ErrPermissionDenied
    }
    return audio.AddContent(title, contentType, genre, duration, ageRating, artist, album, trackNumber)
}

// Obtengo calificaciones individuales para contenido audiovisual
func GetAudiovisualIndividualRatings(adminUserID, contentID int) ([]categories.UserRating, error) {
    if !IsAdmin(adminUserID) {
        return nil, errors.ErrPermissionDenied
    }
    return audiovisual.GetIndividualRatings(contentID)
}

// Obtengo calificaciones individuales para contenido de audio
func GetAudioIndividualRatings(adminUserID, contentID int) ([]categories.UserRating, error) {
    if !IsAdmin(adminUserID) {
        return nil, errors.ErrPermissionDenied
    }
    return audio.GetIndividualRatings(contentID)
}