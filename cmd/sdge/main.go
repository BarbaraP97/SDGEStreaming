package main

import (
	"SDGEStreaming/internal/admin"
	"SDGEStreaming/internal/audio"
	"SDGEStreaming/internal/audiovisual"
	"SDGEStreaming/internal/categories"
	"SDGEStreaming/internal/contentclass"
	"SDGEStreaming/internal/errors"
	"SDGEStreaming/internal/history"
	"SDGEStreaming/internal/profiles"
	"SDGEStreaming/internal/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Variables globales para la sesión
var (
	currentUser      *categories.User
	currentSessionID string
	lastActivity     time.Time
	sessionTimeout   = 5 * time.Minute
)

func main() {
	fmt.Print("\033[H\033[2J") // Limpiar pantalla

	lastActivity = time.Now()

	for {
		// Verificar expiración de sesión
		if currentUser != nil && time.Since(lastActivity) > sessionTimeout {
			fmt.Println("Sesión expirada por inactividad. Por favor inicie sesión nuevamente.")
			currentUser = nil
			waitForEnter()
			continue
		}

		if currentUser == nil {
			showAuthMenu()
		} else {
			showMainMenu()
		}
	}
}

func showHeader() {
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║ SDGEStreaming Versión 1.0.0-AA2 Paso 1                 ║")
	fmt.Println("║ Sistema de Gestión de Contenido Audiovisual y Audio    ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func waitForEnter() {
	fmt.Println("Presione Enter para continuar...")
	bufio.NewScanner(os.Stdin).Scan()
}

// Mostrar menú de autenticación
func showAuthMenu() {
	fmt.Print("\033[H\033[2J") // Limpiar pantalla
	showHeader()

	fmt.Println("Bienvenido a SDGEStreaming")
	fmt.Println("════════════════════════════")
	fmt.Println()
	fmt.Println("1. Iniciar Sesión")
	fmt.Println("2. Registrarse")
	fmt.Println("3. Explorar como Invitado")
	fmt.Println("4. Salir")
	fmt.Println("────────────────────────────────────────────────────────────")

	option := readInput("Seleccione una opción: ")

	switch option {
	case "1":
		login()
	case "2":
		register()
	case "3":
		currentUser = nil
		showContentMenu(true)
	case "4":
		fmt.Print("\033[H\033[2J")
		fmt.Println("Gracias por usar SDGEStreaming. ¡Hasta luego!")
		os.Exit(0)
	default:
		if option != "" {
			fmt.Println("Opción inválida. Por favor seleccione una opción del menú.")
			waitForEnter()
		}
	}
}

// Iniciar sesión
func login() {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Iniciar Sesión")
	fmt.Println("═══════════════")

	email := readInput("Email: ")
	if email == "0" {
		return
	}

	password := readInput("Contraseña: ")
	if password == "0" {
		return
	}

	user, err := profiles.FindByEmail(email)
	if err != nil {
		fmt.Println("✗ Usuario no encontrado")
		waitForEnter()
		return
	}

	if user.Password != password {
		fmt.Println("✗ Contraseña incorrecta")
		waitForEnter()
		return
	}

	profiles.UpdateLastLogin(user.ID)
	currentUser = user
	currentSessionID = fmt.Sprintf("sess_%d_%d", user.ID, time.Now().Unix())
	lastActivity = time.Now()

	fmt.Printf(" ¡Bienvenido, %s!\n", user.Name)
	waitForEnter()
}

// Registrar nuevo usuario
func register() {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Registro de Nuevo Usuario")
	fmt.Println("═════════════════════════")

	name := readInput("Nombre completo: ")
	if name == "0" {
		return
	}

	ageStr := readInput("Edad: ")
	if ageStr == "0" {
		return
	}
	age, err := strconv.Atoi(ageStr)
	if err != nil || age < 13 || age > 120 {
		fmt.Println("Edad inválida")
		waitForEnter()
		return
	}

	email := readInput("Email: ")
	if email == "0" {
		return
	}

	password := readInput("Contraseña (6+ caracteres): ")
	if password == "0" {
		return
	}

	if len(password) < 6 {
		fmt.Println("Contraseña muy corta")
		waitForEnter()
		return
	}

	// Mostrar clasificaciones
	fmt.Println()
	fmt.Println("Clasificación por Edad")
	fmt.Println("───────────────────────")
	ratings := contentclass.GetAllRatings()
	for i, r := range ratings {
		fmt.Printf("%d. %s - %s\n", i+1, r.Name, r.Description)
	}

	ratingStr := readInput("Seleccione su clasificación (1-3): ")
	if ratingStr == "0" {
		return
	}
	ratingNum, err := strconv.Atoi(ratingStr)
	if err != nil || ratingNum < 1 || ratingNum > len(ratings) {
		fmt.Println("Opción inválida")
		waitForEnter()
		return
	}

	ageRating := ratings[ratingNum-1].Name

	_, err = profiles.AddUser(name, age, email, password, "Free", ageRating, false)
	if err != nil {
		errors.HandleAppError(err)
		waitForEnter()
		return
	}

	fmt.Println(" Usuario registrado exitosamente")
	waitForEnter()
}

// Mostrar menú principal
func showMainMenu() {
	fmt.Print("\033[H\033[2J") // Limpiar pantalla
	showHeader()

	fmt.Printf("Menú Principal - %s (%s)\n", currentUser.Name, currentUser.Plan)
	fmt.Println("════════════════════════════════")
	fmt.Println()

	if currentUser.IsAdmin {
		fmt.Println("1. Mi Perfil")
		fmt.Println("2. Explorar Contenido")
		fmt.Println("3. Gestionar Usuarios")
		fmt.Println("4. Gestionar Contenido Audiovisual")
		fmt.Println("5. Gestionar Contenido de Audio")
		fmt.Println("6. Cerrar Sesión")
		fmt.Println("7. Salir")
	} else {
		fmt.Println("1. Mi Perfil")
		fmt.Println("2. Explorar Contenido")
		fmt.Println("3. Mi Lista")
		fmt.Println("4. Historial de Reproducción")
		fmt.Println("5. Configuraciones")
		fmt.Println("6. Cerrar Sesión")
		fmt.Println("7. Salir")
	}

	fmt.Println("────────────────────────────────────────────────────────────")

	option := readInput("Seleccione una opción: ")

	switch option {
	case "1":
		showUserProfile()
	case "2":
		showContentMenu(false)
	case "3":
		if currentUser.IsAdmin {
			showUserManagement()
		} else {
			showMyList()
		}
	case "4":
		if currentUser.IsAdmin {
			showAudiovisualManagement()
		} else {
			showPlaybackHistory()
		}
	case "5":
		if currentUser.IsAdmin {
			showAudioManagement()
		} else {
			fmt.Println("Funcionalidad para AA2")
			waitForEnter()
		}
	case "6":
		currentUser = nil
		fmt.Println("Sesión cerrada")
		waitForEnter()
	case "7":
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Hasta luego, %s\n", currentUser.Name)
		os.Exit(0)
	default:
		if option != "" {
			fmt.Println("Opción inválida")
			waitForEnter()
		}
	}
}
// Mostrar historial de reproducción
func showPlaybackHistory() {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Historial de Reproducción")
	fmt.Println("════════════════════════")

	entries := history.GetHistory(currentUser.ID)
	if len(entries) == 0 {
		fmt.Println("No tienes historial de reproducción.")
	} else {
for _, e := range entries {
			timestamp := e.Timestamp.Format("02/01/2006 15:04")
			if e.ContentType == "audiovisual" {
				c, err := audiovisual.GetByID(e.ContentID)
				if err != nil {
					fmt.Printf("• [Audiovisual] ID %d - NO DISPONIBLE (%s)\n", e.ContentID, timestamp)
					continue
				}
				if !contentclass.CanAccessContent(currentUser.Age, c.AgeRating) {
					fmt.Printf("• [Audiovisual] ID %d - RESTRINGIDO (%s)\n", e.ContentID, timestamp)
					continue
				}
				fmt.Printf("• [Audiovisual] %s (%s)\n", c.Title, timestamp)
				fmt.Printf("   ID: %d | Tipo: %s | Género: %s\n", c.ID, c.Type, c.Genre)
				fmt.Printf("   Duración: %s | Año: %d | Director: %s\n", utils.FormatDuration(c.Duration), c.ReleaseYear, c.Director)
				fmt.Printf("   Clasificación: %s | Rating: %s\n", c.AgeRating, utils.FormatRating(c.AverageRating))
				fmt.Println("────────────────────────────────────────────────────────────")
			} else if e.ContentType == "audio" {
				c, err := audio.GetByID(e.ContentID)
				if err != nil {
					fmt.Printf("• [Audio] ID %d - NO DISPONIBLE (%s)\n", e.ContentID, timestamp)
					continue
				}
				if !contentclass.CanAccessContent(currentUser.Age, c.AgeRating) {
					fmt.Printf("• [Audio] ID %d - RESTRINGIDO (%s)\n", e.ContentID, timestamp)
					continue
				}
				fmt.Printf("• [Audio] %s (%s)\n", c.Title, timestamp)
				fmt.Printf("   ID: %d | Tipo: %s | Género: %s\n", c.ID, c.Type, c.Genre)
				fmt.Printf("   Duración: %s | Artista: %s | Álbum: %s\n", utils.FormatDuration(c.Duration), c.Artist, c.Album)
				fmt.Printf("   Clasificación: %s | Rating: %s\n", c.AgeRating, utils.FormatRating(c.AverageRating))
				fmt.Println("────────────────────────────────────────────────────────────")
			}
		}
	}

	fmt.Println("────────────────────────────────────────────────────────────")
	waitForEnter()
}
// Mostrar "Mi Lista" (favoritos)
func showMyList() {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Mi Lista (Favoritos)")
	fmt.Println("═══════════════════")

	favIDs := profiles.GetFavorites(currentUser.ID)
	if len(favIDs) == 0 {
		fmt.Println("No tienes contenido en tu lista.")
	} else {
		favorites := profiles.GetFavorites(currentUser.ID)
		if len(favorites) == 0 {
			fmt.Println("No tienes contenido en tu lista.")
		} else {
			for _, fav := range favorites {
				if fav.ContentType == "audiovisual" {
					c, err := audiovisual.GetByID(fav.ContentID)
					if err != nil {
						fmt.Printf("★ [Audiovisual] ID %d - NO DISPONIBLE\n", fav.ContentID)
						continue
					}
					if !contentclass.CanAccessContent(currentUser.Age, c.AgeRating) {
						fmt.Printf("★ [Audiovisual] ID %d - RESTRINGIDO\n", fav.ContentID)
						continue
					}
					fmt.Printf("★ [Audiovisual] %s\n", c.Title)
					fmt.Printf("   ID: %d | Tipo: %s | Género: %s\n", c.ID, c.Type, c.Genre)
					fmt.Printf("   Duración: %s | Año: %d | Director: %s\n", utils.FormatDuration(c.Duration), c.ReleaseYear, c.Director)
					fmt.Printf("   Clasificación: %s | Rating: %s\n", c.AgeRating, utils.FormatRating(c.AverageRating))
					fmt.Println("────────────────────────────────────────────────────────────")
				} else if fav.ContentType == "audio" {
					c, err := audio.GetByID(fav.ContentID)
					if err != nil {
						fmt.Printf("★ [Audio] ID %d - NO DISPONIBLE\n", fav.ContentID)
						continue
					}
					if !contentclass.CanAccessContent(currentUser.Age, c.AgeRating) {
						fmt.Printf("★ [Audio] ID %d - RESTRINGIDO\n", fav.ContentID)
						continue
					}
					fmt.Printf("★ [Audio] %s\n", c.Title)
					fmt.Printf("   ID: %d | Tipo: %s | Género: %s\n", c.ID, c.Type, c.Genre)
					fmt.Printf("   Duración: %s | Artista: %s | Álbum: %s\n", utils.FormatDuration(c.Duration), c.Artist, c.Album)
					fmt.Printf("   Clasificación: %s | Rating: %s\n", c.AgeRating, utils.FormatRating(c.AverageRating))
					fmt.Println("────────────────────────────────────────────────────────────")
				}
			}
		}
	}

	fmt.Println("\n1. Añadir contenido por ID")
	fmt.Println("2. Volver")
	fmt.Println("────────────────────────────────────────────────────────────")

	option := readInput("Opción: ")
	if option == "1" {
		fmt.Println("\n¿Qué tipo de contenido deseas añadir?")
		fmt.Println("1. Audiovisual")
		fmt.Println("2. Audio")
		typeOption := readInput("Seleccione (1 o 2): ")

		if typeOption == "1" {
			// Mostrar audiovisual
			contents := audiovisual.ListAll()
			if len(contents) == 0 {
				fmt.Println("\nNo hay contenido audiovisual disponible.")
				waitForEnter()
				return
			}
			fmt.Println("\nContenido Audiovisual Disponible:")
			fmt.Println("═════════════════════════════════")
			for _, c := range contents {
				if contentclass.CanAccessContent(currentUser.Age, c.AgeRating) {
					fmt.Printf("ID: %d | %s • %s\n", c.ID, c.Title, utils.FormatDuration(c.Duration))
				}
			}

			contentIDStr := readInput("\nIngrese el ID del contenido a añadir: ")
			contentID, err := strconv.Atoi(contentIDStr)
			if err != nil || contentID <= 0 {
				fmt.Println("ID inválido.")
				waitForEnter()
				return
			}

			_, err = audiovisual.GetByID(contentID)
			if err != nil {
				fmt.Println("Contenido no encontrado.")
				waitForEnter()
				return
			}

			profiles.AddFavorite(currentUser.ID, contentID, "audiovisual")
			fmt.Println("\n Audiovisual añadido a Mi Lista.")
			waitForEnter()

		} else if typeOption == "2" {
			// Mostrar audio
			contents := audio.ListAll()
			if len(contents) == 0 {
				fmt.Println("\nNo hay contenido de audio disponible.")
				waitForEnter()
				return
			}
			fmt.Println("\nContenido de Audio Disponible:")
			fmt.Println("═════════════════════════════")
			for _, c := range contents {
				if contentclass.CanAccessContent(currentUser.Age, c.AgeRating) {
					fmt.Printf("ID: %d | %s • %s\n", c.ID, c.Title, utils.FormatDuration(c.Duration))
				}
			}

			contentIDStr := readInput("\nIngrese el ID del contenido a añadir: ")
			contentID, err := strconv.Atoi(contentIDStr)
			if err != nil || contentID <= 0 {
				fmt.Println("ID inválido.")
				waitForEnter()
				return
			}

			_, err = audio.GetByID(contentID)
			if err != nil {
				fmt.Println("Contenido no encontrado.")
				waitForEnter()
				return
			}

			profiles.AddFavorite(currentUser.ID, contentID, "audio")
			fmt.Println("\n Contenido de audio añadido a Mi Lista.")
			waitForEnter()

		} else {
			fmt.Println("Opción inválida.")
			waitForEnter()
		}
	}
}
// Mostrar perfil de usuario
func showUserProfile() {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Mi Perfil")
	fmt.Println("═════════")

	fmt.Printf("Nombre: %s\n", currentUser.Name)
	fmt.Printf("Email: %s\n", currentUser.Email)
	fmt.Printf("Plan: %s\n", currentUser.Plan)
	fmt.Printf("Edad: %d años\n", currentUser.Age)
	fmt.Printf("Clasificación: %s\n", currentUser.AgeRating)
	fmt.Printf("Último acceso: %s\n", currentUser.LastLogin.Format("02/01/2006 15:04"))

	fmt.Println("────────────────────────────────────────────────────────────")
	waitForEnter()
}

// Mostrar menú de contenido
func showContentMenu(isGuest bool) {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Explorar Contenido")
	fmt.Println("══════════════════")
	fmt.Println()
	fmt.Println("1. Contenido Audiovisual")
	fmt.Println("2. Contenido de Audio")
	fmt.Println("3. Volver al Menú Principal")
	fmt.Println("────────────────────────────────────────────────────────────")

	option := readInput("Seleccione una opción: ")

	switch option {
	case "1":
		showAudiovisualContent(isGuest)
	case "2":
		showAudioContent(isGuest)
	case "3":
		return
	default:
		if option != "" {
			fmt.Println("Opción inválida")
			waitForEnter()
		}
	}
}

// Mostrar contenido audiovisual
func showAudiovisualContent(isGuest bool) {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Contenido Audiovisual")
	fmt.Println("═════════════════════")

	contents := audiovisual.ListAll()
	if len(contents) == 0 {
		fmt.Println("No hay contenido disponible")
		waitForEnter()
		return
	}

	for _, c := range contents {
		// Verificar clasificación
		if !isGuest && !contentclass.CanAccessContent(currentUser.Age, c.AgeRating) {
			continue
		}

		fmt.Printf("ID: %d | %s\n", c.ID, c.Title)
		fmt.Printf("   %s • %s • %s\n", c.Type, c.Genre, utils.FormatDuration(c.Duration))
		fmt.Printf("   Clasificación: %s • Rating: %s\n", c.AgeRating, utils.FormatRating(c.AverageRating))
		fmt.Println("────────────────────────────────────────────────────────────")
	}

	if !isGuest {
		contentIDStr := readInput("ID para calificar (0 para volver): ")
		if contentIDStr != "0" {
			contentID, err := strconv.Atoi(contentIDStr)
			if err == nil && contentID > 0 {
				// Primero verificamos que el contenido exista y sea accesible
				c, err := audiovisual.GetByID(contentID)
				if err != nil {
					fmt.Println("Contenido no encontrado.")
					waitForEnter()
					return
				}
				if !contentclass.CanAccessContent(currentUser.Age, c.AgeRating) {
					fmt.Println("No tienes acceso a este contenido por tu clasificación de edad.")
					waitForEnter()
					return
				}
				// Solo si es válido, registramos y calificamos
				history.AddPlayback(currentUser.ID, contentID, "audiovisual")
				rateAudiovisualContent(contentID)
			} else {
				fmt.Println("ID inválido.")
				waitForEnter()
			}
		}
	} else {
		waitForEnter()
	}
}

// Mostrar contenido de audio
func showAudioContent(isGuest bool) {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Contenido de Audio")
	fmt.Println("══════════════════")

	contents := audio.ListAll()
	if len(contents) == 0 {
		fmt.Println("No hay contenido disponible")
		waitForEnter()
		return
	}

	for _, c := range contents {
		// Verificar clasificación
		if !isGuest && !contentclass.CanAccessContent(currentUser.Age, c.AgeRating) {
			continue
		}

		fmt.Printf("ID: %d | %s\n", c.ID, c.Title)
		fmt.Printf("   %s • %s • %s\n", c.Type, c.Genre, utils.FormatDuration(c.Duration))
		fmt.Printf("   Clasificación: %s • Rating: %s\n", c.AgeRating, utils.FormatRating(c.AverageRating))
		fmt.Println("────────────────────────────────────────────────────────────")
	}

	if !isGuest {
		contentIDStr := readInput("ID para calificar (0 para volver): ")
		if contentIDStr != "0" {
			contentID, err := strconv.Atoi(contentIDStr)
			if err == nil && contentID > 0 {
				// Verificar existencia y acceso
				c, err := audio.GetByID(contentID)
				if err != nil {
					fmt.Println("Contenido no encontrado.")
					waitForEnter()
					return
				}
				if !contentclass.CanAccessContent(currentUser.Age, c.AgeRating) {
					fmt.Println("No tienes acceso a este contenido por tu clasificación de edad.")
					waitForEnter()
					return
				}
				// Registrar y calificar
				history.AddPlayback(currentUser.ID, contentID, "audio")
				rateAudioContent(contentID)
			} else {
				fmt.Println("ID inválido.")
				waitForEnter()
			}
		}
	} else {
		waitForEnter()
	}
}

// Calificar contenido audiovisual
func rateAudiovisualContent(contentID int) {
	c, err := audiovisual.GetByID(contentID)
	if err != nil {
		fmt.Println("Contenido no encontrado")
		waitForEnter()
		return
	}

	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Printf("Calificar: %s\n", c.Title)
	fmt.Println("══════════════")

	ratingStr := readInput("Calificación (1.0 - 10.0): ")
	rating, err := utils.ToFloat(ratingStr)
	if err != nil || rating < 1.0 || rating > 10.0 {
		fmt.Println("Calificación inválida")
		waitForEnter()
		return
	}

	message, err := audiovisual.RateContent(contentID, currentUser.ID, rating)
	if err != nil {
		fmt.Println("Error al calificar")
	} else {
		fmt.Printf(" %s\n", message)
	}
	waitForEnter()
}

// Calificar contenido de audio
func rateAudioContent(contentID int) {
	c, err := audio.GetByID(contentID)
	if err != nil {
		fmt.Println("Contenido no encontrado")
		waitForEnter()
		return
	}

	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Printf("Calificar: %s\n", c.Title)
	fmt.Println("══════════════")

	ratingStr := readInput("Calificación (1.0 - 10.0): ")
	rating, err := utils.ToFloat(ratingStr)
	if err != nil || rating < 1.0 || rating > 10.0 {
		fmt.Println("Calificación inválida")
		waitForEnter()
		return
	}

	message, err := audio.RateContent(contentID, currentUser.ID, rating)
	if err != nil {
		fmt.Println("Error al calificar")
	} else {
		fmt.Printf(" %s\n", message)
	}
	waitForEnter()
}

// Gestión de usuarios (admin)
func showUserManagement() {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Gestión de Usuarios")
	fmt.Println("═══════════════════")

	users, err := admin.GetAllUsers(currentUser.ID)
	if err != nil {
		fmt.Println("No tienes permisos")
		waitForEnter()
		return
	}

	for _, u := range users {
		adminTag := ""
		if u.IsAdmin {
			adminTag = " [ADMIN]"
		}
		fmt.Printf("ID: %d | %s%s\n", u.ID, u.Name, adminTag)
		fmt.Printf("   %s • %d años • %s\n", u.Email, u.Age, u.Plan)
		fmt.Println("────────────────────────────────────────────────────────────")
	}

	waitForEnter()
}

// Gestión de contenido audiovisual (admin)
func showAudiovisualManagement() {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Gestión de Contenido Audiovisual")
	fmt.Println("═══════════════════════════════")

	fmt.Println("1. Listar Contenido")
	fmt.Println("2. Agregar Contenido")
	fmt.Println("3. Volver al Menú Principal")
	fmt.Println("────────────────────────────────────────────────────────────")

	option := readInput("Seleccione una opción: ")

	switch option {
	case "1":
		showAudiovisualContent(false)
	case "2":
		addAudiovisualContent()
	case "3":
		return
	default:
		if option != "" {
			fmt.Println("Opción inválida")
			waitForEnter()
		}
	}
}

// Gestión de contenido de audio (admin)
func showAudioManagement() {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Gestión de Contenido de Audio")
	fmt.Println("═════════════════════════════")

	fmt.Println("1. Listar Contenido")
	fmt.Println("2. Agregar Contenido")
	fmt.Println("3. Volver al Menú Principal")
	fmt.Println("────────────────────────────────────────────────────────────")

	option := readInput("Seleccione una opción: ")

	switch option {
	case "1":
		showAudioContent(false)
	case "2":
		addAudioContent()
	case "3":
		return
	default:
		if option != "" {
			fmt.Println("Opción inválida")
			waitForEnter()
		}
	}
}

// Agregar contenido audiovisual
func addAudiovisualContent() {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Agregar Contenido Audiovisual")
	fmt.Println("════════════════════════════")

	title := readInput("Título: ")
	if title == "0" {
		return
	}

	fmt.Println("Tipos: 1. Película  2. Serie  3. Documental")
	typeStr := readInput("Tipo (1-3): ")
	if typeStr == "0" {
		return
	}

	typeNum, err := strconv.Atoi(typeStr)
	if err != nil || typeNum < 1 || typeNum > 3 {
		fmt.Println("Tipo inválido")
		waitForEnter()
		return
	}

	contentTypes := []string{"Película", "Serie", "Documental"}
	contentType := contentTypes[typeNum-1]

	durationStr := readInput("Duración (minutos): ")
	if durationStr == "0" {
		return
	}
	duration, err := strconv.Atoi(durationStr)
	if err != nil || duration <= 0 {
		fmt.Println("Duración inválida")
		waitForEnter()
		return
	}

	// Clasificaciones
	fmt.Println("Clasificaciones:")
	ratings := contentclass.GetAllRatings()
	for i, r := range ratings {
		fmt.Printf("%d. %s\n", i+1, r.Name)
	}

	ratingStr := readInput("Clasificación (1-3): ")
	if ratingStr == "0" {
		return
	}
	ratingNum, err := strconv.Atoi(ratingStr)
	if err != nil || ratingNum < 1 || ratingNum > len(ratings) {
		fmt.Println("Clasificación inválida")
		waitForEnter()
		return
	}

	ageRating := ratings[ratingNum-1].Name

	err = audiovisual.AddContent(title, contentType, "Acción", duration, ageRating, "Sinopsis", 2024, "Director")
	if err != nil {
		fmt.Println("Error al agregar contenido")
	} else {
		fmt.Println(" Contenido agregado")
	}
	waitForEnter()
}

// Agregar contenido de audio
func addAudioContent() {
	fmt.Print("\033[H\033[2J")
	showHeader()
	fmt.Println("Agregar Contenido de Audio")
	fmt.Println("═════════════════════════")

	title := readInput("Título: ")
	if title == "0" {
		return
	}

	fmt.Println("Tipos: 1. Música  2. Podcast  3. Audiolibro")
	typeStr := readInput("Tipo (1-3): ")
	if typeStr == "0" {
		return
	}

	typeNum, err := strconv.Atoi(typeStr)
	if err != nil || typeNum < 1 || typeNum > 3 {
		fmt.Println("Tipo inválido")
		waitForEnter()
		return
	}

	contentTypes := []string{"Música", "Podcast", "Audiolibro"}
	contentType := contentTypes[typeNum-1]

	durationStr := readInput("Duración (minutos): ")
	if durationStr == "0" {
		return
	}
	duration, err := strconv.Atoi(durationStr)
	if err != nil || duration <= 0 {
		fmt.Println("Duración inválida")
		waitForEnter()
		return
	}

	// Clasificaciones
	fmt.Println("Clasificaciones:")
	ratings := contentclass.GetAllRatings()
	for i, r := range ratings {
		fmt.Printf("%d. %s\n", i+1, r.Name)
	}

	ratingStr := readInput("Clasificación (1-3): ")
	if ratingStr == "0" {
		return
	}
	ratingNum, err := strconv.Atoi(ratingStr)
	if err != nil || ratingNum < 1 || ratingNum > len(ratings) {
		fmt.Println("Clasificación inválida")
		waitForEnter()
		return
	}

	ageRating := ratings[ratingNum-1].Name

	err = audio.AddContent(title, contentType, "Música", duration, ageRating, "Artista", "Álbum", 1)
	if err != nil {
		fmt.Println("Error al agregar contenido")
	} else {
		fmt.Println(" Contenido agregado")
	}
	waitForEnter()
}

// Leer entrada del usuario
func readInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	return ""
}
