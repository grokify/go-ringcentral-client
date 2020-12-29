package clientutil

import (
	"fmt"
	"net/http"

	"github.com/grokify/simplego/mime/multipartutil"
	"github.com/grokify/simplego/net/httputilmore"
	"github.com/grokify/simplego/net/urlutil"
)

func UploadUserProfileImage(httpClient *http.Client, serverUrl, accountId, extensionId, imgFilepath string) (*http.Response, error) {
	builder := multipartutil.NewMultipartBuilder()
	err := builder.WriteFilePathPlus("image", imgFilepath, false)
	if err != nil {
		return nil, err
	}
	err = builder.Close()
	if err != nil {
		return nil, err
	}
	//fmt.Println(builder.Buffer.String())

	urlPath := fmt.Sprintf(`restapi/v1.0/account/%s/extension/%s/profile-image`, accountId, extensionId)
	apiURL := urlutil.JoinAbsolute(serverUrl, urlPath)

	req, err := http.NewRequest(http.MethodPost, apiURL, builder.Buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Add(httputilmore.HeaderContentType, builder.ContentType())
	return httpClient.Do(req)
}
