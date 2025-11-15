package profiles

import (
    "time"
    "SDGEStreaming/internal/categories"
    "SDGEStreaming/internal/errors"
    "SDGEStreaming/internal/utils"
)

// Variables globales para almacenamiento en memoria
var (
    users  = make(map[int]categories.User)
    nextID = 1
)

// Inicializo usuarios predeterminados para pruebas
func init() {
    // Usuario administrador
    AddUser("Administrador", 35, "admin@sdge.com", "admin123", "Premium", "Adulto", true)
    // Usuario de ejemplo
    AddUser("Usuario Demo", 28, "user@demo.com", "demo123", "Free", "Adulto", false)
}

// Agrego un nuevo usuario al sistema
func AddUser(name string, age int, email string, password string, plan string, ageRating string, isAdmin bool) (*categories.User, error) {
    // Valido datos de entrada
    if !utils.IsValidName(name) {
        return nil, errors.ErrInvalidName
    }
    
    if age < 13 || age > 120 {
        return nil, errors.NewAppError("USER_001", "Edad inválida", "Debe estar entre 13 y 120 años")
    }
    
    if !utils.IsValidEmail(email) {
        return nil, errors.ErrInvalidEmail
    }
    
    if !utils.IsValidPassword(password) {
        return nil, errors.ErrInvalidPassword
    }
    
    // Verifico que el email no exista
    for _, u := range users {
        if u.Email == email {
            return nil, errors.ErrEmailExists
        }
    }
    
    // Creo el nuevo usuario
    newUser := categories.User{
        ID:          nextID,
        Name:        name,
        Age:         age,
        Email:       email,
        Password:    password,
        Plan:        plan,
        AgeRating:   ageRating,
        CreatedAt:   time.Now(),
        LastLogin:   time.Now(),
        Preferences: make(map[string]string),
        IsAdmin:     isAdmin,
    }
    
    users[nextID] = newUser
    nextID++
    return &newUser, nil
}

// Busco un usuario por email
func FindByEmail(email string) (*categories.User, error) {
    for _, u := range users {
        if u.Email == email {
            return &u, nil
        }
    }
    return nil, errors.ErrUserNotFound
}

// Busco un usuario por ID
func FindByID(id int) (*categories.User, error) {
    user, exists := users[id]
    if !exists {
        return nil, errors.ErrInvalidUserID
    }
    return &user, nil
}

// Obtengo todos los usuarios (solo para administradores)
func GetAllUsers() []categories.User {
    var allUsers []categories.User
    for _, user := range users {
        allUsers = append(allUsers, user)
    }
    return allUsers
}

// Actualizo las preferencias de un usuario
func UpdatePreferences(userID int, key, value string) error {
    user, err := FindByID(userID)
    if err != nil {
        return err
    }
    
    user.Preferences[key] = value
    users[userID] = *user
    return nil
}

// Actualizo el último inicio de sesión
func UpdateLastLogin(userID int) error {
    user, err := FindByID(userID)
    if err != nil {
        return err
    }
    
    user.LastLogin = time.Now()
    users[userID] = *user
    return nil
}