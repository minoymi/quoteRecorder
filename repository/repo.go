package repository

import (
	"errors"
	"math/rand/v2"
)

type Quote struct {
	ID     int
	Author string
	Quote  string
}

var next_id int = 0
var quotes []*Quote //это все для быстрого поиска
var quotes_by_ID map[int]*Quote
var quotes_by_Author map[string][]*Quote

func Initialize_repo() {
	quotes_by_ID = make(map[int]*Quote)
	quotes_by_Author = make(map[string][]*Quote)
}

func AddEntry(q Quote) {
	q.ID = next_id //игнорирует предоставленный в запросе ID и присваивает свой
	quotes_by_ID[next_id] = &q
	next_id += 1

	quotes_by_Author[q.Author] = append(quotes_by_Author[q.Author], &q)
	quotes = append(quotes, &q)
}

func GetAll() []*Quote {
	return quotes
}

func GetAllByAuthor(author string) []*Quote {
	return quotes_by_Author[author]
}

func GetRandom() *Quote {
	i := rand.IntN(len(quotes))
	for quotes[i] == nil {
		i = rand.IntN(len(quotes))
	}
	return quotes[i]
}

func RemoveAtIndex(index int) error {
	if index > len(quotes) {
		return errors.New("invalid index")
	}

	entry := quotes[index]
	quotes_by_ID[entry.ID] = nil
	quotes[index] = nil
	return nil
}
