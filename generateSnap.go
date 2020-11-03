package snap

import (
	"github.com/jeremywohl/flatten"
	"time"
)

func generateFileTree() *Tree {
	root := &Tree{}
	root.generateRoot(nil)

	//read from json
	jsonMap := make(map[string]interface{})
	generateFileTreeRecursion(root, jsonMap)
	return root
}

func generateFileTreeRecursion(root *Tree, jsonMap map[string]interface{}) {
	flat, err := flatten.Flatten(jsonMap, "", flatten.RailsStyle)
	if err != nil {
		panic(err)
	}
	if msg, ok := flat["msg"]; ok {
		msgMap := msg.(map[string]interface{})
		updateTime, err := time.ParseInLocation("2006-01-02 15:04:05", msgMap["updatetime"].(string), time.Local)
		createTime, err := time.ParseInLocation("2006-01-02 15:04:05", msgMap["createtime"].(string), time.Local)
		if err != nil {
			panic("can not parse in location with update time and create time")
		}
		file := &File{
			Uuid:       msgMap["uuid"].(string),
			PUuid:      msgMap["puuid"].(string),
			Name:       msgMap["name"].(string),
			Md5:        msgMap["md5"].(string),
			Size:       int64(msgMap["size"].(float64)),
			Times:      int64(msgMap["times"].(float64)),
			UpdateTime: updateTime,
			CreateTime: createTime,
			Type:       msgMap["size"].(int8),
		}
		root.file = file

		for k, v := range flat {
			if k != "msg" {
				generateFileTreeRecursion(root.addChild(), v.(map[string]interface{}))
			}
		}
	} else {
		panic(root.file.Name + " do not have msg field in json")
	}
}
