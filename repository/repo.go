package repository

import (
	"maps"
	"math/rand/v2"
	"slices"
)

type Quote struct {
	ID     int
	Author string
	Quote  string
}

var next_id int = 0
var quotes_by_ID map[int]Quote

func Initialize_repo() {
	quotes_by_ID = make(map[int]Quote)
}

func AddEntry(q Quote) {
	q.ID = next_id //игнорирует предоставленный в запросе ID и присваивает свой
	quotes_by_ID[next_id] = q
	next_id += 1
}

func GetAll() []Quote {
	return slices.Collect(maps.Values(quotes_by_ID))
}

func GetAllByAuthor(author string) []Quote {
	var quotes_by_auth []Quote
	for _, v := range quotes_by_ID { //не очень эффективно, SQL база данных (или даже redis) была бы лучше
		if v.Author == author {
			quotes_by_auth = append(quotes_by_auth, v)
		}
	}
	return quotes_by_auth
}

func GetRandom() Quote {
	all_quotes := GetAll()
	i := rand.IntN(len(all_quotes))
	return all_quotes[i]
}

func RemoveAtID(id int) {
	delete(quotes_by_ID, id)
}
