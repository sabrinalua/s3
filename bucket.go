package gos3 

import (
	"encoding/xml"
	"time"
)

/*https://docs.aws.amazon.com/AmazonS3/latest/API/RESTServiceGET.html*/
type ListBuckets struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListAllMyBucketsResult" json:"-"`
	Owner initiator
	Buckets []Bucket `xml:"Buckets>Bucket"`
}

type Bucket struct {
	Name string 
	CreationDate time.Time
}

/*https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETacl.html */
type BucketAccessControlPolicy struct {
	XMLName xml.Name `xml:"AccessControlPolicy" json:"-"`
	Owner initiator
	AccessControlList []Grant `xml:"AccessControlList>Grant"`
}

type Grant struct {
}

/*https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETlocation.html */
type BucketLocation struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ LocationConstraint" json:"-"`
	Location    string   `xml:",chardata"`
}

/*https://docs.aws.amazon.com/AmazonS3/latest/API/v2-RESTBucketGET.html */
type ListObjectsV2Response struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListBucketResult" json:"-"`

	Name       string
	Prefix     string
	StartAfter string `xml:"StartAfter,omitempty"`
	// When response is truncated (the IsTruncated element value in the response
	// is true), you can use the key name in this field as marker in the subsequent
	// request to get next set of objects. Server lists objects in alphabetical
	// order Note: This element is returned only if you have delimiter request parameter
	// specified. If response does not include the NextMaker and it is truncated,
	// you can use the value of the last Key in the response as the marker in the
	// subsequent request to get the next set of object keys.
	ContinuationToken     string `xml:"ContinuationToken,omitempty"`
	NextContinuationToken string `xml:"NextContinuationToken,omitempty"`

	KeyCount  int
	MaxKeys   int
	Delimiter string
	// A flag that indicates whether or not ListObjects returned all of the results
	// that satisfied the search criteria.
	IsTruncated bool

	Contents       []Object
	CommonPrefixes []CommonPrefix

	// Encoding type used to encode object keys in the response.
	EncodingType string `xml:"EncodingType,omitempty"`
}

type Object struct {
	Key          string
	LastModified time.Time // time string of format "2006-01-02T15:04:05.000Z"
	ETag         string
	Size         int64

	// Owner of the object.
	Owner initiator

	// The class of storage used to store the object.
	StorageClass string
}