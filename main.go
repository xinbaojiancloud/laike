package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	title := "视频Title,发布时间,成交券数,点赞数,成交金额,成交订单数\n"
	douYinHaoIndex := 1
	for i := 1; i <= 50; i++ {
		awemeUserPage, err := getDouYinHao(i)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for _, awemeUser := range awemeUserPage.AwemeUsers {
			fmt.Println("---------------------------------------------------------------------")
			fmt.Printf("正在获取第%d个抖音号[%s]的视频数据\n", douYinHaoIndex, awemeUser.NickName)
			fileName := "csv/" + awemeUser.NickName + ".csv"
			var videoList []List
			videoIndex := 1
			for {
				jsonData, err := getVideosJson(videoIndex, strconv.FormatInt(awemeUser.AwemeUserID, 10), awemeUser.LifeAccountID)
				if err != nil {
					fmt.Println(err)
					continue
				}
				var videoPage VideoPage
				err = json.Unmarshal([]byte(jsonData), &videoPage)
				if len(videoPage.Data.List) == 0 {
					fmt.Printf("第%d个抖音号[%s]的视频数据获取完毕\n", douYinHaoIndex, awemeUser.NickName)
					break
				}
				videoList = append(videoList, videoPage.Data.List...)
				videoIndex = videoIndex + 1
			}
			if len(videoList) > 0 {
				fmt.Printf("抓取了%d个视频\n", len(videoList))
				fmt.Println("正在写入文件" + fileName)
				WriteToFile(fileName, title)
				for _, video := range videoList {
					str := video.Title + "," + video.Time + "," + video.Measures.CERTNumAll.Num + "," + video.Measures.LikeCntAll.Num + "," + video.Measures.PayGmvAll.Num + "," + video.Measures.PayOrderCntAll.Num + "\n"
					fmt.Println(str)
					WriteToFile(fileName, str)
				}
				fmt.Println("写入文件" + fileName + "成功")
			}
			douYinHaoIndex = douYinHaoIndex + 1
		}
		if !awemeUserPage.HasMore {
			fmt.Println("所有抖音号视频数据获取完毕")
			break
		}
	}
}

func getDouYinHao(index int) (AwemeUserPage, error) {
	dyhUrl := "https://life.douyin.com/life/gate/v2/account/bc_bind_relationships"
	params := url.Values{
		"page_index":           {strconv.Itoa(index)},
		"page_count":           {"10"},
		"search_keyword":       {""},
		"root_life_account_id": {"7101194713696307240"},
	}
	req, err := http.NewRequest("GET", dyhUrl+"?"+params.Encode(), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return AwemeUserPage{}, err
	}
	req.Header.Set("Ac-Tag", "ka_1h")
	req.Header.Set("Cookie", getCookie())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return AwemeUserPage{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return AwemeUserPage{}, err
	}
	var awemeUserPage AwemeUserPage
	err = json.Unmarshal(body, &awemeUserPage)
	if err != nil {
		fmt.Println("Error:", err)
		return AwemeUserPage{}, err
	}
	return awemeUserPage, nil
}

func getVideosJson(index int, awemeUserId string, lifeAccountIds string) (string, error) {
	videoUrl := "https://life.douyin.com/life/infra/v1/content/video/get_video_list"
	params := url.Values{
		"page_index":           {strconv.Itoa(index)},
		"page_size":            {"10"},
		"sort_key":             {"publish_time"},
		"is_asc":               {"false"},
		"aweme_id":             {awemeUserId},
		"content_tag":          {"2"},
		"poi_life_account_ids": {lifeAccountIds},
		"root_life_account_id": {"7101194713696307240"},
	}

	request, err := http.NewRequest("GET", videoUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	request.Header.Set("Ac-Tag", "ka_1h")
	request.Header.Set("Cookie", "s_v_web_id=verify_ljcnkqg2_cVTxZb2M_H0o8_4l3R_8Mht_gBAdT4ZjCxm0; csrf_session_id=5fa8c294aa95cbb1236f08b90498e850; passport_csrf_token=83d57adee803cf055d26abb7c602fb46; passport_csrf_token_default=83d57adee803cf055d26abb7c602fb46; odin_tt=1c4a032c944069a53112b70bc8281fa353bd8ee602b30aed1ec3e3c617771f2de24b0eced14e4a2a081d5c72077f2a5412f4e647018360c0d6196bdbb5af69d2; sid_guard_ls=feef383db0dc5f1ed97ba8054761c4f0%7C1687771548%7C5184002%7CFri%2C+25-Aug-2023+09%3A25%3A50+GMT; uid_tt_ls=d191137dfa4e1b3f39c9666c4ede382a; uid_tt_ss_ls=d191137dfa4e1b3f39c9666c4ede382a; sid_tt_ls=feef383db0dc5f1ed97ba8054761c4f0; sessionid_ls=feef383db0dc5f1ed97ba8054761c4f0; sessionid_ss_ls=feef383db0dc5f1ed97ba8054761c4f0; sid_ucp_v1_ls=1.0.0-KDUxNzQ3NjJjMDY3NjJkODQxOGM0NjMyZGRmNWYzZThkMmUzYjJmZTgKGgjNz-DEpYyJBBCcs-WkBhjRwRIgDDgBQOsHGgJobCIgZmVlZjM4M2RiMGRjNWYxZWQ5N2JhODA1NDc2MWM0ZjA; ssid_ucp_v1_ls=1.0.0-KDUxNzQ3NjJjMDY3NjJkODQxOGM0NjMyZGRmNWYzZThkMmUzYjJmZTgKGgjNz-DEpYyJBBCcs-WkBhjRwRIgDDgBQOsHGgJobCIgZmVlZjM4M2RiMGRjNWYxZWQ5N2JhODA1NDc2MWM0ZjA; store-region=cn-cq; store-region-src=uid; ttwid=1%7CAV_EhZH5_kGemuMTgOXCFcE-lmlS9HiqF3bJeLQDzY0%7C1687772850%7Cd228454397fb93b080661b9f4e0542f8b8b29cfd96aaf3c24157f5b4c3b9fd3e")
	request.URL.RawQuery = params.Encode()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}
	return string(body), nil
}

func getCookie() string {
	body, err := os.ReadFile("cookie.txt")
	if err != nil {
		fmt.Println("Error reading cookie")
	}
	return string(body)
}

func WriteToFile(filename string, content string) error {
	createDirectoryIfNotExists("csv")
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write content to file: %v", err)
	}
	return nil
}

func createDirectoryIfNotExists(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}
	return nil
}
