// YApi QuickType插件生成，具体参考文档:https://plugins.jetbrains.com/plugin/18847-yapi-quicktype/documentation
package main

type AwemeUserPage struct {
	LogID      string      `json:"log_id"`
	AwemeUsers []AwemeUser `json:"aweme_users"`
	StatusCode int64       `json:"status_code"`
	TotalCount int64       `json:"total_count"`
	Now        string      `json:"now"`
	StatusMsg  string      `json:"status_msg"`
	HasMore    bool        `json:"has_more"`
}

type AwemeUser struct {
	AwemeUserAvatar string `json:"aweme_user_avatar"`
	AwemeID         string `json:"aweme_id"`
	AwemeUserID     int64  `json:"aweme_user_id"`
	IsEnable        bool   `json:"is_enable"`
	KeyAccountID    int64  `json:"key_account_id"`
	NickName        string `json:"nick_name"`
	IsBluev         bool   `json:"is_bluev"`
	BindType        int64  `json:"bind_type"`
	LifeAccountID   string `json:"life_account_id"`
}
