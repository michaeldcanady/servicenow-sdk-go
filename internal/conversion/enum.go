package conversion

// EnumString looks up v in m and returns the matching string, or fallback if v has no entry.
func EnumString[T comparable](m map[T]string, v T, fallback string) string {
	if str, ok := m[v]; ok {
		return str
	}
	return fallback
}
