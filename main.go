package main

func main() {

	// Create jsonwebhook
	var incomingWH *JSONWebhook
	incomingWH = new(JSONWebhook)

	// TODO this logic doesnt work, spotty, need unit tests
	incomingWH.AddExecutable("whoami", "after re 1481a2de7b2a7d02428ad93446ab166be7793fbb")
	incomingWH.AddExecutable("date", "commits.0.author.email eq lolwut@noway.biz")

	Startup(incomingWH)

}
