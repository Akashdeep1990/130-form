package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func main() {

	mySlice := []contact{
		{"Good to see you", "Akash"},
		{"Good to see you", "Aman"},
		{"Good to see you", "Anshul"},
		{"Good to see you", "Ankita"},
		{"Good to see you", "Shruti"},
		{"Good to see you", "Shriya"},


	}
i:=1
	for _, currentEntry := range mySlice {
		fmt.Println(currentEntry.greeting + ", " + currentEntry.name)


	myGreeting := map[int]string{
		1:     "Good morning!",
		2:   "Bonjour!",
		3:   "Buenos dias!",
		4:  "Bongiorno!",
		5:  "Ohayo!",
		7: "Selamat pagi!",
		6:    "Gutten morgen!",
	}

	fmt.Println(myGreeting[i])
	i++
	}}
