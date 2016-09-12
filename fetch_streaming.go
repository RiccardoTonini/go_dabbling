// +build fetch

package main

// Prints the content found at a URL.
import (
 	"fmt";
 	// "io/ioutil";
 	"io";
 	"net/http";
 	"os";
 	"strings";
)

func normalize_url (url string)  string {
	// fmt.Printf("The type of url is: %t", url)

	if strings.Contains(url, "://"){
		return url
	}

	url_bits := []string{"http://", url}
	return strings.Join(url_bits, "")
}

func check(err *error, extra string) {

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error. %v\t (%s) \n", err, extra)
		os.Exit(1)
	}
}

// checkClose is used to:
// 1) close resource
// 2) check the return from Close in a defer statement
func checkClose(c io.Closer, err *error) {
	cerr := c.Close()
	if *err == nil {
		*err = cerr
	}

	fmt.Println("checkClose *err is ", *err)
}


func main () {

	for _, url := range os.Args[1:] {
		normalizedUrl := ""
		normalizedUrl = normalize_url(url)
		// Get from net/http package returns the result in a response struct
		resp, err := http.Get(normalizedUrl)
		check(err, "Error happened fetching url")

		defer resp.Body.Close() // Closing to avoid leaking resources

		// resp.Body is a readable stream
		// use io.Copy(out, resp.Body) instead of ioutil.ReadAll(resp.Body)
		// Give io.Copy an io.Reader and an io.Writer
		// it copies the data over, one small chunk at a time.
		// b, err := ioutil.ReadAll(resp.Body)

		_, copyErr := io.Copy(os.Stdout, resp.Body);
		check(copyErr, "Error happened copying reader resp.Body to writer os.Stdout")

	}
}