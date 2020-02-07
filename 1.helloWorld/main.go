package main // project / workspace collection of files
import "fmt"

/*
2 types of packages
executable type => splits out executable
reusable type  => code used as a helpers, dependencies

filename main responsible for splitting out runnable file.

whenever we make executable package make sure to add func main().
*/

// used to give our package access to all the code that contained inside fmt (format)
/*http://golang.org/pkg/ */

func main() { // func = function.
	fmt.Println("Hello Go!") //
}

/* Structure

package declaration. package ...
import other packages that are needed. import ....
body

 */

/*
help : go help

Run:
go run main.go => compile and run
go build => compile and build binary

go fmt => formats
go install => handle dependencies in proj
go get =>  handle dependencies in proj
go test => run and exec tests
*/
