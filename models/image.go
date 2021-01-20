package models

import (
	"io/ioutil"
	"wedding-invitation-go/utils"
)

type ImageResponse struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	Category string `json:"category"`
}

type Image struct {
}

func (i *Image) GetIndexImagesByUserId(userId int64) []ImageResponse {
	var results []ImageResponse

	assetsPath := utils.GetConf().String("download::assetspath")
	imageFileList := getFileList(assetsPath + "/comer" + utils.Int64ToString(userId))

	for _, imageFile := range imageFileList {
		results = append(results, ImageResponse{imageFile, "/" + imageFile, "index"})
	}

	return results
}

func getFileList(path string) []string {
	var fileList []string

	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
		}
	}

	return fileList
}
