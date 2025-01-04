//go:build !webview

package desktop

import "github.com/manfromdownunder/eKonyv/util"

func StartWebView(url string, debug bool) {
	util.OpenBrowser(url)

	<-make(chan struct{})
}
