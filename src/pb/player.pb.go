// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: player.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 玩家信息
type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId          string                    `protobuf:"bytes,1,opt,name=PlayerId,proto3" json:"PlayerId,omitempty" bson:"playerId"`                                                                                                  // 玩家ID       @gotags: bson:"playerId"
	AccountId         string                    `protobuf:"bytes,2,opt,name=AccountId,proto3" json:"AccountId,omitempty" bson:"accountId"`                                                                                                // 账户ID       @gotags: bson:"accountId"
	Nickname          string                    `protobuf:"bytes,3,opt,name=Nickname,proto3" json:"Nickname,omitempty" bson:"nickname"`                                                                                                  // 昵称         @gotags: bson:"nickname"
	Avatar            string                    `protobuf:"bytes,4,opt,name=Avatar,proto3" json:"Avatar,omitempty" bson:"avatar"`                                                                                                      // 头像         @gotags: bson:"avatar"
	Level             int64                     `protobuf:"varint,5,opt,name=Level,proto3" json:"Level,omitempty" bson:"level"`                                                                                                       // 等级         @gotags: bson:"level"
	Exp               int64                     `protobuf:"varint,6,opt,name=Exp,proto3" json:"Exp,omitempty" bson:"exp"`                                                                                                           // 经验         @gotags: bson:"exp"
	WorldLevel        int64                     `protobuf:"varint,7,opt,name=WorldLevel,proto3" json:"WorldLevel,omitempty" bson:"worldLevel"`                                                                                             // 世界等级      @gotags: bson:"worldLevel"
	Roles             []*UserRole               `protobuf:"bytes,8,rep,name=Roles,proto3" json:"Roles,omitempty" bson:"roles"`                                                                                                        // 用户角色      @gotags: bson:"roles"
	Teams             []*UserTeam               `protobuf:"bytes,9,rep,name=Teams,proto3" json:"Teams,omitempty" bson:"teams"`                                                                                                        // 用户编队      @gotags: bson:"teams"
	Props             []*UserProp               `protobuf:"bytes,10,rep,name=Props,proto3" json:"Props,omitempty" bson:"props"`                                                                                                       // 用户道具      @gotags: bson:"props"
	CreateAt          int64                     `protobuf:"varint,11,opt,name=CreateAt,proto3" json:"CreateAt,omitempty" bson:"createAt"`                                                                                                // 创建时间      @gotags: bson:"createAt"
	LastHeartbeatTime int64                     `protobuf:"varint,12,opt,name=LastHeartbeatTime,proto3" json:"LastHeartbeatTime,omitempty" bson:"lastHeartbeatTime"`                                                                              // 最后心跳时间   @gotags: bson:"lastHeartbeatTime"
	StageMission      map[string]*UnlockMission `protobuf:"bytes,13,rep,name=StageMission,proto3" json:"StageMission,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"stageMission"` // 关卡信息 @gotags: bson:"stageMission"
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *Player) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Player) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Player) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Player) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Player) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *Player) GetWorldLevel() int64 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *Player) GetRoles() []*UserRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Player) GetTeams() []*UserTeam {
	if x != nil {
		return x.Teams
	}
	return nil
}

func (x *Player) GetProps() []*UserProp {
	if x != nil {
		return x.Props
	}
	return nil
}

func (x *Player) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Player) GetLastHeartbeatTime() int64 {
	if x != nil {
		return x.LastHeartbeatTime
	}
	return 0
}

func (x *Player) GetStageMission() map[string]*UnlockMission {
	if x != nil {
		return x.StageMission
	}
	return nil
}

// 用户角色
type UserRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRoleId    string           `protobuf:"bytes,1,opt,name=UserRoleId,proto3" json:"UserRoleId,omitempty" bson:"userRoleId"`                                                                                              // 用户角色ID    @gotags: bson:"userRoleId"
	RoleId        int64            `protobuf:"varint,2,opt,name=RoleId,proto3" json:"RoleId,omitempty" bson:"roleId"`                                                                                                     // 角色ID       @gotags: bson:"roleId"
	Level         int64            `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty" bson:"level"`                                                                                                       // 等级         @gotags: bson:"level"
	SkinId        int64            `protobuf:"varint,4,opt,name=SkinId,proto3" json:"SkinId,omitempty" bson:"skinId"`                                                                                                     // 皮肤ID       @gotags: bson:"skinId"
	CareerId      int64            `protobuf:"varint,5,opt,name=CareerId,proto3" json:"CareerId,omitempty" bson:"careerId"`                                                                                                 // 职业ID       @gotags: bson:"careerId"
	AdvancementId int64            `protobuf:"varint,6,opt,name=AdvancementId,proto3" json:"AdvancementId,omitempty" bson:"advancementId"`                                                                                       // 阶段ID       @gotags: bson:"advancementId"
	MaxLevel      int64            `protobuf:"varint,7,opt,name=MaxLevel,proto3" json:"MaxLevel,omitempty" bson:"maxLevel"`                                                                                                 // 等级上线      @gotags: bson:"maxLevel"
	Exp           int64            `protobuf:"varint,8,opt,name=Exp,proto3" json:"Exp,omitempty" bson:"exp"`                                                                                                           // 当前经验      @gotags: bson:"exp"
	UnlockCareer  map[string]int64 `protobuf:"bytes,9,rep,name=UnlockCareer,proto3" json:"UnlockCareer,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3" bson:"unlockCareer"` // 已解锁职业    @gotags: bson:"unlockCareer"
}

func (x *UserRole) Reset() {
	*x = UserRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRole) ProtoMessage() {}

func (x *UserRole) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRole.ProtoReflect.Descriptor instead.
func (*UserRole) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{1}
}

func (x *UserRole) GetUserRoleId() string {
	if x != nil {
		return x.UserRoleId
	}
	return ""
}

func (x *UserRole) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *UserRole) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *UserRole) GetSkinId() int64 {
	if x != nil {
		return x.SkinId
	}
	return 0
}

func (x *UserRole) GetCareerId() int64 {
	if x != nil {
		return x.CareerId
	}
	return 0
}

func (x *UserRole) GetAdvancementId() int64 {
	if x != nil {
		return x.AdvancementId
	}
	return 0
}

func (x *UserRole) GetMaxLevel() int64 {
	if x != nil {
		return x.MaxLevel
	}
	return 0
}

func (x *UserRole) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *UserRole) GetUnlockCareer() map[string]int64 {
	if x != nil {
		return x.UnlockCareer
	}
	return nil
}

type UserProp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId string `protobuf:"bytes,1,opt,name=ItemId,proto3" json:"ItemId,omitempty" bson:"itemId"`  // 道具ID   @gotags: bson:"itemId"
	Counts int64  `protobuf:"varint,2,opt,name=Counts,proto3" json:"Counts,omitempty" bson:"counts"` // 道具数量  @gotags: bson:"counts"
}

func (x *UserProp) Reset() {
	*x = UserProp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProp) ProtoMessage() {}

func (x *UserProp) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProp.ProtoReflect.Descriptor instead.
func (*UserProp) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{2}
}

func (x *UserProp) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *UserProp) GetCounts() int64 {
	if x != nil {
		return x.Counts
	}
	return 0
}

// 用户编队
type UserTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId      string  `protobuf:"bytes,1,opt,name=TeamId,proto3" json:"TeamId,omitempty" bson:"teamId"`                   // 编队ID @gotags: bson:"teamId"
	TeamName    string  `protobuf:"bytes,2,opt,name=TeamName,proto3" json:"TeamName,omitempty" bson:"teamName"`               // 编队名称 @gotags: bson:"teamName"
	UserRoleIds []int64 `protobuf:"varint,3,rep,packed,name=UserRoleIds,proto3" json:"UserRoleIds,omitempty" bson:"userRoleIds"` // 用户角色ID @gotags: bson:"userRoleIds"
}

func (x *UserTeam) Reset() {
	*x = UserTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTeam) ProtoMessage() {}

func (x *UserTeam) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTeam.ProtoReflect.Descriptor instead.
func (*UserTeam) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{3}
}

func (x *UserTeam) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *UserTeam) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *UserTeam) GetUserRoleIds() []int64 {
	if x != nil {
		return x.UserRoleIds
	}
	return nil
}

type UnlockMission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MissionID int64  `protobuf:"varint,1,opt,name=MissionID,proto3" json:"MissionID,omitempty" bson:"missionID"` //@gotags: bson:"missionID"
	ChapterID int64  `protobuf:"varint,2,opt,name=ChapterID,proto3" json:"ChapterID,omitempty" bson:"chapterID"` //@gotags: bson:"chapterID"
	StarLevel uint32 `protobuf:"varint,3,opt,name=StarLevel,proto3" json:"StarLevel,omitempty" bson:"starLevel"` //@gotags: bson:"starLevel"
}

func (x *UnlockMission) Reset() {
	*x = UnlockMission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockMission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockMission) ProtoMessage() {}

func (x *UnlockMission) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockMission.ProtoReflect.Descriptor instead.
func (*UnlockMission) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{4}
}

func (x *UnlockMission) GetMissionID() int64 {
	if x != nil {
		return x.MissionID
	}
	return 0
}

func (x *UnlockMission) GetChapterID() int64 {
	if x != nil {
		return x.ChapterID
	}
	return 0
}

func (x *UnlockMission) GetStarLevel() uint32 {
	if x != nil {
		return x.StarLevel
	}
	return 0
}

var File_player_proto protoreflect.FileDescriptor

var file_player_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x9e, 0x04, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x78, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x45, 0x78, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x26, 0x0a, 0x05, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x05, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x05, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x4c,
	0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x4c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0c, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a,
	0x56, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x55,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe9, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61,
	0x72, 0x65, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x61,
	0x72, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x41,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x4d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x4d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x78, 0x70, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x45, 0x78, 0x70, 0x12, 0x46, 0x0a, 0x0c, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x65,
	0x65, 0x72, 0x1a, 0x3f, 0x0a, 0x11, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x65,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x3a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22,
	0x60, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x54,
	0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64,
	0x73, 0x22, 0x69, 0x0a, 0x0d, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_player_proto_rawDescOnce sync.Once
	file_player_proto_rawDescData = file_player_proto_rawDesc
)

func file_player_proto_rawDescGZIP() []byte {
	file_player_proto_rawDescOnce.Do(func() {
		file_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_player_proto_rawDescData)
	})
	return file_player_proto_rawDescData
}

var file_player_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_player_proto_goTypes = []interface{}{
	(*Player)(nil),        // 0: player.Player
	(*UserRole)(nil),      // 1: player.UserRole
	(*UserProp)(nil),      // 2: player.UserProp
	(*UserTeam)(nil),      // 3: player.UserTeam
	(*UnlockMission)(nil), // 4: player.UnlockMission
	nil,                   // 5: player.Player.StageMissionEntry
	nil,                   // 6: player.UserRole.UnlockCareerEntry
}
var file_player_proto_depIdxs = []int32{
	1, // 0: player.Player.Roles:type_name -> player.UserRole
	3, // 1: player.Player.Teams:type_name -> player.UserTeam
	2, // 2: player.Player.Props:type_name -> player.UserProp
	5, // 3: player.Player.StageMission:type_name -> player.Player.StageMissionEntry
	6, // 4: player.UserRole.UnlockCareer:type_name -> player.UserRole.UnlockCareerEntry
	4, // 5: player.Player.StageMissionEntry.value:type_name -> player.UnlockMission
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_player_proto_init() }
func file_player_proto_init() {
	if File_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRole); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTeam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockMission); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_player_proto_goTypes,
		DependencyIndexes: file_player_proto_depIdxs,
		MessageInfos:      file_player_proto_msgTypes,
	}.Build()
	File_player_proto = out.File
	file_player_proto_rawDesc = nil
	file_player_proto_goTypes = nil
	file_player_proto_depIdxs = nil
}