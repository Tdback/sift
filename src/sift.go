package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"
	// "path/filepath"
)

// TODO:
// - Need to clean up big time before beginning to implement features
// - Use goroutine to grab full path to directory
// - Alphabatize results

// Maybe make this generic
func showing_it (show bool, c []fs.DirEntry) {
	for _, entry := range c {
		if !show {
			if !strings.HasPrefix(entry.Name(), ".") {
				fmt.Printf("%s\n", entry.Name())
			}
		} else {
			fmt.Printf("%s\n", entry.Name())
		}
	}
}

func helpPage() {
	// TODO: Add more formatting to this and better descriptions.
	fmt.Fprintln(os.Stderr, "Description...")
	fmt.Fprintln(os.Stderr, "\nUsage:")
	fmt.Fprintln(os.Stderr, "\n\tsift [arguments]...")
	fmt.Fprintln(os.Stderr, "\nThe arguments are:")
	fmt.Fprintln(os.Stderr, "\n\t-., -hidden\n\t\tShow hidden dotfiles and directories.\n\t\tDisabled by default.")
	fmt.Fprintln(os.Stderr, "\n\t-t, -type string\n\t\tType of file to return.")
	fmt.Fprintln(os.Stderr, "\n\t-e, -ext  string\n\t\tOnly return files ending in [string].")
}

func main() {
	// TODO:
	// - Still need to implement the following arguments:
	var show_dots bool
	flag.BoolVar(&show_dots, ".", false, "Include dotfiles\n")
	flag.BoolVar(&show_dots, "hidden", false, "Include dotfiles\n")

	var file_type string
	flag.StringVar(&file_type, "t", "", "Type of file returned (shorthand)")
	flag.StringVar(&file_type, "type", "", "Type of file returned")

	var ext_type string
	flag.StringVar(&ext_type, "e", "", "Only return files ending in [string] (shorthand)")
	flag.StringVar(&ext_type, "ext", "", "Only return files ending in [string]")

	var recursive bool
	flag.BoolVar(&recursive, "r", false, "Recursively walk through directories (shorthand)")
	flag.BoolVar(&recursive, "recurse", false, "Recursively walk through directories")

	var count bool
	flag.BoolVar(&count, "c", false, "Count number of files returned (shorthand)")
	flag.BoolVar(&count, "count", false, "Count number of files returned")

	flag.Usage = helpPage
	flag.Parse()

	// Set to current directory if we don't pass one in

	// Check if we're accepting from stdin
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Read from stdin pipe
		scanner := bufio.NewScanner(os.Stdin)

		// Always show dotfiles when reading from stdin
		show_dots = true

		for scanner.Scan() {

			fmt.Println(scanner.Text())
			// Need to implement a way to parse the data read from stdin,
			// especially if it's from a tool like `walk'. Should have the
			// function be generic enough to accept a list of files and maybe
			// parse through them.
			
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	} else {
		if c, err := os.ReadDir("."); err != nil {
			fmt.Sprintln(err)
		} else {
			showing_it(show_dots, c)
		}
	}
	
	if pager := os.Getenv("PAGER"); pager != "" {
		fmt.Println(pager)
	} else {
		fmt.Println("Defaulting to `less'")
	}

	if file_type != "" {
		fmt.Println("Filetype:", file_type)
	}
}
