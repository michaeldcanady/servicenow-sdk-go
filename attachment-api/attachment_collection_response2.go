package attachmentapi

import (
	"errors"
	"regexp"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

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

type attachmentCollectionResponse2 struct {
	backingStore store.BackingStore
}

func NewAttachmentCollectionResponse2(node serialization.ParseNode) (serialization.Parsable, error) {
	return &attachmentCollectionResponse2{
		backingStore: store.NewInMemoryBackingStore(),
	}, nil
}

func CreateAttachmentCollectionResponse2FromDiscriminatorValue() serialization.ParsableFactory {
	return NewAttachmentCollectionResponse2
}

func (tE *attachmentCollectionResponse2) GetBackingStore() store.BackingStore {
	if internal.IsNil(tE) {
		return nil
	}

	if internal.IsNil(tE.backingStore) {
		tE.backingStore = store.BackingStoreFactoryInstance()
	}

	return tE.backingStore
}

// Serialize writes the objects properties to the current writer.
func (tE *attachmentCollectionResponse2) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(tE) {
		return nil
	}
	return nil
}

func (tE *attachmentCollectionResponse2) GetResult() ([]Attachmentable, error) {
	if internal.IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get("Result")
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

func (tE *attachmentCollectionResponse2) setResult(result []Attachmentable) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set("Result", result)
}

func (tE *attachmentCollectionResponse2) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(tE) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		"result": func(pn serialization.ParseNode) error {
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

func (tE *attachmentCollectionResponse2) setNextLink(nextLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(nextLinkHeaderKey, nextLink)
}

func (tE *attachmentCollectionResponse2) setPreviousLink(previousLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(prevLinkHeaderKey, previousLink)
}

func (tE *attachmentCollectionResponse2) setFirstLink(firstLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(firstLinkHeaderKey, firstLink)
}

func (tE *attachmentCollectionResponse2) setLastLink(lastLink *string) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(lastLinkHeaderKey, lastLink)
}

const (
	firstLinkHeaderKey = "first"
	prevLinkHeaderKey  = "prev"
	nextLinkHeaderKey  = "next"
	lastLinkHeaderKey  = "last"
)

var (
	linkHeaderRegex = regexp.MustCompile(`<([^>]+)>;rel="([^"]+)"`)
)

// parseNavLinkHeaders parses navigational links and applies the to the provided response.
func parseNavLinkHeaders(hearderLinks []string, resp AttachmentCollectionResponse2) error {
	for _, header := range hearderLinks {
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
