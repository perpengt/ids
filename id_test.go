package ids

import "testing"

const generateCount = 1000000

func TestNew(t *testing.T) {

}

func TestGenerateID(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < generateCount; i++ {
		id := GenerateID(1)
		m[id.String()] = true
	}

	l := len(m)
	if l != generateCount {
		t.Errorf("id conflicted: %d out of %d IDs are conflicted", generateCount-l, generateCount)
	}
}
