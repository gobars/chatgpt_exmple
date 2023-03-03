package main

import (
	"bufio"
	"chatgpt/Global"
	"chatgpt/Viper"
	"chatgpt/openai"
	"fmt"
	"os"
	"strings"
	"time"
)

var config openai.Request

func main() {
	Viper.Config() //加载配置文件
	ReqData := &config
	ReqData.Stop[0] = Global.GptConfig.GptRuq.Stop[0]
	ReqData.Stop[1] = Global.GptConfig.GptRuq.Stop[1]
	ReqData.N = Global.GptConfig.GptRuq.N
	ReqData.FrequencyPenalty = float64(Global.GptConfig.GptRuq.Frequency_penalty)
	ReqData.MaxTokens = Global.GptConfig.GptRuq.Max_tokens
	ReqData.Temperature = Global.GptConfig.GptRuq.Temperature
	ReqData.TopP = float64(Global.GptConfig.GptRuq.Top_p)
	Cache := "" //缓存区，方标连续对话
	inpu := bufio.NewReader(os.Stdin)
	re := false
	client := openai.NewClient(Global.GptConfig.ApiKey)
	fmt.Println("********************************************")
	fmt.Println("作者：Liujinliang")
	fmt.Printf("\n调用Gpt-3 API，开始和ChatGPT聊天吧!\n输入exit退出程序\n")
	fmt.Printf("*********************************************\n\n")

	for !re {
		fmt.Print(ReqData.Stop[0])

		input, _ := inpu.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}
		if input == "" {
			continue
		}
		if input == "clear" {
			Cache = ""
			continue
		}
		message := Cache + "\n" + config.Stop[0] + input + "\n" + config.Stop[1]

		ReqData.Model = openai.ModelGpt35Turbo
		ReqData.Messages = []*openai.Message{
			{
				Role:    openai.RoleUser,
				Content: message,
			},
		}
		Cache = message
		aiRspText := ""
		retry := 0
		for {

			resp, err := client.GetChat(ReqData)
			if err != nil {
				break
			}

			if resp.Error != nil && resp.Error.Message != "" && retry < 10 {
				fmt.Printf("ChatGPT连接出错:%v\n", resp.Error.Message)
				retry++
				time.Sleep(time.Second)
				continue
			}

			Cache += resp.Choices[0].Message.Content
			aiRspText += resp.Choices[0].Message.Content

			// 判断回答是否结束
			if resp.Choices[0].FinishReason != "length" {
				break
			}

			// 继续循环
			ReqData.Messages = []*openai.Message{
				{
					Role:    openai.RoleUser,
					Content: Cache,
				},
			}
			//Global.GptConfig.GptRuq.N++
		}
		if re {
			break
		}
		fmt.Printf("%v%v\n", strings.TrimSpace(Global.GptConfig.GptRuq.Stop[1]), aiRspText) //打印返回的文本

	}
}
