// YApi QuickType插件生成，具体参考文档:https://plugins.jetbrains.com/plugin/18847-yapi-quicktype/documentation
package main

type VideoPage struct {
	LogID      string `json:"log_id"`
	StatusCode int64  `json:"status_code"`
	Data       Data   `json:"data"`
	Now        string `json:"now"`
	StatusMsg  string `json:"status_msg"`
}

type Data struct {
	Pagination Pagination `json:"pagination"`
	List       []List     `json:"list"`
}

type List struct {
	Measures     Measures `json:"measures"`
	PlayURL      string   `json:"play_url"`
	UserID       string   `json:"user_id"`
	UserImage    string   `json:"user_image"`
	UserNickname string   `json:"user_nickname"`
	CoverImage   string   `json:"cover_image"`
	ID           string   `json:"id"`
	Time         string   `json:"time"`
	Title        string   `json:"title"`
}

type Measures struct {
	CERTNumAll     All `json:"cert_num_all"`
	PayOrderCntAll All `json:"pay_order_cnt_all"`
	PayGmvAll      All `json:"pay_gmv_all"`
	LikeCntAll     All `json:"like_cnt_all"`
}

type All struct {
	Unit string `json:"unit"`
	Num  string `json:"num"`
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Pagination struct {
	IsAsc      bool   `json:"is_asc"`
	TotalCount int64  `json:"total_count"`
	PageIndex  int64  `json:"page_index"`
	SortKey    string `json:"sort_key"`
	PageCount  int64  `json:"page_count"`
	PageSize   int64  `json:"page_size"`
}
