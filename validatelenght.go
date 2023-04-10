package validatelength

const (
	ge = iota // greater or equal
	le        // lower or equal
	e         // equal
	l         // lower
	g         // greater
)

func ValidateLen(s string, m uint8, constraint int) (bool, error) {
	switch m {
	case ge:
		return len(s) >= constraint, nil
	case le:
		return len(s) <= constraint, nil
	case e:
		return len(s) == constraint, nil
	case l:
		return len(s) < constraint, nil
	case g:
		return len(s) > constraint, nil
	}
	return false, nil

}
