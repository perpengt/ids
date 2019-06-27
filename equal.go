package ids

import "bytes"

func Equal(a ID, b ID) bool {
	return bytes.Equal(a, b)
}
