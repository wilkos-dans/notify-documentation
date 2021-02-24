package main

import (
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func resetVolatileFolder(folderPath string) error {
	var err error
	err = os.RemoveAll(folderPath)
	if err != nil {
		return err
	}
	err = os.MkdirAll(folderPath, os.ModePerm)
	return err
}

func readBytesFromFileOrUrl(source string) ([]byte,error) {
	var content []byte
	var err error
	if strings.HasPrefix(source,"http") {
		zapLogger.Debug("Loading from source",zap.String("URL",source))
		response, err := http.Get(source)
		if err != nil {
			zapLogger.Error(err.Error())
			return content,err
		}
		defer response.Body.Close()
		content, err = ioutil.ReadAll(response.Body)
		if err != nil {
			zapLogger.Error(err.Error())
			return content,err
		}
	} else {
		zapLogger.Debug("Loading from source",zap.String("Local file path",source))
		return ioutil.ReadFile(source)
	}
	return content,err
}