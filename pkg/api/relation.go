package api

import "context"

type RelationActionRequest struct {
	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ToUserId   int64  `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`
	ActionType int32  `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
}

type RelationActionResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

type RelationFollowListRequest struct {
	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

type RelationFollowListResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	UserList   []User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

type RelationFollowerListRequest struct {
	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

type RelationFollowerListResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	UserList   []User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

type RelationFriendListRequest struct {
	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

type RelationFriendListResponse struct {
	StatusCode int32        `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string       `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	UserList   []FriendUser `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

type FriendUser struct {
	User
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	MsgType int64  `protobuf:"varint,2,opt,name=msgType,proto3" json:"msgType,omitempty"`
}

type RelationActionApiParam struct{}

func NewRelationActionApi(param RelationActionApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/relation/action/",
		Handler: WrapperFunc(func(ctx context.Context, req *RelationActionRequest) (*RelationActionResponse, error) {
			return nil, nil
		}),
	}
}

type RelationFollowListApiParam struct{}

func NewRelationFollowListApi(param RelationFollowListApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/relation/follow_list/",
		Handler: WrapperFunc(func(ctx context.Context, req *RelationFollowListRequest) (*RelationFollowListResponse, error) {
			return nil, nil
		}),
	}
}

type RelationFollowerListApiParam struct{}

func NewRelationFollowerListApi(param RelationFollowerListApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/relation/follower_list/",
		Handler: WrapperFunc(func(ctx context.Context, req *RelationFollowerListRequest) (*RelationFollowerListResponse, error) {
			return nil, nil
		}),
	}
}

type RelationFriendListApiParam struct{}

func NewRelationFriendListApi(param RelationFriendListApiParam) *Api {
	return &Api{
		Method: "POST",
		Path:   "/relation/friend_list/",
		Handler: WrapperFunc(func(ctx context.Context, req *RelationFriendListRequest) (*RelationFriendListResponse, error) {
			return nil, nil
		}),
	}
}
