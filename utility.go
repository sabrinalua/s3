package gos3
import (
	"encoding/xml"
)


func EncodeXml( is_resp bool, args...interface{})(resp []byte, err error) {
	resp, err= xml.Marshal(args[0])
	if is_resp{
		resp = []byte (xml.Header+ string(resp))
	}
	return
}