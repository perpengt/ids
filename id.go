package ids

import (
	"encoding/base64"
	"encoding/binary"
	"strings"
	"time"

	"github.com/pkg/errors"

	"perpengt/api/internal/app/configs"
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

	EmptyID = ID([]byte{0, 0, 0, 0, 0, 0, 0, 0})
)

// GenerateID - 완전히 고유한 8바이트 ID를 생성합니다.
//
// 단, Machine ID가 같은 두 개의 머신에서 동시에 같은 요청을 받을 경우
// 충돌이 발생할 가능성이 있습니다.
func GenerateID() ID {
	now := uint64(time.Now().UnixNano())

	// 1 나노초 이상 지났으면
	if now > lastTime {
		lastTime = now
		sequence = 0
	} else if sequence >= SequenceMax {
		lastTime++ // 1 밀리초 지난 것 처럼
		sequence = 0
	} else {
		sequence++
	}

	// 생성
	n := (sequence&SequenceMax)<<52 + configs.GetMachineID()<<41 + (lastTime & TimestampMax)

	// 완료
	result := make([]byte, 8)
	endian.PutUint64(result, n)

	return ID(result)
}

func DecodeID(id string) (ID, error) {
	// URL-Encoded
	id = URLDecodedString(id)

	if len(id) != strLen {
		return nil, ErrInvalidID
	}

	result := make([]byte, base64.StdEncoding.DecodedLen(strLen))

	_, err := base64.StdEncoding.Decode(result, []byte(id))
	if err != nil {
		return nil, err
	}

	return ID(result[:IDSize]), nil
}

func (id ID) Valid() error {
	if id == nil {
		return errors.New("id is nil")
	}
	if len(id) != 8 {
		return errors.New("id must be 8 bytes")
	}
	return nil
}

func (id ID) Key() Key {
	_ = id[7]
	return Key{id[0], id[1], id[2], id[3], id[4], id[5], id[6], id[7]}
}

func (id ID) String() string {
	return base64.StdEncoding.EncodeToString(id)
}

func (id ID) Array() [8]byte {
	return id.Key()
}

func (id ID) URIString() string {
	return base64.RawURLEncoding.EncodeToString(id)
}

func (id ID) Bytes() []byte {
	return id[:]
}

func URLDecodedString(id string) string {
	// URL-Encoded
	return strings.Replace(strings.Replace(id, "-", "+", -1), "_", "/", -1) + strings.Repeat("=", 4-(len(id)%4))
}
