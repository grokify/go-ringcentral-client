#perl -p -i -e 's/\s+\[Beta\]\s*$/\n/g' openapi-spec.json
go run merge.go -d specs-voice_v3.0.0 -v 3
export GO_POST_PROCESS_FILE="/usr/local/go/bin/gofmt -s -w"

java -jar openapi-generator-cli.jar generate -i specs-engagevoice_v3.0.0.json -g go -o engagevoice --package-name engagevoice

#java -jar openapi-generator-cli.jar generate -i partial-specs_v3.0.0/openapi-spec_teams.json -g go -o engagedigital --package-name engagedigital

#perl -p -i -e 's/(delimiter,\s+-1\),\s+"\[\]"\))/$1\n} else if t, ok := obj.(time.Time); ok {\nreturn t.Format(time.RFC3339)/g' client/client.go
echo "\n\nfunc (apiClient *APIClient) HTTPClient() *http.Client { return apiClient.cfg.HTTPClient }" >> engagevoice/client.go
gofmt -s -w engagevoice/*.go
