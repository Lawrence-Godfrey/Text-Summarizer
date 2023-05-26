package processing

import (
	"context"
	"github.com/nlpodyssey/cybertron/pkg/tasks"
	"github.com/nlpodyssey/cybertron/pkg/tasks/tokenclassification"
	"text_summarizer/internal/utils"
)

func ExtractKeywords(text string) (tokenclassification.Response, error) {
	model, err := tasks.Load[tokenclassification.Interface](&tasks.Config{
		ModelsDir: utils.GetModelsDir(),
		ModelName: "dslim/bert-base-NER",
	})

	if err != nil {
		return tokenclassification.Response{}, err
	}

	defer tasks.Finalize(model)

	params := tokenclassification.Parameters{
		AggregationStrategy: tokenclassification.AggregationStrategySimple,
	}

	result, err := model.Classify(context.Background(), text, params)

	if err != nil {
		return tokenclassification.Response{}, err
	}

	return result, nil
}
