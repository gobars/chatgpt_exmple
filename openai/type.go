package openai

type RoleType string
type ModelType string

type Base struct {
	Model            string    `mapstructure:"model" json:"model"`             //gpt模型
	Temperature      float64   `mapstructure:"temperature" json:"temperature"` //温度参数，用于控制生成文本的随机程度
	TopP             float64   `json:"top_p,omitempty"`
	N                int       `json:"n,omitempty"`
	MaxTokens        int       `json:"max_tokens,omitempty"`
	PresencePenalty  float64   `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64   `json:"frequency_penalty,omitempty"`
	Stop             [2]string `json:"stop,omitempty"`
	//返回信息的开头结尾

}
type Gpt3Req struct {
	Base
	Prompt string `mapstructure:"prompt" json:"prompt"` //输入给chat主体内容
}

type Request struct {
	Base
	Prompt    string      `mapstructure:"prompt" json:"prompt"` //输入给chat主体内容
	Messages  []*Message  `json:"messages"`
	Stream    bool        `json:"stream,omitempty"`
	LogitBias interface{} `json:"logit_bias,omitempty"`
	User      string      `json:"user,omitempty"`
}

type Response struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int64     `json:"created"`
	Choices []*Choice `json:"choices"`
	Usage   *Usage    `json:"usage"`
	Error   *Error    `json:"error,omitempty"`
}

type Message struct {
	Role    RoleType `json:"role,omitempty"`
	Content string   `json:"content"`
}

type Choice struct {
	Index        int      `json:"index"`
	Message      *Message `json:"message"`
	FinishReason string   `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Error struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    string `json:"code"`
}
