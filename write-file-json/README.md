
## ファイルに文字列を複数行書き込む

- ファイルの保存はプレーンテキストと同様に[WriteFile](https://pkg.go.dev/io/ioutil#example-WriteFile)を使う
- 第一引数: ファイル名
- 第二引数: 書き込む内容("\n"で区切ることで複数行書き込み可能)
- 第三引数: ファイルの権限設定。例えばos.ModePermは0777と同義。

JSON形式にするには[json.MarshalIndent](https://pkg.go.dev/encoding/json#MarshalIndent)を使用する
> マーシャリングとは、異なる技術基盤で実装されたコンピュータプログラム間で、データの交換や機能の呼び出しができるようデータ形式の変換などを行うこと。

json.MarshalIndentを使うと、オブジェクトを指定したインデント付きで変換（つまり整形）できる
戻り値が[]byteなのでそのままWriteFileの第二引数に渡せる

例: 
```golang
contents :=
    `{
        "id": 1,
        "name": "Hello",
        "job": "World!"
    }`
ioutil.WriteFile("hello.json", []byte(contents), os.ModePerm)
```

出力されるファイル
```:hello.json
{
	"id": 1,
	"name": "Hello",
	"job": "World!"
}
```

## プログラム実行例:


```
go run main.go 

Enter file name
> favorites.json
filename is  favorites.json

Enter task name (or press [Enter-Key] to finish):
> test1
> test2
> test3
> 

Successfully saved!
```

ファイル内容
```:favorites.json
[
    {
        "ID": 1,
        "Title": "test1",
        "CreatedAt": "2023-11-18T19:13:11.217689+09:00"
    },
    {
        "ID": 2,
        "Title": "test2",
        "CreatedAt": "2023-11-18T19:13:12.304812+09:00"
    },
    {
        "ID": 3,
        "Title": "test3",
        "CreatedAt": "2023-11-18T19:13:13.552073+09:00"
    }
]
```