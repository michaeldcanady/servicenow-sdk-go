package internal

import "strings"

type ServiceNowResponseType int64

const (
	ServiceNowResponseUnknown ServiceNowResponseType = iota - 1
	ServiceNowResponseTypeCollection
	ServiceNowResponseTypeItem
)

func ParseServiceNowResponseType(str string) ServiceNowResponseType {
	str = strings.ToLower(str)

	for _, displayValue := range []ServiceNowResponseType{ServiceNowResponseUnknown, ServiceNowResponseTypeCollection, ServiceNowResponseTypeItem} {
		displayValueString := strings.ToLower(displayValue.String())
		if str == displayValueString {
			return displayValue
		}
	}

	return ServiceNowResponseUnknown
}

// String return string representation
func (p ServiceNowResponseType) String() string {
	value, ok := map[ServiceNowResponseType]string{
		ServiceNowResponseUnknown:        "unknown",
		ServiceNowResponseTypeCollection: "collection",
		ServiceNowResponseTypeItem:       "item",
	}[p]

	if !ok {
		return ServiceNowResponseUnknown.String()
	}

	return value
}
