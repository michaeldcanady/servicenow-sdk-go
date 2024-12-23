package attachmentapi

import (
	"errors"
	"regexp"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	resultKey          = "result"
	firstLinkHeaderKey = "first"
	prevLinkHeaderKey  = "prev"
	nextLinkHeaderKey  = "next"
	lastLinkHeaderKey  = "last"
)

var (
	linkHeaderRegex = regexp.MustCompile(`<([^>]+)>;rel="([^"]+)"`)
)

// AttachmentCollectionResponse2 represents a Service-Now attachment collection response
type AttachmentCollectionResponse2 interface {
	GetResult() ([]Attachmentable, error)
	GetNextLink() (*string, error)
	GetPreviousLink() (*string, error)
	GetFirstLink() (*string, error)
	GetLastLink() (*string, error)
	setNextLink(*string) error
	setPreviousLink(*string) error
	setFirstLink(*string) error
	setLastLink(*string) error
	setResult([]Attachmentable) error
	serialization.Parsable
	store.BackedModel
}

// attachmentCollectionResponse2 implementation of AttachmentCollectionResponse2
type attachmentCollectionResponse2 struct {
	// backingStoreFactory factory to create backingStore
	backingStoreFactory store.BackingStoreFactory
	// backingStore the store backing the model
	backingStore store.BackingStore
}

// NewAttachmentCollectionResponse2 creates a new AttachmentCollectionResponse2
func NewAttachmentCollectionResponse2() AttachmentCollectionResponse2 {
	return &attachmentCollectionResponse2{
		backingStore:        store.NewInMemoryBackingStore(),
		backingStoreFactory: store.NewInMemoryBackingStore,
	}
}

// CreateAttachmentCollectionResponse2FromDiscriminatorValue is a factory for creating an AttachmentCollectionResponse2
func CreateAttachmentCollectionResponse2FromDiscriminatorValue(node serialization.ParseNode) (serialization.Parsable, error) {
	return NewAttachmentCollectionResponse2(), nil
}

// GetBackingStore returns the backing store of the record
func (tE *attachmentCollectionResponse2) GetBackingStore() store.BackingStore {
	if internal.IsNil(tE) {
		return nil
	}

	if internal.IsNil(tE.backingStore) {
		tE.backingStore = tE.backingStoreFactory()
	}

	return tE.backingStore
}

// Serialize writes the objects properties to the current writer
func (tE *attachmentCollectionResponse2) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(tE) {
		return nil
	}
	return errors.New("not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object
func (tE *attachmentCollectionResponse2) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(tE) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		resultKey: func(pn serialization.ParseNode) error {
			elem, err := pn.GetCollectionOfObjectValues(CreateAttachmentFromDiscriminatorValue)
			if err != nil {
				return err
			}
			if elem != nil {
				res := make([]Attachmentable, len(elem))
				for i, v := range elem {
					if v != nil {
						res[i] = v.(Attachmentable)
					}
				}
				return tE.setResult(res)
			}
			return nil
		},
	}
}

// GetResult returns results, a slice of attachments
func (tE *attachmentCollectionResponse2) GetResult() ([]Attachmentable, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(resultKey)
	if err != nil {
		return nil, err
	}
	if internal.IsNil(val) {
		return []Attachmentable{}, nil
	}

	typedVal, ok := val.([]Attachmentable)
	if !ok {
		return nil, errors.New("val is not Attachmentable")
	}

	return typedVal, nil
}

// setResult sets the result to the provided value
func (tE *attachmentCollectionResponse2) setResult(result []Attachmentable) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(resultKey, result)
}

// GetNextLink returns next link, if it exists
func (tE *attachmentCollectionResponse2) GetNextLink() (*string, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(nextLinkHeaderKey)
	if err != nil {
		return nil, err
	}
	if internal.IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// GetPreviousLink returns previous link, if it exists
func (tE *attachmentCollectionResponse2) GetPreviousLink() (*string, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(prevLinkHeaderKey)
	if err != nil {
		return nil, err
	}
	if internal.IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// GetFirstLink returns first link, if it exists
func (tE *attachmentCollectionResponse2) GetFirstLink() (*string, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(firstLinkHeaderKey)
	if err != nil {
		return nil, err
	}
	if internal.IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// GetLastLink returns last link, if it exists
func (tE *attachmentCollectionResponse2) GetLastLink() (*string, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(lastLinkHeaderKey)
	if err != nil {
		return nil, err
	}
	if internal.IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setNextLink sets next link
func (tE *attachmentCollectionResponse2) setNextLink(nextLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(nextLinkHeaderKey, nextLink)
}

// setPreviousLink sets previous link
func (tE *attachmentCollectionResponse2) setPreviousLink(previousLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(prevLinkHeaderKey, previousLink)
}

// setFirstLink sets first link
func (tE *attachmentCollectionResponse2) setFirstLink(firstLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(firstLinkHeaderKey, firstLink)
}

// setLastLink sets last link
func (tE *attachmentCollectionResponse2) setLastLink(lastLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(lastLinkHeaderKey, lastLink)
}

// parseNavLinkHeaders parses navigational links and applies the to the provided response.
func parseNavLinkHeaders(headerLinks []string, resp AttachmentCollectionResponse2) error {
	for _, header := range headerLinks {
		linkMatches := linkHeaderRegex.FindAllStringSubmatch(header, -1)

		for _, match := range linkMatches {
			link := match[1]
			rel := match[2]

			var err error
			// Determine the type of link based on the 'rel' attribute
			switch rel {
			case firstLinkHeaderKey:
				err = resp.setFirstLink(&link)
			case prevLinkHeaderKey:
				err = resp.setPreviousLink(&link)
			case nextLinkHeaderKey:
				err = resp.setNextLink(&link)
			case lastLinkHeaderKey:
				err = resp.setLastLink(&link)
			}
			if err != nil {
				return err
			}
		}
	}

	return nil
}
