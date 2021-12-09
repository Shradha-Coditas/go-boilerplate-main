package tagredis

import (
	RedisWrapper "boilerplate/library/redis"
	"boilerplate/util/logwrapper"
	"encoding/json"
)

type TagSegment []struct {
	Tag     string `json:"tag"`
	Segment string `json:"segment"`
	Symbols []struct {
		TradeSymbol string `json:"tradeSymbol"`
		Value       int    `json:"value"`
	} `json:"symbols"`
}

//TagSegmentResponse - TagSegmentPost Response structure.
type TagSegementResponse struct {
	Message string `form:"message" json:"string"`
}

func (i TagSegment) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
}

func AddTagSegment(data TagSegment) error {
	for _, v := range data {
		i, err1 := json.Marshal(v.Symbols)
		if err1 != nil {
			return err1
		}
		err := RedisWrapper.Client.Conn.Set(RedisKey(v.Tag, v.Segment), i, 0).Err()
		logwrapper.Logger.Debugln(err)
		if err != nil {
			return err
		}
	}
	return nil
}

func RedisKey(tagName string, segmentName string) string {
	return tagName + "." + segmentName
}
