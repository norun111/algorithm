package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func returnInt(result []int) int {
	sum := 0
	for _, r := range result {
		if r == 1 {
			sum++
		}
	}
	return sum
}

func main() {
	// 自分の得意な言語で
	// Let's チャレンジ！！
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	tate, _ := strconv.Atoi(list[0])
	yoko, _ := strconv.Atoi(list[1])

	preList := make([]int, yoko)

	result := 0

	for i := 0; i < tate; i++ {
		// fmt.Println("**************")
		sc.Scan()
		kadanList := strings.Split(sc.Text(), "")

		sum := 0

		for j, k := range kadanList {
			// fmt.Println(j,k)

			//横が1つより多い時 j+1がruntimeerrorになるから そしてj+1 j-1を調べるには3つ以上ではないとだめ
			if yoko > 2 {
				// ・ or #
				//  # ・ or # #
				//一番端の場合かつ花壇の場合
				if j == 0 && k == "#" {

					//上に花壇がない場合
					if preList[j] == 0 {

						if kadanList[j+1] == "#" {
							//.
							//# #
							preList[j] = 1 //次の調査の為に花壇の存在を入力
							sum += 2
						} else {
							// .
							// # ・
							sum += 2
							preList[j] = 1
						}
					} else {
						// preListが1だった場合 上に花壇がある
						if kadanList[j+1] == "#" {
							//#
							//# #
							sum++
							preList[j] = 1
						} else {
							//#
							// # .
							sum++
							preList[j] = 1
						}
					}

					//一番端の場合かつ花壇ではない場合
				} else if j == 0 && k == "." {
					//上に花壇がない場合
					if preList[j] == 0 {

						if kadanList[j+1] == "#" {
							// .
							// . #
							preList[j] = 0 //次の調査の為に花壇の存在を入力
						} else {
							// .
							// . .
							preList[j] = 0
						}
					} else {
						// preListが1だった場合 上に花壇がある
						if kadanList[j+1] == "#" {
							// #
							// . #
							sum++ // .の上にロープ
							preList[j] = 0
						} else {
							// #
							// . .
							sum++ // .の上にロープ
							preList[j] = 0
						}
					}
				} else if j == yoko-1 && k == "#" {
					// 最後の時かつ花壇がある時 ・・・# <-ココ 最後

					//上に花壇がない場合
					if preList[j] == 0 {

						//左隣が"."
						if kadanList[j-1] == "." {
							//   .
							// . #
							preList[j] = 1 //次の調査の為に花壇の存在を入力
							sum += 3
						} else {
							//   .
							// # #
							sum += 2
							preList[j] = 1
						}
					} else {
						// preListが1だった場合 上に花壇がある
						if kadanList[j-1] == "." {
							//  #
							//. #
							sum += 2
							preList[j] = 1
						} else {
							//   #
							// # #
							sum++
							preList[j] = 1
						}
					}
					//一番最後の場合かつ花壇ではない場合
				} else if j == yoko-1 && k == "." {
					//上に花壇がない場合
					if preList[j] == 0 {

						//左隣が"."
						if kadanList[j-1] == "." {
							//   .
							// . .
							preList[j] = 0 //次の調査の為に花壇の存在を入力
						} else {
							//   .
							// #|.
							sum++
							preList[j] = 0
						}
					} else {
						// preListが1だった場合 上に花壇がある
						if kadanList[j-1] == "." {
							//  #
							//. .
							sum++
							preList[j] = 0
						} else {
							//   #
							// # .
							sum += 2
							preList[j] = 0
						}
					}

					//中間かつ"#"
				} else if j != 0 && j != yoko-1 && k == "#" {
					//上に花壇がない場合
					if preList[j] == 0 {

						//左隣が"."
						if kadanList[j-1] == "." && kadanList[j+1] == "." {
							//   .
							// . # .
							preList[j] = 1 //次の調査の為に花壇の存在を入力
							sum += 2       //右側のロープは無視するので＋2
						} else if kadanList[j-1] == "." && kadanList[j+1] == "#" {
							//   .
							// . # #
							sum += 2
							preList[j] = 1
						} else if kadanList[j-1] == "#" && kadanList[j+1] == "." {
							//   .
							// # # .
							sum++ //右側のロープは無視するので＋1
							preList[j] = 1
						} else if kadanList[j-1] == "#" && kadanList[j+1] == "#" {
							//   .
							// # # #
							sum++
							preList[j] = 1
						}
					} else {
						// preListが1だった場合 上に花壇がある
						//左隣が"."
						if kadanList[j-1] == "." && kadanList[j+1] == "." {
							//   #
							// . # .
							preList[j] = 1 //次の調査の為に花壇の存在を入力
							sum++
						} else if kadanList[j-1] == "." && kadanList[j+1] == "#" {
							//   #
							// . # #
							sum++
							preList[j] = 1
						} else if kadanList[j-1] == "#" && kadanList[j+1] == "." {
							//   #
							// # # .
							//右側を無視するのでプラスしない
							preList[j] = 1
						} else if kadanList[j-1] == "#" && kadanList[j+1] == "#" {
							//   #
							// # # #
							preList[j] = 1
						}
					}
					//中間かつ"."
				} else if j != 0 && j != yoko-1 && k == "." {
					//上に花壇がない場合
					if preList[j] == 0 {

						//左隣が"."
						if kadanList[j-1] == "." && kadanList[j+1] == "." {
							//   .
							// . . .
							preList[j] = 0 //次の調査の為に花壇の存在を入力
						} else if kadanList[j-1] == "." && kadanList[j+1] == "#" {
							//   .
							// . . #
							preList[j] = 0
						} else if kadanList[j-1] == "#" && kadanList[j+1] == "." {
							//   .
							// # . .
							sum++ //左側
							preList[j] = 0
						} else if kadanList[j-1] == "#" && kadanList[j+1] == "#" {
							//   .
							// # . #
							sum++ //左側
							preList[j] = 0
						}
					} else {
						// preListが1だった場合 上に花壇がある
						//左隣が"."
						if kadanList[j-1] == "." && kadanList[j+1] == "." {
							//   #
							// . . .
							sum++
							preList[j] = 0 //次の調査の為に花壇の存在を入力
							sum++
						} else if kadanList[j-1] == "." && kadanList[j+1] == "#" {
							//   #
							// . . #
							sum++
							preList[j] = 0
						} else if kadanList[j-1] == "#" && kadanList[j+1] == "." {
							//   #
							// # . .
							sum += 2 //左と上
							preList[j] = 0
						} else if kadanList[j-1] == "#" && kadanList[j+1] == "#" {
							//   #
							// # . #
							sum += 2 //右無視
							preList[j] = 0
						}
					}
				}
			} else if yoko == 2 {
				if j == 0 && k == "#" {

					// 上に花壇がない場合
					if preList[j] == 0 {
						if kadanList[j+1] == "#" {
							// .
							// # #
							sum += 2
							preList[j] = 1
						} else {
							// .
							// # .
							sum += 2 //右側無視
							preList[j] = 1
						}
						//上に花壇がある場合
					} else {
						if kadanList[j+1] == "#" {
							// #
							// # #
							sum++
							preList[j] = 1
						} else {
							// #
							// # .
							sum++ //右側無視
							preList[j] = 1
						}
					}

				} else if j == 0 && k == "." {

					if preList[j] == 0 {
						if kadanList[j+1] == "#" {
							// #
							// . #
							//右側無視
							sum++ //上側のロープ
							preList[j] = 0
						} else {
							// #
							// . .
							sum++
							preList[j] = 0
						}
					}

				} else if j == 1 && k == "." {

					//上に花壇がない
					if preList[j] == 0 {
						if kadanList[j-1] == "#" {
							//   .
							// # .
							sum++
							preList[j] = 0
						} else {
							//   .
							// . .
							preList[j] = 0
						}

						//上に花壇がある
					} else {
						if kadanList[j-1] == "#" {
							//   #
							// # .
							sum += 2
							preList[j] = 0
						} else {
							//   #
							// . .
							sum++
							preList[j] = 0
						}
					}
				}
			} else if yoko == 1 {
				if k == "#" {
					//上に花壇がない
					if preList[j] == 0 {
						// .
						// #
						sum += 3
						preList[j] = 1
						//上に花壇がある
					} else {
						// #
						// #
						sum += 2

						preList[j] = 1
					}
				} else if k == "." {
					//上に花壇がない
					if preList[j] == 0 {
						// .
						// .
						preList[j] = 0
						//上に花壇がある
					} else {
						// #
						// .
						sum++
						preList[j] = 0
					}
				}
			}

		}
		result += sum

	}
	fmt.Println(result + returnInt(preList))
	// fmt.Println(preList)

}
