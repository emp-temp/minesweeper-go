package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/mattn/go-tty"
)

const FIELD_HEIGHT = 4
const FIELD_WIDTH = 20
const MINE = -1;

func main() {

  rand.Seed(time.Now().UnixNano())

  var field [FIELD_WIDTH*FIELD_HEIGHT]int 
  var show_field [FIELD_WIDTH*FIELD_HEIGHT]string
  mine_num := 12
  set_mine := 0
  player_position := 0
  for i:=0; i<FIELD_HEIGHT*FIELD_HEIGHT; i++ {
    field[i] = 0
  }
  for i:=0; i<FIELD_HEIGHT*FIELD_WIDTH; i++ {
    show_field[i] = "■"
  }
  for ;set_mine!=mine_num; {
    j := rand.Intn(FIELD_HEIGHT)
    i := rand.Intn(FIELD_WIDTH)

    position := j*FIELD_WIDTH+i

    if field[position] != MINE {
      field[position] = MINE
      set_mine++
    }
  }

  for i:=0; i<4; i++ {
    for j:=0; j<20; j++ {
      fmt.Printf("%2d, ", field[j+i*20])
    }
    println()
  }

  fmt.Scan()

  for i:=0; i<FIELD_HEIGHT*FIELD_WIDTH; i++ {
    num := 0

    // fmt.Println(i)
    // fmt.Println(i%FIELD_WIDTH)
    // fmt.Println(i/FIELD_WIDTH)

    if field[i] == MINE {
      continue
    }

    if i%FIELD_WIDTH == 0 && i/FIELD_WIDTH == 0 {
      num = 0
      if field[i+1] == MINE {
        num++
      }
      if field[i+FIELD_WIDTH] == MINE {
        num++
      }
      field[i] = num
      continue
    }

    if i%FIELD_WIDTH == 0 && i/FIELD_WIDTH == FIELD_HEIGHT-1 {
      num = 0
      if field[i+1] == MINE {
        num++
      }
      if field[i-FIELD_WIDTH] == MINE {
        num++
      }
      field[i] = num
      continue
    }

    if i%FIELD_WIDTH == FIELD_WIDTH-1 && i/FIELD_WIDTH == 0 {
      num = 0
      if field[i-1] == MINE {
        num++
      }
      if field[i+FIELD_WIDTH] == MINE {
        num++
      }
      field[i] = num
      continue
    }
    if i%FIELD_WIDTH == FIELD_WIDTH-1 && i/FIELD_WIDTH == FIELD_HEIGHT-1 {
      num = 0
      if field[i-1] == MINE {
        num++
      }
      if field[i-FIELD_WIDTH] == MINE {
        num++
      }
      field[i] = num
      continue
    }

    if i%FIELD_WIDTH == 0 {
      num = 0
      if field[i+FIELD_WIDTH] == MINE {
        num++
      }
      if field[i+1] == MINE {
        num++
      }
      if field[i-FIELD_WIDTH] == MINE {
        num++
      }

      field[i] = num
      continue
    }

    if i/FIELD_WIDTH == 0 {
      num = 0
      if field[i-1] == MINE {
        num++
      }
      if field[i+1] == MINE {
        num++
      }
      if field[i+FIELD_WIDTH] == MINE {
        num++
      }

      field[i] = num
      continue
    }

    if i%FIELD_WIDTH == FIELD_WIDTH-1 {
      num = 0
      if field[i-FIELD_WIDTH] == MINE {
        num++
      }
      if field[i-1] == MINE {
        num++
      }
      if field[i+FIELD_WIDTH] == MINE {
        num++
      }

      field[i] = num
      continue
    }

    if i/FIELD_WIDTH == FIELD_HEIGHT-1 {
      num = 0
      if field[i-1] == MINE {
        num++
      }
      if field[i+1] == MINE {
        num++
      }
      if field[i-FIELD_WIDTH] == MINE {
        num++
      }

      field[i] = num
      continue
    }

    num = 0
    if field[i-1] == MINE {
      num++
    }
    if field[i+1] == MINE {
      num++
    }
    if field[i-FIELD_WIDTH] == MINE {
      num++
    }
    if field[i+FIELD_WIDTH] == MINE {
      num++
    }
    // fmt.Println(num)
    field[i] = num
  }

  tty, err := tty.Open()
  if err != nil {
    log.Fatal(err)
  }

  defer tty.Close()

  for {
    r, err := tty.ReadRune()
    if err != nil {
      log.Fatal(err)
    }
    // key := string(r)
    key := int(r)
    fmt.Println("Key press => ", key)
    
    if key == 104 && player_position%FIELD_WIDTH != 0 {
      player_position--
    }  
    if key == 108 && player_position%FIELD_WIDTH != FIELD_WIDTH-1 {
      player_position++
    }
    if key == 107 && player_position/FIELD_WIDTH != 0 {
      player_position -= FIELD_WIDTH
    }
    if key == 106 && player_position/FIELD_WIDTH != FIELD_HEIGHT-1 {
      player_position += FIELD_WIDTH
    }

    

    
    fmt.Println(player_position)

    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()

  
    for i:=0; i<80; i++ {
      if player_position == i {
        fmt.Printf("%2s, ", "s")
      } else {
        fmt.Printf("%2s, ", show_field[i])
      }
      if i%20 == 19 {
        fmt.Println()
      }
    }

    if key == 13 {
      if field[player_position] == MINE {
        fmt.Println("game over")
        os.Exit(1)
      } else if field[player_position] == 0 {
        fmt.Println("nonn")           
      } else {
        show_field[player_position] = fmt.Sprintf("%d", field[player_position])
      }
    }

  }
}
