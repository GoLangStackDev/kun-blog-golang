package utils

import (
	"bufio"
	"bytes"
	"gopkg.in/yaml.v3"
	"io"
	"kun-blog-golang/pkg/store/models"
)

func ParseFileToPost(file io.Reader) (*models.Post, error) {
	var err error

	// 两个缓存
	var yamlStrBuffer bytes.Buffer
	var mdStrBuffer bytes.Buffer

	writeToMd := true //切换开关
	changeCount := 0  //记录切换次数，只切换2次

	spit := "---" //分割字符

	// 扫描器
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fTxt := scanner.Text()
		if fTxt == spit && changeCount < 2 { //判断是不是分割字符 同时是否已经切换够了
			writeToMd = !writeToMd
			changeCount = changeCount + 1
			continue
		}
		// 根据目前写入类型 写入内容
		if writeToMd {
			mdStrBuffer.WriteString(fTxt + "\n")
		} else {
			yamlStrBuffer.WriteString(fTxt + "\n")
		}
	}

	// 解析成配置文件
	var config *models.PostConfig
	err = yaml.Unmarshal(yamlStrBuffer.Bytes(), &config)
	if err != nil {
		return nil, err
	}

	// 组装返回的数据
	result := &models.Post{
		PostConfig: config,
		//FilePath:   path,
		Md: mdStrBuffer.String(),
		//Html:       string(blackfriday.MarkdownBasic(mdStrBuffer.Bytes())),
	}
	return result, nil
}
