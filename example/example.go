package main

import wh "../webhook"

var exArr [][]string

// ExecuteThisWhen helper method for building an arr of executable statements
func ExecuteThisWhen(this string, when string) {
	exStatement := []string{this, when}
	exArr = append(exArr, exStatement)
}

func main() {

	// Execute 'date' when the authors email is lolwut@noway.biz
	ExecuteThisWhen("date", "commits.0.author.email eq lolwut@noway.biz")

	// More logical tests
	ExecuteThisWhen("whoami", "after eq 1481")
	ExecuteThisWhen("uname", "commits.1.committer.username eq octokitty")

	// Start web server on http://localhost:8080/
	wh.Startup(exArr)

}
