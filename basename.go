package main

func main() {

}

/*If input is a string - a/b/c.go then we remove the directory component and a . suffix
then output is c
*/
func basename(path string) string {
	var s string
	// Step 1 : Reverse the string
	for i := len(path)0; i > 0; i-- {
		s += path[]
	}
}
