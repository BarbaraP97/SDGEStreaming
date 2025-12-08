# 📺 SDGEStreaming – Sistema de Gestión de Contenido Audiovisual y de Audio

> **Fase AA2 – Paso 1**  
> Implementación de **Mi Lista (Favoritos)** e **Historial de Reproducción**  
> Proyecto académico – Programación orientada a objetos en Go

---

## 📌 Descripción General

**SDGEStreaming** es un sistema de gestión de contenido audiovisual y de audio desarrollado en **Go (Golang)**, diseñado con una arquitectura modular y separación clara de responsabilidades. Durante la **fase AA1**, se establecieron los cimientos del sistema: registro y login de usuarios, exploración y calificación del contenido.

En la **fase AA2 – Paso 1 y Paso 2**, se han activado e implementado dos funcionalidades clave anunciadas previamente:

- ✅ **Mi Lista**: gestión de favoritos por tipo de contenido (audio / audiovisual).
- ✅ **Historial de Reproducción**: registro de reproducciones con validación de acceso

El sistema se implementa **protección a menores** (clasificación por edad)
Se implenta interfaz intuitiva, segura y validada.

## 🧩 Funcionalidades Implementadas

| Funcionalidad              | Descripción                                                                 |
|---------------------------|-----------------------------------------------------------------------------|
| **Autenticación segura**  | Registro e inicio de sesión con validaciones y hashing de contraseñas.       |
| **Catálogo filtrado**     | Listado de contenido accesible según la edad del usuario activo.             |
| **Mi Lista (Favoritos)**  | Añadir y visualizar contenido favorito, diferenciando por tipo (audio/visual). |
| **Historial de Reproducción** | Registro automático tras reproducción exitosa. |
| **Calificación de ítems** | Calificación inmediata y promedio ajustado. |
| **Tendencias**            | Muestra contenido más popular según calificaciones.                         |
| **Menú administrador**    | Opciones ocultas accesibles solo para usuarios con rol `admin`.              |

---

## ⚙️ Estilo de Código
  
- **Estructuras de datos**: `struct` usadas únicamente para agrupar datos (sin comportamiento).
- **Flujo de control**: Validaciones tempranas, retorno explícito, sin excepciones.

## 📦 Dependencias Externas

El proyecto utiliza las siguientes librerías de Go:

| Librería                                      | Propósito                                                                 |
|----------------------------------------------|---------------------------------------------------------------------------|
| [`golang.org/x/crypto/bcrypt`](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | Hashing seguro de contraseñas (`bcrypt`).                                |
| [`github.com/mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3) | Soporte para persistencia en SQLite. |

> ⚠️ **Nota**: `go-sqlite3` requiere un compilador C para su correcto funciónamiento. Asegúrese de tener uno configurado en su entorno de desarrollo.

## 🛠️ Requisitos del Entorno

- **Lenguaje**: Go
- **Editor recomendado**: Visual Studio Code (con terminal integrada)
- **Control de versiones**: Git
- **Compilador C**: Requerido al momento de integrar SQLite (para `gcc` en `go-sqlite3`)

## 📖 Bibliografía

Chacon, S., & Straub, B. (2023). Pro Git (7.ª ed.). Apress. https://git-scm.com/book/en/v2
The Go Authors. (2025). Package bcrypt – golang.org/x/crypto. Go Documentation. https://pkg.go.dev/golang.org/x/crypto/bcrypt
MSYS2 Project. (2025). MSYS2 – Software distribution and building platform for Windows. https://www.msys2.org/
mattn. (2025). go-sqlite3: SQLite3 driver for Go using database/sql (Versión 1.14.16) [Código fuente]. GitHub. https://github.com/mattn/go-sqlite3

 ## 🚀 Instrucciones para Ejecutar el Proyecto
   **En el terminal integrado de VS Code**:
1. **Clonar el repositorio**:
   git clone https://github.com/IsraelRiveraSxEc/SDGEStreaming.git
2. **Asegurarse de tener Go instalado**:
   En un terminal, ejecutar:
   go version **para verificar que go esté instalado correctamente.**
3. **Instalar dependencias**
   Instalar el compilador de c de preferencia para su sistema operativo y reiniciar el sistema. `En este caso AA2 se uso MSYS2 para Windows.`
   Ejecutar el comando gcc --version para verificar la instalación.
4. **Ejecutar la aplicación**
   go run cmd/sdgestreaming/main.go

© 2025 – Proyecto Académico SDGEStreaming – Fase AA2