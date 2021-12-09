package scrip

import (
	RedisWrapper "boilerplate/library/redis"
	"boilerplate/util/logwrapper"
	"encoding/json"

	"github.com/fatih/structs"
)

//ScripCompareables - ScripCompareables post request structure.
type ScripComparables struct {
	Id            string `form:"id" json:"id" `                      //id of scrip/stock (Primary Key)
	FiftyTwoWH    int64  `form:"FiftyTwoWH" json:"FiftyTwoWH"`       //Fifty Two Week high of Stock
	FiftyTwoWL    int64  `form:"FiftyTwoWL" json:"FiftyTwoWL"`       //Fifty Two Week Low of Stock
	ATH           int64  `form:"ATH" json:"ATH"`                     //All Time high of Stock
	ATL           int64  `form:"ATL" json:"ATL"`                     //All Time Low of Stock
	PreviousClose int64  `form:"PreviousClose" json:"PreviousClose"` //Previous closing price of Stock
}

//ScripCompareablesResponse - ScripCompareables Post Response structure.
type ScripComparablesResponse struct {
	Message string `form:"message" json:"string"`
}

func (i ScripComparables) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
}

func AddScripComparables(data ScripComparables) error {
	err := RedisWrapper.Client.Conn.HMSet("ScripComparables:"+data.Id, structs.Map(&data)).Err()
	logwrapper.Logger.Debugln(err)
	if err != nil {
		return err
	}
	return nil
}
