package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main(){
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	seconds, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	tate, _ := strconv.Atoi(list[0])
	yoko, _ := strconv.Atoi(list[1])

	DummyDurtyList := make([][]int, 0)

	for i:=0; i<tate; i++ {
		durty := make([]int, yoko, yoko)
		DummyDurtyList = append(DummyDurtyList, durty)
	}

	durtyList := make([][]string, 0)

	for i:=0; i<tate; i++ {
		sc.Scan()
		slice := strings.Split(sc.Text(), "")

		durtyList = append(durtyList, slice)
	}

	nowPlace := []int{0, 0} // 今の場所 [nowPlace[0]][nowPlace[1]]
	nowDirection := "Right"
	cleanSum := 0 // 掃除した数

	for i:=seconds; i>0; i-- {
		// fmt.Println(nowPlace)
		// fmt.Println(durtyList[nowPlace[0]][nowPlace[1]], nowDirection)

		if durtyList[nowPlace[0]][nowPlace[1]] == "." {
			DummyDurtyList[nowPlace[0]][nowPlace[1]] = 1 // 通ったところを1にする

			if nowDirection == "Right" {

				if nowPlace[1] == yoko-1 {
					//突き当たりか掃除済みか 以下の条件分岐で上か下かなどを決める

					if nowPlace[0] == 0 {
						// {0, } 上は突き当たりなので下に行くしかない
						nowDirection = "Down"
						nowPlace[0]++ //下に行くので+1

					} else if nowPlace[0] == yoko-1 {
						// {3,0} 下が突き当たり
						nowDirection = "Up"
						nowPlace[0]--
					} else {
						// nowPlace[0]が1,2,3
						// 上のマスが掃除済みなら"Down"
						if DummyDurtyList[nowPlace[0]-1][nowPlace[1]] == 1 {
							nowDirection = "Down"
							nowPlace[0]++
						} else if DummyDurtyList[nowPlace[0]+1][nowPlace[1]] == 1 {
							// 下のマスが掃除済みなら"Up"
							nowDirection = "Up"
							nowPlace[0]--
						}
					}
				} else if DummyDurtyList[nowPlace[0]][nowPlace[1]+1] == 1 {
					//突き当たりか掃除済みか 以下の条件分岐で上か下かなどを決める

					if nowPlace[0] == 0 {
						// {0, } 上は突き当たりなので下に行くしかない
						nowDirection = "Down"
						nowPlace[0]++ //下に行くので+1

					} else if nowPlace[0] == yoko-1 {
						// {3,0} 下が突き当たり
						nowDirection = "Up"
						nowPlace[0]--
					} else {
						// nowPlace[0]が1,2,3
						// 上のマスが掃除済みなら"Down"
						if DummyDurtyList[nowPlace[0]-1][nowPlace[1]] == 1 {
							nowDirection = "Down"
							nowPlace[0]++
						} else if DummyDurtyList[nowPlace[0]+1][nowPlace[1]] == 1 {
							// 下のマスが掃除済みなら"Up"
							nowDirection = "Up"
							nowPlace[0]--
						}
					}
				} else {
					nowPlace[1]++ //右に進む
				}

			} else if nowDirection == "Left" {


				if nowPlace[1] == 0 {
					//掃除済みか
					if nowPlace[0] == 0 {
						// {0, } 上は突き当たりなので下に行くしかない
						nowDirection = "Down"
						nowPlace[0]++ //下に行くので+1

					} else if nowPlace[0] == yoko-1 {
						// {3,0} 下が突き当たり上に行くしかない
						nowDirection = "Up"
						nowPlace[0]--
					} else {
						// nowPlace[0]が1,2,3
						// 上のマスが掃除済みなら"Down"
						if DummyDurtyList[nowPlace[0]-1][nowPlace[1]] == 1 {
							nowDirection = "Down"
							nowPlace[0]++
						} else {
							nowDirection = "Up"
							nowPlace[0]--
						}
						// } else if DummyDurtyList[nowPlace[0]+1][nowPlace[1]] == 1 {
						//     // 下のマスが掃除済みなら"Up"
						//     nowDirection = "Up"
						//     nowPlace[0]--
						// }
					}
				} else if DummyDurtyList[nowPlace[0]][nowPlace[1]-1] == 1 {
					//掃除済みか
					if nowPlace[0] == 0 {
						// {0, } 上は突き当たりなので下に行くしかない
						nowDirection = "Down"
						nowPlace[0]++ //下に行くので+1

					} else if nowPlace[0] == yoko-1 {
						// {3,0} 下が突き当たり上に行くしかない
						nowDirection = "Up"
						nowPlace[0]--
					} else {
						// nowPlace[0]が1,2,3
						// 上のマスが掃除済みなら"Down"
						if DummyDurtyList[nowPlace[0]-1][nowPlace[1]] == 1 {
							nowDirection = "Down"
							nowPlace[0]++
						} else if DummyDurtyList[nowPlace[0]+1][nowPlace[1]] == 1 {
							// 下のマスが掃除済みなら"Up"
							nowDirection = "Up"
							nowPlace[0]--
						}
					}
				} else {
					nowPlace[1]-- //左に進む
				}

			} else if nowDirection == "Up" {

				if nowPlace[0] == 0 {
					//上に突き当たりか 次の上のマスが1であれば DummyDurtyList[nowPlace[0]-1]
					// RightかLeftかを決める
					if nowPlace[1] == 0 {
						// 左側が壁なら右に方向転換
						nowDirection = "Right"
						nowPlace[1]++
					} else if nowPlace[1] == yoko-1 {
						// 右側が壁なら左に行くしかない
						nowDirection = "Left"
						nowPlace[1]--
					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]+1] == 1 {
						// 右側が掃除済み
						nowDirection = "Left"
						nowPlace[1]--

					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]-1] == 1 {
						//左側が掃除済み
						nowDirection = "Right"
						nowPlace[1]++
					}
				} else if DummyDurtyList[nowPlace[0]-1][nowPlace[1]] == 1 {
					if nowPlace[1] == 0 {
						// 左側が壁なら右に方向転換
						nowDirection = "Right"
						nowPlace[1]++
					} else if nowPlace[1] == yoko-1 {
						// 右側が壁なら左に行くしかない
						nowDirection = "Left"
						nowPlace[1]--
					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]+1] == 1 {
						// 右側が掃除済み
						nowDirection = "Left"
						nowPlace[1]--

					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]-1] == 1 {
						//左側が掃除済み
						nowDirection = "Right"
						nowPlace[1]++
					}
				} else {
					nowPlace[0]-- // 上に進む
				}

			} else if nowDirection == "Down" {

				if nowPlace[0] == tate-1 {
					//下に突き当たりか 進行方向が下なのでDummyDurtyList[nowPlace[0]+1] 下の次のますが1
					if nowPlace[1] == 0 {
						// 左側が壁なら右に方向転換
						nowDirection = "Right"
						nowPlace[1]++
					} else if nowPlace[1] == yoko-1 {
						// 右側が壁なら左に行くしかない
						nowDirection = "Left"
						nowPlace[1]--
					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]+1] == 1 {
						// 右側が掃除済み
						nowDirection = "Left"
						nowPlace[1]--

					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]-1] == 1 {
						//左側が掃除済み
						nowDirection = "Right"
						nowPlace[1]++
					}

				} else if DummyDurtyList[nowPlace[0]+1][nowPlace[1]] == 1 {
					//下に突き当たりか 進行方向が下なのでDummyDurtyList[nowPlace[0]+1] 下の次のますが1
					if nowPlace[1] == 0 {
						// 左側が壁なら右に方向転換
						nowDirection = "Right"
						nowPlace[1]++
					} else if nowPlace[1] == yoko-1 {
						// 右側が壁なら左に行くしかない
						nowDirection = "Left"
						nowPlace[1]--
					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]+1] == 1 {
						// 右側が掃除済み
						nowDirection = "Left"
						nowPlace[1]--

					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]-1] == 1 {
						//左側が掃除済み
						nowDirection = "Right"
						nowPlace[1]++
					}
				} else {
					nowPlace[0]++ // 下に進む
				}
			}

		} else if durtyList[nowPlace[0]][nowPlace[1]]=="#"{
			cleanSum++ //掃除カウントを+1
			DummyDurtyList[nowPlace[0]][nowPlace[1]] = 1 // 1は掃除済みとして扱う
			if nowDirection == "Right" {

				if nowPlace[1] == yoko-1 {
					//突き当たりか掃除済みか 以下の条件分岐で上か下かなどを決める

					if nowPlace[0] == 0 {
						// {0, } 上は突き当たりなので下に行くしかない
						nowDirection = "Down"
						nowPlace[0]++ //下に行くので+1

					} else if nowPlace[0] == yoko-1 {
						// {3,0} 下が突き当たり
						nowDirection = "Up"
						nowPlace[0]--
					} else {
						// nowPlace[0]が1,2,3
						// 上のマスが掃除済みなら"Down"
						if DummyDurtyList[nowPlace[0]-1][nowPlace[1]] == 1 {
							nowDirection = "Down"
							nowPlace[0]++
						} else if DummyDurtyList[nowPlace[0]+1][nowPlace[1]] == 1 {
							// 下のマスが掃除済みなら"Up"
							nowDirection = "Up"
							nowPlace[0]--
						}
					}
				} else if DummyDurtyList[nowPlace[0]][nowPlace[1]+1] == 1 {
					//突き当たりか掃除済みか 以下の条件分岐で上か下かなどを決める

					if nowPlace[0] == 0 {
						// {0, } 上は突き当たりなので下に行くしかない
						nowDirection = "Down"
						nowPlace[0]++ //下に行くので+1

					} else if nowPlace[0] == yoko-1 {
						// {3,0} 下が突き当たり
						nowDirection = "Up"
						nowPlace[0]--
					} else {
						// nowPlace[0]が1,2,3
						// 上のマスが掃除済みなら"Down"
						if DummyDurtyList[nowPlace[0]-1][nowPlace[1]] == 1 {
							nowDirection = "Down"
							nowPlace[0]++
						} else if DummyDurtyList[nowPlace[0]+1][nowPlace[1]] == 1 {
							// 下のマスが掃除済みなら"Up"
							nowDirection = "Up"
							nowPlace[0]--
						}
					}
				} else {
					nowPlace[1]++ //右に進む
				}
			} else if nowDirection == "Left" {

				if nowPlace[1] == 0 {
					//掃除済みか
					if nowPlace[0] == 0 {
						// {0, } 上は突き当たりなので下に行くしかない
						nowDirection = "Down"
						nowPlace[0]++ //下に行くので+1

					} else if nowPlace[0] == yoko-1 {
						// {3,0} 下が突き当たり上に行くしかない
						nowDirection = "Up"
						nowPlace[0]--
					} else {
						// nowPlace[0]が1,2,3
						// 上のマスが掃除済みなら"Down"
						if DummyDurtyList[nowPlace[0]-1][nowPlace[1]] == 1 {
							nowDirection = "Down"
							nowPlace[0]++
						} else {
							nowDirection = "Up"
							nowPlace[0]--
						}
						// else if DummyDurtyList[nowPlace[0]+1][nowPlace[1]] == 1 {
						//     // 下のマスが掃除済みなら"Up"
						//     nowDirection = "Up"
						//     nowPlace[0]--
						// }
					}
				} else if DummyDurtyList[nowPlace[0]][nowPlace[1]-1] == 1 {
					//掃除済みか
					if nowPlace[0] == 0 {
						// {0, } 上は突き当たりなので下に行くしかない
						nowDirection = "Down"
						nowPlace[0]++ //下に行くので+1

					} else if nowPlace[0] == yoko-1 {
						// {3,0} 下が突き当たり上に行くしかない
						nowDirection = "Up"
						nowPlace[0]--
					} else {
						// nowPlace[0]が1,2,3
						// 上のマスが掃除済みなら"Down"
						if DummyDurtyList[nowPlace[0]-1][nowPlace[1]] == 1 {
							nowDirection = "Down"
							nowPlace[0]++
						} else if DummyDurtyList[nowPlace[0]+1][nowPlace[1]] == 1 {
							// 下のマスが掃除済みなら"Up"
							nowDirection = "Up"
							nowPlace[0]--
						}
					}
				} else {
					nowPlace[1]-- //左に進む
				}

			} else if nowDirection == "Up" {

				if nowPlace[0] == 0 {
					//上に突き当たりか 次の上のマスが1であれば DummyDurtyList[nowPlace[0]-1]
					// RightかLeftかを決める
					if nowPlace[1] == 0 {
						// 左側が壁なら右に方向転換
						nowDirection = "Right"
						nowPlace[1]++
					} else if nowPlace[1] == yoko-1 {
						// 右側が壁なら左に行くしかない
						nowDirection = "Left"
						nowPlace[1]--
					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]+1] == 1 {
						// 右側が掃除済み
						nowDirection = "Left"
						nowPlace[1]--

					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]-1] == 1 {
						//左側が掃除済み
						nowDirection = "Right"
						nowPlace[1]++
					}
				} else if DummyDurtyList[nowPlace[0]-1][nowPlace[1]] == 1 {
					//上に突き当たりか 次の上のマスが1であれば DummyDurtyList[nowPlace[0]-1]
					// RightかLeftかを決める
					if nowPlace[1] == 0 {
						// 左側が壁なら右に方向転換
						nowDirection = "Right"
						nowPlace[1]++
					} else if nowPlace[1] == yoko-1 {
						// 右側が壁なら左に行くしかない
						nowDirection = "Left"
						nowPlace[1]--
					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]+1] == 1 {
						// 右側が掃除済み
						nowDirection = "Left"
						nowPlace[1]--

					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]-1] == 1 {
						//左側が掃除済み
						nowDirection = "Right"
						nowPlace[1]++
					}
				} else {
					nowPlace[0]-- // 上に進む
				}
			} else if nowDirection == "Down" {

				if nowPlace[0] == tate-1 || DummyDurtyList[nowPlace[0]+1][nowPlace[1]] == 1  {
					//下に突き当たりか 進行方向が下なのでDummyDurtyList[nowPlace[0]+1] 下の次のますが1
					if nowPlace[1] == 0 {
						// 左側が壁なら右に方向転換
						nowDirection = "Right"
						nowPlace[1]++
					} else if nowPlace[1] == yoko-1 {
						// 右側が壁なら左に行くしかない
						nowDirection = "Left"
						nowPlace[1]--
					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]+1] == 1 {
						// 右側が掃除済み
						nowDirection = "Left"
						nowPlace[1]--

					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]-1] == 1 {
						//左側が掃除済み
						nowDirection = "Right"
						nowPlace[1]++
					}

				} else if DummyDurtyList[nowPlace[0]+1][nowPlace[1]] == 1 {
					if nowPlace[1] == 0 {
						// 左側が壁なら右に方向転換
						nowDirection = "Right"
						nowPlace[1]++
					} else if nowPlace[1] == yoko-1 {
						// 右側が壁なら左に行くしかない
						nowDirection = "Left"
						nowPlace[1]--
					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]+1] == 1 {
						// 右側が掃除済み
						nowDirection = "Left"
						nowPlace[1]--

					} else if DummyDurtyList[nowPlace[0]][nowPlace[1]-1] == 1 {
						//左側が掃除済み
						nowDirection = "Right"
						nowPlace[1]++
					} else {
						nowPlace[0]++ // 下に進む
					}
				}
			}
		}
	}

	fmt.Println(cleanSum)
}