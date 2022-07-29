package mongo

import "testing"

func TestMongo(t *testing.T) {
	//info, err := mongo.Collection(mongo.GAME_DB, mongo.PLAYERS_COLLECTION).UpdateAll(bson.M{"playerId": 13532180605083655, "userteams.teamid": 13532180605083654}, bson.M{"$set": bson.M{"userteams.$.teamname": "demo"}})
	//if err != nil {
	//	fmt.Println("error:", err.Error())
	//}
	//fmt.Printf("info:%#v \n", info)

	//playerId := 16721836268101640
	//teamId := 16721836268101638
	//team := pb.UserTeam{
	//	TeamId:      111111,
	//	TeamName:    "test",
	//	UserRoleIds: []int64{1, 2, 3, 4},
	//}
	//_ = playerId
	//_ = teamId
	//_ = team

	//	内嵌子集合
	//	更新 $set, 条件中添加自己和查询条件
	//selector := bson.M{"playerId": playerId, "teams.teamId": teamId}
	//update := bson.M{"$set": bson.M{"teams.$": team}}
	//if err := mongo.Collection(mongo.GAME_DB, mongo.PLAYERS_COLLECTION).Update(selector, update); err != nil {
	//	fmt.Println("error:", err.Error())
	//}

	//	插入 $push
	//selector := bson.M{"playerId": playerId}
	//update := bson.M{"$push": bson.M{"teams": team}}
	//if err := mongo.Collection(mongo.GAME_DB, mongo.PLAYERS_COLLECTION).Update(selector, update); err != nil {
	//	fmt.Println("error:", err.Error())
	//}

	//	删除 $unset
	//selector := bson.M{"playerId": playerId}
	//update := bson.M{"$unset": bson.M{"teams": 1}}
	//if err := mongo.Collection(mongo.GAME_DB, mongo.PLAYERS_COLLECTION).Update(selector, update); err != nil {
	//	fmt.Println("error:", err.Error())
	//}

	//selector := bson.M{"playerId": playerId}
	//update := bson.M{"$pop": bson.M{"roles": 1}}
	//if err := mongo.Collection(mongo.GAME_DB, mongo.PLAYERS_COLLECTION).Update(selector, update); err != nil {
	//	fmt.Println("error:", err.Error())
	//}

	// 从数组field内删除一个等于value值
	//selector := bson.M{"name": "ls"}
	//update := bson.M{"$pull": bson.M{"tags": "a"}}
	//if err := mongo.Collection(mongo.GAME_DB, "person").Update(selector, update); err != nil {
	//	fmt.Println("error:", err.Error())
	//}

	//	https://www.runoob.com/mongodb/mongodb-atomic-operations.html

}
