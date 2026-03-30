package main

import (
	"fmt"
)

func pointerExample() {
	// ポインタを使用して、変数のアドレスを取得し、値を変更する例。
	var x int = 10
	fmt.Println("元の値:", x)
	p := &x // xのアドレスを取得
	*p = 20 // ポインタ経由で値を変更
	fmt.Println("変更後の値:", x)
}

func gotoExample() {
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
