package main

import "fmt"

func main() {
	states := make(map[string]string)

	states["New York"] = "New York City"
	states["California"] = "Los Angeles"
	states["Illinois"] = "Chicago"
	states["Texas"] = "Houston"
	states["Arizona"] = "Phoenix"

	//Get Map Length
	fmt.Println("len:", len(states))

	// Delete from map
	delete(states, "Texas")

	//Print map
	fmt.Println("map:", states)

	//Check if key exists
	_, prs := states["Arizona"]
	fmt.Println("prs:", prs)

	//Loop over keys and values
	for state, city := range states {
		fmt.Println("key:", state)
		fmt.Println("value:", city)
	}
}
