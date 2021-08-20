package controller

import (
	"Backlog2Teams/src/model/backlog"
	"Backlog2Teams/src/model/backlogRequest"
	response "Backlog2Teams/src/model/response"
	"Backlog2Teams/src/model/teamsMessage"
	utils "Backlog2Teams/src/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type NoticeController struct{}

func (nc NoticeController) DoNotice(c *gin.Context) {
	log.Printf("notice/do called")
	res := new(response.Response)

	var req backlogRequest.BacklogRequest
	err := c.BindJSON(&req)
	if err != nil {
		log.Printf(err.Error())
		res.SetMessage(err.Error())
		c.String(http.StatusBadRequest, utils.JsonToString(res))
		return
	}
	log.Print(req)

	log.Printf("Project Key is :" + req.Project.Projectkey)

	env := os.Getenv("ENV")

	baseURL := os.Getenv("BACKLOG_BASE_URL")

	msg := new(teamsMessage.TeamsMessage)
	msg.Type = "MessageCard"
	msg.Context = "http://schema.org/extensions"
	msg.Themecolor = "0076D7"
	msg.Summary = "Backlog更新通知"

	sec := new(teamsMessage.Sections)
	// BacklogロゴBase64画像
	sec.Activityimage = "data:image/jpeg;base64,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KKKACiiigAoPSiigD9Rvhj/yTvw3/wBgy1/9Fiu6ryn4Jazbax8NNFMDhns4fscyg8pJAduD9Vww9iK9WoP7v4drQq5VhqkHdOEP/SUQy28E+3z40k2nK71DYPqM1Kqqo2qAAOw4paKD2OVXuFFFct4l8a+FvCFubjxDqUFpxlY2bdK/+7GuXb8BQY4nFUcPTdbETUYrdt2S+bOpr5X+Ofxmi0y3uPBfhO43X0oMd9dRHi3Q8GNGH/LQjgkfcH+108/+I37ROq6/HLpHg1JNMsXBV7p+LqVfRcZEQPsS3uK+Z+pyeSeTQfz/AMf+K9OrSll2SSvfSVTbTqo9f+3vu7oooooP5+CiiigAooooAKKKKACiiigD/9D3yiiig/z3CiiigAooooAKKKKAO48DfEPxL8Pr57vQZl8qbHn20wLQTbemQCCGHZgQfw4r3pP2q9SCASeHYC+OSt0wBPsDGT+tfJtFB9TlHGud5XR+r4HEOMO1k0vTmTt8j6sl/ap1oj9z4ftV/wB64dv5IKw739p7x3OpWzsdNtc9G2SSMPzkA/Svm+ig7a3iPxLVVpYuXyUV+SR6hrHxn+Jmtq0dzrk0EbcFLRVthj6xgN/49Xmc001xK09xI8srnLO7FmY+5PJqOig+Yx2aYzGy58ZVlN/3pN/mFFFFBwBRRRQAUUUUAFFFFABRRRQAUUUUAf/R98ooooP89wooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD//0vfKKKKD/PcKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooA//9P3yiiig/z3CiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAP//Z"
	tar := new(teamsMessage.Targets)
	tar.Os = "default"
	tar.URI = baseURL + "projects/" + req.Project.Projectkey
	fact := new(teamsMessage.Facts)

	fact = new(teamsMessage.Facts)
	fact.Name = "プロジェクト名"
	fact.Value = req.Project.Name
	sec.Facts = append(sec.Facts, *fact)

	if 1 <= req.Type && req.Type <= 4 {
		// 課題更新系
		issueKey := req.Project.Projectkey + "-" + strconv.Itoa(req.Content.KeyID)
		fact = new(teamsMessage.Facts)
		fact.Name = "課題キー"
		fact.Value = issueKey
		sec.Facts = append(sec.Facts, *fact)

		fact = new(teamsMessage.Facts)
		fact.Name = "課題名"
		fact.Value = req.Content.Summary
		sec.Facts = append(sec.Facts, *fact)

		fact = new(teamsMessage.Facts)
		fact.Name = "課題種別"
		fact.Value = req.Content.Issuetype.Name
		sec.Facts = append(sec.Facts, *fact)

		if req.Type == 1 {
			sec.Activitytitle = "課題が追加されました"

			fact = new(teamsMessage.Facts)
			fact.Name = "状態"
			fact.Value = req.Content.Status.Name
			sec.Facts = append(sec.Facts, *fact)

			fact = new(teamsMessage.Facts)
			fact.Name = "内容"
			fact.Value = req.Content.Description
			fact.Value = strings.Replace(fact.Value, "\n", "<br />", -1)
			sec.Facts = append(sec.Facts, *fact)

			tar.URI = baseURL + "view/" + issueKey
		} else if req.Type == 2 {
			sec.Activitytitle = "課題が更新されました"

			fact = new(teamsMessage.Facts)
			fact.Name = "状態"
			fact.Value = req.Content.Status.Name
			sec.Facts = append(sec.Facts, *fact)

			if len(req.Content.Changes) > 0 {
				for i, v := range req.Content.Changes {
					fact = new(teamsMessage.Facts)
					fact.Name = "変更点" + strconv.Itoa(i+1)
					var oldValue, newValue string
					if v.Field == "status" {
						oldValue = getStatusValue(v.OldValue)
						newValue = getStatusValue(v.NewValue)
					} else {
						oldValue = v.OldValue
						newValue = v.NewValue
					}

					fact.Value = "(" + v.Field + ") " + oldValue + "→" + newValue
					sec.Facts = append(sec.Facts, *fact)
				}
			}

			if req.Content.Comment.Content != "" {
				fact = new(teamsMessage.Facts)
				fact.Name = "コメント"
				fact.Value = req.Content.Comment.Content
				fact.Value = strings.Replace(fact.Value, "\n", "<br />", -1)
				sec.Facts = append(sec.Facts, *fact)
			}

			tar.URI = baseURL + "view/" + issueKey
		} else if req.Type == 3 {
			sec.Activitytitle = "課題にコメントされました"

			fact = new(teamsMessage.Facts)
			fact.Name = "状態"
			fact.Value = req.Content.Status.Name
			sec.Facts = append(sec.Facts, *fact)

			fact = new(teamsMessage.Facts)
			fact.Name = "コメント"
			if req.Content.Comment.Content != "" {
				fact.Value = req.Content.Comment.Content
				fact.Value = strings.Replace(fact.Value, "\n", "<br />", -1)
			} else {
				fact.Value = "[コメントの取得に失敗しました]"
			}

			sec.Facts = append(sec.Facts, *fact)

			tar.URI = baseURL + "view/" + issueKey +
				"#comment-" + strconv.Itoa(req.Content.Comment.ID)

		} else if req.Type == 4 {
			sec.Activitytitle = "課題が削除されました"
		}

	} else if 5 <= req.Type && req.Type <= 7 {
		fact = new(teamsMessage.Facts)
		fact.Name = "タイトル"
		fact.Value = req.Content.Name
		sec.Facts = append(sec.Facts, *fact)

		if req.Type == 5 {
			sec.Activitytitle = "wikiが追加されました"
			fact = new(teamsMessage.Facts)
			fact.Name = "内容"
			fact.Value = req.Content.Content
			fact.Value = strings.Replace(fact.Value, "\n", "<br />", -1)
			sec.Facts = append(sec.Facts, *fact)
			tar.URI = baseURL + "alias/wili/" + strconv.Itoa(req.Content.ID)
		} else if req.Type == 6 {
			sec.Activitytitle = "wikiが更新されました"

			fact = new(teamsMessage.Facts)
			fact.Name = "内容"
			fact.Value = req.Content.Content
			fact.Value = strings.Replace(fact.Value, "\n", "<br />", -1)
			sec.Facts = append(sec.Facts, *fact)

			tar.URI = baseURL + "alias/wili/" + strconv.Itoa(req.Content.ID)

		} else if req.Type == 7 {
			sec.Activitytitle = "wikiが削除されました"

			fact = new(teamsMessage.Facts)
			fact.Name = "タイトル"
			fact.Value = req.Content.Name
			sec.Facts = append(sec.Facts, *fact)
		}
	} else if 8 <= req.Type && req.Type <= 10 {
		// ファイル系
		if req.Type == 8 {
			sec.Activitytitle = "ファイルが追加されました"
			tar.URI = baseURL + "alias/file/" + strconv.Itoa(req.Content.ID)
		} else if req.Type == 9 {
			sec.Activitytitle = "ファイルが更新されました"
			tar.URI = baseURL + "alias/file/" + strconv.Itoa(req.Content.ID)
		} else if req.Type == 10 {
			sec.Activitytitle = "ファイルが削除されました"
		}
	} else if req.Type == 11 {
		// Subversion
		sec.Activitytitle = "Subversionにコミットがありました"
		tar.URI = baseURL + "subversion/" + req.Project.Projectkey

		fact = new(teamsMessage.Facts)
		fact.Name = "注記"
		fact.Value = "通知詳細未実装"
		sec.Facts = append(sec.Facts, *fact)

	} else if 12 <= req.Type && req.Type <= 13 {
		// Git系
		if req.Type == 12 {
			sec.Activitytitle = "Gitにプッシュがありました"
			tar.URI = baseURL + "git/" + req.Project.Projectkey + "/" + req.Content.Repository.Name
		} else if req.Type == 13 {
			sec.Activitytitle = "Gitリポジトリが作成されました"
			tar.URI = baseURL + "git/" + req.Project.Projectkey + "/" + req.Content.Repository.Name
		}

		fact = new(teamsMessage.Facts)
		fact.Name = "注記"
		fact.Value = "通知詳細未実装"
		sec.Facts = append(sec.Facts, *fact)

	} else if req.Type == 14 {
		// まとめて更新
		sec.Activitytitle = "課題がまとめて更新されました"
		fact = new(teamsMessage.Facts)
		fact.Name = "注記"
		fact.Value = "通知詳細未実装"
		sec.Facts = append(sec.Facts, *fact)

	} else if 15 <= req.Type && req.Type <= 16 {
		// メンバー系
		fact = new(teamsMessage.Facts)
		fact.Value = ""
		if req.Type == 15 {
			sec.Activitytitle = "プロジェクトに参加しました"
			fact.Name = "参加者"
			if len(req.Content.Users) > 0 {
				for _, v := range req.Content.Users {
					fact.Value += v.Name + "(" + v.Nulabaccount.Name + ")" + ","
				}
			}
			sec.Facts = append(sec.Facts, *fact)

		} else if req.Type == 16 {
			sec.Activitytitle = "プロジェクトから脱退しました"
			fact.Name = "脱退者"
		}

		fact.Value = ""
		if len(req.Content.Users) > 0 {
			for _, v := range req.Content.Users {
				fact.Value += v.Name + "(" + v.Nulabaccount.Name + ")" + ","
			}
		}
		sec.Facts = append(sec.Facts, *fact)
	} else if req.Type == 17 {
		// お知らせ追加
		sec.Activitytitle = "お知らせに追加されました"
		fact = new(teamsMessage.Facts)
		fact.Name = "注記"
		fact.Value = "通知詳細未実装"
		sec.Facts = append(sec.Facts, *fact)

	} else if 18 <= req.Type && req.Type <= 20 {
		// プルリクエスト
		if req.Type == 18 {
			sec.Activitytitle = "プルリクエストがありました"
		} else if req.Type == 19 {
			sec.Activitytitle = "プルリクエスト更新されました"
		} else if req.Type == 20 {
			sec.Activitytitle = "プルリクエストにコメントがありました"
		}

		fact = new(teamsMessage.Facts)
		fact.Name = "注記"
		fact.Value = "通知詳細未実装"
		sec.Facts = append(sec.Facts, *fact)

	} else if 22 <= req.Type && req.Type <= 24 {
		// 発生バージョン/マイルストーン
		if req.Type == 22 {
			sec.Activitytitle = "発生バージョン/マイルストーンの追加"
		} else if req.Type == 23 {
			sec.Activitytitle = "発生バージョン/マイルストーンの更新"
		} else if req.Type == 24 {
			sec.Activitytitle = "発生バージョン/マイルストーンの削除"
		}
		fact = new(teamsMessage.Facts)
		fact.Name = "注記"
		fact.Value = "通知詳細未実装"
		sec.Facts = append(sec.Facts, *fact)

	} else {
		sec.Activitytitle = "この更新タイプは未対応です。 Type:" + strconv.Itoa(req.Type)
	}

	msg.Sections = append(msg.Sections, *sec)

	pot := new(teamsMessage.Potentialaction)
	pot.Targets = append(pot.Targets, *tar)
	pot.Type = "OpenUri"
	pot.Name = "Backlogで開く"
	msg.Potentialaction = append(msg.Potentialaction, *pot)

	// 送信
	url := ""
	if env == "local" {
		url = os.Getenv("TEAMS_TEST_URL")
		teamsPostSimpleMsg(url, msg)
	} else {
		b := new(backlog.Backlog)
		data, err := b.GetItem(req.Project.Projectkey)
		if err != nil {
			log.Printf(err.Error())
			res.SetMessage(err.Error())
			c.String(http.StatusBadRequest, utils.JsonToString(res))
			return
		}

		for _, v := range data.GetConnections() {
			log.Printf("webhookURL: " + v.GetTeamsWebhookURL())
			url = v.GetTeamsWebhookURL()
			teamsPostSimpleMsg(url, msg)
		}
	}

	res.SetMessage("Success")

	c.JSON(http.StatusOK, res)
	return
}

func teamsPostSimpleMsg(url string, msg *teamsMessage.TeamsMessage) {
	fmt.Println(msg)
	jsonString, err := json.Marshal(msg)
	fmt.Println(jsonString)
	if err != nil {
		fmt.Printf("send msg convert Error, " + err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	if err != nil {
		fmt.Printf("Error, " + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error, " + err.Error())
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error, " + err.Error())
	}
	fmt.Printf("%#v", string(byteArray))

}

func getStatusValue(code string) string {

	if code == "1" {
		return "未対応"
	} else if code == "2" {
		return "処理中"
	} else if code == "3" {
		return "処理済み"
	} else if code == "4" {
		return "完了"
	} else {
		return "[カスタム状態]"
	}
}
