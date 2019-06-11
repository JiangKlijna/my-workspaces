package main

import (
	"flag"
	"github.com/zserge/webview"
	"os"
	"os/exec"
	"path/filepath"
)

var set = new(webview.Settings)

func init() {
	var (
		help bool
	)
	flag.StringVar(&(set.Title), "t", "JiangKlijnaBrowser", "WebView main window title")
	flag.StringVar(&(set.URL), "u", "index.html", "URL to open in a webview")
	flag.IntVar(&(set.Width), "w", 1000, "Window width in pixels")
	flag.IntVar(&(set.Height), "h", 800, "Window height in pixels")
	flag.BoolVar(&(set.Resizable), "r", true, "Allows/disallows window resizing")
	flag.BoolVar(&(set.Debug), "d", false, "Enable debugging tools (Linux/BSD/MacOS, on Windows use Firebug)")

	flag.BoolVar(&help, "help", false, "this help")
	flag.Parse()
	if help {
		println(`Usage: browser [-t title] [-u url] [-w width] [-h height] [-r resizable] [-d debug]
Example: browser -t JiangKlijnaBrowser -u index.html -w 800 -h 800 -r true -d false

Options:`)
		flag.PrintDefaults()
		os.Exit(1)
	}

	url, err := exec.LookPath(set.URL)
	if err == nil {
		set.URL, _ = filepath.Abs(url)
	}
}

func main() {
	webview.New(*set).Run()
}