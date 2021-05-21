// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trace.proto

package tipb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Event int32

const (
	Event_Unknown                       Event = 0
	Event_TiKvCoprGetRequest            Event = 1000
	Event_TiKvCoprHandleRequest         Event = 1001
	Event_TiKvCoprScheduleTask          Event = 1002
	Event_TiKvCoprGetSnapshot           Event = 1003
	Event_TiKvCoprExecuteDagRunner      Event = 1004
	Event_TiKvCoprExecuteBatchDagRunner Event = 1005
)

var Event_name = map[int32]string{
	0:    "Unknown",
	1000: "TiKvCoprGetRequest",
	1001: "TiKvCoprHandleRequest",
	1002: "TiKvCoprScheduleTask",
	1003: "TiKvCoprGetSnapshot",
	1004: "TiKvCoprExecuteDagRunner",
	1005: "TiKvCoprExecuteBatchDagRunner",
}
var Event_value = map[string]int32{
	"Unknown":                       0,
	"TiKvCoprGetRequest":            1000,
	"TiKvCoprHandleRequest":         1001,
	"TiKvCoprScheduleTask":          1002,
	"TiKvCoprGetSnapshot":           1003,
	"TiKvCoprExecuteDagRunner":      1004,
	"TiKvCoprExecuteBatchDagRunner": 1005,
}

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}
func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (x *Event) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Event_value, data, "Event")
	if err != nil {
		return err
	}
	*x = Event(value)
	return nil
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptorTrace, []int{0} }

type ResourceGroupTag struct {
	SqlDigest        []byte `protobuf:"bytes,1,opt,name=sql_digest,json=sqlDigest" json:"sql_digest,omitempty"`
	PlanDigest       []byte `protobuf:"bytes,2,opt,name=plan_digest,json=planDigest" json:"plan_digest,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ResourceGroupTag) Reset()                    { *m = ResourceGroupTag{} }
func (m *ResourceGroupTag) String() string            { return proto.CompactTextString(m) }
func (*ResourceGroupTag) ProtoMessage()               {}
func (*ResourceGroupTag) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{0} }

func (m *ResourceGroupTag) GetSqlDigest() []byte {
	if m != nil {
		return m.SqlDigest
	}
	return nil
}

func (m *ResourceGroupTag) GetPlanDigest() []byte {
	if m != nil {
		return m.PlanDigest
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceGroupTag)(nil), "tipb.ResourceGroupTag")
	proto.RegisterEnum("tipb.Event", Event_name, Event_value)
}
func (m *ResourceGroupTag) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceGroupTag) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SqlDigest != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTrace(dAtA, i, uint64(len(m.SqlDigest)))
		i += copy(dAtA[i:], m.SqlDigest)
	}
	if m.PlanDigest != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTrace(dAtA, i, uint64(len(m.PlanDigest)))
		i += copy(dAtA[i:], m.PlanDigest)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTrace(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResourceGroupTag) Size() (n int) {
	var l int
	_ = l
	if m.SqlDigest != nil {
		l = len(m.SqlDigest)
		n += 1 + l + sovTrace(uint64(l))
	}
	if m.PlanDigest != nil {
		l = len(m.PlanDigest)
		n += 1 + l + sovTrace(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTrace(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTrace(x uint64) (n int) {
	return sovTrace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceGroupTag) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrace
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResourceGroupTag: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceGroupTag: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SqlDigest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SqlDigest = append(m.SqlDigest[:0], dAtA[iNdEx:postIndex]...)
			if m.SqlDigest == nil {
				m.SqlDigest = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanDigest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlanDigest = append(m.PlanDigest[:0], dAtA[iNdEx:postIndex]...)
			if m.PlanDigest == nil {
				m.PlanDigest = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrace
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTrace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrace
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTrace
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTrace
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTrace
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTrace
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTrace(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTrace = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrace   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("trace.proto", fileDescriptorTrace) }

var fileDescriptorTrace = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0xcd, 0x4e, 0x32, 0x31,
	0x14, 0xc6, 0x71, 0xfa, 0xe6, 0x35, 0x8d, 0x07, 0x17, 0x4d, 0x85, 0x88, 0x24, 0x8c, 0x86, 0x95,
	0xba, 0x18, 0xef, 0x01, 0x21, 0x98, 0xb8, 0x31, 0x03, 0xae, 0x4d, 0xe9, 0x9c, 0x94, 0x09, 0x63,
	0x5b, 0xfa, 0x81, 0x5e, 0x8a, 0xb7, 0xe3, 0xce, 0xa5, 0x97, 0x60, 0x70, 0xe3, 0xe7, 0x3d, 0x18,
	0x20, 0x13, 0x8d, 0xbb, 0xe6, 0xff, 0x7b, 0xd2, 0xc5, 0x81, 0x7a, 0x70, 0x42, 0x62, 0x6a, 0x9d,
	0x09, 0x86, 0xff, 0x0f, 0x85, 0x9d, 0xb4, 0x1b, 0xca, 0x28, 0xb3, 0x0e, 0xa7, 0xab, 0xd7, 0xc6,
	0xba, 0x19, 0xb0, 0x0c, 0xbd, 0x89, 0x4e, 0xe2, 0xd0, 0x99, 0x68, 0xc7, 0x42, 0xf1, 0x0e, 0x80,
	0x9f, 0x97, 0xd7, 0x79, 0xa1, 0xd0, 0x87, 0x16, 0x39, 0x24, 0x47, 0x3b, 0xd9, 0xb6, 0x9f, 0x97,
	0xfd, 0x75, 0xe0, 0x07, 0x50, 0xb7, 0xa5, 0xd0, 0x95, 0xff, 0x5b, 0x3b, 0xac, 0xd2, 0x66, 0x70,
	0xf2, 0x40, 0x60, 0x6b, 0xb0, 0x40, 0x1d, 0x78, 0x1d, 0xe8, 0x95, 0x9e, 0x69, 0x73, 0xab, 0x59,
	0x8d, 0xef, 0x01, 0x1f, 0x17, 0x17, 0x8b, 0x33, 0x63, 0xdd, 0x10, 0x43, 0x86, 0xf3, 0x88, 0x3e,
	0xb0, 0x57, 0xca, 0xdb, 0xd0, 0xac, 0xe0, 0x5c, 0xe8, 0xbc, 0xc4, 0xca, 0xde, 0x28, 0xdf, 0x87,
	0x46, 0x65, 0x23, 0x39, 0xc5, 0x3c, 0x96, 0x38, 0x16, 0x7e, 0xc6, 0xde, 0x29, 0x6f, 0xc1, 0xee,
	0xaf, 0xff, 0x46, 0x5a, 0x58, 0x3f, 0x35, 0x81, 0x7d, 0x50, 0xde, 0x81, 0x56, 0x25, 0x83, 0x3b,
	0x94, 0x31, 0x60, 0x5f, 0xa8, 0x2c, 0x6a, 0x8d, 0x8e, 0x7d, 0x52, 0xde, 0x85, 0xce, 0x1f, 0xee,
	0x89, 0x20, 0xa7, 0x3f, 0x9b, 0x2f, 0xda, 0x3b, 0x7e, 0x5c, 0x26, 0xe4, 0x69, 0x99, 0x90, 0xe7,
	0x65, 0x42, 0xee, 0x5f, 0x92, 0x1a, 0x34, 0xa5, 0xb9, 0x49, 0x6d, 0xa1, 0x95, 0x14, 0x36, 0x0d,
	0x45, 0x3e, 0x49, 0x57, 0x67, 0xbd, 0x24, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x49, 0x33,
	0x02, 0x6c, 0x01, 0x00, 0x00,
}
