[
# // // package main

# // // import "fmt"

# // // func main() {
# // //   b := NewBoard()

# // //   for !b.IsFinished() {
# // //     player := "p1"
# // //     if !b.turn {
# // //       player = "p2"
# // //     }

# // //     fmt.Printf("turn of %s\n", player)
# // //     fmt.Printf("p1 pits: %v total: %d\n", b.pits[:6], b.pits[6])
# // //     fmt.Printf("p2 pits: %v total: %d\n", b.pits[7:13], b.pits[13])

# // //     fmt.Println("please select pit")

# // //     var pit int
# // //     _, _ = fmt.Scanln(&pit)

# // //     if err := b.Turn(pit); err != nil {
# // //       fmt.Println(err)
# // //       continue
# // //     }
# // //   }
  
# // //   b.Totalize()

# // //   if b.pits[6] == b.pits[13] {
# // //     fmt.Println("standoff")
# // //   }

# // //   if b.pits[6] > b.pits[13] {
# // //     fmt.Println("p1 is winner")
# // //   }

# // //   fmt.Println("p2 is winner")
# // // }

# // // type Board struct {
# // //   pits [14]int
# // //   turn bool
# // // }

# // // func NewBoard() Board {
# // //   return Board{
# // //     pits: [14]int{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0},
# // //     turn: true,
# // //   }
# // // }

# // // func (b *Board) Turn(pit int) error {
# // //   if pit > 5 {
# // //     return fmt.Errorf("bad pit")
# // //   }

# // //   idx := pit
# // //   if !b.turn {
# // //     idx = pit + 7
# // //   }

# // //   count := b.pits[idx]
# // //   if count == 0 {
# // //     return fmt.Errorf("pit %d is empty", pit)
# // //   }

# // //   b.pits[idx] = 0

# // //   idx++

# // //   for count > 0 {
# // //     if idx == 14 {
# // //       idx = 0
# // //     }

# // //     if idx == 6 && !b.turn {
# // //       idx++
# // //       continue
# // //     }

# // //     if idx == 13 && b.turn {
# // //       idx = 0
# // //       continue
# // //     }

# // //     if count == 1 && b.pits[idx] == 0 && idx < 6 && b.turn {
# // //       b.pits[6] += 1 + b.pits[idx+7]
# // //       b.pits[idx+7] = 0
# // //       break
# // //     }

# // //     if count == 1 && b.pits[idx] == 0 && idx > 6 && idx != 13 && !b.turn {
# // //       b.pits[13] += 1 + b.pits[idx-7]
# // //       b.pits[idx-7] = 0
# // //       break
# // //     }

# // //     if count == 1 && idx == 6 && b.turn {
# // //       b.pits[6]++
# // //       return nil
# // //     }

# // //     if count == 1 && idx == 13 && !b.turn {
# // //       b.pits[13]++
# // //       return nil
# // //     }

# // //     b.pits[idx]++

# // //     idx++
# // //     count--
# // //   }

# // //   b.turn = !b.turn

# // //   return nil
# // // }

# // // func (b *Board) IsFinished() bool {
# // //   finished := true
# // //   for i := 0; i < 6; i++ {
# // //     if b.pits[i] > 0 {
# // //       finished = false
# // //     }
# // //   }

# // //   if finished {
# // //     return true
# // //   }

# // //   finished = true
# // //   for i := 7; i < 13; i++ {
# // //     if b.pits[i] > 0 {
# // //       finished = false
# // //     }
# // //   }

# // //   return finished
# // // }

# // // func (b *Board) Totalize() {
# // //   for i := 0; i < 13; i++ {
# // //     if i < 6 {
# // //       b.pits[6] += b.pits[i]
# // //       b.pits[i] = 0
# // //     }
    
# // //     if i > 6 {
# // //       b.pits[13] += b.pits[i]
# // //       b.pits[i] = 0
# // //     }
# // //   }
# // // }](url)
