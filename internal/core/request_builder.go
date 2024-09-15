package core

type RequestBuilder interface {
	SendDelete2(config *RequestConfigurationImpl) error
	SendGet2(config *RequestConfigurationImpl) error
	SendPost3(config *RequestConfigurationImpl) error
	SendPut2(config *RequestConfigurationImpl) error
	ToDeleteRequestInformation(params interface{}) (*RequestInformation, error)
	ToDeleteRequestInformation2(config *RequestConfigurationImpl) (*RequestInformation, error)
	ToGetRequestInformation(params interface{}) (*RequestInformation, error)
	ToGetRequestInformation2(config *RequestConfigurationImpl) (*RequestInformation, error)
	ToHeadRequestInformation() (*RequestInformation, error)
	ToPostRequestInformation(data map[string]string, params interface{}) (*RequestInformation, error)
	ToPostRequestInformation2(data interface{}, params interface{}) (*RequestInformation, error)
	ToPostRequestInformation3(config interface{}) (*RequestInformation, error)
	ToPutRequestInformation(data map[string]string, params interface{}) (*RequestInformation, error)
	ToPutRequestInformation2(config interface{}) (*RequestInformation, error)
	ToRequestInformation(method HttpMethod, data map[string]string, params interface{}) (*RequestInformation, error)
	ToRequestInformation2(method HttpMethod, rawData interface{}, params interface{}) (*RequestInformation, error)
	ToRequestInformation3(method HttpMethod, config *interface{}) (*RequestInformation, error)
}
