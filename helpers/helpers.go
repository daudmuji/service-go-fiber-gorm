package helpers

import (
	"errors"
	"os"
)

var (
	RC = map[string]string{
		"NODATA":            "14",
		"TIMEOUT":           "68",
		"SUCCESS":           "00",
		"GENERAL_ERROR":     "05",
		"DUPLICATE_STAN":    "94",
		"FORMAT_DATA_ERROR": "30",
	}
)

const (
	FieldReferer       = "referer"
	FieldProtocol      = "protocol"
	FieldPID           = "pid"
	FieldPort          = "port"
	FieldIP            = "ip"
	FieldIPs           = "ips"
	FieldHost          = "host"
	FieldPath          = "path"
	FieldURL           = "url"
	FieldUserAgent     = "ua"
	FieldLatency       = "latency"
	FieldStatus        = "status"
	FieldResBody       = "resBody"
	FieldQueryParams   = "queryParams"
	FieldBody          = "body"
	FieldBytesReceived = "bytesReceived"
	FieldBytesSent     = "bytesSent"
	FieldRoute         = "route"
	FieldMethod        = "method"
	FieldRequestID     = "requestId"
	FieldError         = "error"
	FieldReqHeaders    = "reqHeaders"
)

func CheckFolderPath(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func AppendIfLess(dataSlice []string, value int) []string {
	lengthArray := len(dataSlice)
	if lengthArray < value {
		minusAppenda := value - lengthArray
		for i := 0; i < minusAppenda; i++ {
			dataSlice = append(dataSlice, "")
		}
	}
	return dataSlice
}
