package authentication

type authorizationCodeQueryParameters struct {
	responseType string `uri:"response_type"`
	redirectURI  string `uri:"redirect_uri"`
	clientID     string `uri:"client_id"`
	state        string `uri:"state"`
}
