package duolingo

import (
	"sort"
	"sync"
)

func FetchUsers() []User {
	println("Starting users fetching")
	var wg sync.WaitGroup
	var users []User
	wg.Add(len(UserIds))
	for i := 0; i < len(UserIds); i++ {
		go GetUser(UserIds[i], &wg, &users)
	}
	wg.Wait()
	println("\nDone fetching users")
	sort.Slice(users, func(i, j int) bool {
		return users[i].GainedXp > users[j].GainedXp
	})
	return users
}
