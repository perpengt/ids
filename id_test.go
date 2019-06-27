package ids

import "testing"

const generateCount = 1000000

func TestGenerateID(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < generateCount; i++ {
		id := GenerateID()
		m[id.String()] = true
	}

	l := len(m)
	if l != generateCount {
		t.Errorf("id conflicted: %d out of %d IDs are conflicted", generateCount-l, generateCount)
	}
}
