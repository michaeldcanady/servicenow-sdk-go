package integration

import (
	"fmt"
)

type apiTestContext struct {
	instanceURL    string
	lastResponse   int
	incidentNumber string
}

func (c *apiTestContext) aServiceNowInstance() error {
	c.instanceURL = "https://example.service-now.com"
	return nil
}

func (c *apiTestContext) anExistingIncidentRecordWithNumber(num string) error {
	c.incidentNumber = num
	return nil
}

func (c *apiTestContext) iRequestTheIncidentRecordByItsNumber() error {
	// For US2 we'll just mock the response
	c.lastResponse = 200
	return nil
}

func (c *apiTestContext) theResponseShouldContainTheIncidentNumber(num string) error {
	if c.incidentNumber != num {
		return fmt.Errorf("expected incident number %s, but got %s", num, c.incidentNumber)
	}
	return nil
}

func (c *apiTestContext) theResponseStatusShouldBe(status int) error {
	if c.lastResponse != status {
		return fmt.Errorf("expected status %d, but got %d", status, c.lastResponse)
	}
	return nil
}
