package main

import (
	"fmt"
)

func changeValue(val *int) {
	// ポインタを受け取って、値を変更する関数の例。
	fmt.Println("関数内の元の値:", *val)
	*val = 50 // ポインタ経由で値を変更
	fmt.Println("関数内の変更後の値:", *val)
	val = nil // ポインタをnilに設定（この操作は呼び出し元には影響しない）
	fmt.Println("関数内のポインタをnilに設定しました。", val)
}

func pointerExampleWithFunction() {
	// ポインタを関数に渡して、関数内で値を変更する例。
	var y int = 30
	fmt.Println("元の値:", y)
	changeValue(&y) // yのアドレスを渡す
	fmt.Println("変更後の値:", y)
}

func SecureArea() {
	// 領域確保の例。Goでは、new関数やmake関数を使用して、メモリを確保することができます。
	// new関数は、指定した型のゼロ値を持つ新しい変数を作成し、そのアドレスを返します。
	// make関数は、スライス、マップ、チャネルなどの組み込み型の初期化に使用されます。
	fmt.Println("領域確保の例")
	// new関数を使用して、int型の変数を作成し、そのアドレスを取得。
	var p *int = new(int)
	fmt.Println("new関数で確保された領域のアドレス:", p)
	fmt.Println("new関数で確保された領域の初期値:", *p)
	*p = 100 // ポインタを介して値を設定
	fmt.Println("new関数で確保された領域の変更後の値:", *p)

	// make関数を使用して、スライスを作成。
	slice := make([]int, 5)
	fmt.Println("make関数で確保されたスライス:", slice)
	for i := range slice {
		slice[i] = i * 10 // スライスの各要素に値を設定
	}
	fmt.Println("make関数で確保されたスライス (変更後):", slice)
}

func gotoExample() {
	SecureArea() // 領域確保の例を実行する関数を呼び出し。f
	// goto文を使用して、コードの特定の位置にジャンプする例。
	fmt.Println("goto文の例")
	var i int = 0
Start:
	fmt.Println(i)
	i++
	fmt.Println("goto文でループします。")
	if i < 5 {
		goto Start // Startラベルにジャンプ
	}
	fmt.Println("goto文のループが終了しました。")
}

/*
goto文は、コードの特定の位置にジャンプするための構文です。
ただし、goto文はコードの可読性を低下させる可能性があるため、使用は推奨されません。
goto文を使用する代わりに、for文やif文などの制御構造を使用して、同様のことができます。
while文は、条件が真の間、繰り返し処理を行うための構文です。
go言語にはwhile文はありませんが、for文を使って同様のことができます。
*/
