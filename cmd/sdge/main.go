<<<<<<< HEAD
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

=======
/* @Programa principal del proyecto SDGEStreaming - Programación orientada a objetos
   @Autores: Nelson Espinosa, Barbara Peñaherrera
   @Domingo 7 de diciembre de 2025. Quito - Ecuador
   @Punto de entrada del sistema. Contiene el menú interactivo y la lógica de control principal que orquesta las interacciones con los módulos y la base de datos.*/
// cmd/sdge/main.go
package main

import (
	"SDGEStreaming/internal/db"
	"SDGEStreaming/internal/models"
	"SDGEStreaming/internal/repositories"
	"SDGEStreaming/internal/security"
	"SDGEStreaming/internal/services"
	"SDGEStreaming/internal/utils"
	"fmt"
	"os"
	"time"
)
// Variables globales
var (
	currentUser *CurrentUser
)
// Inicializacion de servicios globales
var (
	userService         *services.UserService
	contentService      *services.ContentService
	subscriptionService *services.SubscriptionService
	playbackService     *services.PlaybackService

	userRepo repositories.UserRepo
)
// Estructura para almacenar información del usuario actual
type CurrentUser struct {
	ID        int
	Name      string
	Email     string
	PlanID    int
	PlanName  string
	Age       int
	AgeRating string
	IsAdmin   bool
}
// Función principal
func main() {
	if err := db.InitDB("sdgestreaming.db"); err != nil {
		fmt.Printf("Error fatal al iniciar la base de datos: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	// Inicialización de repositorios
	userRepo = repositories.NewUserRepo()
	contentRepo := repositories.NewContentRepo()
	subscriptionRepo := repositories.NewSubscriptionRepo()
	playbackHistoryRepo := repositories.NewPlaybackHistoryRepo()
	favoriteRepo := repositories.NewFavoriteRepo()

	// Crear usuario admin si no existe
	adminUser, err := userRepo.FindByEmail("admin@sdge.com")
	if err != nil {
		fmt.Printf("Error buscando usuario admin: %v\n", err)
	}
	if adminUser == nil {
		hashedPass, err := security.HashPassword("admin123")
		if err != nil {
			fmt.Printf("Error generando contraseña del admin: %v\n", err)
		} else {
			now := time.Now()
			adminModel := &models.User{
				Name:         "Admin",
				Email:        "admin@sdge.com",
				Age:          30,
				PlanID:       3,
				AgeRating:    "Adulto",
				IsAdmin:      true,
				PasswordHash: hashedPass,
				CreatedAt:    now,
				LastLogin:    now,
			}
			if err := userRepo.Create(adminModel); err != nil {
				fmt.Printf("Error creando usuario admin: %v\n", err)
			}
		}
	}

	userService = services.NewUserService(userRepo, subscriptionRepo)
	contentService = services.NewContentService(contentRepo)
	subscriptionService = services.NewSubscriptionService(subscriptionRepo, userRepo)
	playbackService = services.NewPlaybackService(playbackHistoryRepo, favoriteRepo, contentRepo)

	utils.ClearScreen()
	runApplication()
}

func runApplication() {
	for {
>>>>>>> b9e3b62 (AA2_CULMINADO)
		if currentUser == nil {
			showAuthMenu()
		} else {
			showMainMenu()
		}
	}
}

<<<<<<< HEAD
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

=======
func showAuthMenu() {
	utils.ClearScreen()
	fmt.Println("╔════════════════════════════════╗")
	fmt.Println("║    SDGEStreaming - Inicio    ║")
	fmt.Println("╚════════════════════════════════╝")
	fmt.Println()
	fmt.Println("1. Iniciar Sesión")
	fmt.Println("2. Registrarse")
	fmt.Println("3. Salir")
	fmt.Print("\nSeleccione una opción: ")

	option := utils.ReadLine("")
>>>>>>> b9e3b62 (AA2_CULMINADO)
	switch option {
	case "1":
		login()
	case "2":
		register()
	case "3":
<<<<<<< HEAD
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
=======
		fmt.Println("¡Gracias por usar SDGEStreaming!")
		os.Exit(0)
	default:
		fmt.Println("Opción inválida.")
		utils.WaitForEnter()
	}
}

func login() {
	utils.ClearScreen()
	fmt.Println("Iniciar Sesión")
	fmt.Println("══════════════")
	email := utils.ReadLine("Email: ")
	password := utils.ReadLine("Contraseña: ")

	user, err := userService.Login(email, password)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		utils.WaitForEnter()
		return
	}

	planName := getPlanName(user.PlanID)

	currentUser = &CurrentUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		PlanID:    user.PlanID,
		PlanName:  planName,
		Age:       user.Age,
		AgeRating: user.AgeRating,
		IsAdmin:   user.IsAdmin,
	}
	fmt.Printf("¡Bienvenido, %s!\n", user.Name)
	utils.WaitForEnter()
}

func register() {
	utils.ClearScreen()
	fmt.Println("Registro de Nuevo Usuario")
	fmt.Println("══════════════════════════")
	name := utils.ReadLine("Nombre completo: ")
	ageStr := utils.ReadLine("Edad (mínimo 13): ")
	age, err := utils.ToInt(ageStr)
	if err != nil || age < 13 {
		fmt.Println("Edad inválida. Debe ser un número entero mayor o igual a 13.")
		utils.WaitForEnter()
		return
	}

	email := utils.ReadLine("Email: ")
	password := utils.ReadLine("Contraseña (mínimo 6 caracteres): ")

	if !utils.IsValidEmail(email) {
		fmt.Println("Formato de email inválido.")
		utils.WaitForEnter()
		return
	}
	if !utils.IsValidPassword(password) {
		fmt.Println("La contraseña debe tener al menos 6 caracteres.")
		utils.WaitForEnter()
		return
	}

	_, err = userService.Register(name, age, email, password, false)
	if err != nil {
		fmt.Printf("Error en el registro: %v\n", err)
	} else {
		fmt.Println("¡Registro exitoso! Ahora puede iniciar sesión.")
	}
	utils.WaitForEnter()
}

func showMainMenu() {
	utils.ClearScreen()
	fmt.Println("Menú Principal")
	fmt.Println("══════════════")
	fmt.Printf("Hola, %s (%s)\n", currentUser.Name, currentUser.PlanName)
	fmt.Println()
	fmt.Println("1. Inicio")
	fmt.Println("2. Tendencias")
	fmt.Println("3. Explorar Contenido")
	fmt.Println("4. Mi Lista")
	fmt.Println("5. Perfil y Cuenta")
	if currentUser.IsAdmin {
		fmt.Println("6. Panel de Administración")
		fmt.Println("7. Cerrar Sesión")
	} else {
		fmt.Println("6. Cerrar Sesión")
	}
	fmt.Print("\nSeleccione una opción: ")

	option := utils.ReadLine("")
	switch option {
	case "1":
		showHome()
	case "2":
		showTrending()
	case "3":
		browseContent(false)
	case "4":
		showMyList()
	case "5":
		showProfileMenu()
	case "6":
		if currentUser.IsAdmin {
			showAdminPanel()
		} else {
			logout()
		}
	case "7":
		if currentUser.IsAdmin {
			logout()
		}
	default:
		fmt.Println("Opción inválida.")
	}
	utils.WaitForEnter()
}

func showHome() {
	utils.ClearScreen()
	fmt.Println("Inicio")
	fmt.Println("══════")
	fmt.Println("¡Bienvenido a tu página de inicio!\n")

	fmt.Println("► Continuar viendo:")
	continueWatching, _ := playbackService.GetContinueWatching(currentUser.ID)
	if len(continueWatching) == 0 {
		fmt.Println("  No tienes nada en progreso.")
	} else {
		for _, entry := range continueWatching {
			var title string
			if entry.ContentType == "audiovisual" {
				content, _ := contentService.GetAudiovisualByID(entry.ContentID)
				if content != nil {
					title = content.Title
				}
			} else {
				content, _ := contentService.GetAudioByID(entry.ContentID)
				if content != nil {
					title = content.Title
				}
			}
			if title != "" {
				fmt.Printf("  * %s (ID: %d)\n", title, entry.ContentID)
>>>>>>> b9e3b62 (AA2_CULMINADO)
			}
		}
	}

<<<<<<< HEAD
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

=======
	fmt.Println()
	utils.WaitForEnter()
}

func showTrending() {
	utils.ClearScreen()
	fmt.Println("Tendencias")
	fmt.Println("══════════")

	fmt.Println("\n🎬 Contenido Audiovisual Popular:")
	audiovisuals, err := contentService.GetAllAudiovisual()
	if err != nil {
		fmt.Printf("Error al cargar contenido: %v\n", err)
	} else {
		for i, av := range audiovisuals {
			if i >= 3 {
				break
			}
			fmt.Printf("  %d. %s (%.1f⭐)\n", i+1, av.Title, av.AverageRating)
		}
	}

	fmt.Println("\n🎵 Contenido de Audio Popular:")
	audios, err := contentService.GetAllAudio()
	if err != nil {
		fmt.Printf("Error al cargar contenido: %v\n", err)
	} else {
		for i, a := range audios {
			if i >= 3 {
				break
			}
			fmt.Printf("  %d. %s - %s (%.1f⭐)\n", i+1, a.Artist, a.Title, a.AverageRating)
		}
	}

	utils.WaitForEnter()
}

func browseContent(isGuest bool) {
	for {
		utils.ClearScreen()
		fmt.Println("Explorar Contenido")
		fmt.Println("══════════════════")
		fmt.Println("1. Contenido Audiovisual")
		fmt.Println("2. Contenido de Audio")
		fmt.Println("3. Volver")
		fmt.Print("\nSeleccione una opción: ")

		option := utils.ReadLine("")
		switch option {
		case "1":
			browseAudiovisual(isGuest)
		case "2":
			browseAudio(isGuest)
		case "3":
			return
		default:
			fmt.Println("Opción inválida.")
			utils.WaitForEnter()
		}
	}
}

func browseAudiovisual(isGuest bool) {
	contents, err := contentService.GetAllAudiovisualForUser(currentUser.AgeRating)
	if err != nil {
		fmt.Printf("Error al cargar contenido: %v\n", err)
		utils.WaitForEnter()
		return
	}

	if len(contents) == 0 {
		fmt.Println("No hay contenido audiovisual disponible para tu clasificación de edad.")
		utils.WaitForEnter()
		return
	}

	utils.ClearScreen()
	fmt.Println("\n🎬 Contenido Audiovisual Disponible:")
	for _, c := range contents {
		fmt.Printf("ID: %d | %s (%s)\n", c.ID, c.Title, c.Type)
		fmt.Printf("   Género: %s | Duración: %d min | Clasificación: %s\n", c.Genre, c.Duration, c.AgeRating)
		fmt.Printf("   Promedio: %.1f⭐\n", c.AverageRating)
		fmt.Println("────────────────────────────────────────")
	}

	if !isGuest {
		contentIDStr := utils.ReadLine("\nIngrese el ID del contenido para ver detalles (0 para volver): ")
		if contentIDStr == "0" {
			return
		}
		contentID, err := utils.ToInt(contentIDStr)
		if err != nil {
			fmt.Println("ID inválido.")
			utils.WaitForEnter()
			return
		}

		content, err := contentService.GetAudiovisualByID(contentID)
		if err != nil {
			fmt.Println("Contenido no encontrado.")
			utils.WaitForEnter()
			return
		}

		utils.ClearScreen()
		fmt.Printf("═══ %s ═══\n", content.Title)
		fmt.Printf("Tipo: %s\n", content.Type)
		fmt.Printf("Género: %s\n", content.Genre)
		fmt.Printf("Sinopsis: %s\n", content.Synopsis)
		fmt.Printf("Director: %s\n", content.Director)
		fmt.Printf("Año: %d\n", content.ReleaseYear)
		fmt.Printf("Duración: %d minutos\n", content.Duration)
		fmt.Printf("Clasificación: %s\n", content.AgeRating)
		fmt.Printf("Promedio de calificación: %.1f⭐\n", content.AverageRating)
		fmt.Println("\n1. Reproducir")
		fmt.Println("2. Marcar como favorito")
		fmt.Println("3. Calificar")
		fmt.Println("4. Volver")
		action := utils.ReadLine("Seleccione una acción: ")

		switch action {
		case "1":
			playAudiovisual(contentID)
		case "2":
			err = playbackService.AddFavorite(currentUser.ID, contentID, "audiovisual")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("¡Agregado a Mi Lista!")
			}
			utils.WaitForEnter()
		case "3":
			rateContent(contentID, "audiovisual")
		case "4":
			return
		}
	}
}

func browseAudio(isGuest bool) {
	contents, err := contentService.GetAllAudioForUser(currentUser.AgeRating)
	if err != nil {
		fmt.Printf("Error al cargar contenido: %v\n", err)
		utils.WaitForEnter()
		return
	}

	if len(contents) == 0 {
		fmt.Println("No hay contenido de audio disponible para tu clasificación de edad.")
		utils.WaitForEnter()
		return
	}

	utils.ClearScreen()
	fmt.Println("\n🎵 Contenido de Audio Disponible:")
	for _, c := range contents {
		fmt.Printf("ID: %d | %s - %s\n", c.ID, c.Artist, c.Title)
		fmt.Printf("   Tipo: %s | Género: %s | Álbum: %s\n", c.Type, c.Genre, c.Album)
		fmt.Printf("   Duración: %d min | Clasificación: %s\n", c.Duration, c.AgeRating)
		fmt.Printf("   Promedio: %.1f⭐\n", c.AverageRating)
		fmt.Println("────────────────────────────────────────")
	}

	if !isGuest {
		contentIDStr := utils.ReadLine("\nIngrese el ID del contenido para ver detalles (0 para volver): ")
		if contentIDStr == "0" {
			return
		}
		contentID, err := utils.ToInt(contentIDStr)
		if err != nil {
			fmt.Println("ID inválido.")
			utils.WaitForEnter()
			return
		}

		content, err := contentService.GetAudioByID(contentID)
		if err != nil {
			fmt.Println("Contenido no encontrado.")
			utils.WaitForEnter()
			return
		}

		utils.ClearScreen()
		fmt.Printf("═══ %s ═══\n", content.Title)
		fmt.Printf("Artista: %s\n", content.Artist)
		fmt.Printf("Álbum: %s\n", content.Album)
		fmt.Printf("Género: %s\n", content.Genre)
		fmt.Printf("Duración: %d minutos\n", content.Duration)
		fmt.Printf("Clasificación: %s\n", content.AgeRating)
		fmt.Printf("Promedio de calificación: %.1f⭐\n", content.AverageRating)
		fmt.Println("\n1. Reproducir")
		fmt.Println("2. Marcar como favorito")
		fmt.Println("3. Calificar")
		fmt.Println("4. Volver")
		action := utils.ReadLine("Seleccione una acción: ")

		switch action {
		case "1":
			playAudio(contentID)
		case "2":
			err = playbackService.AddFavorite(currentUser.ID, contentID, "audio")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("¡Agregado a Mi Lista!")
			}
			utils.WaitForEnter()
		case "3":
			rateContent(contentID, "audio")
		case "4":
			return
		}
	}
}

func showMyList() {
	utils.ClearScreen()
	fmt.Println("Mi Lista")
	fmt.Println("════════")

	favorites, err := playbackService.GetFavorites(currentUser.ID)
	if err != nil {
		fmt.Printf("Error al cargar favoritos: %v\n", err)
		utils.WaitForEnter()
		return
	}

	if len(favorites) == 0 {
		fmt.Println("No tienes ningún contenido en tu lista.")
		utils.WaitForEnter()
		return
	}

	fmt.Println("Contenido en tu lista:")
	for _, fav := range favorites {
		var title, details string
		if fav.ContentType == "audiovisual" {
			content, _ := contentService.GetAudiovisualByID(fav.ContentID)
			if content != nil {
				title = content.Title
				details = fmt.Sprintf("[%s] %s", content.Type, content.Genre)
			}
		} else {
			content, _ := contentService.GetAudioByID(fav.ContentID)
			if content != nil {
				title = fmt.Sprintf("%s - %s", content.Artist, content.Title)
				details = fmt.Sprintf("[%s] %s", content.Type, content.Genre)
			}
		}
		if title != "" {
			fmt.Printf("* %s\n", title)
			fmt.Printf("  %s\n", details)
			fmt.Println("────────────────────────────────────────")
		}
	}
	utils.WaitForEnter()
}

func showProfileMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("Mi Perfil")
		fmt.Println("═════════")
		fmt.Printf("Nombre: %s\n", currentUser.Name)
		fmt.Printf("Email: %s\n", currentUser.Email)
		fmt.Printf("Plan actual: %s\n", currentUser.PlanName)
		fmt.Printf("Edad: %d\n", currentUser.Age)
		fmt.Printf("Clasificación: %s\n", currentUser.AgeRating)
		fmt.Println()
		fmt.Println("1. Cambiar Plan de Suscripción")
		fmt.Println("2. Ver Métodos de Pago")
		fmt.Println("3. Ver Historial de Reproducción")
		fmt.Println("4. Volver al Menú Principal")
		fmt.Print("\nSeleccione una opción: ")

		option := utils.ReadLine("")
		switch option {
		case "1":
			upgradePlan()
		case "2":
			viewPaymentMethods()
		case "3":
			viewPlaybackHistory()
		case "4":
			return
		default:
			fmt.Println("Opción inválida.")
			utils.WaitForEnter()
		}
	}
}

func viewPlaybackHistory() {
	utils.ClearScreen()
	fmt.Println("Historial de Reproducción")
	fmt.Println("══════════════════════════")

	history, err := playbackService.GetHistory(currentUser.ID)
	if err != nil {
		fmt.Printf("Error al cargar el historial: %v\n", err)
		utils.WaitForEnter()
		return
	}

	if len(history) == 0 {
		fmt.Println("No tienes historial de reproducción.")
		utils.WaitForEnter()
		return
	}

	fmt.Println("Tus últimas reproducciones:")
	for _, entry := range history {
		var title string
		if entry.ContentType == "audiovisual" {
			content, _ := contentService.GetAudiovisualByID(entry.ContentID)
			if content != nil {
				title = content.Title
			}
		} else {
			content, _ := contentService.GetAudioByID(entry.ContentID)
			if content != nil {
				title = content.Title
			}
		}
		if title != "" {
			fmt.Printf("* %s (%s)\n", title, entry.ContentType)
		}
	}
	utils.WaitForEnter()
}

func upgradePlan() {
	utils.ClearScreen()
	fmt.Println("Cambiar Plan de Suscripción")
	fmt.Println("════════════════════════════")

	plans, err := subscriptionService.GetAvailablePlans()
	if err != nil {
		fmt.Printf("Error al cargar planes: %v\n", err)
		utils.WaitForEnter()
		return
	}

	fmt.Println("Planes disponibles:")
	for _, p := range plans {
		fmt.Printf("%d. %s - $%.2f/mes | Calidad: %s | Dispositivos: %d\n",
			p.ID, p.Name, p.Price, p.MaxQuality, p.MaxDevices)
	}

	planIDStr := utils.ReadLine("\nSeleccione el número del plan deseado (0 para cancelar): ")
	if planIDStr == "0" {
		return
	}
	planID, err := utils.ToInt(planIDStr)
	if err != nil {
		fmt.Println("Selección inválida.")
		utils.WaitForEnter()
		return
	}

	if planID == currentUser.PlanID {
		fmt.Println("Ya está suscrito a este plan.")
		utils.WaitForEnter()
		return
	}

	if planID == 1 {
		err = userService.UpdateUserPlan(currentUser.ID, 1)
		if err != nil {
			fmt.Printf("Error al actualizar el plan: %v\n", err)
			utils.WaitForEnter()
			return
		}
		currentUser.PlanID = 1
		currentUser.PlanName = "Free"
		fmt.Println("Su plan ha sido cambiado a Free.")
		utils.WaitForEnter()
		return
	}

	fmt.Println("\n--- Información de Pago ---")
	cardHolder := utils.ReadLine("Nombre del titular de la tarjeta: ")
	cardNumber := utils.ReadLine("Número de tarjeta (16 dígitos): ")
	expiry := utils.ReadLine("Fecha de vencimiento (MM/AAAA): ")
	cvvStr := utils.ReadLine("CVV (3 dígitos): ")

	var expiryMonth, expiryYear int
	if len(expiry) == 7 && expiry[2] == '/' {
		expiryMonth, _ = utils.ToInt(expiry[0:2])
		expiryYear, _ = utils.ToInt(expiry[3:7])
	} else {
		fmt.Println("Formato de fecha de vencimiento inválido (MM/AAAA).")
		utils.WaitForEnter()
		return
	}

	cvv, err := utils.ToInt(cvvStr)
	if err != nil {
		fmt.Println("CVV inválido.")
		utils.WaitForEnter()
		return
	}

	err = subscriptionService.ProcessPayment(currentUser.ID, planID, cardHolder, cardNumber, expiryMonth, expiryYear, cvv)
	if err != nil {
		fmt.Printf("Error en el pago: %v\n", err)
		utils.WaitForEnter()
		return
	}

	currentUser.PlanID = planID
	currentUser.PlanName = getPlanName(planID)
	fmt.Println("¡Su plan ha sido actualizado exitosamente!")
	utils.WaitForEnter()
}

func viewPaymentMethods() {
	utils.ClearScreen()
	fmt.Println("Métodos de Pago")
	fmt.Println("════════════════")
	method, err := userService.GetDefaultPaymentMethod(currentUser.ID)
	if err != nil {
		fmt.Println("No tiene métodos de pago guardados.")
	} else {
		fmt.Printf("Tarjeta predeterminada: **** **** **** %s\n", method.Last4)
		fmt.Printf("Titular: %s\n", method.CardHolder)
		fmt.Printf("Vence: %02d/%d\n", method.ExpiryMonth, method.ExpiryYear)
	}
	utils.WaitForEnter()
}

func showAdminPanel() {
	utils.ClearScreen()
	fmt.Println("Panel de Administración")
	fmt.Println("════════════════════════")
	fmt.Println("1. Gestionar Usuarios")
	fmt.Println("2. Gestionar Contenido")
	fmt.Println("3. Generar Reportes")
	fmt.Println("4. Volver")
	fmt.Print("\nSeleccione una opción: ")

	option := utils.ReadLine("")
	switch option {
	case "1":
		manageUsers()
	case "2":
		manageContent()
	case "3":
		generateReports()
	case "4":
		return
	default:
		fmt.Println("Opción inválida.")
	}
	utils.WaitForEnter()
}

func manageUsers() {
	utils.ClearScreen()
	fmt.Println("Gestión de Usuarios")
	fmt.Println("════════════════════")
	users, err := userService.GetAllUsers()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		utils.WaitForEnter()
		return
	}
>>>>>>> b9e3b62 (AA2_CULMINADO)
	for _, u := range users {
		adminTag := ""
		if u.IsAdmin {
			adminTag = " [ADMIN]"
		}
<<<<<<< HEAD
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
=======
		fmt.Printf("- %s%s (%s) | Edad: %d | Clasificación: %s\n", u.Name, adminTag, u.Email, u.Age, u.AgeRating)
	}
	utils.WaitForEnter()
}

func manageContent() {
	for {
		utils.ClearScreen()
		fmt.Println("Gestión de Contenido")
		fmt.Println("════════════════════")
		fmt.Println("1. Agregar Contenido Audiovisual")
		fmt.Println("2. Agregar Contenido de Audio")
		fmt.Println("3. Listar Contenido Audiovisual")
		fmt.Println("4. Listar Contenido de Audio")
		fmt.Println("5. Volver")
		fmt.Print("\nSeleccione una opción: ")

		option := utils.ReadLine("")
		switch option {
		case "1":
			addAudiovisualContent()
		case "2":
			addAudioContent()
		case "3":
			listAudiovisualAdmin()
		case "4":
			listAudioAdmin()
		case "5":
			return
		default:
			fmt.Println("Opción inválida.")
			utils.WaitForEnter()
		}
	}
}

func addAudiovisualContent() {
	utils.ClearScreen()
	fmt.Println("Agregar Contenido Audiovisual")
	fmt.Println("══════════════════════════════")

	title := utils.ReadLine("Título: ")
	contentType := utils.ReadLine("Tipo (movie/series/documentary): ")
	genre := utils.ReadLine("Género: ")
	durationStr := utils.ReadLine("Duración (minutos): ")
	duration, err := utils.ToInt(durationStr)
	if err != nil {
		fmt.Println("Duración inválida.")
		utils.WaitForEnter()
		return
	}

	ageRating := utils.ReadLine("Clasificación (G/PG/PG-13/R): ")
	synopsis := utils.ReadLine("Sinopsis: ")
	yearStr := utils.ReadLine("Año de lanzamiento: ")
	year, err := utils.ToInt(yearStr)
	if err != nil {
		fmt.Println("Año inválido.")
		utils.WaitForEnter()
		return
	}
	director := utils.ReadLine("Director: ")

	err = contentService.CreateAudiovisual(title, contentType, genre, duration, ageRating, synopsis, year, director)
	if err != nil {
		fmt.Printf("Error al agregar contenido: %v\n", err)
	} else {
		fmt.Println("¡Contenido agregado exitosamente!")
	}
	utils.WaitForEnter()
}

func addAudioContent() {
	utils.ClearScreen()
	fmt.Println("Agregar Contenido de Audio")
	fmt.Println("═══════════════════════════")

	title := utils.ReadLine("Título: ")
	contentType := utils.ReadLine("Tipo (song/podcast/audiobook): ")
	genre := utils.ReadLine("Género: ")
	durationStr := utils.ReadLine("Duración (minutos): ")
	duration, err := utils.ToInt(durationStr)
	if err != nil {
		fmt.Println("Duración inválida.")
		utils.WaitForEnter()
		return
	}

	ageRating := utils.ReadLine("Clasificación (General/Explicit): ")
	artist := utils.ReadLine("Artista: ")
	album := utils.ReadLine("Álbum: ")
	trackStr := utils.ReadLine("Número de pista: ")
	trackNumber, err := utils.ToInt(trackStr)
	if err != nil {
		trackNumber = 1
	}

	err = contentService.CreateAudio(title, contentType, genre, duration, ageRating, artist, album, trackNumber)
	if err != nil {
		fmt.Printf("Error al agregar contenido: %v\n", err)
	} else {
		fmt.Println("¡Contenido agregado exitosamente!")
	}
	utils.WaitForEnter()
}

func listAudiovisualAdmin() {
	utils.ClearScreen()
	fmt.Println("Lista de Contenido Audiovisual")
	fmt.Println("═══════════════════════════════")

	contents, err := contentService.GetAllAudiovisual()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		utils.WaitForEnter()
		return
	}

	for _, c := range contents {
		fmt.Printf("ID: %d | %s (%s) - %d min | Clasificación: %s\n", c.ID, c.Title, c.Type, c.Duration, c.AgeRating)
	}
	utils.WaitForEnter()
}

func listAudioAdmin() {
	utils.ClearScreen()
	fmt.Println("Lista de Contenido de Audio")
	fmt.Println("════════════════════════════")

	contents, err := contentService.GetAllAudio()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		utils.WaitForEnter()
		return
	}

	for _, c := range contents {
		fmt.Printf("ID: %d | %s - %s (%s) - %d min | Clasificación: %s\n", c.ID, c.Artist, c.Title, c.Type, c.Duration, c.AgeRating)
	}
	utils.WaitForEnter()
}

func generateReports() {
	utils.ClearScreen()
	fmt.Println("Generación de Reportes")
	fmt.Println("═══════════════════════")
	users, err := userService.GetAllUsers()
	if err != nil {
		fmt.Printf("Error al cargar usuarios: %v\n", err)
	} else {
		fmt.Printf("* Total de Usuarios: %d\n", len(users))
	}

	audiovisuals, err := contentService.GetAllAudiovisual()
	if err != nil {
		fmt.Printf("Error al cargar contenido audiovisual: %v\n", err)
	} else {
		fmt.Printf("* Total de Contenido Audiovisual: %d\n", len(audiovisuals))
		if len(audiovisuals) > 0 {
			fmt.Printf("* Contenido más popular: '%s' (%.1f⭐)\n", audiovisuals[0].Title, audiovisuals[0].AverageRating)
		}
	}

	audios, err := contentService.GetAllAudio()
	if err != nil {
		fmt.Printf("Error al cargar contenido de audio: %v\n", err)
	} else {
		fmt.Printf("* Total de Contenido de Audio: %d\n", len(audios))
		if len(audios) > 0 {
			fmt.Printf("* Audio más popular: '%s - %s' (%.1f⭐)\n", audios[0].Artist, audios[0].Title, audios[0].AverageRating)
		}
	}

	utils.WaitForEnter()
}

func logout() {
	currentUser = nil
	fmt.Println("Sesión cerrada correctamente.")
	utils.WaitForEnter()
}

func getPlanName(planID int) string {
	switch planID {
	case 1:
		return "Free"
	case 2:
		return "Estándar"
	case 3:
		return "Premium 4K"
	default:
		return "Desconocido"
	}
}

func playAudiovisual(contentID int) {
	content, err := contentService.GetAudiovisualByID(contentID)
	if err != nil {
		fmt.Println("Error al cargar contenido.")
		utils.WaitForEnter()
		return
	}

	utils.ClearScreen()
	fmt.Printf("▶ Reproduciendo: %s\n", content.Title)
	fmt.Println("══════════════════════════════════════")
	fmt.Println("Simulando reproducción...")
	fmt.Println("[████████████████████] 100%")
	fmt.Printf("Duración total: %d minutos\n", content.Duration)
	fmt.Println("══════════════════════════════════════")

	// Registrar en historial
	if err := playbackService.AddToHistory(currentUser.ID, contentID, "audiovisual"); err != nil {
		fmt.Printf("No se pudo registrar en historial: %v\n", err)
	}

	// Simular progreso (50% visto)
	progressSeconds := (content.Duration * 60) / 2
	if err := playbackService.UpdateProgress(currentUser.ID, contentID, "audiovisual", progressSeconds); err != nil {
		fmt.Printf("No se pudo actualizar progreso: %v\n", err)
	}

	fmt.Println("\n✓ Reproducción finalizada")
	fmt.Println("Se ha guardado tu progreso.")
	utils.WaitForEnter()
}

func playAudio(contentID int) {
	content, err := contentService.GetAudioByID(contentID)
	if err != nil {
		fmt.Println("Error al cargar contenido.")
		utils.WaitForEnter()
		return
	}

	utils.ClearScreen()
	fmt.Printf("♪ Reproduciendo: %s - %s\n", content.Artist, content.Title)
	fmt.Println("══════════════════════════════════════")
	fmt.Println("Simulando reproducción...")
	fmt.Println("[████████████████████] 100%")
	fmt.Printf("Duración total: %d minutos\n", content.Duration)
	fmt.Println("══════════════════════════════════════")

	// Registrar en historial
	if err := playbackService.AddToHistory(currentUser.ID, contentID, "audio"); err != nil {
		fmt.Printf("No se pudo registrar en historial: %v\n", err)
	}

	// Simular progreso (70% escuchado)
	progressSeconds := (content.Duration * 60) * 7 / 10
	if err := playbackService.UpdateProgress(currentUser.ID, contentID, "audio", progressSeconds); err != nil {
		fmt.Printf("No se pudo actualizar progreso: %v\n", err)
	}

	fmt.Println("\n✓ Reproducción finalizada")
	fmt.Println("Se ha guardado tu progreso.")
	utils.WaitForEnter()
}

func rateContent(contentID int, contentType string) {
	utils.ClearScreen()
	fmt.Println("Calificar Contenido")
	fmt.Println("════════════════════")
	fmt.Println("Ingrese su calificación (1.0 - 10.0)")
	ratingStr := utils.ReadLine("Calificación: ")

	rating, err := utils.ToFloat(ratingStr)
	if err != nil || rating < 1.0 || rating > 10.0 {
		fmt.Println("Calificación inválida. Debe ser entre 1.0 y 10.0")
		utils.WaitForEnter()
		return
	}

	// Guardar calificación
	err = contentService.RateContent(currentUser.ID, contentID, contentType, rating)
	if err != nil {
		fmt.Printf("Error al calificar: %v\n", err)
	} else {
		fmt.Printf("¡Gracias! Has calificado este contenido con %.1f⭐\n", rating)
	}
	utils.WaitForEnter()
>>>>>>> b9e3b62 (AA2_CULMINADO)
}
