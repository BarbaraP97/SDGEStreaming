<<<<<<< HEAD
=======
// internal/errors/errors.go
>>>>>>> b9e3b62 (AA2_CULMINADO)
package errors

import "fmt"

<<<<<<< HEAD
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
=======
// AppError represents a standardized application error.
// It is used to provide structured and consistent error handling.
type AppError struct {
	Code    string // Unique code for type of error
	Message string // Human-readable message
	Err     error  // Original error (optional)
}

// Error implements the Go error interface.
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Wrap wraps an existing error into an AppError.
func Wrap(code, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// New creates a new AppError without an underlying error.
func New(code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

//
// Common predefined error constructors
//

// ErrNotFound represents a "resource not found" scenario.
func ErrNotFound(resource string) *AppError {
	return New("NOT_FOUND", fmt.Sprintf("%s no encontrado", resource))
}

// ErrInvalidInput represents invalid user input.
func ErrInvalidInput(field string) *AppError {
	return New("INVALID_INPUT", fmt.Sprintf("El campo '%s' no es válido", field))
}

// ErrUnauthorized indicates a missing or invalid auth.
func ErrUnauthorized() *AppError {
	return New("UNAUTHORIZED", "No autorizado")
}

// ErrForbidden indicates lack of permissions.
func ErrForbidden() *AppError {
	return New("FORBIDDEN", "Acceso denegado")
}

// ErrConflict indicates duplicated values (ej: email ya existe).
func ErrConflict(msg string) *AppError {
	return New("CONFLICT", msg)
}

// ErrInternal indicates unexpected server or DB crashes.
func ErrInternal(err error) *AppError {
	return Wrap("INTERNAL_ERROR", "Error interno del servidor", err)
}

// ErrDatabase indicates errors originating from SQL/DB.
func ErrDatabase(err error) *AppError {
	return Wrap("DATABASE_ERROR", "Error de base de datos", err)
}
>>>>>>> b9e3b62 (AA2_CULMINADO)
