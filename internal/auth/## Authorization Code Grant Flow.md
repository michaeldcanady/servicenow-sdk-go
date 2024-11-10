## Authorization Code Grant Flow
https://docs.servicenow.com/bundle/vancouver-platform-security/page/administer/security/concept/c_OAuthAuthorizationCodeFlow.html

1. send GET request to: https://myinstance.service-now.com/oauth_auth.do?response_type=code&redirect_uri={the_redirect_url}&client_id={the_client_identifier}
2. which will respond with: https/http://{callbackURL}?code={the actual auth code}
3. send POST request to: https://myinstance.service-now.com/oauth_token.do?grant_type=authorization_code&code={the auth code}&redirect_uri={the_same_redirect_url}&client_id={the_same_client_identifier}&client_secret={client_secret_value}
4. which will respond with: https://myinstance.service-now.com/api/now/table/incident?access_token={the_token}

## Implicit Grant Flow
https://docs.servicenow.com/bundle/vancouver-platform-security/page/administer/security/concept/c_OAuthImplicitGrants.html

1. Send GET request to: https://myinstance.servicenow.com/oauth_auth.do?response_type=token&redirect_uri={the_redirect_url}&client_id={the_client_identifier}
2. which will respond with: https/http://{callbackURL}?access_token={the_token}

### External clients (Machine to machine)
https://docs.servicenow.com/bundle/vancouver-platform-security/page/administer/security/task/create-jwt-endpoint.html
