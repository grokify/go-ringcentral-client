java -jar swagger-codegen-cli.jar generate -c swagger_codegen_config.json -i swagger_spec.yaml -l go -o client
grep '*os.File,' client/*.go
perl -p -i -e 's/, ([^\s]+) \*os\.File,/, localVarFile *os.File,/g' `grep -ril ' *os.File,' client/*.go`
grep '*os.File,' client/*.go
grep '*interface{}' client/*.go
perl -p -i -e 's/\*(interface\{\})/$1/g' `grep -ril '*interface{}' client/*.go`
gofmt -s -w client/*.go