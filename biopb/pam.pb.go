// Code generated by protoc-gen-gogo.
// source: proto/bio/pam.proto
// DO NOT EDIT!

package biopb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PAMBlockHeader struct {
	Offset     uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	BlobOffset uint32 `protobuf:"varint,3,opt,name=blob_offset,json=blobOffset,proto3" json:"blob_offset,omitempty"`
}

func (m *PAMBlockHeader) Reset()                    { *m = PAMBlockHeader{} }
func (m *PAMBlockHeader) String() string            { return proto.CompactTextString(m) }
func (*PAMBlockHeader) ProtoMessage()               {}
func (*PAMBlockHeader) Descriptor() ([]byte, []int) { return fileDescriptorPam, []int{0} }

func (m *PAMBlockHeader) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *PAMBlockHeader) GetBlobOffset() uint32 {
	if m != nil {
		return m.BlobOffset
	}
	return 0
}

type PAMBlockIndexEntry struct {
	FileOffset uint64 `protobuf:"varint,1,opt,name=file_offset,json=fileOffset,proto3" json:"file_offset,omitempty"`
	NumRecords uint32 `protobuf:"varint,3,opt,name=num_records,json=numRecords,proto3" json:"num_records,omitempty"`
	StartAddr  Coord  `protobuf:"bytes,4,opt,name=start_addr,json=startAddr" json:"start_addr"`
	EndAddr    Coord  `protobuf:"bytes,5,opt,name=end_addr,json=endAddr" json:"end_addr"`
}

func (m *PAMBlockIndexEntry) Reset()                    { *m = PAMBlockIndexEntry{} }
func (m *PAMBlockIndexEntry) String() string            { return proto.CompactTextString(m) }
func (*PAMBlockIndexEntry) ProtoMessage()               {}
func (*PAMBlockIndexEntry) Descriptor() ([]byte, []int) { return fileDescriptorPam, []int{1} }

func (m *PAMBlockIndexEntry) GetFileOffset() uint64 {
	if m != nil {
		return m.FileOffset
	}
	return 0
}

func (m *PAMBlockIndexEntry) GetNumRecords() uint32 {
	if m != nil {
		return m.NumRecords
	}
	return 0
}

func (m *PAMBlockIndexEntry) GetStartAddr() Coord {
	if m != nil {
		return m.StartAddr
	}
	return Coord{}
}

func (m *PAMBlockIndexEntry) GetEndAddr() Coord {
	if m != nil {
		return m.EndAddr
	}
	return Coord{}
}

type PAMShardIndex struct {
	Magic            uint64     `protobuf:"fixed64,1,opt,name=magic,proto3" json:"magic,omitempty"`
	Version          string     `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Range            CoordRange `protobuf:"bytes,4,opt,name=range" json:"range"`
	EncodedBamHeader []byte     `protobuf:"bytes,15,opt,name=encoded_bam_header,json=encodedBamHeader,proto3" json:"encoded_bam_header,omitempty"`
}

func (m *PAMShardIndex) Reset()                    { *m = PAMShardIndex{} }
func (m *PAMShardIndex) String() string            { return proto.CompactTextString(m) }
func (*PAMShardIndex) ProtoMessage()               {}
func (*PAMShardIndex) Descriptor() ([]byte, []int) { return fileDescriptorPam, []int{2} }

func (m *PAMShardIndex) GetMagic() uint64 {
	if m != nil {
		return m.Magic
	}
	return 0
}

func (m *PAMShardIndex) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PAMShardIndex) GetRange() CoordRange {
	if m != nil {
		return m.Range
	}
	return CoordRange{}
}

func (m *PAMShardIndex) GetEncodedBamHeader() []byte {
	if m != nil {
		return m.EncodedBamHeader
	}
	return nil
}

type PAMFieldIndex struct {
	Magic   uint64               `protobuf:"fixed64,1,opt,name=magic,proto3" json:"magic,omitempty"`
	Version string               `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Field   int32                `protobuf:"varint,4,opt,name=field,proto3" json:"field,omitempty"`
	Blocks  []PAMBlockIndexEntry `protobuf:"bytes,16,rep,name=blocks" json:"blocks"`
}

func (m *PAMFieldIndex) Reset()                    { *m = PAMFieldIndex{} }
func (m *PAMFieldIndex) String() string            { return proto.CompactTextString(m) }
func (*PAMFieldIndex) ProtoMessage()               {}
func (*PAMFieldIndex) Descriptor() ([]byte, []int) { return fileDescriptorPam, []int{3} }

func (m *PAMFieldIndex) GetMagic() uint64 {
	if m != nil {
		return m.Magic
	}
	return 0
}

func (m *PAMFieldIndex) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PAMFieldIndex) GetField() int32 {
	if m != nil {
		return m.Field
	}
	return 0
}

func (m *PAMFieldIndex) GetBlocks() []PAMBlockIndexEntry {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto.RegisterType((*PAMBlockHeader)(nil), "grail.proto.bio.PAMBlockHeader")
	proto.RegisterType((*PAMBlockIndexEntry)(nil), "grail.proto.bio.PAMBlockIndexEntry")
	proto.RegisterType((*PAMShardIndex)(nil), "grail.proto.bio.PAMShardIndex")
	proto.RegisterType((*PAMFieldIndex)(nil), "grail.proto.bio.PAMFieldIndex")
}
func (m *PAMBlockHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PAMBlockHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Offset != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPam(dAtA, i, uint64(m.Offset))
	}
	if m.BlobOffset != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPam(dAtA, i, uint64(m.BlobOffset))
	}
	return i, nil
}

func (m *PAMBlockIndexEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PAMBlockIndexEntry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FileOffset != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPam(dAtA, i, uint64(m.FileOffset))
	}
	if m.NumRecords != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPam(dAtA, i, uint64(m.NumRecords))
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintPam(dAtA, i, uint64(m.StartAddr.Size()))
	n1, err := m.StartAddr.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x2a
	i++
	i = encodeVarintPam(dAtA, i, uint64(m.EndAddr.Size()))
	n2, err := m.EndAddr.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *PAMShardIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PAMShardIndex) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Magic != 0 {
		dAtA[i] = 0x9
		i++
		i = encodeFixed64Pam(dAtA, i, uint64(m.Magic))
	}
	if len(m.Version) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPam(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintPam(dAtA, i, uint64(m.Range.Size()))
	n3, err := m.Range.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.EncodedBamHeader) > 0 {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintPam(dAtA, i, uint64(len(m.EncodedBamHeader)))
		i += copy(dAtA[i:], m.EncodedBamHeader)
	}
	return i, nil
}

func (m *PAMFieldIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PAMFieldIndex) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Magic != 0 {
		dAtA[i] = 0x9
		i++
		i = encodeFixed64Pam(dAtA, i, uint64(m.Magic))
	}
	if len(m.Version) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPam(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	if m.Field != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintPam(dAtA, i, uint64(m.Field))
	}
	if len(m.Blocks) > 0 {
		for _, msg := range m.Blocks {
			dAtA[i] = 0x82
			i++
			dAtA[i] = 0x1
			i++
			i = encodeVarintPam(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Pam(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Pam(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPam(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PAMBlockHeader) Size() (n int) {
	var l int
	_ = l
	if m.Offset != 0 {
		n += 1 + sovPam(uint64(m.Offset))
	}
	if m.BlobOffset != 0 {
		n += 1 + sovPam(uint64(m.BlobOffset))
	}
	return n
}

func (m *PAMBlockIndexEntry) Size() (n int) {
	var l int
	_ = l
	if m.FileOffset != 0 {
		n += 1 + sovPam(uint64(m.FileOffset))
	}
	if m.NumRecords != 0 {
		n += 1 + sovPam(uint64(m.NumRecords))
	}
	l = m.StartAddr.Size()
	n += 1 + l + sovPam(uint64(l))
	l = m.EndAddr.Size()
	n += 1 + l + sovPam(uint64(l))
	return n
}

func (m *PAMShardIndex) Size() (n int) {
	var l int
	_ = l
	if m.Magic != 0 {
		n += 9
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovPam(uint64(l))
	}
	l = m.Range.Size()
	n += 1 + l + sovPam(uint64(l))
	l = len(m.EncodedBamHeader)
	if l > 0 {
		n += 1 + l + sovPam(uint64(l))
	}
	return n
}

func (m *PAMFieldIndex) Size() (n int) {
	var l int
	_ = l
	if m.Magic != 0 {
		n += 9
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovPam(uint64(l))
	}
	if m.Field != 0 {
		n += 1 + sovPam(uint64(m.Field))
	}
	if len(m.Blocks) > 0 {
		for _, e := range m.Blocks {
			l = e.Size()
			n += 2 + l + sovPam(uint64(l))
		}
	}
	return n
}

func sovPam(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPam(x uint64) (n int) {
	return sovPam(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PAMBlockHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPam
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
			return fmt.Errorf("proto: PAMBlockHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PAMBlockHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlobOffset", wireType)
			}
			m.BlobOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlobOffset |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPam
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PAMBlockIndexEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPam
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
			return fmt.Errorf("proto: PAMBlockIndexEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PAMBlockIndexEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileOffset", wireType)
			}
			m.FileOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FileOffset |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRecords", wireType)
			}
			m.NumRecords = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumRecords |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartAddr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPam
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StartAddr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndAddr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPam
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EndAddr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPam
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PAMShardIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPam
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
			return fmt.Errorf("proto: PAMShardIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PAMShardIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Magic", wireType)
			}
			m.Magic = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.Magic = uint64(dAtA[iNdEx-8])
			m.Magic |= uint64(dAtA[iNdEx-7]) << 8
			m.Magic |= uint64(dAtA[iNdEx-6]) << 16
			m.Magic |= uint64(dAtA[iNdEx-5]) << 24
			m.Magic |= uint64(dAtA[iNdEx-4]) << 32
			m.Magic |= uint64(dAtA[iNdEx-3]) << 40
			m.Magic |= uint64(dAtA[iNdEx-2]) << 48
			m.Magic |= uint64(dAtA[iNdEx-1]) << 56
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPam
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Range", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPam
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Range.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncodedBamHeader", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPam
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
				return ErrInvalidLengthPam
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncodedBamHeader = append(m.EncodedBamHeader[:0], dAtA[iNdEx:postIndex]...)
			if m.EncodedBamHeader == nil {
				m.EncodedBamHeader = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPam
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PAMFieldIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPam
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
			return fmt.Errorf("proto: PAMFieldIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PAMFieldIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Magic", wireType)
			}
			m.Magic = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.Magic = uint64(dAtA[iNdEx-8])
			m.Magic |= uint64(dAtA[iNdEx-7]) << 8
			m.Magic |= uint64(dAtA[iNdEx-6]) << 16
			m.Magic |= uint64(dAtA[iNdEx-5]) << 24
			m.Magic |= uint64(dAtA[iNdEx-4]) << 32
			m.Magic |= uint64(dAtA[iNdEx-3]) << 40
			m.Magic |= uint64(dAtA[iNdEx-2]) << 48
			m.Magic |= uint64(dAtA[iNdEx-1]) << 56
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPam
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			m.Field = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Field |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPam
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocks = append(m.Blocks, PAMBlockIndexEntry{})
			if err := m.Blocks[len(m.Blocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPam
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPam(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPam
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
					return 0, ErrIntOverflowPam
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
					return 0, ErrIntOverflowPam
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
				return 0, ErrInvalidLengthPam
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPam
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
				next, err := skipPam(dAtA[start:])
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
	ErrInvalidLengthPam = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPam   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("proto/bio/pam.proto", fileDescriptorPam) }

var fileDescriptorPam = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0xad, 0xe9, 0x64, 0x4a, 0xef, 0x50, 0x5a, 0x99, 0x52, 0x45, 0x45, 0x4c, 0xa3, 0x61, 0x93,
	0x05, 0x24, 0x52, 0x59, 0x74, 0xc1, 0x6a, 0x06, 0x81, 0xe8, 0xa2, 0x6a, 0x65, 0x76, 0x6c, 0x22,
	0x3b, 0x76, 0x32, 0x16, 0x89, 0x3d, 0x72, 0x32, 0x08, 0xfe, 0x82, 0x25, 0x5f, 0xc0, 0xb7, 0x74,
	0xc1, 0x82, 0x2f, 0x40, 0x68, 0xf8, 0x11, 0xe4, 0xc7, 0x00, 0x62, 0x58, 0xa0, 0x2e, 0x22, 0xe5,
	0x9c, 0x7b, 0xcf, 0xf1, 0xb9, 0xbe, 0x86, 0x7b, 0x0b, 0xa3, 0x7b, 0x9d, 0x33, 0xa9, 0xf3, 0x05,
	0x6d, 0x33, 0x87, 0xf0, 0x7e, 0x6d, 0xa8, 0x6c, 0x3c, 0xc8, 0x98, 0xd4, 0xc7, 0x4f, 0x6a, 0xd9,
	0xcf, 0x97, 0x2c, 0x2b, 0x75, 0x9b, 0xd7, 0xba, 0xd6, 0xb9, 0x2b, 0xb1, 0x65, 0xe5, 0x90, 0xb7,
	0xb0, 0x7f, 0x5e, 0x72, 0x7c, 0xff, 0xb7, 0x69, 0xa9, 0xb5, 0xe1, 0x9e, 0x9e, 0x9c, 0xc3, 0xdd,
	0xab, 0xe9, 0xc5, 0xac, 0xd1, 0xe5, 0xdb, 0x57, 0x82, 0x72, 0x61, 0xf0, 0x11, 0x0c, 0x75, 0x55,
	0x75, 0xa2, 0x8f, 0x6f, 0x25, 0x28, 0xdd, 0x23, 0x01, 0xe1, 0x13, 0x18, 0xb1, 0x46, 0xb3, 0x22,
	0x14, 0xb7, 0x5d, 0x11, 0x2c, 0x75, 0xe9, 0x98, 0xc9, 0x17, 0x04, 0x78, 0xed, 0x75, 0xae, 0xb8,
	0x78, 0xff, 0x42, 0xf5, 0xe6, 0x83, 0xd5, 0x55, 0xb2, 0x11, 0x6b, 0x1d, 0x4a, 0x50, 0x3a, 0x20,
	0x60, 0xa9, 0xcb, 0x5f, 0xc6, 0x6a, 0xd9, 0x16, 0x46, 0x94, 0xda, 0xf0, 0x6e, 0x6d, 0xac, 0x96,
	0x2d, 0xf1, 0x0c, 0x7e, 0x06, 0xd0, 0xf5, 0xd4, 0xf4, 0x05, 0xe5, 0xdc, 0xc4, 0x83, 0x04, 0xa5,
	0xa3, 0xd3, 0xa3, 0xec, 0xaf, 0xfb, 0xc8, 0x9e, 0xdb, 0xa9, 0x66, 0x83, 0xeb, 0x6f, 0x27, 0x5b,
	0x64, 0xd7, 0xf5, 0x4f, 0x39, 0x37, 0xf8, 0x0c, 0x6e, 0x0b, 0xc5, 0xbd, 0x34, 0xfa, 0x0f, 0xe9,
	0x8e, 0x50, 0xdc, 0x0a, 0x27, 0x9f, 0x11, 0xec, 0x5d, 0x4d, 0x2f, 0x5e, 0xcf, 0xa9, 0xe1, 0x6e,
	0x1c, 0x7c, 0x08, 0x51, 0x4b, 0x6b, 0x59, 0xba, 0x19, 0x86, 0xc4, 0x03, 0x1c, 0xc3, 0xce, 0x3b,
	0x61, 0x3a, 0xa9, 0x95, 0x8b, 0xbe, 0x4b, 0xd6, 0x10, 0x9f, 0x41, 0x64, 0xa8, 0xaa, 0x45, 0x88,
	0xfc, 0xe0, 0xdf, 0xe7, 0x12, 0xdb, 0x12, 0x0e, 0xf7, 0xfd, 0xf8, 0x31, 0x60, 0xa1, 0x4a, 0xcd,
	0x05, 0x2f, 0x18, 0x6d, 0x8b, 0xb9, 0x5b, 0x4c, 0xbc, 0x9f, 0xa0, 0xf4, 0x0e, 0x39, 0x08, 0x95,
	0x19, 0x6d, 0xfd, 0xc2, 0x26, 0x9f, 0x7c, 0xd0, 0x97, 0x52, 0x34, 0x37, 0x0c, 0x7a, 0x08, 0x51,
	0x65, 0xd5, 0x2e, 0x68, 0x44, 0x3c, 0xc0, 0x53, 0x18, 0x32, 0xbb, 0xcb, 0x2e, 0x3e, 0x48, 0xb6,
	0xd3, 0xd1, 0xe9, 0xa3, 0x8d, 0xfc, 0x9b, 0xdb, 0x0e, 0x73, 0x04, 0xe1, 0x2c, 0xbf, 0x5e, 0x8d,
	0xd1, 0xd7, 0xd5, 0x18, 0x7d, 0x5f, 0x8d, 0xd1, 0xc7, 0x1f, 0xe3, 0xad, 0x37, 0x0f, 0xff, 0x7c,
	0xb5, 0xd6, 0xce, 0x3e, 0xc8, 0xf0, 0x2d, 0x18, 0x1b, 0x3a, 0xf3, 0xa7, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x37, 0x41, 0xa4, 0x36, 0x03, 0x03, 0x00, 0x00,
}
