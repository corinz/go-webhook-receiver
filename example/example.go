package main

import wh "../webhook"

func main() {

	// Execute 'date' when the authors email is lolwut@noway.biz
	wh.ExecuteThisWhen("date", "commits.0.author.email eq lolwut@noway.biz")

	// More logical tests
	wh.ExecuteThisWhen("whoami", "after eq 1481")
	wh.ExecuteThisWhen("uname", "commits.1.committer.username eq octokitty")

	// Start web server on http://localhost:8080/
	wh.Startup("/example")

}
