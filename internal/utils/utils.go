package utils

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Valido que un email tenga formato correcto mínimo
func IsValidEmail(email string) bool {
    return strings.Contains(email, "@") && strings.Contains(email, ".") && len(email) > 5
}

// Valido que una contraseña cumpla con requisitos mínimos
func IsValidPassword(password string) bool {
    return len(password) >= 6 && len(password) <= 32
}

// Leo entrada segura del usuario con mensaje personalizado
func ReadInput(prompt string) string {
    fmt.Print(prompt)
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        return strings.TrimSpace(scanner.Text())
    }
    return ""
}

// Convierto string a int con manejo de errores
func ToInt(value string) (int, error) {
    return strconv.Atoi(value)
}

// Convierto string a float con soporte para coma y punto
func ToFloat(value string) (float64, error) {
    value = strings.Replace(value, ",", ".", 1)
    return strconv.ParseFloat(value, 64)
}

// Limpio la pantalla de la consola (compatible con múltiples sistemas)
func ClearScreen() {
    fmt.Print("\033[H\033[2J")
}

// Pauso la ejecución hasta que el usuario presione Enter
func WaitForEnter() {
    fmt.Println("\nPresione Enter para continuar...")
    bufio.NewScanner(os.Stdin).Scan()
}

// Valido que una cadena sea numérica
func IsNumeric(s string) bool {
    for _, r := range s {
        if r < '0' || r > '9' {
            return false
        }
    }
    return true
}

// Formateo un tiempo en minutos a formato legible (1h 30min)
func FormatDuration(minutes int) string {
    if minutes < 60 {
        return fmt.Sprintf("%d min", minutes)
    }
    hours := minutes / 60
    remaining := minutes % 60
    if remaining == 0 {
        return fmt.Sprintf("%d h", hours)
    }
    return fmt.Sprintf("%d h %d min", hours, remaining)
}

// Valido que un nombre sea válido (mínimo 2 caracteres, sin números)
func IsValidName(name string) bool {
    if len(name) < 2 {
        return false
    }
    for _, r := range name {
        if (r >= '0' && r <= '9') || r == '@' || r == '#' || r == '$' {
            return false
        }
    }
    return true
}

// Formateo un rating para mostrar con un decimal, excepto 10 que muestro como entero
func FormatRating(rating float64) string {
    if rating == 10.0 {
        return "10"
    }
    return fmt.Sprintf("%.1f", rating)
}
// internal/utils/utils.go
package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ClearScreen clears the terminal screen (LIMPIA PANTALLA).
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// WaitForEnter waits for the user to press Enter.
func WaitForEnter() {
	fmt.Print("Presione Enter para continuar...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// ReadLine prompts the user and reads input from stdin.
func ReadLine(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

// ToInt converts a string to int.
func ToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ToFloat converts a string to float64.
func ToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// Normalize trims spaces and standardizes lowercase.
func Normalize(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

// IsEmpty checks if a string is empty after trimming.
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
