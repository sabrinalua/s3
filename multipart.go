package gos3
//
// import (
// 	"time"
// 	"encoding/xml"
// )
// //data struct for initializing multipart upload
// type InitiateMultipartUploadResult struct {
// 	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ InitiateMultipartUploadResult" json:"-"`
// 	Bucket string `xml:"Bucket"`
// 	Key string `xml:"Key"`
// 	UploadId string `xml:"UploadId"`
// }
//
// //request body for complete multipart upload
// type CompleteMultipartUpload struct{
// 	Parts []PartInfo `xml:"Part" json:"parts"`
// }
//
// type PartInfo struct {
// 	PartNumber int `xml:"PartNumber" json:"part_number"`
// 	LastModified time.Time `xml:"LastModified,omitempty" json:"last_modified,omitempty"`
// 	ETag string `xml:"ETag" json:"etag"`
// 	Size int64 `xml:"Size,omitempty" json:"size,omitempty"`
// }
//
// //response body for complete multipart upload
// type CompleteMultipartUploadResult struct {
// 	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CompleteMultipartUploadResult" json:"-"`
// 	Location string `xml:"Location"`
// 	Bucket string `xml:"Bucket"`
// 	Key string  `xml:"Key"`
// 	ETag string `xml:"ETag"`
// }
//
// //response body for list parts of object
// type ListPartsResult struct {
// 	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListPartsResult" json:"-"`
// 	Bucket string
// 	EncodingType string `xml:"EncodingType,omitempty"`
// 	Key string
// 	UploadId string
// 	Initiator initiator
// 	Owner initiator
// 	StorageClass string
// 	PartNumberMarker int
// 	NextPartNumberMarker int
// 	MaxParts int
// 	IsTruncated bool
// 	Parts []PartInfo `xml:"Part"`
// }
//
// /*https://docs.aws.amazon.com/AmazonS3/latest/API/mpUploadListMPUpload.html*/
// type ListMultipartUploadsResult struct {
// 	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListMultipartUploadsResult" json:"-"`
// 	Bucket             string
// 	KeyMarker          string
// 	UploadIDMarker     string `xml:"UploadIdMarker"`
// 	NextKeyMarker      string
// 	NextUploadIDMarker string `xml:"NextUploadIdMarker"`
// 	EncodingType       string
// 	MaxUploads         int64
// 	IsTruncated        bool
// 	Uploads            []ObjectMultipartInfo `xml:"Upload"`
// 	Prefix             string
// 	Delimiter          string
// 	CommonPrefixes []CommonPrefix
// }
//
// type ObjectMultipartInfo struct {
// 	Key int
// 	UploadId int
// 	Initiator initiator
// 	Owner initiator
// 	StorageClass string
// 	Initiated time.Time
// }
