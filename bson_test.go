package bson

import (
	"fmt"
	"testing"
	"time"
)

func TestBSON(t *testing.T) {
	tm, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 12:11:10")
	id := NewObjectIdWithTime(tm.Add(time.Hour * -8)).Hex()

	fmt.Println(id)
	fmt.Println(ObjectIdHex(id).Time())
}
