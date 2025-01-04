//go:build !windows && webview

package desktop

import "github.com/webview/webview"

func StartWebView(url string, debug bool) {
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("eKonyv")
	w.SetSize(1200, 800, webview.HintNone)
	w.Navigate(url)
	w.Run()
}
