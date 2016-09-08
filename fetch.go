// +build fetch

package main

// Prints the content found at a URL.
import (
 	"fmt";
 	"io/ioutil";
 	"net/http";
 	"os";
 	"strings";
)

func normalize_url (url string)  string {
	fmt.Printf("The type of url is: %t", url)

	if strings.Contains(url, "://"){
		return url
	}

	url_bits := []string{"http://", url}
	return strings.Join(url_bits, "")
}

func main () {

	for _, url := range os.Args[1:] {
		normalized_url := ""
		normalized_url = normalize_url(url)
		// Get from net/http package returns the result in a response struct
		resp, err := http.Get(normalized_url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1) // Causes the process to exit with a status code of 1.
		}

		// resp.Body is a readable stream
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close() // Closing to avoid leaking resources

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		// Writes to standard output
		fmt.Printf("%s", b)
	}
}