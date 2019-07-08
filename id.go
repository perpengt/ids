package ids

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"strings"
	"time"
)

const (
	TimestampMax uint64 = 0x1FFFFFFFFFF
	SequenceMax  uint64 = 0xFFF
	IDSize              = 8
)

type ID []byte

var (
	strLen = base64.StdEncoding.EncodedLen(IDSize)
	endian = binary.BigEndian

	lastTime uint64 = 0
	sequence uint64 = 0
)

// New creates new id from []byte.
func New(data []byte) *ID {
	_ = data[7]
	id := &ID{data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7]}
	return id
}

// GenerateID creates and returns a completely unique ID.
//
// However, if two machines with the same Machine ID receive
// the same request at the same time, there is a possibility of a conflict.
func GenerateID(machineID uint64) *ID {
	now := uint64(time.Now().UnixNano())

	if now > lastTime {
		lastTime = now
		sequence = 0
	} else if sequence >= SequenceMax {
		lastTime++
		sequence = 0
	} else {
		sequence++
	}

	n := (sequence&SequenceMax)<<52 + machineID<<41 + (lastTime & TimestampMax)

	result := make([]byte, IDSize)
	endian.PutUint64(result, n)

	return (*ID)(&result)
}

func DecodeID(id string) (*ID, error) {
	// URL-Encoded ID?
	id = strings.Replace(strings.Replace(id, "-", "+", -1), "_", "/", -1) + strings.Repeat("=", 4-(len(id)%4))

	if len(id) != strLen {
		return nil, ErrInvalidID
	}

	result := make([]byte, base64.StdEncoding.DecodedLen(strLen))

	_, err := base64.StdEncoding.Decode(result, []byte(id))
	if err != nil {
		return nil, err
	}

	result = result[:IDSize]

	return (*ID)(&result), nil
}

func (id *ID) Valid() error {
	if id == nil {
		return ErrNilID
	}
	if len(*id) != IDSize {
		return ErrWrongSize
	}
	return nil
}

func (id *ID) Key() Key {
	arr := *id
	_ = arr[7]
	return Key{arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7]}
}

func (id *ID) String() string {
	return base64.StdEncoding.EncodeToString(*id)
}

func (id *ID) Array() [8]byte {
	return id.Key()
}

func (id *ID) URIString() string {
	return base64.RawURLEncoding.EncodeToString(*id)
}

func (id *ID) Bytes() []byte {
	return *id
}

func Equal(a *ID, b *ID) bool {
	return bytes.Equal(*a, *b)
}
