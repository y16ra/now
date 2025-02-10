package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/y16ra/now/internal/timeutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// フォーマットのプリセット
var formatPresets = map[string]string{
	"Default (2006-01-02 15:04:05)": "2006-01-02 15:04:05",
	"US (MM/DD/YYYY)": "01/02/2006",
	"US with Time (MM/DD/YYYY hh:mm:ss AM/PM)": "01/02/2006 03:04:05 PM",
	"UK/EU (DD/MM/YYYY)": "02/01/2006",
	"UK/EU with Time (DD/MM/YYYY HH:mm:ss)": "02/01/2006 15:04:05",
	"EU Standard (DD.MM.YYYY)": "02.01.2006",
	"EU with Time (DD.MM.YYYY HH:mm:ss)": "02.01.2006 15:04:05",
	"Date Only (YYYY-MM-DD)": "2006-01-02",
	"Time Only (HH:mm:ss)": "15:04:05",
	"Time with AM/PM": "03:04:05 PM",
	"RFC3339": time.RFC3339,
	"ISO8601": "2006-01-02T15:04:05-07:00",
	"Custom": "",
}

// タイムゾーンのプリセット
var timezonePresets = map[string]string{
	"Local Time": "Local",
	"UTC": "UTC",
	"Asia/Tokyo (JST)": "Asia/Tokyo",
	"America/New_York (EST/EDT)": "America/New_York",
	"America/Los_Angeles (PST/PDT)": "America/Los_Angeles",
	"Europe/London (GMT/BST)": "Europe/London",
	"Europe/Paris (CET/CEST)": "Europe/Paris",
	"Asia/Shanghai": "Asia/Shanghai",
	"Asia/Singapore": "Asia/Singapore",
	"Australia/Sydney": "Australia/Sydney",
}

func runGUI(formatter *timeutil.Formatter) {
	myApp := app.New()
	window := myApp.NewWindow("Now")

	// 現在のタイムゾーンを保持
	currentLocation, _ := time.LoadLocation("Local")

	// 時刻表示用のラベル
	timeLabel := widget.NewLabel("")
	
	// タイムゾーン表示用のラベル
	zoneLabel := widget.NewLabel("")
	
	// フォーマットのオプションを準備
	formatOptions := make([]string, 0, len(formatPresets))
	for k := range formatPresets {
		formatOptions = append(formatOptions, k)
	}

	// タイムゾーンのオプションを準備
	timezoneOptions := make([]string, 0, len(timezonePresets))
	for k := range timezonePresets {
		timezoneOptions = append(timezoneOptions, k)
	}
	
	// フォーマット検索用のエントリー
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Type to search formats...")
	
	// タイムゾーン検索用のエントリー
	tzSearchEntry := widget.NewEntry()
	tzSearchEntry.SetPlaceHolder("Type to search timezones...")
	
	// クリアボタン
	clearButton := widget.NewButton("Clear", nil)
	
	// タイムゾーン検索用のクリアボタン
	tzClearButton := widget.NewButton("Clear", nil)
	
	// 検索エリアのコンテナ
	searchContainer := container.NewBorder(
		nil, nil, nil, clearButton,
		searchEntry,
	)

	// タイムゾーン検索エリアのコンテナ
	tzSearchContainer := container.NewBorder(
		nil, nil, nil, tzClearButton,
		tzSearchEntry,
	)
	
	// フォーマット入力用のエントリー（カスタム用）
	formatEntry := widget.NewEntry()
	formatEntry.SetText(formatter.GetFormat())
	formatEntry.Hide()
	
	// フィルタリングされたオプションを保持する変数
	var filteredOptions []string
	var filteredTZOptions []string
	
	// 時刻更新関数
	updateTime := func() {
		now := time.Now().In(currentLocation)
		formatted, zone := formatter.FormatWithZone(now)
		timeLabel.SetText(formatted)
		zoneLabel.SetText("Timezone: " + zone)
	}
	
	// フォーマット選択用のセレクト
	formatSelect := widget.NewSelect(formatOptions, func(selected string) {
		if selected == "Custom" {
			formatEntry.Show()
			formatEntry.SetText(formatter.GetFormat())
		} else {
			formatEntry.Hide()
			if format, ok := formatPresets[selected]; ok {
				formatter.SetFormat(format)
				updateTime()
			}
		}
	})

	// タイムゾーン選択用のセレクト
	tzSelect := widget.NewSelect(timezoneOptions, func(selected string) {
		if tz, ok := timezonePresets[selected]; ok {
			if loc, err := time.LoadLocation(tz); err == nil {
				currentLocation = loc
				updateTime()
			}
		}
	})
	
	// 検索エントリーの変更時のイベントハンドラ
	searchEntry.OnChanged = func(text string) {
		if text == "" {
			formatSelect.Options = formatOptions
			return
		}
		
		searchText := strings.ToLower(text)
		
		filteredOptions = nil
		for _, opt := range formatOptions {
			if strings.Contains(strings.ToLower(opt), searchText) {
				filteredOptions = append(filteredOptions, opt)
			}
		}
		
		formatSelect.Options = filteredOptions
		if len(filteredOptions) > 0 {
			formatSelect.SetSelectedIndex(0)
		}
	}

	// タイムゾーン検索エントリーの変更時のイベントハンドラ
	tzSearchEntry.OnChanged = func(text string) {
		if text == "" {
			tzSelect.Options = timezoneOptions
			return
		}
		
		searchText := strings.ToLower(text)
		
		filteredTZOptions = nil
		for _, opt := range timezoneOptions {
			if strings.Contains(strings.ToLower(opt), searchText) {
				filteredTZOptions = append(filteredTZOptions, opt)
			}
		}
		
		tzSelect.Options = filteredTZOptions
		if len(filteredTZOptions) > 0 {
			tzSelect.SetSelectedIndex(0)
		}
	}
	
	// クリアボタンのイベントハンドラ
	clearButton.OnTapped = func() {
		searchEntry.SetText("")
		formatSelect.Options = formatOptions
	}

	// タイムゾーンクリアボタンのイベントハンドラ
	tzClearButton.OnTapped = func() {
		tzSearchEntry.SetText("")
		tzSelect.Options = timezoneOptions
	}
	
	// フォーマット変更時のイベントハンドラ（カスタム用）
	formatEntry.OnChanged = func(format string) {
		formatter.SetFormat(format)
		updateTime()
	}
	
	// 初期選択
	formatSelect.SetSelected("Default (2006-01-02 15:04:05)")
	tzSelect.SetSelected("Local Time")
	
	// 初期表示
	updateTime()
	
	// レイアウトの設定
	formatGroup := widget.NewCard("Format", "", container.NewVBox(
		searchContainer,
		formatSelect,
		formatEntry,
	))

	timezoneGroup := widget.NewCard("Timezone", "", container.NewVBox(
		tzSearchContainer,
		tzSelect,
	))

	content := container.NewVBox(
		formatGroup,
		timezoneGroup,
		timeLabel,
		zoneLabel,
	)
	
	window.SetContent(content)
	
	// 定期的な時刻更新
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			updateTime()
		}
	}()
	
	window.Resize(fyne.NewSize(400, 300))
	window.ShowAndRun()
}

func runCLI(formatter *timeutil.Formatter) {
	formatted, zone := formatter.FormatWithZone(time.Now())
	fmt.Printf("%s (%s)\n", formatted, zone)
}

func main() {
	// コマンドラインフラグの定義
	guiMode := flag.Bool("gui", false, "GUIモードで起動")
	format := flag.String("f", "", "時刻フォーマット (例: 2006-01-02 15:04:05)")
	flag.Parse()

	// フォーマッタの初期化
	formatter := timeutil.NewFormatter(*format)

	// モードに応じて実行
	if *guiMode {
		runGUI(formatter)
	} else {
		runCLI(formatter)
	}
}
