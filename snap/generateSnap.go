package snap

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)



func GenerateFileTree() *Tree {
	root := &Tree{}
	root.generateRoot(nil)

	filePtr, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}

	var jsonMap map[string]interface{}

	err = json.Unmarshal([]byte(filePtr), &jsonMap)
	if err != nil {
		panic(err)
	}
	generateFileTreeRecursion(root, jsonMap["root"].(map[string]interface{}))
	return root
}

func generateFileTreeRecursion(root *Tree, jsonMap map[string]interface{}) {
	if msg, ok := jsonMap["msg"]; ok {
		msgMap := msg.(map[string]interface{})
		updateTime, err := time.ParseInLocation("2006-01-02 15:04:05", msgMap["updatetime"].(string), time.Local)
		createTime, err := time.ParseInLocation("2006-01-02 15:04:05", msgMap["createtime"].(string), time.Local)
		if err != nil {
			panic("can not parse in location with update time and create time")
		}
		var size int64
		if msgMap["size"] != nil {
			str := msgMap["size"].(string)
			size, err = strconv.ParseInt(str, 10, 64)
		}

		var times int64
		if msgMap["times"] != nil {
			str := msgMap["times"].(string)
			size, err = strconv.ParseInt(str, 10, 64)
		}

		var ftype int8
		if msgMap["type"] != nil {
			str := msgMap["type"].(string)
			size, err = strconv.ParseInt(str, 10, 8)
		}

		file := &File{
			Uuid:       msgMap["uuid"].(string),
			PUuid:      msgMap["puuid"].(string),
			Name:       msgMap["name"].(string),
			Md5:        msgMap["md5"].(string),
			Size:       size,
			Times:      times,
			UpdateTime: updateTime,
			CreateTime: createTime,
			Type:       ftype,
		}
		root.file = file

		for k, v := range jsonMap {
			if k != "msg" {
				generateFileTreeRecursion(root.addChild(), v.(map[string]interface{}))
			}
		}
	} else {
		panic(root.file.Name + " do not have msg field in json")
	}
}
