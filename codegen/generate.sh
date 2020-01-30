#perl -p -i -e 's/\s+\[Beta\]\s*$/\n/g' openapi-spec.json
export GO_POST_PROCESS_FILE="/usr/local/go/bin/gofmt -s -w"

java -jar openapi-generator-cli.jar generate -i openapi-spec_v3.0.0.json -g go -o engagedigital --package-name engagedigital

#java -jar openapi-generator-cli.jar generate -i partial-specs_v3.0.0/openapi-spec_teams.json -g go -o engagedigital --package-name engagedigital

#perl -p -i -e 's/(delimiter,\s+-1\),\s+"\[\]"\))/$1\n} else if t, ok := obj.(time.Time); ok {\nreturn t.Format(time.RFC3339)/g' client/client.go
echo "\n\nfunc (apiClient *APIClient) HTTPClient() *http.Client { return apiClient.cfg.HTTPClient }" >> engagedigital/client.go
gofmt -s -w engagedigital/*.go
