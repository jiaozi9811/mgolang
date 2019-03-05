package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)
const(
	rock,scissors,paper int=1,2,3

	rockName,scissorsName,paperName string="石头","剪刀","布"

	win,lost,draw int=1,-1,0
)
type Finger struct{
	value int
	name string
}

func main(){
	var aiFinger,userFinger Finger
	reader:=bufio.NewReader(os.Stdin)

	fmt.Printf("输入要出的拳: %d->石头,%d->剪刀,%d->布,9->退出\n",rock,scissors,paper)
	for{
		data,_,err:=reader.ReadLine()
		if err!=nil{
			fmt.Println("No input")
			break
		}

		input,err:=strconv.Atoi(string(data))
		if err!=nil{
			fmt.Println("please input data")
			continue
		}
		if input==9{
			break
		}

		switch input{
		case rock,scissors,paper:
			aiFinger=randFinger()
			fmt.Println("compute tach:",aiFinger.name)

			userFinger=createFinger(input)
			fmt.Println("you tech",userFinger.name)

			aiWin:=isAiWin(aiFinger,userFinger)
			if aiWin==win{
				fmt.Println("you are lost")
			}else if aiWin==lost{
				fmt.Println("you are win")
			}else {
				fmt.Println("equel")
			}
		default:
			fmt.Println("input error")
		}
	}
}

func randFinger()(finger Finger){
	rand:=rand.New(rand.NewSource(time.Now().UnixNano()))
	switch rand.Intn(3){
	case 0:
		finger.value=rock
		finger.name=rockName
	case 1:
		finger.value=scissors
		finger.name=scissorsName
	case 2:
		finger.value=paper
		finger.name=paperName
	}
	return
}

func createFinger(n int) (finger Finger){
	switch n{
	case rock:
		finger.value=rock
		finger.name=rockName
	case scissors:
		finger.value=scissors
		finger.name=scissorsName
	case paper:
		finger.value=paper
		finger.name=paperName
	}
	return
}

func isAiWin(ai Finger,user Finger) int{
	result:=ai.value-user.value
	if int(math.Abs(float64(result)))==paper-rock{
		result=-(result)
	}
	if result<0{
		return win
	}else if result>0{
		return lost
	}
	return draw
}

/*
package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Finger struct {
	ID    int
	name  string
	value int
}
type Player struct {
	name   string
	value  int
	finger Finger
}
type Compute Player

func main() {
	var compute Compute
	var human Player
	rock, scissors, paper := Finger{1, "石头", -1}, Finger{2, "剪刀", 1}, Finger{3, "布", 0}

	read := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("输入你出的拳:%d->石头,%d->剪刀,%d->布,9->Quit\n", rock.ID, scissors.ID, paper.ID)
		data, _, err := read.ReadLine()
		if err != nil {
			fmt.Println("input error")
			continue
		}

		input, err := strconv.Atoi(string(data))
		if err != nil {
			fmt.Println("input error")
			continue
		}

		if input == 9 {
			break
		}

		switch input {
		case rock.ID:
			human.getfinger(input, rock, "Human")
		case scissors.ID:
			human.getfinger(input, paper, "Human")
		case paper.ID:
			human.getfinger(input, paper, "Human")
		default:
			continue
		}
		compute.getfinger("Compute", rock, scissors, paper)
		diff(Player(compute), human, paper, rock)
	}
}

func diff(compute, human Player, paper, rock Finger) {
	resu := compute.finger.value - human.finger.value
	if int(math.Abs(float64(resu))) == paper.value-rock.value {
		resu = -(resu)
	}
	fmt.Println("compute:", compute.finger.name, ";", "human:", human.finger.name)
	if resu < 0 {
		fmt.Println("you are lost")
	} else if resu > 0 {
		fmt.Println("you are win")
	} else {
		fmt.Println("equel")
	}
}

func (h *Player) getfinger(n int, finger Finger, s string) *Player {
	h.name = "Human"
	h.value = n
	h.finger = finger
	return h
}
func (c *Compute) getfinger(s string, rock, scissors, paper Finger) *Compute {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c.name = s
	c.value = r.Intn(3)
	switch c.value {
	case rock.ID:
		c.finger = rock
	case scissors.ID:
		c.finger = scissors
	case paper.ID:
		c.finger = paper
	}
	return c
}
*/