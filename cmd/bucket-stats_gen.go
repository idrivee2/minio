package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BucketReplicationStat) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "PendingSize":
			z.PendingSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "PendingSize")
				return
			}
		case "ReplicatedSize":
			z.ReplicatedSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "ReplicatedSize")
				return
			}
		case "ReplicaSize":
			z.ReplicaSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "ReplicaSize")
				return
			}
		case "FailedSize":
			z.FailedSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "FailedSize")
				return
			}
		case "PendingCount":
			z.PendingCount, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "PendingCount")
				return
			}
		case "FailedCount":
			z.FailedCount, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "FailedCount")
				return
			}
		case "Latency":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Latency")
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					err = msgp.WrapError(err, "Latency")
					return
				}
				switch msgp.UnsafeString(field) {
				case "UploadHistogram":
					err = z.Latency.UploadHistogram.DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "Latency", "UploadHistogram")
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						err = msgp.WrapError(err, "Latency")
						return
					}
				}
			}
		case "BandWidthLimitInBytesPerSecond":
			z.BandWidthLimitInBytesPerSecond, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "BandWidthLimitInBytesPerSecond")
				return
			}
		case "CurrentBandwidthInBytesPerSecond":
			z.CurrentBandwidthInBytesPerSecond, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "CurrentBandwidthInBytesPerSecond")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BucketReplicationStat) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "PendingSize"
	err = en.Append(0x89, 0xab, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.PendingSize)
	if err != nil {
		err = msgp.WrapError(err, "PendingSize")
		return
	}
	// write "ReplicatedSize"
	err = en.Append(0xae, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.ReplicatedSize)
	if err != nil {
		err = msgp.WrapError(err, "ReplicatedSize")
		return
	}
	// write "ReplicaSize"
	err = en.Append(0xab, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.ReplicaSize)
	if err != nil {
		err = msgp.WrapError(err, "ReplicaSize")
		return
	}
	// write "FailedSize"
	err = en.Append(0xaa, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.FailedSize)
	if err != nil {
		err = msgp.WrapError(err, "FailedSize")
		return
	}
	// write "PendingCount"
	err = en.Append(0xac, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.PendingCount)
	if err != nil {
		err = msgp.WrapError(err, "PendingCount")
		return
	}
	// write "FailedCount"
	err = en.Append(0xab, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.FailedCount)
	if err != nil {
		err = msgp.WrapError(err, "FailedCount")
		return
	}
	// write "Latency"
	err = en.Append(0xa7, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79)
	if err != nil {
		return
	}
	// map header, size 1
	// write "UploadHistogram"
	err = en.Append(0x81, 0xaf, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d)
	if err != nil {
		return
	}
	err = z.Latency.UploadHistogram.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Latency", "UploadHistogram")
		return
	}
	// write "BandWidthLimitInBytesPerSecond"
	err = en.Append(0xbe, 0x42, 0x61, 0x6e, 0x64, 0x57, 0x69, 0x64, 0x74, 0x68, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.BandWidthLimitInBytesPerSecond)
	if err != nil {
		err = msgp.WrapError(err, "BandWidthLimitInBytesPerSecond")
		return
	}
	// write "CurrentBandwidthInBytesPerSecond"
	err = en.Append(0xd9, 0x20, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x49, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.CurrentBandwidthInBytesPerSecond)
	if err != nil {
		err = msgp.WrapError(err, "CurrentBandwidthInBytesPerSecond")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BucketReplicationStat) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "PendingSize"
	o = append(o, 0x89, 0xab, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.PendingSize)
	// string "ReplicatedSize"
	o = append(o, 0xae, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.ReplicatedSize)
	// string "ReplicaSize"
	o = append(o, 0xab, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.ReplicaSize)
	// string "FailedSize"
	o = append(o, 0xaa, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.FailedSize)
	// string "PendingCount"
	o = append(o, 0xac, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt64(o, z.PendingCount)
	// string "FailedCount"
	o = append(o, 0xab, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt64(o, z.FailedCount)
	// string "Latency"
	o = append(o, 0xa7, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79)
	// map header, size 1
	// string "UploadHistogram"
	o = append(o, 0x81, 0xaf, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d)
	o, err = z.Latency.UploadHistogram.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Latency", "UploadHistogram")
		return
	}
	// string "BandWidthLimitInBytesPerSecond"
	o = append(o, 0xbe, 0x42, 0x61, 0x6e, 0x64, 0x57, 0x69, 0x64, 0x74, 0x68, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64)
	o = msgp.AppendInt64(o, z.BandWidthLimitInBytesPerSecond)
	// string "CurrentBandwidthInBytesPerSecond"
	o = append(o, 0xd9, 0x20, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x49, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64)
	o = msgp.AppendFloat64(o, z.CurrentBandwidthInBytesPerSecond)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BucketReplicationStat) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "PendingSize":
			z.PendingSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PendingSize")
				return
			}
		case "ReplicatedSize":
			z.ReplicatedSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ReplicatedSize")
				return
			}
		case "ReplicaSize":
			z.ReplicaSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ReplicaSize")
				return
			}
		case "FailedSize":
			z.FailedSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "FailedSize")
				return
			}
		case "PendingCount":
			z.PendingCount, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PendingCount")
				return
			}
		case "FailedCount":
			z.FailedCount, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "FailedCount")
				return
			}
		case "Latency":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Latency")
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					err = msgp.WrapError(err, "Latency")
					return
				}
				switch msgp.UnsafeString(field) {
				case "UploadHistogram":
					bts, err = z.Latency.UploadHistogram.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Latency", "UploadHistogram")
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						err = msgp.WrapError(err, "Latency")
						return
					}
				}
			}
		case "BandWidthLimitInBytesPerSecond":
			z.BandWidthLimitInBytesPerSecond, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BandWidthLimitInBytesPerSecond")
				return
			}
		case "CurrentBandwidthInBytesPerSecond":
			z.CurrentBandwidthInBytesPerSecond, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CurrentBandwidthInBytesPerSecond")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BucketReplicationStat) Msgsize() (s int) {
	s = 1 + 12 + msgp.Int64Size + 15 + msgp.Int64Size + 12 + msgp.Int64Size + 11 + msgp.Int64Size + 13 + msgp.Int64Size + 12 + msgp.Int64Size + 8 + 1 + 16 + z.Latency.UploadHistogram.Msgsize() + 31 + msgp.Int64Size + 34 + msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BucketReplicationStats) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Stats":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if z.Stats == nil {
				z.Stats = make(map[string]*BucketReplicationStat, zb0002)
			} else if len(z.Stats) > 0 {
				for key := range z.Stats {
					delete(z.Stats, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 *BucketReplicationStat
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Stats")
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "Stats", za0001)
						return
					}
					za0002 = nil
				} else {
					if za0002 == nil {
						za0002 = new(BucketReplicationStat)
					}
					err = za0002.DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "Stats", za0001)
						return
					}
				}
				z.Stats[za0001] = za0002
			}
		case "PendingSize":
			z.PendingSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "PendingSize")
				return
			}
		case "ReplicatedSize":
			z.ReplicatedSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "ReplicatedSize")
				return
			}
		case "ReplicaSize":
			z.ReplicaSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "ReplicaSize")
				return
			}
		case "FailedSize":
			z.FailedSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "FailedSize")
				return
			}
		case "PendingCount":
			z.PendingCount, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "PendingCount")
				return
			}
		case "FailedCount":
			z.FailedCount, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "FailedCount")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BucketReplicationStats) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "Stats"
	err = en.Append(0x87, 0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Stats)))
	if err != nil {
		err = msgp.WrapError(err, "Stats")
		return
	}
	for za0001, za0002 := range z.Stats {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Stats")
			return
		}
		if za0002 == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = za0002.EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Stats", za0001)
				return
			}
		}
	}
	// write "PendingSize"
	err = en.Append(0xab, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.PendingSize)
	if err != nil {
		err = msgp.WrapError(err, "PendingSize")
		return
	}
	// write "ReplicatedSize"
	err = en.Append(0xae, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.ReplicatedSize)
	if err != nil {
		err = msgp.WrapError(err, "ReplicatedSize")
		return
	}
	// write "ReplicaSize"
	err = en.Append(0xab, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.ReplicaSize)
	if err != nil {
		err = msgp.WrapError(err, "ReplicaSize")
		return
	}
	// write "FailedSize"
	err = en.Append(0xaa, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.FailedSize)
	if err != nil {
		err = msgp.WrapError(err, "FailedSize")
		return
	}
	// write "PendingCount"
	err = en.Append(0xac, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.PendingCount)
	if err != nil {
		err = msgp.WrapError(err, "PendingCount")
		return
	}
	// write "FailedCount"
	err = en.Append(0xab, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.FailedCount)
	if err != nil {
		err = msgp.WrapError(err, "FailedCount")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BucketReplicationStats) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Stats"
	o = append(o, 0x87, 0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Stats)))
	for za0001, za0002 := range z.Stats {
		o = msgp.AppendString(o, za0001)
		if za0002 == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = za0002.MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "Stats", za0001)
				return
			}
		}
	}
	// string "PendingSize"
	o = append(o, 0xab, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.PendingSize)
	// string "ReplicatedSize"
	o = append(o, 0xae, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.ReplicatedSize)
	// string "ReplicaSize"
	o = append(o, 0xab, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.ReplicaSize)
	// string "FailedSize"
	o = append(o, 0xaa, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.FailedSize)
	// string "PendingCount"
	o = append(o, 0xac, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt64(o, z.PendingCount)
	// string "FailedCount"
	o = append(o, 0xab, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt64(o, z.FailedCount)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BucketReplicationStats) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Stats":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if z.Stats == nil {
				z.Stats = make(map[string]*BucketReplicationStat, zb0002)
			} else if len(z.Stats) > 0 {
				for key := range z.Stats {
					delete(z.Stats, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 *BucketReplicationStat
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Stats")
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					za0002 = nil
				} else {
					if za0002 == nil {
						za0002 = new(BucketReplicationStat)
					}
					bts, err = za0002.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Stats", za0001)
						return
					}
				}
				z.Stats[za0001] = za0002
			}
		case "PendingSize":
			z.PendingSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PendingSize")
				return
			}
		case "ReplicatedSize":
			z.ReplicatedSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ReplicatedSize")
				return
			}
		case "ReplicaSize":
			z.ReplicaSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ReplicaSize")
				return
			}
		case "FailedSize":
			z.FailedSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "FailedSize")
				return
			}
		case "PendingCount":
			z.PendingCount, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PendingCount")
				return
			}
		case "FailedCount":
			z.FailedCount, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "FailedCount")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BucketReplicationStats) Msgsize() (s int) {
	s = 1 + 6 + msgp.MapHeaderSize
	if z.Stats != nil {
		for za0001, za0002 := range z.Stats {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001)
			if za0002 == nil {
				s += msgp.NilSize
			} else {
				s += za0002.Msgsize()
			}
		}
	}
	s += 12 + msgp.Int64Size + 15 + msgp.Int64Size + 12 + msgp.Int64Size + 11 + msgp.Int64Size + 13 + msgp.Int64Size + 12 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BucketStats) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ReplicationStats":
			err = z.ReplicationStats.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "ReplicationStats")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BucketStats) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "ReplicationStats"
	err = en.Append(0x81, 0xb0, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73)
	if err != nil {
		return
	}
	err = z.ReplicationStats.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "ReplicationStats")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BucketStats) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "ReplicationStats"
	o = append(o, 0x81, 0xb0, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73)
	o, err = z.ReplicationStats.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "ReplicationStats")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BucketStats) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ReplicationStats":
			bts, err = z.ReplicationStats.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "ReplicationStats")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BucketStats) Msgsize() (s int) {
	s = 1 + 17 + z.ReplicationStats.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BucketStatsMap) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Stats":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if z.Stats == nil {
				z.Stats = make(map[string]BucketStats, zb0002)
			} else if len(z.Stats) > 0 {
				for key := range z.Stats {
					delete(z.Stats, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 BucketStats
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Stats")
					return
				}
				var zb0003 uint32
				zb0003, err = dc.ReadMapHeader()
				if err != nil {
					err = msgp.WrapError(err, "Stats", za0001)
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						err = msgp.WrapError(err, "Stats", za0001)
						return
					}
					switch msgp.UnsafeString(field) {
					case "ReplicationStats":
						err = za0002.ReplicationStats.DecodeMsg(dc)
						if err != nil {
							err = msgp.WrapError(err, "Stats", za0001, "ReplicationStats")
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							err = msgp.WrapError(err, "Stats", za0001)
							return
						}
					}
				}
				z.Stats[za0001] = za0002
			}
		case "Timestamp":
			z.Timestamp, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "Timestamp")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BucketStatsMap) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Stats"
	err = en.Append(0x82, 0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Stats)))
	if err != nil {
		err = msgp.WrapError(err, "Stats")
		return
	}
	for za0001, za0002 := range z.Stats {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Stats")
			return
		}
		// map header, size 1
		// write "ReplicationStats"
		err = en.Append(0x81, 0xb0, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73)
		if err != nil {
			return
		}
		err = za0002.ReplicationStats.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Stats", za0001, "ReplicationStats")
			return
		}
	}
	// write "Timestamp"
	err = en.Append(0xa9, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	if err != nil {
		return
	}
	err = en.WriteTime(z.Timestamp)
	if err != nil {
		err = msgp.WrapError(err, "Timestamp")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BucketStatsMap) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Stats"
	o = append(o, 0x82, 0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Stats)))
	for za0001, za0002 := range z.Stats {
		o = msgp.AppendString(o, za0001)
		// map header, size 1
		// string "ReplicationStats"
		o = append(o, 0x81, 0xb0, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73)
		o, err = za0002.ReplicationStats.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Stats", za0001, "ReplicationStats")
			return
		}
	}
	// string "Timestamp"
	o = append(o, 0xa9, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o = msgp.AppendTime(o, z.Timestamp)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BucketStatsMap) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Stats":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if z.Stats == nil {
				z.Stats = make(map[string]BucketStats, zb0002)
			} else if len(z.Stats) > 0 {
				for key := range z.Stats {
					delete(z.Stats, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 BucketStats
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Stats")
					return
				}
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Stats", za0001)
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "Stats", za0001)
						return
					}
					switch msgp.UnsafeString(field) {
					case "ReplicationStats":
						bts, err = za0002.ReplicationStats.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "Stats", za0001, "ReplicationStats")
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							err = msgp.WrapError(err, "Stats", za0001)
							return
						}
					}
				}
				z.Stats[za0001] = za0002
			}
		case "Timestamp":
			z.Timestamp, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Timestamp")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BucketStatsMap) Msgsize() (s int) {
	s = 1 + 6 + msgp.MapHeaderSize
	if z.Stats != nil {
		for za0001, za0002 := range z.Stats {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + 1 + 17 + za0002.ReplicationStats.Msgsize()
		}
	}
	s += 10 + msgp.TimeSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ReplicationLatency) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "UploadHistogram":
			err = z.UploadHistogram.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "UploadHistogram")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ReplicationLatency) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "UploadHistogram"
	err = en.Append(0x81, 0xaf, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d)
	if err != nil {
		return
	}
	err = z.UploadHistogram.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "UploadHistogram")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ReplicationLatency) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "UploadHistogram"
	o = append(o, 0x81, 0xaf, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d)
	o, err = z.UploadHistogram.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "UploadHistogram")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ReplicationLatency) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "UploadHistogram":
			bts, err = z.UploadHistogram.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "UploadHistogram")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ReplicationLatency) Msgsize() (s int) {
	s = 1 + 16 + z.UploadHistogram.Msgsize()
	return
}
