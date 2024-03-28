package storage

type Storage interface {
	Save(shortURL, originalURL string)
	Load(shortURL string) (string, bool)
}