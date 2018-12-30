package gos3
// import (
//   "encoding/xml"
// )
//
// /*https://docs.aws.amazon.com/AmazonS3/latest/API/v2-RESTBucketGET.html */
// type ListObjectsV2Response struct {
// 	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListBucketResult" json:"-"`
//
// 	Name       string
// 	Prefix     string
// 	StartAfter string `xml:"StartAfter,omitempty"`
// 	// When response is truncated (the IsTruncated element value in the response
// 	// is true), you can use the key name in this field as marker in the subsequent
// 	// request to get next set of objects. Server lists objects in alphabetical
// 	// order Note: This element is returned only if you have delimiter request parameter
// 	// specified. If response does not include the NextMaker and it is truncated,
// 	// you can use the value of the last Key in the response as the marker in the
// 	// subsequent request to get the next set of object keys.
// 	ContinuationToken     string `xml:"ContinuationToken,omitempty"`
// 	NextContinuationToken string `xml:"NextContinuationToken,omitempty"`
//
// 	KeyCount  int
// 	MaxKeys   int
// 	Delimiter string
// 	// A flag that indicates whether or not ListObjects returned all of the results
// 	// that satisfied the search criteria.
// 	IsTruncated bool
//
// 	Contents       []ObjectInfo
// 	CommonPrefixes []CommonPrefix
//
// 	// Encoding type used to encode object keys in the response.
// 	EncodingType string `xml:"EncodingType,omitempty"`
// }
