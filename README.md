# My WorkSpaces

![GitHub stars](https://img.shields.io/github/stars/JiangKlijna/my-workspaces)
![GitHub forks](https://img.shields.io/github/forks/JiangKlijna/my-workspaces)

> A collection of utility code snippets and learning examples in multiple programming languages.

## 📁 Project Structure

This repository contains various utility code snippets and learning examples across multiple programming languages.

```
my-workspaces/
├── go/           # Go utilities and servers
├── java/         # Java utilities and algorithms
├── cpp/          # C++ data structures and algorithms
├── html/         # HTML/JS demos and games
├── python/       # Python utilities
├── nodejs/       # Node.js utilities
├── php/          # PHP utilities
├── pwsh/         # PowerShell scripts
├── kotlin/       # Kotlin examples
├── clojure/      # Clojure examples
├── setting/      # Configuration files
└── .vscode/      # VS Code settings
```

---

## 🌍 Go

### Server Utilities

| File | Description |
|------|-------------|
| [server/ClusterServer.go](go/server/ClusterServer.go) | Cluster server - run multiple HTTP servers on different ports |
| [server/DynamicServer.go](go/server/DynamicServer.go) | Dynamic server with custom routing |
| [server/ProxyServer.go](go/server/ProxyServer.go) | Proxy server |
| [server/StaticServer.go](go/server/StaticServer.go) | Static file server |
| [server/ForwordServer.go](go/server/ForwordServer.go) | Forward server |
| [server/MonitorServer.go](go/server/MonitorServer.go) | Monitor server |
| [server/ResourceServer.go](go/server/ResourceServer.go) | Resource server |
| [server/UploadServer.go](go/server/UploadServer.go) | File upload server |

### Tools

| File | Description |
|------|-------------|
| [tool/compile.go](go/tool/compile.go) | Online code compiler using Judge0 API |
| [tool/datauri.go](go/tool/datauri.go) | Convert images to Base64 data URI |
| [tool/localip.go](go/tool/localip.go) | Get local IP address |
| [tool/logcmd.go](go/tool/logcmd.go) | Log command execution results |
| [tool/messagebox.go](go/tool/messagebox.go) | Cross-platform message box |
| [tool/recmd.go](go/tool/recmd.go) | Retry command execution |
| [tool/sort.go](go/tool/sort.go) | Sorting algorithms |
| [tool/bingimg.go](go/tool/bingimg.go) | Fetch Bing daily wallpaper |
| [tool/browser.go](go/tool/browser.go) | Browser automation |
| [tool/json2sql.go](go/tool/json2sql.go) | Convert JSON to SQL |
| [tool/sheet2sql.go](go/tool/sheet2sql.go) | Convert spreadsheet to SQL |
| [tool/multi_cpu.go](go/tool/multi_cpu.go) | Multi-core CPU utilities |
| [image/color-builder.go](go/image/color-builder.go) | Color builder for images |

---

## ☕ Java

### Core Utilities

| File | Description |
|------|-------------|
| [Hex.java](java/Hex.java) | Hexadecimal encoding/decoding |
| [RomanNumber.java](java/RomanNumber.java) | Roman numeral conversion |
| [Tuple.java](java/Tuple.java) | Tuple implementation |
| [ObjectMap.java](java/ObjectMap.java) | Object to Map conversion |
| [Compression.java](java/Compression.java) | Compress JAR files |
| [PngBuilder.java](java/PngBuilder.java) | Generate PNG with specified colors |

### Concurrency

| File | Description |
|------|-------------|
| [ThreadPool.java](java/ThreadPool.java) | Thread pool implementation |
| [AlternateThread.java](java/AlternateThread.java) | Alternate thread execution |
| [ProducerConsumer.java](java/ProducerConsumer.java) | Producer-Consumer pattern |
| [AssemblyLineFactory.java](java/AssemblyLineFactory.java) | Assembly line pattern (disruptor) |
| [Times.java](java/Times.java) | Timing utilities - timing wheel pattern |

### Algorithms

| File | Description |
|------|-------------|
| [Sort.java](java/Sort.java) | 8 sorting algorithms |

### JNI

| File | Description |
|------|-------------|
| [jni/](java/jni/) | Java Native Interface examples |

---

## 🖥️ C++

### Data Structures & Algorithms

| File | Description |
|------|-------------|
| [sort/](cpp/sort/) | Sorting algorithms implementation |
| [list/](cpp/list/) | Linked list implementation |
| [hex/](cpp/hex/) | Hex utilities |

### Platform Specific

| File | Description |
|------|-------------|
| [win32/](cpp/win32/) | Windows API examples |
| [termux/](cpp/termux/) | Termux (Android) utilities |
| [test/](cpp/test/) | Test files |
| [fpoo/](cpp/fpoo/) | Learning exercises |

---

## 🌐 HTML

### Games

| File | Description |
|------|-------------|
| [2048.html](html/2048.html) | 2048 game |
| [Shudu.html](html/Shudu.html) | Sudoku game |

### Utilities

| File | Description |
|------|-------------|
| [hex.html](html/hex.html) | Hex converter |
| [calendar.html](html/calendar.html) | Calendar display |
| [color_obj.html](html/color_obj.html) | Color format converter |
| [length.html](html/length.html) | Length unit converter |
| [urlform.html](html/urlform.html) | HTML form to API |
| [TopButton.html](html/TopButton.html) | Back to top button |
| [MouseMove.html](html/MouseMove.html) | Drag and drop div |
| [repeatpie.html](html/repeatpie.html) | Repeating pie chart |
| [sousuo.html](html/sousuo.html) | Search engine |

---

## 🐍 Python

| File | Description |
|------|-------------|
| [latlng.py](python/latlng.py) | Coordinate conversion (WGS/GCJ/BD) |
| [sort.py](python/sort.py) | Sorting algorithms |
| [pip_update.py](python/pip_update.py) | Batch pip update utility |

---

## 🟢 Node.js

| File | Description |
|------|-------------|
| [family.js](nodejs/family.js) | Family birth statistics simulation |
| [latlng.js](nodejs/latlng.js) | Coordinate conversion |
| [sort.js](nodejs/sort.js) | Sorting algorithms |
| [rename.js](nodejs/rename.js) | Batch file rename |

---

## 🔴 PHP

| File | Description |
|------|-------------|
| [methods.php](php/methods.php) | Common utility methods |
| [sort.php](php/sort.php) | Sorting algorithms |
| [upload.php](php/upload.php) | File upload handler |

---

## ⚡ PowerShell

| File | Description |
|------|-------------|
| [sort.ps1](pwsh/sort.ps1) | Sorting algorithms |
| [pwshrc.ps1](pwsh/pwshrc.ps1) | PowerShell configuration |

---

## 🟡 Kotlin

| File | Description |
|------|-------------|
| [Sort.kt](kotlin/Sort.kt) | Sorting algorithms |

---

## ⚙️ Clojure

| File | Description |
|------|-------------|
| [learn.clj](clojure/learn.clj) | Learning examples |

---

## ⚙️ Settings

| File | Description |
|------|-------------|
| [setting/.bashrc](setting/.bashrc) | Bash configuration |
| [setting/soft.yml](setting/soft.yml) | Software list |

---

## 🔧 VS Code Settings

| File | Description |
|------|-------------|
| [.vscode/launch.json](.vscode/launch.json) | Debug configurations |
| [.vscode/tasks.json](.vscode/tasks.json) | Build tasks |

---

## 📝 License

[MIT](LICENSE)

---

> Last updated: 2026-03-02
