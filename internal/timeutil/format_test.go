package timeutil

import (
	"testing"
	"time"
)

func TestFormatter_Format(t *testing.T) {
	// テスト用の固定時刻を設定
	fixedTime := time.Date(2025, 2, 8, 16, 51, 40, 0, time.Local)

	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{
			name:     "デフォルトフォーマット",
			format:   "",
			expected: "2025-02-08 16:51:40",
		},
		{
			name:     "日付のみ",
			format:   "2006-01-02",
			expected: "2025-02-08",
		},
		{
			name:     "時刻のみ",
			format:   "15:04:05",
			expected: "16:51:40",
		},
		{
			name:     "日本語形式",
			format:   "2006年01月02日 15時04分05秒",
			expected: "2025年02月08日 16時51分40秒",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(tt.format)
			got := f.Format(fixedTime)
			if got != tt.expected {
				t.Errorf("Format() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFormatter_SetFormat(t *testing.T) {
	f := NewFormatter("")
	originalFormat := f.GetFormat()

	// 新しいフォーマットを設定
	newFormat := "15:04:05"
	f.SetFormat(newFormat)
	if got := f.GetFormat(); got != newFormat {
		t.Errorf("SetFormat() = %v, want %v", got, newFormat)
	}

	// 空文字列を設定した場合、フォーマットは変更されない
	f.SetFormat("")
	if got := f.GetFormat(); got != newFormat {
		t.Errorf("SetFormat() with empty string changed format: got %v, want %v", got, newFormat)
	}

	// 元のフォーマットと異なることを確認
	if originalFormat == f.GetFormat() {
		t.Error("Format should have changed from original")
	}
}

func TestNewFormatter(t *testing.T) {
	// デフォルトフォーマットのテスト
	f1 := NewFormatter("")
	if f1.GetFormat() != "2006-01-02 15:04:05" {
		t.Errorf("NewFormatter() with empty string didn't set default format")
	}

	// カスタムフォーマットのテスト
	customFormat := "2006/01/02"
	f2 := NewFormatter(customFormat)
	if f2.GetFormat() != customFormat {
		t.Errorf("NewFormatter() didn't set custom format correctly")
	}
}
