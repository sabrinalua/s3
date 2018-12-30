package gos3

import (
	"encoding/xml"
	"time"
)

/*https://docs.aws.amazon.com/AmazonS3/latest/API/RESTServiceGET.html*/
type ListBuckets struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListAllMyBucketsResult" json:"-"`
	Owner initiator
	Buckets []BucketInfo `xml:"Buckets>Bucket"`
}

type BucketInfo struct {
	// The name of the bucket.
	Name string `json:"name"`
	// Date the bucket was created.
	CreationDate time.Time `json:"creationDate"`
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

// ListBucketV2Result container for listObjects response version 2.
type ListBucketV2Result struct {
	// A response can contain CommonPrefixes only if you have
	// specified a delimiter.
	CommonPrefixes []CommonPrefix
	// Metadata about each object returned.
	Contents  []ObjectInfo
	Delimiter string

	// Encoding type used to encode object keys in the response.
	EncodingType string

	// A flag that indicates whether or not ListObjects returned all of the results
	// that satisfied the search criteria.
	IsTruncated bool
	MaxKeys     int64
	Name        string

	// Hold the token that will be sent in the next request to fetch the next group of keys
	NextContinuationToken string

	ContinuationToken string
	Prefix            string

	// FetchOwner and StartAfter are currently not used
	FetchOwner string
	StartAfter string
}

// ListBucketResult container for listObjects response.
type ListBucketResult struct {
	// A response can contain CommonPrefixes only if you have
	// specified a delimiter.
	CommonPrefixes []CommonPrefix
	// Metadata about each object returned.
	Contents  []ObjectInfo
	Delimiter string

	// Encoding type used to encode object keys in the response.
	EncodingType string

	// A flag that indicates whether or not ListObjects returned all of the results
	// that satisfied the search criteria.
	IsTruncated bool
	Marker      string
	MaxKeys     int64
	Name        string

	// When response is truncated (the IsTruncated element value in
	// the response is true), you can use the key name in this field
	// as marker in the subsequent request to get next set of objects.
	// Object storage lists objects in alphabetical order Note: This
	// element is returned only if you have delimiter request
	// parameter specified. If response does not include the NextMaker
	// and it is truncated, you can use the value of the last Key in
	// the response as the marker in the subsequent request to get the
	// next set of object keys.
	NextMarker string
	Prefix     string
}
