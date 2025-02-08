package timeutil

import "time"

// Formatter は時刻のフォーマット処理を行う構造体です
type Formatter struct {
	format string
}

// NewFormatter は新しいFormatterを作成します
func NewFormatter(format string) *Formatter {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return &Formatter{format: format}
}

// Format は指定されたフォーマットで時刻を文字列に変換します
func (f *Formatter) Format(t time.Time) string {
	return t.Format(f.format)
}

// FormatWithZone は時刻とタイムゾーン情報を含む文字列を返します
func (f *Formatter) FormatWithZone(t time.Time) (string, string) {
	formatted := f.Format(t)
	name, offset := t.Zone()
	// オフセットを時間と分に変換
	hours := offset / 3600
	minutes := (offset % 3600) / 60
	zone := name
	if name == "" {
		// タイムゾーン名が空の場合は、UTC±HH:MMの形式で表示
		sign := "+"
		if hours < 0 {
			sign = "-"
			hours = -hours
			minutes = -minutes
		}
		zone = "UTC" + sign + formatTwoDigits(hours) + ":" + formatTwoDigits(minutes)
	}
	return formatted, zone
}

// SetFormat はフォーマット文字列を設定します
func (f *Formatter) SetFormat(format string) {
	if format != "" {
		f.format = format
	}
}

// GetFormat は現在のフォーマット文字列を返します
func (f *Formatter) GetFormat() string {
	return f.format
}

// formatTwoDigits は数値を2桁の文字列に変換します
func formatTwoDigits(n int) string {
	if n < 10 {
		return "0" + string(rune('0'+n))
	}
	return string(rune('0'+n/10)) + string(rune('0'+n%10))
}
