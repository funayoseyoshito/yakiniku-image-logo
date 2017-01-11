package lib

import (
	"fmt"

	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

//Config 設定ファイルのグローバル変数
var Config Configs

//Configs 設定ファイル
type Configs struct {
	Database DatabaseConfig
	Logo     LogoConfig
	Cooking  CookingConfig
	Other    OtherConfig
}

//DatabaseConfig database 設定ファイル
type DatabaseConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

//LogoConfig logo画像の設定
type LogoConfig struct {
	MediumName string
	LargeName  string
	OriginName string
}

//CookingConfig 料理画像の設定
type CookingConfig struct {
	OriginID int
	LargeID  int
	MediumID int
	SmallID  int
	MicroID  int
}

//OtherConfig その他画像の設定
type OtherConfig struct {
	OriginID int
	LargeID  int
	MediumID int
	SmallID  int
	MicroID  int
}

//GetMediumLogoPath mediumLogo画像のフルパスを取得する
func (con Configs) GetMediumLogoPath() string {
	basePath, _ := os.Getwd()
	return filepath.Join(basePath, "assets", con.Logo.MediumName)
}

//GetLargeLogoPath largeLogo画像のフルパスを取得する
func (con Configs) GetLargeLogoPath() string {
	basePath, _ := os.Getwd()
	return filepath.Join(basePath, "assets", con.Logo.LargeName)
}

//GetOriginLogoPath originLogo画像のフルパスを取得する
func (con Configs) GetOriginLogoPath() string {
	basePath, _ := os.Getwd()
	return filepath.Join(basePath, "assets", con.Logo.OriginName)
}

func init() {
	_, err := toml.DecodeFile("./config.toml", &Config)
	if err != nil {
		fmt.Println(err)
		panic("設定ファイル読み込み失敗")
	}
}
