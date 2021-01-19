package main

import wh "github.com/corinz/go-webhook-receiver"

func main() {

	// Create jsonwebhook
	var incomingWH *wh.JSONWebhook
	incomingWH = new(wh.JSONWebhook)

	// Execute when logic test passes
	incomingWH.AddExecutable("whoami", "after eq 1481a2de7b2a7d02428ad93446ab166be7793fbb")
	incomingWH.AddExecutable("date", "commits.0.author.email eq lolwut@noway.biz")
	incomingWH.AddExecutable("uname", "commits.1.committer.username eq octokitty")

	// Start web server on http://localhost:8080/
	Startup(incomingWH)

}
