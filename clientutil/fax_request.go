package clientutil

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/grokify/gotilla/mime/multipartutil"
	hum "github.com/grokify/gotilla/net/httputilmore"
)

const (
	attachmentFieldName = "attachment"
	FaxUrl              = "/restapi/v1.0/account/~/extension/~/fax"
)

// FaxRequest is a fax request helper that can send more than one attachment.
// The core Swagger Codegen appears to only send one file at a time with a fixed
// MIME part name.
type FaxRequest struct {
	CoverIndex          *int
	CoverPageText       string
	FaxResolution       string
	SendTime            *time.Time
	IsoCode             string
	To                  []string
	AttachmentFilepaths []string
}

func (fax *FaxRequest) builder() (multipartutil.MultipartBuilder, error) {
	builder := multipartutil.NewMultipartBuilder()

	if fax.CoverIndex != nil {
		if err := builder.Writer.WriteField("coverIndex", strconv.Itoa(*fax.CoverIndex)); err != nil {
			return builder, err
		}
	}
	if len(strings.TrimSpace(fax.CoverPageText)) > 0 {
		if err := builder.Writer.WriteField("coverPageText", fax.CoverPageText); err != nil {
			return builder, err
		}
	}
	if len(strings.TrimSpace(fax.FaxResolution)) > 0 {
		if err := builder.Writer.WriteField("faxResolution", fax.FaxResolution); err != nil {
			return builder, err
		}
	}
	if fax.SendTime != nil {
		if err := builder.Writer.WriteField("sendTime", fax.SendTime.Format(time.RFC3339)); err != nil {
			return builder, err
		}
	}
	for _, to := range fax.To {
		to := strings.TrimSpace(to)
		if len(to) > 0 {
			if err := builder.Writer.WriteField("to", to); err != nil {
				return builder, err
			}
		}
	}

	for _, attachmentName := range fax.AttachmentFilepaths {
		err := builder.WriteFile(attachmentFieldName, attachmentName)
		if err != nil {
			return builder, err
		}
	}

	err := builder.Close()
	return builder, err
}

func (fax *FaxRequest) Post(httpClient *http.Client, url string) (*http.Response, error) {
	builder, err := fax.builder()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, ioutil.NopCloser(builder.Buffer))
	req.Header.Set(hum.HeaderContentType, builder.Writer.FormDataContentType())
	if err != nil {
		return nil, err
	}
	return httpClient.Do(req)
}
