# Unclipper

## 概要
`Unclipper`は、クリップボードにコピーされたパスを読み取り、関連するアプリケーションで開くためのWindows向けプログラムです。

## 使用方法
1. パスをクリップボードにコピーします。
1. `unclipper.exe`を実行します。プログラムがクリップボードの内容を解析し、適切なアプリケーションでパスを開きます。

## パスの例
- ファイルは基本的に既定のアプリで開かれます。フォルダーはエクスプローラーで開きます。
    - `C:\Users\sample\Downloads\sample.txt`
    - `C:\Users\sample\Desktop\sample\`
- `""`や`<>`で囲んでいても開きます。
    - `"C:\Users\sample\Desktop\sample.jpg"`
    - `<C:\Users\sample\Desktop\sample.pdf>`
- WordファイルとExcelファイルは新しいインスタンスで開きます。Microsoft WordやMicrosoft Excelがインストールされている必要があります。
    - `D:\Documents\sample.docx`
    - `D:\Documents\sample.xlsx`
- 複数行の場合でも開きます。
    ```
    C:\Users\sample\Downloads\sample.txt
    C:\Users\sample\Desktop\sample
    ```
- パスの区切り文字が`/`でも開きます。
    - `E:/Videos/sample.mp4`
- UNCパスも開きます。
    - `\\192.168.1.10\public`
- Git Bashなどで使われるMSYS2形式のパスも開きます。ただし`/c/...`や`/d/...`などのドライブレター形式に限られます。`/usr/...`のような仮想ルート形式には対応していません。
    - `/d/Projects/sample.txt`
- 先頭に`- `があっても開きます。これはMarkdownのリストでパスをメモしている場合に役立ちます。ネストされていても開きます。
    - `- "C:\Users\sample\Music\sample.wav"`
    - `    - C:\Users\sample\Music\sample.flac`
- HTMLコメントされている場合は開きません。
    - `<!-- C:\Users\sample\Pictures\sample.png -->`
- 以下のようなパスは開かれます。これはSumatraPDF v3.6でPDFファイルに埋め込まれたUNCパスのリンクを正常に開けない問題に対応するための機能です。リンク先のアドレスをコピーし、`unclipper.exe`を実行してください。
    - `\192.168.1.100\share\sample.txt`

## ビルド方法
`go build -ldflags "-s -w -H windowsgui" -trimpath -o unclipper.exe`

## ライセンス
このプログラムはMITライセンスの下で公開されています。
