package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type TradeData struct {
	EventSymbol   string  `json:"EventSymbol"`
	Time          string  `json:"Time"`
	Sequence      int     `json:"Sequence"`
	Count         int     `json:"Count"`
	Open          float64 `json:"Open"`
	High          float64 `json:"High"`
	Low           float64 `json:"Low"`
	Close         float64 `json:"Close"`
	Volume        int     `json:"Volume"`
	VWAP          float64 `json:"VWAP"`
	BidVolume     int     `json:"BidVolume"`
	AskVolume     int     `json:"AskVolume"`
	ImpVolatility string  `json:"ImpVolatility"`
}

var tradeData = TradeData{
	EventSymbol:   "AAPL",
	Time:          "20190904-170000+0300",
	Sequence:      0,
	Count:         298,
	Open:          207.73,
	High:          207.89,
	Low:           207.61,
	Close:         207.85,
	Volume:        90897,
	VWAP:          207.79158,
	BidVolume:     45583,
	AskVolume:     45314,
	ImpVolatility: "NaN",
}

func main() {
	app := fiber.New()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// preparing a sample
		tradeData.BidVolume = (rand.Intn(100000))
		tradeData.AskVolume = (rand.Intn(100000))
		tradeData.Volume = (rand.Intn(100000))
		tradeData.Low = (rand.Float64())
		tradeData.High = (rand.Float64())

		data, _ := json.Marshal(tradeData)
		for {
			if err := c.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println("write:", err)
				break
			}

			time.Sleep(time.Second)
		}

	}))

	log.Fatal(app.Listen(":3000"))

}
