perl -p -i -e 's/\s+\[Beta\]\s*$/\n/g' swagger_spec.yaml
java -jar openapi-generator-cli.jar generate -i swagger_spec.yaml -g go -o client -D packageName=ringcentral
#perl -p -i -e 's/(delimiter,\s+-1\),\s+"\[\]"\))/$1\n} else if t, ok := obj.(time.Time); ok {\nreturn t.Format(time.RFC3339)/g' client/client.go
echo "\n\nfunc (apiClient *APIClient) HTTPClient() *http.Client { return apiClient.cfg.HTTPClient }" >> client/client.go
gofmt -s -w client/*.go
