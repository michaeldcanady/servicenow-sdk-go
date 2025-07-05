package mocking

import (
	"time"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"

	"github.com/stretchr/testify/mock"
)

type MockFile struct {
	mock.Mock
}

func NewMockFile() *MockFile {
	return &MockFile{
		Mock: mock.Mock{},
	}
}

func (mock *MockFile) GetAverageImageColor() (*string, error) {
	args := mock.Called()

	return args.Get(0).(*string), args.Error(1)
}

func (mock *MockFile) SetAverageImageColor(averageImageColor *string) error {
	args := mock.Called(averageImageColor)

	return args.Error(0)
}

func (mock *MockFile) GetCompressed() (*bool, error) {
	args := mock.Called()

	return args.Get(0).(*bool), args.Error(1)
}

func (mock *MockFile) SetCompressed(compressed *bool) error {
	args := mock.Called(compressed)

	return args.Error(0)
}

func (mock *MockFile) GetContentType() (*string, error) {
	args := mock.Called()

	return args.Get(0).(*string), args.Error(1)
}

func (mock *MockFile) SetContentType(contentType *string) error {
	args := mock.Called(contentType)

	return args.Error(0)
}

func (mock *MockFile) GetCreatedByName() (*string, error) {
	args := mock.Called()

	return args.Get(0).(*string), args.Error(1)
}

func (mock *MockFile) SetCreatedByName(createdByName *string) error {
	args := mock.Called(createdByName)

	return args.Error(0)
}

func (mock *MockFile) GetDownloadLink() (*string, error) {
	args := mock.Called()

	return args.Get(0).(*string), args.Error(1)
}

func (mock *MockFile) SetDownloadLink(downloadLink *string) error {
	args := mock.Called(downloadLink)

	return args.Error(0)
}

func (mock *MockFile) GetFileName() (*string, error) {
	args := mock.Called()

	return args.Get(0).(*string), args.Error(1)
}

func (mock *MockFile) SetFileName(fileName *string) error {
	args := mock.Called(fileName)

	return args.Error(0)
}

func (mock *MockFile) GetImageHeight() (*float64, error) {
	args := mock.Called()

	return args.Get(0).(*float64), args.Error(1)
}

func (mock *MockFile) SetImageHeight(imageHeight *float64) error {
	args := mock.Called(imageHeight)

	return args.Error(0)
}

func (mock *MockFile) GetImageWidth() (*float64, error) {
	args := mock.Called()

	return args.Get(0).(*float64), args.Error(1)
}

func (mock *MockFile) SetImageWidth(imageWidth *float64) error {
	args := mock.Called(imageWidth)

	return args.Error(0)
}

func (mock *MockFile) GetSizeBytes() (*int64, error) {
	args := mock.Called()

	return args.Get(0).(*int64), args.Error(1)
}

func (mock *MockFile) SetSizeBytes(sizeBytes *int64) error {
	args := mock.Called(sizeBytes)

	return args.Error(0)
}

func (mock *MockFile) GetSizeCompressed() (*int64, error) {
	args := mock.Called()

	return args.Get(0).(*int64), args.Error(1)
}

func (mock *MockFile) SetSizeCompressed(compressedSize *int64) error {
	args := mock.Called(compressedSize)

	return args.Error(0)
}

func (mock *MockFile) GetSysCreatedBy() (*string, error) {
	args := mock.Called()

	return args.Get(0).(*string), args.Error(1)
}

func (mock *MockFile) SetSysCreatedBy(sysCreatedBy *string) error {
	args := mock.Called(sysCreatedBy)

	return args.Error(0)
}

func (mock *MockFile) GetSysCreatedOn() (*time.Time, error) {
	args := mock.Called()

	return args.Get(0).(*time.Time), args.Error(1)
}

func (mock *MockFile) SetSysCreatedOn(sysCreatedOn *time.Time) error {
	args := mock.Called(sysCreatedOn)

	return args.Error(0)
}

func (mock *MockFile) GetSysID() (*string, error) {
	args := mock.Called()

	return args.Get(0).(*string), args.Error(1)
}

func (mock *MockFile) SetSysID(sysID *string) error {
	args := mock.Called(sysID)

	return args.Error(0)
}

func (mock *MockFile) GetSysModCount() (*int64, error) {
	args := mock.Called()

	return args.Get(0).(*int64), args.Error(1)
}

func (mock *MockFile) SetSysModCount(sysModCount *int64) error {
	args := mock.Called(sysModCount)

	return args.Error(0)
}

func (mock *MockFile) GetSysTags() ([]string, error) {
	args := mock.Called()

	return args.Get(0).([]string), args.Error(1)
}

func (mock *MockFile) SetSysTags(sysTags []string) error {
	args := mock.Called(sysTags)

	return args.Error(0)
}

func (mock *MockFile) GetSysUpdatedBy() (*string, error) {
	args := mock.Called()

	return args.Get(0).(*string), args.Error(1)
}

func (mock *MockFile) SetSysUpdatedBy(sysUpdatedBy *string) error {
	args := mock.Called(sysUpdatedBy)

	return args.Error(0)
}

func (mock *MockFile) GetSysUpdatedOn() (*time.Time, error) {
	args := mock.Called()

	return args.Get(0).(*time.Time), args.Error(1)
}

func (mock *MockFile) SetSysUpdatedOn(sysUpdatedOn *time.Time) error {
	args := mock.Called(sysUpdatedOn)

	return args.Error(0)
}

func (mock *MockFile) GetTableName() (*string, error) {
	args := mock.Called()

	return args.Get(0).(*string), args.Error(1)
}

func (mock *MockFile) SetTableName(tableName *string) error {
	args := mock.Called(tableName)

	return args.Error(0)
}

func (mock *MockFile) GetTableSysID() (*string, error) {
	args := mock.Called()

	return args.Get(0).(*string), args.Error(1)
}
func (mock *MockFile) SetTableSysID(tableSysID *string) error {
	args := mock.Called(tableSysID)

	return args.Error(0)
}

func (mock *MockFile) GetUpdatedByName() (*string, error) {
	args := mock.Called()

	return args.Get(0).(*string), args.Error(1)
}

func (mock *MockFile) SetUpdatedByName(updatedByName *string) error {
	args := mock.Called(updatedByName)

	return args.Error(0)
}

// Serialize writes the objects properties to the current writer.
func (mock *MockFile) Serialize(writer serialization.SerializationWriter) error {
	args := mock.Called(writer)
	return args.Error(0)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (mock *MockFile) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	args := mock.Called()
	return args.Get(0).(map[string]func(serialization.ParseNode) error)
}

func (mock *MockFile) GetBackingStore() store.BackingStore {
	args := mock.Called()
	return args.Get(0).(store.BackingStore)
}
