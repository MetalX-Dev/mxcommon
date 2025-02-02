// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package protocol

import "strconv"

type Compression byte

const (
	CompressionNone Compression = 0
	CompressionLz4  Compression = 1
)

var EnumNamesCompression = map[Compression]string{
	CompressionNone: "None",
	CompressionLz4:  "Lz4",
}

var EnumValuesCompression = map[string]Compression{
	"None": CompressionNone,
	"Lz4":  CompressionLz4,
}

func (v Compression) String() string {
	if s, ok := EnumNamesCompression[v]; ok {
		return s
	}
	return "Compression(" + strconv.FormatInt(int64(v), 10) + ")"
}
