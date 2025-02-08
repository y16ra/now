package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/y16ra/now/internal/timeutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func runGUI(formatter *timeutil.Formatter) {
	myApp := app.New()
	window := myApp.NewWindow("Now")

	// 時刻表示用のラベル
	timeLabel := widget.NewLabel("")
	
	// タイムゾーン表示用のラベル
	zoneLabel := widget.NewLabel("")
	
	// フォーマット入力用のエントリー
	formatEntry := widget.NewEntry()
	formatEntry.SetText(formatter.GetFormat())
	
	// 時刻更新関数
	updateTime := func() {
		formatted, zone := formatter.FormatWithZone(time.Now())
		timeLabel.SetText(formatted)
		zoneLabel.SetText("Timezone: " + zone)
	}
	
	// フォーマット変更時のイベントハンドラ
	formatEntry.OnChanged = func(format string) {
		formatter.SetFormat(format)
		updateTime()
	}
	
	// 初期表示
	updateTime()
	
	// 定期的な時刻更新
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			updateTime()
		}
	}()
	
	// レイアウトの設定
	content := container.NewVBox(
		widget.NewLabel("Format:"),
		formatEntry,
		timeLabel,
		zoneLabel,
	)
	
	window.SetContent(content)
	window.Resize(fyne.NewSize(300, 180)) // ウィンドウを少し大きくする
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
