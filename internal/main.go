package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
)

func main() {

	if err := filepath.Walk("./internal/selling-partner-api-models/models", func(filePath string, info os.FileInfo, err error) error {
		if filepath.Ext(filePath) == ".json" {

			buf, err := ioutil.ReadFile(filePath)
			if err != nil {
				return err
			}

			var doc2 openapi2.Swagger
			if err = json.Unmarshal(buf, &doc2); err != nil {
				return err
			}

			doc3, err := openapi2conv.ToV3Swagger(&doc2)
			if err != nil {
				return err
			}

			outBuf, err := doc3.MarshalJSON()
			if err != nil {
				return err
			}

			filename := path.Base(filePath)
			err = ioutil.WriteFile(filepath.Join("./internal/selling-partner-api-models-temp", filename), outBuf, os.ModePerm)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		panic(err)
	}

	cmdStr := ""

	if err := filepath.Walk("./internal/selling-partner-api-models-temp", func(filePath string, info os.FileInfo, err error) error {
		if filepath.Ext(filePath) == ".json" {

			buf, err := ioutil.ReadFile(filePath)
			if err != nil {
				return err
			}

			var doc3 openapi3.Swagger
			if err = json.Unmarshal(buf, &doc3); err != nil {
				return err
			}

			packageName := ""

			if len(doc3.Paths) > 0 {
				for _, item := range doc3.Paths {
					if item.Get != nil && len(item.Get.Tags) > 0 {
						packageName = item.Get.Tags[0]
						break
					}
					if item.Post != nil && len(item.Post.Tags) > 0 {
						packageName = item.Post.Tags[0]
						break
					}
				}
				if packageName == "" {
					packageName = strings.Replace(strings.Replace(path.Base(filePath), "V0.json", "", -1), "V0", "", -1)
				}
			}

			cmdStr += fmt.Sprintf("mkdir -p ./selling-partner-api-go-sdk/%s/ & gosdk-codegen -generate types --package=%s -o ./selling-partner-api-go-sdk/%s/types.gen.go ./selling-partner-api-models-temp/%s \n", packageName, packageName, packageName, path.Base(filePath))
			cmdStr += fmt.Sprintf("mkdir -p ./selling-partner-api-go-sdk/%s/ & gosdk-codegen -generate client --package=%s -o ./selling-partner-api-go-sdk/%s/api.gen.go ./selling-partner-api-models-temp/%s \n", packageName, packageName, packageName, path.Base(filePath))

		}
		return nil
	}); err != nil {
		panic(err)
	}

	fmt.Println(cmdStr)

	_ = ioutil.WriteFile("./internal/gencode.sh", []byte(cmdStr), os.ModePerm)
}
