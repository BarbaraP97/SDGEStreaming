package categories

import "time"

// Estructuras comunes para todo el sistema
type UserRating struct {
    UserID int
    Rating float64
}

type ContentRating struct {
    ID          int
    Name        string
    Description string
    MinAge      int
}

type Genre struct {
    ID   int
    Name string
}

type User struct {
    ID          int
    Name        string
    Age         int
    Email       string
    Password    string
    Plan        string
    AgeRating   string
    IsAdmin     bool
    CreatedAt   time.Time    
    LastLogin   time.Time    
    Preferences map[string]string
}