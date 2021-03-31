package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/olahol/melody"
)

const IndexFile = "./public/index.html"

type die struct {
	sides int64
	name  string
}

type roll struct {
	total   int64
	rolls   []int64
	dieType string
}

type inputRoll struct {
	dieType  string `json:"dieType" validate:"required"`
	numRolls uint64 `json:"numRolls" validate:"gte=1,required"`
}

func NewRoll() *roll {
	newR := &roll{0, make([]int64, 0), ""}
	return newR
}

var dice = map[string]*die{
	"d4":   {4, "d4"},
	"d6":   {4, "d4"},
	"d8":   {4, "d4"},
	"d10":  {4, "d4"},
	"d12":  {4, "d4"},
	"d20":  {4, "d4"},
	"d100": {4, "d4"},
}

func rollDice(message []byte) *roll {
	fmt.Println("Rolling dice", string(message))
	var dat inputRoll
	if err := json.Unmarshal(message, &dat); err != nil {
		log.Println(err)
		return NewRoll()
	}
	log.Println(dat)
	// if !jsonIsValid(dat) {
	// 	return NewRoll()
	// }
	// r := NewRoll()
	// numSides := dice[dat.dieType].sides
	// var total int64
	// for i := uint64(0); i < dat.numRolls; i++ {
	// 	result := rand.Int63n(numSides) + 1
	// 	total += result
	// 	r.rolls = append(r.rolls, result)
	// }
	// r.total = total
	// r.dieType = dat.dieType

	// log.Println(r)
	// return r
	return NewRoll()
}

func jsonIsValid(dat inputRoll) bool {
	v := validator.New()
	err := v.Struct(dat)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func main() {
	r := gin.Default()
	m := melody.New()
	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		roll := rollDice(msg)
		log.Println(roll)
		m.Broadcast(msg)
	})

	r.Run("127.0.0.1:8080")
}
