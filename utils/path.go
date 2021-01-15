package utils

import "os"

func InitStatic() error {
	assetsPath := GetConf().String("download::assetspath")
	err := EnsurePath(assetsPath)
	if err != nil {
		return err
	}

	return nil
}

func EnsurePath(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, 0744)
	}
	if err != nil {
		GetLog().Sugar().Errorf("helpers.MakeSurePath: MkdirAll error=%v, path=%v", err, path)
	}

	return err
}
