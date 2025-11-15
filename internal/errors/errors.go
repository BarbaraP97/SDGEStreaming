package errors

import "fmt"

// Tipos de error personalizados para mejor manejo
type AppError struct {
    Code    string
    Message string
    Details string
}

func (e *AppError) Error() string {
    if e.Details != "" {
        return fmt.Sprintf("%s: %s (%s)", e.Code, e.Message, e.Details)
    }
    return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Errores comunes del sistema
var (
    ErrInvalidEmail       = &AppError{Code: "AUTH_001", Message: "Formato de email inválido"}
    ErrInvalidPassword    = &AppError{Code: "AUTH_002", Message: "Contraseña inválida"}
    ErrUserNotFound       = &AppError{Code: "AUTH_003", Message: "Usuario no encontrado"}
    ErrEmailExists        = &AppError{Code: "AUTH_004", Message: "Email ya registrado"}
    ErrInvalidAge         = &AppError{Code: "USER_001", Message: "Edad inválida"}
    ErrInvalidName        = &AppError{Code: "USER_002", Message: "Nombre inválido"}
    ErrContentNotFound    = &AppError{Code: "CONTENT_001", Message: "Contenido no encontrado"}
    ErrInvalidRating      = &AppError{Code: "RATING_001", Message: "Calificación inválida"}
    ErrInvalidContentID   = &AppError{Code: "CONTENT_002", Message: "ID de contenido inválido"}
    ErrInvalidUserID      = &AppError{Code: "USER_003", Message: "ID de usuario inválido"}
    ErrPermissionDenied   = &AppError{Code: "SEC_001", Message: "Permiso denegado"}
    ErrSessionExpired     = &AppError{Code: "SESSION_001", Message: "Sesión expirada"}
    ErrInputTimeout       = &AppError{Code: "INPUT_001", Message: "Tiempo de espera agotado"}
    ErrInvalidDuration    = &AppError{Code: "CONTENT_003", Message: "Duración inválida"}
    ErrInvalidAgeRating   = &AppError{Code: "CONTENT_004", Message: "Clasificación por edad inválida"}
    ErrInvalidGenre       = &AppError{Code: "CONTENT_005", Message: "Género inválido"}
)

// Creo un nuevo error de aplicación con detalles específicos
func NewAppError(code, message, details string) *AppError {
    return &AppError{Code: code, Message: message, Details: details}
}

// Manejo un error de aplicación y muestro mensaje amigable
func HandleAppError(err error) {
    if appErr, ok := err.(*AppError); ok {
        fmt.Printf("️  %s\n", appErr.Message)
        if appErr.Details != "" {
            fmt.Printf("   Detalles: %s\n", appErr.Details)
        }
    } else {
        fmt.Printf("️  Error inesperado: %v\n", err)
    }
}