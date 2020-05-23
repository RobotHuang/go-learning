package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"strconv"
)

// MysqlConfig 配置mysql
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig 配置redis
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

// Config 总的配置累
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(filePath string, data interface{}) (err error) {
	//参数校验
	//如果不是指针类型返回错误
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")
		return
	}
	//判断是否是结构体
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct")
		return
	}
	//读取文件
	contentByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(contentByte), "\n")
	//一行一行读取
	var structName string
	for index, line := range lineSlice {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		//如果是注释就跳过
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		//判断是不是节
		if strings.HasPrefix(line, "[") {
			if line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return
			}
			// 把空格去掉取节点里面内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return
			}
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
				}
			}
		} else {
			// 是键值对
			if strings.Index(line, "=") == -1 {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type()
			if sValue.Kind() != reflect.Struct {
				err = fmt.Errorf("%s should be a struct", structName)
				return
			}
			//遍历一个结构体里面字段，找到对应的字段
			var fieldName string
			var fieldType reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)
				fieldType = field
				if field.Tag.Get("ini") == key {
					fieldName = field.Name
					break
				}
			}
			if len(fieldName) == 0 {
				continue
			}
			//给字段赋值
			fieldObj := sValue.FieldByName(fieldName)
			switch fieldType.Type.Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return
				}
				fieldObj.SetInt(valueInt)
			}
		}
	}
	return nil
}

func main() {
	var config Config
	err := loadIni("./conf.ini", &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)
}
