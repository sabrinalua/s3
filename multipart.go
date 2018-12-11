package gos3

type initiator struct {
	ID string 
	DisplayName string
} 

type InitiateMultipartUploadResult struct {
	Namespace string `xml:"xmlns,attr"`
	Bucket string `xml:"Bucket"`
	Key string `xml:"Key"`
	UploadId string `xml:"UploadId"`
}

type CompleteMultipartUpload struct{
	Parts []PartInfo`xml:"Part" json:"parts"`
}

type CompleteMultipartUploadResult struct {
	Namespace string `xml:"xmlns,attr"`
	Location string `xml:"Location"`
	Bucket string `xml:"Bucket"`
	Key string  `xml:"Key"`
	ETag string `xml:"ETag"`
}

type PartInfo struct {
	PartNumber int `xml:"PartNumber" json:"part_number"`
	ETag string `xml:"ETag" json:"etag"` 
}

type PartInfoWithModded struct {
	PartInfo
	LastModified string 
	Size string
}

type ListPartsResult struct {
	InitiateMultipartUploadResult
	Initiator initiator
	Owner initiator
	StorageClass string 
	PartNumberMarker int
	NextPartNumberMarker int 
	MaxParts int 
	IsTruncated bool
	Part []PartInfoWithModded `xml:"Part"`
}

type ListMultipartUploadsResult struct {
	Namespace string `xml:"xmlns,attr"`
	Bucket string 
	KeyMarker string 
	UploadIdMarker string 
	NextKeyMarker string 
	NextUploadIdMarker string 
	MaxUploads int 
	Delimiter string
	Prefix string 
	CommonPrefixes []string `xml:"CommonPrefixes>Prefix"`
	IsTruncated bool 
	Uploads []Upload `xml:"Upload"`  
}

type Upload struct {
	Key int
	UploadId int 
	Initiator initiator  
	Owner initiator
	StorageClass string
	Initiated string 
}