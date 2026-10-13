// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	tagsKey = "tags"
)

var _ TagListPopulationMethod = (*TagListPopulationMethodModel)(nil)

// TagListPopulationMethod represents a tag-list-based population method that uses tags to discover service relationships.
type TagListPopulationMethod interface {
	PopulationMethod
	GetTags() ([]TagListPopulationMethodTag, error)
	SetTags([]TagListPopulationMethodTag) error
}

// TagListPopulationMethodModel is the backing-store-backed implementation of [TagListPopulationMethod].
type TagListPopulationMethodModel struct {
	PopulationMethod
}

// NewTagListPopulationMethod
func NewTagListPopulationMethod() *TagListPopulationMethodModel {
	return &TagListPopulationMethodModel{
		PopulationMethod: NewPopulationMethod(),
	}
}

// CreateTagListPopulationMethodFromDiscriminatorValue
func CreateTagListPopulationMethodFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewTagListPopulationMethod(), nil
}

// GetFieldDeserializers
func (t *TagListPopulationMethodModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return extendmap(map[string]func(serialization.ParseNode) error{
		tagsKey: internalSerialization.DeserializeCollectionOfObjectValuesFunc[TagListPopulationMethodTag](CreateTagListPopulationMethodTagFromDiscriminatorValue, t.SetTags),
	}, t.PopulationMethod.GetFieldDeserializers())
}

// Serialize implements [TagListPopulationMethod].
// Subtle: this method shadows the method (PopulationMethod).Serialize of TagListPopulationMethodModel.PopulationMethod.
func (t *TagListPopulationMethodModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(t) {
		return nil
	}

	if err := t.PopulationMethod.Serialize(writer); err != nil {
		return err
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeCollectionOfObjectValuesFunc[TagListPopulationMethodTag](tagsKey, t.GetTags),
	)
}

// GetTags
func (t *TagListPopulationMethodModel) GetTags() ([]TagListPopulationMethodTag, error) {
	return store.DefaultBackedModelAccessorFunc[*TagListPopulationMethodModel, []TagListPopulationMethodTag](t, tagsKey)
}

// SetTags
func (t *TagListPopulationMethodModel) SetTags(val []TagListPopulationMethodTag) error {
	return store.DefaultBackedModelMutatorFunc(t, tagsKey, val)
}
