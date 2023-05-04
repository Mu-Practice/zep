package llms

import (
	"context"
	"fmt"

	"github.com/danielchalef/zep/pkg/models"
	"github.com/sashabaranov/go-openai"
)

func EmbedMessages(
	ctx context.Context,
	appState *models.AppState,
	messageContents *[]string,
) (*[]openai.Embedding, error) {

	var embeddingModel openai.EmbeddingModel
	switch appState.Embeddings.Model {
	case "AdaEmbeddingV2":
		embeddingModel = openai.AdaEmbeddingV2
	default:
		return nil, NewLLMError(fmt.Sprintf("invalid embedding model: %s",
			appState.Embeddings.Model), nil)
	}

	embeddingRequest := openai.EmbeddingRequest{
		Input: *messageContents,
		Model: embeddingModel,
		User:  "zep_user",
	}

	response, err := appState.OpenAIClient.CreateEmbeddings(ctx, embeddingRequest)
	if err != nil {
		return nil, NewLLMError("error while creating embedding", err)
	}

	return &response.Data, nil
}