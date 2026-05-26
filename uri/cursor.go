package uri

type cursor struct {
	src string
	pos int
}

func (c *cursor) readUntil(until byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *cursor) readValue(sep byte) (v string, hasNext bool, err error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func (c *cursor) eat(r byte) bool { _ = "STUB: not implemented"; return false }

func (c *cursor) readAll() (string, error) { _ = "STUB: not implemented"; return "", nil }
