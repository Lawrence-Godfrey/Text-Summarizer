package processing

import (
	"context"
	"github.com/nlpodyssey/cybertron/pkg/tasks"
	"github.com/nlpodyssey/cybertron/pkg/tasks/tokenclassification"
	"text_summarizer/internal/utils"
)

// Classifier is a wrapper for the tokenclassification.Interface model.
type Classifier struct {
	model tokenclassification.Interface
}

// ExtractNamedEntities extracts named entities from the given text.
func (c *Classifier) ExtractNamedEntities(text string) ([]tokenclassification.Token, error) {
	return extractNamedEntities(text, c.model)
}

// LoadModelForNER downloads (if necessary) and loads a model for Named Entity Recognition.
func LoadModelForNER() (*Classifier, error) {
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

// extractNamedEntities extracts named entities from the given text.
func extractNamedEntities(text string, model tokenclassification.Interface) ([]tokenclassification.Token, error) {

	params := tokenclassification.Parameters{
		AggregationStrategy: tokenclassification.AggregationStrategySimple,
	}

	result, err := model.Classify(context.Background(), text, params)

	if err != nil {
		return []tokenclassification.Token{}, err
	}

	return result.Tokens, nil
}
