package gos3
import (
	"encoding/xml"
)
type ErrorResponse struct {
	XMLName    xml.Name `xml:"Error" json:"-"`
	Code       string
	Message    string
	BucketName string
	Key        string
	RequestID  string `xml:"RequestId"`
	HostID     string `xml:"HostId"`

	// Region where the bucket is located. This header is returned
	// only in HEAD bucket and ListObjects response.
	Region string

	// Underlying HTTP status code for the returned error
	StatusCode int `xml:"-" json:"-"`

	// Headers of the returned S3 XML error
	Headers http.Header `xml:"-" json:"-"`
}
