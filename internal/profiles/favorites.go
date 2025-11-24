package profiles

// FavoriteEntry almacena el tipo y el ID
type FavoriteEntry struct {
	ContentID   int
	ContentType string // "audio" o "audiovisual"
}

// favorites: userID -> []FavoriteEntry
var favorites = make(map[int][]FavoriteEntry)

// AddFavorite añade un contenido con su tipo
func AddFavorite(userID, contentID int, contentType string) {
	favList := favorites[userID]
	// Evitar duplicados
	for _, f := range favList {
		if f.ContentID == contentID && f.ContentType == contentType {
			return
		}
	}
	favorites[userID] = append(favList, FavoriteEntry{
		ContentID:   contentID,
		ContentType: contentType,
	})
}

// GetFavorites devuelve los favoritos del usuario
func GetFavorites(userID int) []FavoriteEntry {
	return favorites[userID]
}