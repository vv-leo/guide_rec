package utils

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"io"
)

func GetByDoubao(role string, cueWord string) (res string, err error) {
	client := arkruntime.NewClientWithApiKey(
		viper.GetString("doubao.apiKey"),
		arkruntime.WithBaseUrl("https://ark.cn-beijing.volces.com/api/v3"),
		arkruntime.WithRegion("cn-beijing"),
	)

	ctx := context.Background()

	req := model.ChatCompletionRequest{
		Model: viper.GetString("doubao.modelId"),
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleSystem,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String(role),
				},
			},
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String(cueWord),
				},
			},
		},
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Printf("standard chat error: %v\n", err)
		return
	}
	res = *resp.Choices[0].Message.Content.StringValue
	return res, err
}

func GetByDoubaoSSE(role string, prompt string, ch chan<- string) {
	client := arkruntime.NewClientWithApiKey(
		viper.GetString("doubao.apiKey"),
		arkruntime.WithBaseUrl("https://ark.cn-beijing.volces.com/api/v3"),
		arkruntime.WithRegion("cn-beijing"),
	)

	ctx := context.Background()

	req := model.ChatCompletionRequest{
		Model: viper.GetString("doubao.modelId"),
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleSystem,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String(role),
				},
			},
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String(prompt),
				},
			},
		},
	}

	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("stream chat error: %v\n", err)
		return
	}
	defer stream.Close()
	defer close(ch)
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("Stream chat error: %v\n", err)
			return
		}

		if len(recv.Choices) > 0 {
			ch <- fmt.Sprintf(recv.Choices[0].Delta.Content)
		}
	}
}
