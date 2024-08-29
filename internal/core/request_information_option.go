package core

type requestInformationOption func(*requestInformation)

func WithMethod(method HttpMethod) requestInformationOption {
	return func(ri *requestInformation) {
		ri.Method = method
	}
}

func WithPathParams(pathParameters map[string]string) requestInformationOption {
	return func(ri *requestInformation) {
		ri.uri.PathParameters = pathParameters
	}
}

func WithURITemplate(template string) requestInformationOption {
	return func(ri *requestInformation) {
		ri.uri.UrlTemplate = template
	}
}
