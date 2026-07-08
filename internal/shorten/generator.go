package shorten

// Generator creates short codes from database IDs.
// Stub — will be implemented with base-62 encoding.
type Generator struct{}

// NewGenerator creates a new code generator.
func NewGenerator() *Generator {
	return &Generator{}
}

// Encode converts an integer ID to a short code string.
func (g *Generator) Encode(id int64) string {
	// TODO: implement base-62 encoding
	return ""
}

// Decode converts a short code back to an integer ID.
func (g *Generator) Decode(code string) int64 {
	// TODO: implement base-62 decoding
	return 0
}
