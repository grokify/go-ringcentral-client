package lite

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/grokify/simplego/encoding/jsonutil"
	"github.com/grokify/simplego/net/httputilmore"
	"github.com/pkg/errors"
)

// portal.vacd.biz:8081

const (
	EngageVoiceServerURL    string = "https://portal.vacd.biz"
	EngageVoiceLoginURL     string = "https://portal.vacd.biz/api/v1/auth/login"
	EngageVoiceTokenURL     string = "https://portal.vacd.biz/api/v1/admin/token"
	EngageVoiceTokenURLPath string = "/api/v1/admin/token"
	EngageVoiceTokenHeader  string = "x-auth-token"
)

type LoginSuccess struct {
	AuthToken    string    `json:"authToken"`
	PlatformHost string    `json:"platformHost"`
	Accounts     []Account `json:"accounts"`
}

type User struct {
	UserID       int64     `json:"userId"`
	UserName     string    `json:"userName"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	FullName     string    `json:"fullName"`
	CreationDate time.Time `json:"creationDate"`
	Enabled      bool      `json:"enabled"`
	ParentPath   string    `json:"parentPath"`
	RcUserID     string    `json:"rcUserId"`
	Roles        []Role    `json:"roles"`
	RootUser     bool      `json:"rootUser"`
	// PhoneNumber  string    `json:"phoneNumber"`
	// Children []Children `json:"children"`
}

type Role struct {
	RoleType    string `json:"roleType"`
	Description string `json:"description"`
}

type Account struct {
	MainAccountID               string `json:"mainAccountId"`
	AccountID                   string `json:"accountId"`
	AccountName                 string `json:"accountName"`
	EnableMultiUser             bool   `json:"enableMultiUser"`
	EnableSoftphones            bool   `json:"enableSoftphones"`
	TcpaSafeMode                bool   `json:"tcpaSafeMode"`
	EnableVoiceBroadcast        bool   `json:"enableVoiceBroadcast"`
	OutboundPrepay              bool   `json:"outboundPrepay"`
	Enable247Dialing            bool   `json:"enable247Dialing"`
	EnableFifo                  bool   `json:"enableFifo"`
	EnableChat                  bool   `json:"enableChat"`
	EnableInbound               bool   `json:"enableInbound"`
	EnableOutbound              bool   `json:"enableOutbound"`
	EnableVisualIvr             bool   `json:"enableVisualIvr"`
	EnableTracking              bool   `json:"enableTracking"`
	EnableCloudRouting          bool   `json:"enableCloudRouting"`
	UsePowerBy                  bool   `json:"usePowerBy"`
	EnableAgentRankRouting      bool   `json:"enableAgentRankRouting"`
	EnableHciDialer             bool   `json:"enableHciDialer"`
	EnableTcpaSafeMachineDetect bool   `json:"enableTcpaSafeMachineDetect"`
	EmailFromAddress            string `json:"emailFromAddress"`
	DatabaseShardID             string `json:"databaseShardId"`
	DefaultOutdialServerGroupID int64  `json:"defaultOutdialServerGroupId"`
	DefaultIntellidialServerID  int64  `json:"defaultIntellidialServerId"`
	RcAccountAccess             string `json:"rcAccountAccess"`
	EnableGoodData              bool   `json:"enableGoodData"`
	MainAccountName             string `json:"mainAccountName"`
}

type LoginError struct {
	GeneralMessage string `json:"generalMessage"`
	Details        string `json:"details"`
	RequestURI     string `json:"requestUri"`
	Timestamp      int64  `json:"timestamp"`
}

func GenerateAPIToken(username, password string) (string, *LoginSuccess, error) {
	respSucc, _, resp, err := RequestAuthToken(username, password)
	if err != nil {
		return "", nil, err
	} else if resp.StatusCode != 200 {
		return "", nil, fmt.Errorf("E_AUTH_RESPONSE_STATUS_CODE_NOT_200 [%v]", resp.StatusCode)
	}

	if len(strings.TrimSpace(respSucc.AuthToken)) == 0 {
		return "", nil, errors.New("E_EMPTY_AUTH_TOKEN")
	}

	apiToken, err := ExchangeAPIToken(respSucc.AuthToken)
	if err != nil {
		return "", respSucc, err
	}
	return apiToken, respSucc, nil
}

func ExchangeAPIToken(authToken string) (string, error) {
	req, err := http.NewRequest(http.MethodPost, EngageVoiceTokenURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add(EngageVoiceTokenHeader, authToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "ExchangeAPIToken.client.Do()")
	} else if resp.StatusCode != 200 {
		return "", fmt.Errorf("TOKEN_RESP_STATUS_CODE_NOT_200 [%v]", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	token := strings.TrimSpace(string(body))
	if len(token) == 0 {
		return "", fmt.Errorf("E_EMPTY_API_TOKEN")
	}
	return token, nil
}

func RequestAuthToken(username, password string) (*LoginSuccess, *LoginError, *http.Response, error) {
	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)

	resp, err := httputilmore.SendWwwFormUrlEncodedSimple(http.MethodPost, EngageVoiceLoginURL, data)
	if err != nil {
		return nil, nil, resp, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
	if resp.StatusCode != 200 {
		respData := LoginError{}
		_, err := jsonutil.UnmarshalIoReader(resp.Body, &respData)
		if err != nil {
			return nil, &respData, resp, err
		}
		return nil, &respData, resp, fmt.Errorf("AuthUrl Status != 200 [%v]", resp.StatusCode)
	}
	authData := LoginSuccess{}
	_, err = jsonutil.UnmarshalIoReader(resp.Body, &authData)
	if err != nil {
		return nil, nil, resp, err
	}
	return &authData, nil, resp, nil
}

func ListTokens(serverURL, authOrApiToken string) ([]string, error) {
	tokens := []string{}
	headers, apiURL, err := APIInfo(
		serverURL, EngageVoiceTokenURLPath, authOrApiToken)
	if err != nil {
		return tokens, err
	}

	resp, err := httputilmore.GetJsonSimple(apiURL, headers, &tokens)
	if err != nil {
		return tokens, err
	} else if resp.StatusCode >= 300 {
		return tokens, errors.New("E_RESPONSE_STATUS_CODE_NOT_2XX")
	}
	return tokens, nil
}
