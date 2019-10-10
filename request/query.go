package request

import "strings"

type Query map[string][]string

func (q Query) Del(key string) {}

func (q Query) Get(key string) string {
	return strings.Join(q[key], ",")
}

func (q Query) Set(key, value string) {
	q[key] = []string{value}
}

func (q Query) Has(key string) bool {
	_, ok := q[key]
	return ok
}
