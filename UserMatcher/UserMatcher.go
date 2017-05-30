package main

import (
	"os"
	"fmt"
	"bufio"
	"io/ioutil"
	"strings"
	"sort"
)

type Profile struct {
	username string
	answers []bool
	rank int
}

func (prof1 *Profile) Compare(prof2 Profile) {
	score := 0
	for key, answer := range prof2.answers {
		if answer == prof1.answers[key] {
			score++
		}
	}
	prof1.rank = score
}

func (uprof Profile) Rank(profiles []Profile) []Profile {
	for i := 0; i < len(profiles); i++ {
		profiles[i].Compare(uprof)
	}

	sort.Slice(profiles, func(i, j int) bool { return profiles[i].rank > profiles[j].rank })
	return profiles

}

func (profile *Profile) AddUname(uname string) {
	profile.username = uname
}

func (profile *Profile) AddAnswer(ans bool) { 
	profile.answers = append(profile.answers, ans)
}

func (profile Profile) answersString() string {
	answers := []string{}
	for _, answer := range profile.answers {
		if answer {
			answers = append(answers, "Y")
		} else {
			answers = append(answers, "N")
		}
	}

	return strings.Join(answers, ",")
}

func (profile Profile) String() string {
	return fmt.Sprintf("%s,%s", profile.username, profile.answersString()) 
}

func MakeProfile(info string) Profile {
	stuff := strings.Split(info, ",")
	profile := Profile{}
	profile.username = stuff[0]
	
	for _, answer := range stuff[1:] {
		profile.AddAnswer(answer == "Y")
	}

	return profile
}

func createReader() (*bufio.Reader) {
    return bufio.NewReader(os.Stdin) 
} 

func AskQuestion(question string) string {
    fmt.Print(question)
	var reader *bufio.Reader = createReader()
	var response string
    response, _ = reader.ReadString('\n')
    response = strings.TrimSpace(response)
	return response
    
}

func IsYN(response string) bool {
    if response == "Y" {
        return true
    } else {
        return false
	}   
}

func getQuestionsArray()  []string {
    questions := []string {
        "Please enter a username: ",
        "Are you interested in Engineering Competitions? (Y/N): ",
        "Do you like pizza? (Y/N): ",
		"Are you interested in Automotive Engineering Projects? (Y/N): ? ",
     }
    return questions
}

func boolsToYN(boolean bool) string {
	if boolean {
		return "Y"
	} else {
		return "N"
	}

}

func UpdateDatabase(user Profile, fname string) {
	strProf := string(user.username) + "," + boolsToYN(user.answers[0]) + "," + boolsToYN(user.answers[1]) + "," + boolsToYN(user.answers[2])
	f, err := os.OpenFile(string(fname) + ".txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
    	panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString("\r\n" + strProf); err != nil {
   		panic(err)
	}
}


func main() {
    var questions = getQuestionsArray()
    var User Profile
    User.AddUname(AskQuestion(questions[0]))
    for questionIndex := 1; questionIndex < len(questions); questionIndex ++ {
            User.AddAnswer(IsYN(AskQuestion(questions[questionIndex])))
    }
    
	b, err := ioutil.ReadFile("database.txt")
	if err != nil {
		panic(err)

	}
	k := strings.Split(string(b),"\r\n")

	profiles := []Profile{}
	for _, value := range k {
		profiles = append(profiles, MakeProfile(value))
	}
	
	User.Rank(profiles)

	fmt.Println("You have the most in common with " + profiles[0].username + "!\nWhy don't you strike up a conversation about your common interests?")

	UpdateDatabase(User, "database")
}