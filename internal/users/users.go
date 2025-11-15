/* @titulo: modulo de gestion de usuarios
   @descripcion: Define la estructura de datos para usuarios y proporciona funciones para agregar, listar y manipular usuarios en memoria. */

package users // Nombre del paquete, esta en internal/users para que sea no sea importado por otros paquetes del proyecto

import "fmt"

// Defino la estructura que representara a cada usuario en el sistema
type User struct {
	ID   int // ID del usuario
	Name string // Nombre del usuario
	Age  int // Edad del usuario
	Email string // Correo electrónico del usuario
	Password string // Contraseña del usuario
}
// Variable que almacena los usuarios registrados en el sistema en memoria
// Para el AA2 se reemplazara por la base de datos
var users []User
// Variable que asigna los ID's unicos a cada usuario registrado
//comienza en 1 y se incrementa en 1 por cada usuario registrado
var userID int = 1
// funcion que crea un nuevo usuario y lo añade a la lista 'usersInMemory'
// Recibe el nombre, la edad, el correo y la contraseña del usuario.
func AddUser(name string, age int, email string, password string) {
	user := User{
		ID:       userID,
		Name:     name,
		Age:      age,
		Email:    email,
		Password: password,
	}
// Agrego el nuevo usuario a la lista global 'usersInMemory'
	usersInMemory = append(usersInMemory, newuser)
	userID++ // Incrementa el ID para el siguiente usuario
}
// Funcion que imprime en consola todos los usuarios almacenados en 'usersInMemory'
func ListUsers() { // Primero verifico si la lista de usuarios esta vacia
	if len(usersInMemory) == 0 {
		fmt.Println("No hay usuarios registrados.")
		return // Si esta vacia, salgo de la func
	}
	fmt.Println("--- Lista de usuarios ---") // si hay, los imprimo
	for _, user := range usersInMemory { // Itero sobre la lista 'usersInMemory'
		fmt.Println("ID: %d, Nombre: %s, Edad: %d, Correo: %s", user.ID, user.Name, user.Age, user.Email)
	}
}
// Funcion que busca un usuario en la lista 'usersInMemory' por su ID
func imprimirUsuariosDetallados(users []User) {
	for _, user := range users {
		fmt.Printf("Detalle - ID: %d, Nombre: %s, Edad: %d, Correo: %s\n", user.ID, user.Name, user.Age, user.Email)
		}
}
// Funcion anonima que uso como parámetro en 'filtrarUsuarios', recibe una lista y una funcion de filtro
func filtrarUsuarios(users []User, filter func(user) bool) []User {
// Itero sobre cada usuario recibido en la lista variádica 'users'
	var filtrados []User
	for _, u := range users {
		if filter(u) { // Si el filtro devuelve true, agrego el usuario a la lista 'filtrados'
			filtrados = append(filtrados, u)
	}
	return filtrados
}
// Funcion que recibe una lista de usuaros y una función de operación
// Luego aplica esa operación a cada usuario de la lista
func AplicarOperacionUsuario(list []User, operation func(User)) {
	for _, user := range list { //Aplico la función de operación al usuario actual
		operation(u)
	}
}