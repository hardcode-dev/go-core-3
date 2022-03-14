package index

// Обратный индекс отсканированных документов.

import "go-dev-v3/GoSearch/pkg/crawler"

// Interface определяет контракт службы индексирования документов.
type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}
