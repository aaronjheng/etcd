// Code generated by protoc-gen-gogo.
// source: kv.proto
// DO NOT EDIT!

/*
	Package storagepb is a generated protocol buffer package.

	It is generated from these files:
		kv.proto

	It has these top-level messages:
		KeyValue
		Event
*/
package storagepb

import proto "github.com/coreos/etcd/Godeps/_workspace/src/github.com/gogo/protobuf/proto"

// discarding unused import gogoproto "github.com/coreos/etcd/Godeps/_workspace/src/gogoproto"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Event_EventType int32

const (
	PUT    Event_EventType = 0
	DELETE Event_EventType = 1
	EXPIRE Event_EventType = 2
)

var Event_EventType_name = map[int32]string{
	0: "PUT",
	1: "DELETE",
	2: "EXPIRE",
}
var Event_EventType_value = map[string]int32{
	"PUT":    0,
	"DELETE": 1,
	"EXPIRE": 2,
}

func (x Event_EventType) String() string {
	return proto.EnumName(Event_EventType_name, int32(x))
}

type KeyValue struct {
	Key            []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	CreateRevision int64  `protobuf:"varint,2,opt,name=create_revision,proto3" json:"create_revision,omitempty"`
	// mod_revision is the last modified revision of the key.
	ModRevision int64 `protobuf:"varint,3,opt,name=mod_revision,proto3" json:"mod_revision,omitempty"`
	// version is the version of the key. A deletion resets
	// the version to zero and any modification of the key
	// increases its version.
	Version int64  `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Value   []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	// lease is the ID of the lease that attached to key.
	// When the attached lease expires, the key will be deleted.
	Lease int64 `protobuf:"varint,6,opt,name=lease,proto3" json:"lease,omitempty"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}

type Event struct {
	Type Event_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=storagepb.Event_EventType" json:"type,omitempty"`
	// a put event contains the current key-value
	// a delete/expire event contains the previous
	// key-value
	Kv *KeyValue `protobuf:"bytes,2,opt,name=kv" json:"kv,omitempty"`
	// watchID is the ID of watching this event is sent to.
	WatchID int64 `protobuf:"varint,3,opt,name=watchID,proto3" json:"watchID,omitempty"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("storagepb.Event_EventType", Event_EventType_name, Event_EventType_value)
}
func (m *KeyValue) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *KeyValue) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key != nil {
		if len(m.Key) > 0 {
			data[i] = 0xa
			i++
			i = encodeVarintKv(data, i, uint64(len(m.Key)))
			i += copy(data[i:], m.Key)
		}
	}
	if m.CreateRevision != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintKv(data, i, uint64(m.CreateRevision))
	}
	if m.ModRevision != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintKv(data, i, uint64(m.ModRevision))
	}
	if m.Version != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintKv(data, i, uint64(m.Version))
	}
	if m.Value != nil {
		if len(m.Value) > 0 {
			data[i] = 0x2a
			i++
			i = encodeVarintKv(data, i, uint64(len(m.Value)))
			i += copy(data[i:], m.Value)
		}
	}
	if m.Lease != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintKv(data, i, uint64(m.Lease))
	}
	return i, nil
}

func (m *Event) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Event) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintKv(data, i, uint64(m.Type))
	}
	if m.Kv != nil {
		data[i] = 0x12
		i++
		i = encodeVarintKv(data, i, uint64(m.Kv.Size()))
		n1, err := m.Kv.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.WatchID != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintKv(data, i, uint64(m.WatchID))
	}
	return i, nil
}

func encodeFixed64Kv(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Kv(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintKv(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *KeyValue) Size() (n int) {
	var l int
	_ = l
	if m.Key != nil {
		l = len(m.Key)
		if l > 0 {
			n += 1 + l + sovKv(uint64(l))
		}
	}
	if m.CreateRevision != 0 {
		n += 1 + sovKv(uint64(m.CreateRevision))
	}
	if m.ModRevision != 0 {
		n += 1 + sovKv(uint64(m.ModRevision))
	}
	if m.Version != 0 {
		n += 1 + sovKv(uint64(m.Version))
	}
	if m.Value != nil {
		l = len(m.Value)
		if l > 0 {
			n += 1 + l + sovKv(uint64(l))
		}
	}
	if m.Lease != 0 {
		n += 1 + sovKv(uint64(m.Lease))
	}
	return n
}

func (m *Event) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovKv(uint64(m.Type))
	}
	if m.Kv != nil {
		l = m.Kv.Size()
		n += 1 + l + sovKv(uint64(l))
	}
	if m.WatchID != 0 {
		n += 1 + sovKv(uint64(m.WatchID))
	}
	return n
}

func sovKv(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozKv(x uint64) (n int) {
	return sovKv(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyValue) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateRevision", wireType)
			}
			m.CreateRevision = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.CreateRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModRevision", wireType)
			}
			m.ModRevision = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ModRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Version |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			m.Lease = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Lease |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipKv(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKv
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *Event) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (Event_EventType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Kv == nil {
				m.Kv = &KeyValue{}
			}
			if err := m.Kv.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WatchID", wireType)
			}
			m.WatchID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.WatchID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipKv(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKv
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipKv(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthKv
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipKv(data[start:])
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
	ErrInvalidLengthKv = fmt.Errorf("proto: negative length found during unmarshaling")
)
