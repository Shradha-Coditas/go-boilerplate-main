package tags

import (
	RedisWrapper "boilerplate/library/redis"
	"boilerplate/util/logwrapper"
	"encoding/json"

	"github.com/fatih/structs"
)

//LastTick - LastTick post request structure.
type LastTick struct {
	Id          string  `form:"id" json:"id" `                  //id of LastTick
	Price       int64   `form:"price" json:"price"`             //Price of LastTick
	Open        float64 `form:"open" json:"open"`               //Open Price of LastTick
	High        float64 `form:"high" json:"high" `              //High Price of LastTick
	Low         float64 `form:"low" json:"low"`                 //Low Price of LastTick
	Close       float64 `form:"close" json:"close"`             //Low Price of LastTick
	TickDate    int64   `form:"tickDate" json:"tickDate"`       //Total Volume of LasTick
	TotalVolume float64 `form:"totalVolume" json:"totalVolume"` //Total Volume of LasTick
}

//LastTickResponse - LastTick Post Response structure.
type LastTickResponse struct {
	Message string `form:"message" json:"string"` //Total Volume of LasTick
}

func (i LastTick) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
}

func AddLastTick(data LastTick) error {
	err := RedisWrapper.Client.Conn.HMSet("LastTick:"+data.Id, structs.Map(&data)).Err()
	logwrapper.Logger.Debugln(err)
	if err != nil {
		return err
	}
	return nil
}
