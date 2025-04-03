package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID   int
	Name string
}

func NewUser(id int, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}

type Server struct {
	db    map[int]*User
	cache map[int]*User
	dbhit int
}

func NewServer() *Server {
	users := make(map[int]*User)
	for i := 1; i <= 100; i++ {
		users[i] = NewUser(i, fmt.Sprintf("user_%d", i))
	}
	return &Server{
		db:    users,
		cache: make(map[int]*User),
		dbhit: 0,
	}
}

func (s *Server) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	strId, _ := strconv.Atoi(id)

	// try to get the user from the cache
	user, ok := s.TryCache(strId)
	if ok {
		json.NewEncoder(w).Encode(user)
		return
	}

	// we hit the db here
	user, ok = s.db[strId]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	s.dbhit++

	// insert into the cache
	s.cache[strId] = user
	json.NewEncoder(w).Encode(user)
}

func (s *Server) TryCache(id int) (*User, bool) {
	user, ok := s.cache[id]
	return user, ok
}

func main() {

}
