# 📺 SDGEStreaming – Sistema de Gestión de Contenido Audiovisual y de Audio

> **Proyecto Final**  
> Programación orientada a objetos en Go – Backend para plataforma de streaming

---

## 📌 Descripción General

**SDGEStreaming** es un backend académico desarrollado en **Go (Golang)** para una plataforma de streaming de contenido audiovisual y de audio.  
El sistema está diseñado con una **arquitectura por capas** (modelos, repositorios, servicios, handlers/API) y una base de datos **SQLite**.

### Evolución del proyecto

- **AA1 – Fase inicial**
  - Registro e inicio de sesión de usuarios.
  - Exploración de contenido.
  - Sistema de calificación de contenido (ratings).

- **AA2 – Paso 1 y Paso 2**
  - Activación de **Mi Lista (favoritos)**.
  - Implementación de **Historial de reproducción**.
  - Refuerzo de la **protección a menores** mediante clasificación por edad.
  - Mejora del flujo de menús de la **aplicación de consola**.

- **Trabajo final (este entregable)**
  - Implementación de **perfiles por usuario** (según plan de suscripción).
  - Exposición de una **API HTTP REST** para registro, login, planes, contenido y valoraciones.
  - Simulación de **cambio de plan con pago** (validación de tarjeta).
  - Corrección y alineación de modelos con el esquema de base de datos.
  - Ajustes para ejecución en Windows con `CGO_ENABLED=1` y `go-sqlite3`.

---

## 🧩 Funcionalidades Implementadas

### A nivel de dominio

| Funcionalidad                        | Descripción                                                                                          |
|-------------------------------------|------------------------------------------------------------------------------------------------------|
| **Autenticación segura**            | Registro e inicio de sesión con validaciones y hashing de contraseñas (`bcrypt`).                   |
| **Perfiles por usuario**            | Cada cuenta puede tener varios perfiles (niño, adolescente, adulto) según el plan de suscripción.  |
| **Clasificación por edad**          | El perfil tiene una clasificación de edad; el contenido se filtra automáticamente según esa regla. |
| **Catálogo de contenido**           | Gestión de contenido audiovisual y de audio con metadatos y rating promedio.                        |
| **Mi Lista (Favoritos)**            | Añadir y visualizar contenido favorito por tipo (audio / audiovisual).                              |
| **Historial de reproducción**       | Registro automático de reproducciones exitosas.                                                      |
| **Calificación de ítems (ratings)** | Valoración de contenido de 1.0 a 10.0, con promedio recalculado.                                    |
| **Planes y suscripciones**          | Planes Free, Estándar y Premium 4K, con límites de calidad y cantidad de perfiles.                  |
| **Simulación de pagos**             | Cambio de plan con validación de tarjeta y almacenamiento no sensible de datos.                     |
| **Menú administrador**              | Opciones adicionales visibles solo para usuarios con rol `admin`.                                  |

### A nivel de interfaz

- **Aplicación de consola**  
  - Menús de texto para:
    - Iniciar sesión / registrarse.
    - Explorar contenido audiovisual y de audio.
    - Reproducir, valorar, ver historial y favoritos.
    - Gestionar perfiles (según plan).
    - Acceder a opciones de administración (admin).

- **API HTTP REST**  
  - Endpoints para:
    - Registro y login.
    - Consulta de planes.
    - Cambio de plan con tarjeta simulada.
    - Listado de contenido audiovisual y de audio.
    - Registro de valoraciones (ratings).

---

## 🧱 Arquitectura del Proyecto

Arquitectura en capas con separación clara de responsabilidades:

- **cmd/**
  - `sdge/` → Aplicación de consola (menús interactivos).
  - `sdge/web/` → Servidor HTTP (API REST).

- **internal/**
  - `db/` → Conexión y creación del esquema SQLite.
  - `models/` → Modelos de dominio (`User`, `Profile`, `Plan`, `Audiovisual`, `Audio`, etc.).
  - `repositories/` → Acceso a datos (`UserRepo`, `ContentRepo`, `SubscriptionRepo`, `PlaybackHistoryRepo`, `FavoriteRepo`, …).
  - `services/` → Lógica de negocio (`UserService`, `ContentService`, `SubscriptionService`, `PlaybackService`, `ProfileService`).
  - `security/` → Hash de contraseñas y utilidades de seguridad.
  - `utils/` → Funciones auxiliares (validaciones, helpers).
  - `httpapi/` → Handlers HTTP de la API (rutas, parseo de JSON, respuestas).

Esta estructura facilita pruebas, mantenimiento y extensiones futuras.

---

## ⚙️ Estilo de Código

- **Programación orientada a objetos en Go**
  - Uso de `struct` para representar entidades de dominio.
  - Interfaces para abstraer repositorios y servicios donde es necesario.
  - Métodos asociados a servicios para encapsular la lógica de negocio.

- **Buenas prácticas**
  - Validaciones tempranas de entrada.
  - Manejo explícito de errores (`error`) y mensajes claros al usuario.
  - Separación de responsabilidades por paquetes y capas.

---

## 📦 Dependencias Externas

El proyecto utiliza las siguientes librerías de Go:

| Librería                                                                              | Propósito                                  |
|---------------------------------------------------------------------------------------|--------------------------------------------|
| [`golang.org/x/crypto/bcrypt`](https://pkg.go.dev/golang.org/x/crypto/bcrypt)        | Hashing seguro de contraseñas (`bcrypt`). |
| [`github.com/mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3)                 | Driver SQLite3 para `database/sql`.       |

> ⚠️ **Nota:** `go-sqlite3` requiere **CGO habilitado** y un **compilador C** instalado.  
> En Windows se utilizó **MSYS2 / mingw-w64**.

---

## 🛠️ Requisitos del Entorno

- **Lenguaje:** Go ≥ 1.20  
- **Editor recomendado:** Visual Studio Code (con terminal integrada)  
- **Control de versiones:** Git  
- **Base de datos:** SQLite (archivo `sdgestreaming.db`)  
- **Compilador C:** necesario para `go-sqlite3`
  - Windows: MSYS2 / mingw-w64.
  - Linux / macOS: `gcc` o `clang`.

Antes de ejecutar:

```powershell
# Windows (PowerShell)
$env:CGO_ENABLED = "1"

📂 **Estructura del Proyecto (simplificada)**
text

Copy

cmd/

  sdge/          -> Aplicación de consola (menús interactivos)
  
    main.go
 
  sdge/web/      -> Servidor HTTP (API REST)
  
    main.go

internal/

  db/            -> Conexión y creación de esquema SQLite
  
  models/        -> Modelos de dominio (User, Plan, Content, Profile, etc.)
 
  repositories/  -> Acceso a datos (UserRepo, ContentRepo, SubscriptionRepo…)
  
  services/      -> Lógica de negocio (UserService, ContentService, PlaybackService…)
  
  security/      -> Hash de contraseñas, validaciones básicas
  
  utils/         -> Funciones auxiliares
 
  httpapi/       -> Handlers HTTP del API

👤 **Usuario administrador por defecto**
Al iniciar la aplicación (consola o API) se crea automáticamente un usuario administrador si no existe:

Email: admin@sdge.com
Contraseña: admin123
Plan: Premium 4K
Clasificación de edad: Adulto
Este usuario sirve para pruebas rápidas de inicio de sesión y acceso a menús de administración.

🚀 **Instrucciones para Ejecutar el Proyecto**
En el terminal integrado de VS Code (o cualquier terminal):

1. Clonar el repositorio
git clone https://github.com/IsraelRiveraSxEc/SDGEStreaming.git
cd SDGEStreaming
2. Verificar instalación de Go
go version   # Debe mostrar una versión válida de Go
3. Instalar compilador C
Instalar el compilador C de preferencia para su sistema operativo y reiniciar el sistema.
En esta fase AA2 se utilizó MSYS2 para Windows.

Verificar instalación con:
gcc --version
4. Descargar dependencias Go
go mod tidy
▶️ Ejecutar la aplicación de consola
Esta es la interfaz principal para el usuario final (menús de texto).

Desde la raíz del proyecto:
go run ./cmd/sdge
La aplicación:

Creará (si no existe) la base de datos sdgestreaming.db.
Creará los planes por defecto y contenido inicial.
Creará el usuario administrador por defecto.
Mostrará el menú principal en la consola.

🌐 **Ejecutar el API HTTP**
El API HTTP expone parte de la funcionalidad para ser usada desde clientes externos.

Desde la raíz del proyecto:
go run ./cmd/sdge/web
En la consola se mostrará:
Servidor HTTP escuchando en http://localhost:8080
El servidor:

Usa la misma base de datos sdgestreaming.db que la aplicación de consola.
Permite probar operaciones de registro, login, consulta de planes, consulta de contenido y valoraciones.

📡 **Endpoints del API**
1. Autenticación / Usuarios
POST /api/register
Registra un nuevo usuario (no administrador).

Body (JSON):
{
  "name": "Juan Pérez",
  "age": 20,
  "email": "juan@example.com",
  "password": "secreto"
}
Respuesta 201:

{
  "message": "Usuario registrado exitosamente",
  "user_id": 2,
  "email": "juan@example.com"
}
POST /api/login
Inicia sesión de un usuario existente.

Body (JSON):

{
  "email": "juan@example.com",
  "password": "secreto"
}
Respuesta 200:

{
  "message": "Inicio de sesión exitoso",
  "user_id": 2,
  "email": "juan@example.com",
  "plan_id": 1,
  "plan_name": "Free",
  "is_admin": false
}
Si las credenciales son incorrectas, devuelve 401.

2. Planes y suscripciones
GET /api/plans
Obtiene la lista de planes disponibles.

Ejemplo de respuesta 200:
[
  {
    "ID": 1,
    "Name": "Free",
    "Price": 0,
    "MaxQuality": "SD",
    "MaxDevices": 1
  },
  {
    "ID": 2,
    "Name": "Estándar",
    "Price": 9.99,
    "MaxQuality": "HD",
    "MaxDevices": 2
  },
  {
    "ID": 3,
    "Name": "Premium 4K",
    "Price": 15.99,
    "MaxQuality": "4K",
    "MaxDevices": 4
  }
]
POST /api/subscriptions/change-plan
Cambia el plan de un usuario simulando un pago con tarjeta.

Body (JSON):

{
  "user_id": 2,
  "plan_id": 3,
  "card_holder": "Juan Pérez",
  "card_number": "4111111111111111",
  "expiry_month": 12,
  "expiry_year": 2030,
  "cvv": 123
}
Respuesta 200:

{
  "message": "Plan actualizado exitosamente"
}
En caso de error de validación, responde con 400 y un mensaje descriptivo.

3. Contenido
GET /api/content/audiovisual
Devuelve todo el contenido audiovisual disponible.

GET http://localhost:8080/api/content/audiovisual
La respuesta es una lista de elementos audiovisuales (puede variar según los datos cargados).

GET /api/content/audio
Devuelve todo el contenido de audio (música, podcasts, etc.).

GET http://localhost:8080/api/content/audio
La respuesta puede ser null si aún no hay contenido de audio registrado.

4. Valoraciones (Ratings)
Permite que un usuario valore contenido audiovisual o de audio
con una nota de 1.0 a 10.0.

POST /api/content/audiovisual/rate
Valora un contenido audiovisual.

POST http://localhost:8080/api/content/audiovisual/rate
Content-Type: application/json
Body (JSON):
{
  "user_id": 3,
  "content_id": 1,
  "rating": 8.5
}
POST /api/content/audio/rate
Valora un contenido de audio.

POST http://localhost:8080/api/content/audio/rate
Content-Type: application/json
Body (JSON):

{
  "user_id": 3,
  "content_id": 1,
  "rating": 9.0
}
Respuesta 200 en ambos casos:

{
  "message": "Valoración registrada correctamente"
}
Reglas de validación:

user_id y content_id deben ser mayores que 0.
rating debe estar entre 1.0 y 10.0.
Si el contenido no existe o hay un problema en la lógica de negocio, responde con 400.

🧪 **Ejemplos rápidos con PowerShell**
Con el servidor corriendo (go run ./cmd/sdge/web):

# Obtener planes
(Invoke-WebRequest `
  -Uri "http://localhost:8080/api/plans" `
  -Method GET `
  -UseBasicParsing).Content

# Valorar audiovisual (content_id=1) como user_id=3
(Invoke-WebRequest `
  -Uri "http://localhost:8080/api/content/audiovisual/rate" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"user_id":3,"content_id":1,"rating":8.5}' `
  -UseBasicParsing).Content

# Valorar audio
(Invoke-WebRequest `
  -Uri "http://localhost:8080/api/content/audio/rate" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"user_id":3,"content_id":1,"rating":9.0}' `
  -UseBasicParsing).Content

## 📖 Bibliografía

Chacon, S., & Straub, B. (2023). *Pro Git* (7.ª ed.). Apress. Recuperado de https://git-scm.com/book/en/v2  

The Go Authors. (s. f.). *The Go Programming Language Specification*. Go Documentation. Recuperado el 2025 de https://go.dev/ref/spec  

The Go Authors. (s. f.). *Package net/http*. Go Documentation. Recuperado el 2025 de https://pkg.go.dev/net/http  

The Go Authors. (s. f.). *Package database/sql*. Go Documentation. Recuperado el 2025 de https://pkg.go.dev/database/sql  

The Go Authors. (s. f.). *Package encoding/json*. Go Documentation. Recuperado el 2025 de https://pkg.go.dev/encoding/json  

The Go Authors. (2025). *Package bcrypt – golang.org/x/crypto/bcrypt*. Go Documentation. Recuperado de https://pkg.go.dev/golang.org/x/crypto/bcrypt  

MSYS2 Project. (2025). *MSYS2 – Software distribution and building platform for Windows*. Recuperado de https://www.msys2.org/  

mattn. (2025). *go-sqlite3: SQLite3 driver for Go using database/sql* (Versión 1.14.16) [Código fuente]. GitHub. Recuperado de https://github.com/mattn/go-sqlite3  

Mozilla Developer Network. (s. f.). *HTTP response status codes*. MDN Web Docs. Recuperado el 2025 de https://developer.mozilla.org/en-US/docs/Web/HTTP/Status  

Fowler, M. (2002). *Patterns of Enterprise Application Architecture*. Addison-Wesley.  

Gamma, E., Helm, R., Johnson, R., & Vlissides, J. (1994). *Design Patterns: Elements of Reusable Object-Oriented Software*. Addison-Wesley.

© 2025 – Proyecto Académico SDGEStreaming – Proyecto Final
