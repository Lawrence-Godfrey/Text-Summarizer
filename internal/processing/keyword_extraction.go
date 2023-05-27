package processing

import (
	"context"
	"github.com/nlpodyssey/cybertron/pkg/tasks"
	"github.com/nlpodyssey/cybertron/pkg/tasks/tokenclassification"
	"text_summarizer/internal/utils"
)

type Classifier struct {
	model tokenclassification.Interface
}

func (c *Classifier) ExtractKeywords(text string) ([]tokenclassification.Token, error) {
	return extractKeywords(text, c.model)
}

func LoadModel() (*Classifier, error) {
	model, err := tasks.Load[tokenclassification.Interface](&tasks.Config{
		ModelsDir: utils.GetModelsDir(),
		ModelName: "dslim/bert-base-NER",
	})

	if err != nil {
		return nil, err
	}

	defer tasks.Finalize(model)

	return &Classifier{model: model}, nil
}

func extractKeywords(text string, model tokenclassification.Interface) ([]tokenclassification.Token, error) {

	params := tokenclassification.Parameters{
		AggregationStrategy: tokenclassification.AggregationStrategySimple,
	}

	result, err := model.Classify(context.Background(), text, params)

	if err != nil {
		return []tokenclassification.Token{}, err
	}

	return result.Tokens, nil
}
