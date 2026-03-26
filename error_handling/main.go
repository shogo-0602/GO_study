package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func init() {
	// エラー処理の例として、存在しないファイルを作成しておきます。
	// これにより、エラーが発生する状況を再現できます。
	f, err := os.Create("dt/test.csv")
	if err != nil {
		fmt.Printf("ファイルの作成に失敗: %v\n", err)
		return
	}
	// ファイルに初期値を書き込む
	var data string = "id,name,age\n1,Alice,30\n2,Bob,25\n"
	_, err = f.WriteString(data)
	if err != nil {
		fmt.Printf("ファイルへの書き込みに失敗: %v\n", err)
	}
	f.Close()
}

func readFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// ファイルの内容を読み取る処理
	return io.ReadAll(f)
}

func main() {
	//　GO言語では、エラー処理が重要で関数ごとにエラーを返すことが一般的です。
	// 戻り値は、値とエラーの2つを返すことが多いです。
	// 例えば、ファイルを開く関数は、ファイルの内容とエラーを返します。
	// エラーがnilでない場合は、エラーが発生したことを意味します。
	f, err := os.Open("dt/test.csv")

	// エラーが発生した場合の処理
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("ファイルが存在しません")
		} else {
			fmt.Printf("ファイルのオープンに失敗: %v\n", err)
		}
	}

	// エラーが発生しなかった場合の処理
	defer f.Close() // ファイルを閉じる処理

	// ファイルを読み取る処理
	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("ファイルの読み込みに失敗: %v\n", err)
	}
	fmt.Printf("test.csvの内容: \n%s", data)

	// 存在しないファイルを開く例
	data, err = readFile("dt/nonexistent.csv")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("nonexistent.csvが存在しません")
			return
		} else {
			fmt.Printf("nonexistent.csvのオープンに失敗: %v\n", err)
			return
		}
	}

}
