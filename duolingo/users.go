package duolingo

import (
	"duolingo/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	TotalXp  int    `json:"totalXp"`
	GainedXp int    `json:"gainedXp"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Streak   int    `json:"streak"`
}

type XPSummary struct {
	GainedXp int `json:"gainedXp"`
}

type XPSummaries struct {
	Summaries []XPSummary `json:"summaries"`
}

func getXPGains(userId int) int {
	req, newRequestErr := http.NewRequest("GET", BaseURL+"/users/"+strconv.Itoa(userId)+"/xp_summaries?startDate="+utils.FormatBeginningOfMonth(utils.GetBeginningOfMonth()), nil)
	if newRequestErr != nil {
		println("Error creating request")
		return 0
	}
	response, doErr := http.DefaultClient.Do(req)
	if doErr != nil {
		println("Error doing request")
		return 0
	}
	var summaries XPSummaries
	decodeErr := json.NewDecoder(response.Body).Decode(&summaries)
	if decodeErr != nil {
		println("Error decoding response")
		return 0
	}
	var totalXp int
	for i := 0; i < len(summaries.Summaries); i++ {
		totalXp += summaries.Summaries[i].GainedXp
	}
	return totalXp
}

func GetUser(userId int, wg *sync.WaitGroup, users *[]User) {
	defer wg.Done()
	req, newRequestErr := http.NewRequest("GET", BaseURL+"/users/"+strconv.Itoa(userId)+"?fields=name,id,totalXp,streak", nil)
	if newRequestErr != nil {
		println("Error creating request")
	}
	response, doErr := http.DefaultClient.Do(req)
	if doErr != nil {
		panic("Error doing request")
	}
	var User User
	decodeErr := json.NewDecoder(response.Body).Decode(&User)
	if decodeErr != nil {
		println("Error decoding response")
	}
	User.GainedXp = getXPGains(User.Id)
	*users = append(*users, User)
	print(".")
}
