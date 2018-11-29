package gos3


type InitiateMultipartUploadResult struct {
	Namespace string `xml:"xmlns,attr"`
	Bucket string `xml:"Bucket"`
	Key string `xml:"Key"`
	UploadId string `xml:"UploadId"`
}

type PartInfo struct {
	PartNumber int `xml:"PartNumber" json:"part_number"`
	ETag string `xml:"ETag" json:"etag"` 
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

type Error struct {
	Code int  
	Message string 
	RequestId string 
	HostId string
}


type PartInfoWithModded struct {
	PartInfo
	LastModified string 
	Size string
}
type ListPartsResult struct {
	InitiateMultipartUploadResult
	Initiator struct{
		ID string 
		DisplayName string
	}
	Owner struct{
		ID string 
		DisplayName string
	}
	StorageClass string 
	PartNumberMarker int
	NextPartNumberMarker int 
	MaxParts int 
	Part []PartInfoWithModded `xml:"Part"`
}