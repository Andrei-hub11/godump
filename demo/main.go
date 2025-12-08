package main

import (
	"fmt"
	"os"

	"github.com/goforj/godump"
)

type Profile struct {
	Age   int
	Email string
}

type User struct {
	Name    string
	Profile Profile
}

func (u User) String() string {
    return fmt.Sprintf("User: %s", u.Name)
}

func main() {
	user := User{
		Name: "Alice",
		Profile: Profile{
			Age:   30,
			Email: "alice@example.com",
		},
	}

	// Pretty-print to stdout
	fmt.Println("--- godump.Dump(user) ---")
	godump.Dump(user)

	// Get dump as string
	fmt.Println("\n--- output := godump.DumpStr(user) ---")
	output := godump.DumpStr(user)
	fmt.Println(output)

	// HTML for web UI output
	fmt.Println("\n--- html := godump.DumpHTML(user) ---")
	html := godump.DumpHTML(user)
	fmt.Println(html)

	// Print JSON directly to stdout
	fmt.Println("\n--- godump.DumpJSON(user) ---")
	godump.DumpJSON(user)

	// Write to any io.Writer (e.g. file, buffer, logger)
	fmt.Fprintln(os.Stderr, "\n--- godump.Fdump(os.Stderr, user) ---")
	godump.Fdump(os.Stderr, user)

	// Dump and exit
	// I'll leave this commented out so we can see all the output
	// godump.Dd(user) // this will print the dump and exit the program
} 