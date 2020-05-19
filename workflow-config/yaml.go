package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.AddConfigPath("../../../../")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func GetString(key string) string {
	value := viper.GetString(key)
	//判断是否需要从环境变量读取
	exp := regexp.MustCompile(`\${(.*?)}`)
	results := exp.FindAllStringSubmatch(value, -1)
	if len(results) > 0 { //需要从环境变量读取
		for _, result := range results {
			defaultValue := ""
			valueArray := strings.Split(result[1], ":")
			if len(valueArray) >= 2 { //有默认值
				defaultValue = strings.Join(valueArray[1:], ":")
			}
			if len(valueArray) > 0 {
				//获取环境变量
				if env := os.Getenv(valueArray[0]); env != "" {
					defaultValue = env
				}
			}
			value = strings.Replace(value, result[0], defaultValue, -1)
		}
	}
	return value
}

func GetInt(key string) int {
	valueStr := GetString(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func GetInt64(key string) int64 {
	valueStr := GetString(key)
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func GetInt32(key string) int32 {
	valueStr := GetString(key)
	value, err := strconv.ParseInt(valueStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int32(value)
}

func GetFloat64(key string) float64 {
	valueStr := GetString(key)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func GetFloat32(key string) float32 {
	valueStr := GetString(key)
	value, err := strconv.ParseFloat(valueStr, 32)
	if err != nil {
		fmt.Println(err)
	}
	return float32(value)
}

func GetBool(key string) bool {
	valueStr := GetString(key)
	if strings.ToUpper(valueStr) == "TRUE" {
		return true
	}
	return false
}
