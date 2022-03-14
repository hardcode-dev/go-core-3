package storage

import "go-dev-v3/GoSearch/pkg/crawler"

// Храниище отсканированных документов.
// Interface определяет контракт хранилища данных.
type Interface interface {
	Docs([]int) []crawler.Document
	StoreDocs([]crawler.Document) error
}
