package authentication

type oauthAuthQueryParameters struct {
	responseType responseType `uri:"response_type"`
	redirectURI  string       `uri:"redirect_uri"`
	clientID     string       `uri:"client_id"`
	state        string       `uri:"state"`
}
