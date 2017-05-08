package history

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const aList = `{"uuid" : "4970494a-110f-11e6-91da-096d89bd2173", "title" : "Editor's Picks", "concept" : null, "listType" : null, "items" : [ { "uuid" : "2ea7868a-15d5-11e6-9d98-00386a18e39d" }, { "uuid" : "b635c7d4-1340-11e6-91da-096d89bd2173" }, { "uuid" : "ac141226-11ae-11e6-91da-096d89bd2173" }, { "uuid" : "d648f2e2-110d-11e6-839f-2922947098f0" }, { "uuid" : "a7bf81e6-110a-11e6-839f-2922947098f0" }, { "uuid" : "8552bf14-0c74-11e6-9456-444ab5211a2f" }, { "uuid" : "87f44872-1080-11e6-91da-096d89bd2173" } ], "layoutHint" : "standard", "publishReference" : "tid_AseH8ewWZ2", "lastModified" : "2017-04-09T11:05:38.587Z" }`

func TestWriteDB(t *testing.T) {
	var aListMap List
	err := json.Unmarshal([]byte(aList), &aListMap)
	assert.NoError(t, err)
	w := newMongoDBWriter("192.168.99.100:32768", "upp-store", "list-history")
	err = w.Write("3f1392e3-79e6-46c5-9d83-59c771b06b33", aListMap, "tid_AseH8ewWZ2")
	assert.NoError(t, err)
}
