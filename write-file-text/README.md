
## ファイルに文字列を複数行書き込む

- [WriteFile](https://pkg.go.dev/io/ioutil#example-WriteFile)でファイルに書き込む
- 第一引数: ファイル名
- 第二引数: 書き込む内容("\n"で区切ることで複数行書き込み可能)
- 第三引数: ファイルの権限設定。例えばos.ModePermは0777と同義。

例: 
```golang
ioutil.WriteFile("hello.txt", []byte("Hello\nWorld\n!"), 0644)
```

出力されるファイル
```:hello.txt
Hello
World
!
```

## プログラム実行例:


```
go run main.go

Enter file name
> favorites.txt

filename is  favorites.txt

Enter text (or press [Enter-Key] to finish): 
> test1
> test2
> test3
> 
Successfully saved!
write-file-text >>>
```

ファイル内容
```:favorites.txt
test1
test2
test3
```