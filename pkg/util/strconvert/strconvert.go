package strconvert

import "time"

// StringOrDefault は s をstring型で返す
func StringOrDefault(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// ToTime は s をtime型に変換する
func ToTime(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

func ToTimePtr(s string) *time.Time {
	t := ToTime(s)
	if t.IsZero() {
		return nil
	}
	return &t
}
