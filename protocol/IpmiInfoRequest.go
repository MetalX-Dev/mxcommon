// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type IpmiInfoRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsIpmiInfoRequest(buf []byte, offset flatbuffers.UOffsetT) *IpmiInfoRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IpmiInfoRequest{}
	x.Init(buf, n+offset)
	return x
}

func FinishIpmiInfoRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsIpmiInfoRequest(buf []byte, offset flatbuffers.UOffsetT) *IpmiInfoRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &IpmiInfoRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedIpmiInfoRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *IpmiInfoRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IpmiInfoRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func IpmiInfoRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func IpmiInfoRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
