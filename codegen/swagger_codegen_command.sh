perl -p -i -e 's/\s+\[Beta\]\s*$/\n/g' swagger_spec.yaml
java -jar swagger-codegen-cli.jar generate -c swagger_codegen_config.json -i swagger_spec.yaml -l go -o client
grep '*os.File,' client/*.go
perl -p -i -e 's/, ([^\s]+) \*os\.File,/, localVarFile *os.File,/g' `grep -ril ' *os.File,' client/*.go`
grep '*os.File,' client/*.go
grep '*interface{}' client/*.go
perl -p -i -e 's/\*(interface\{\})/$1/g' `grep -ril '*interface{}' client/*.go`
perl -p -i -e 's/(delimiter,\s+-1\),\s+"\[\]"\))/$1\n} else if t, ok := obj.(time.Time); ok {\nreturn t.Format(time.RFC3339)/g' client/api_client.go
echo "func (apiClient *APIClient) HTTPClient() *http.Client { return apiClient.cfg.HTTPClient }" >> client/api_client.go
gofmt -s -w client/*.go