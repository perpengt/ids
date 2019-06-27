package ids

type Key [8]byte

var EmptyKey = Key{}

func (k Key) ID() ID {
	return ID(k[:])
}

func (k Key) String() string {
	return k.ID().String()
}

func (k Key) URIString() string {
	return k.ID().URIString()
}

func (k Key) Bytes() []byte {
	return k[:]
}
