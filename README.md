
![Go Tests](https://github.com/Lawrence-Godfrey/Text-Summarizer/actions/workflows/tests.yml/badge.svg)


## Text Summarizer

This is a Go microservice using a gRPC API to expose various NLP functionality with the end goal of summarizing pieces of text.

The live gRPC API can be accessed via [https://summarizer.lawrences.tech](https://summarizer.lawrences.tech)

### Functionality

  - **Segmentation** - Split text into sentences and words. This can be a difficult task due to the complexity of natural language. For example, the period character is used to denote the end of a sentence, but it can also be used to denote an abbreviation or decimal number. The segmentation service will use a combination of rules and machine learning to split text into sentences and words.
  - **Named Entity Recognition** - Identify important people, places, and things in the text. For example, in the sentence "George Washington was the first president of the United States", the named entities are "George Washington" and "United States". The named entity recognition service will use machine learning to identify named entities in the text.
  - **Keyword Extraction** - Identify the main topics of a piece of text. Here we use TF-IDF (term frequency-inverse document frequency) to identify the most important words in the text. TF-IDF is a statistical measure that evaluates how important a word is to a document in a collection of documents by comparing the number of times the word appears in the document to the number of documents in the collection that contain the word.
  - **Summarization** - Create a summary of the text. This can be done in two ways: extractive summarization and abstractive summarization. Extractive summarization involves selecting key sentences from the original text, while abstractive summarization involves generating new sentences that convey the main points.
    - **Extractive Summarization** - This will be done using a combination of the above services.
    - **Abstractive Summarization** - This will be done using an existing model, such as [BART](https://arxiv.org/abs/1910.13461) or [T5](https://arxiv.org/abs/1910.10683).

### Usage
To be able to use the API, you will need to download the protobuf file (api/proto/summarizer.proto) and generate the gRPC client code for your language of choice. To generate the Go client code you can simply use the script in this repository:
```bash
./scripts/genproto.sh
```

### Resources
 - [A list of Go NLP libraries](https://awesome-go.com/natural-language-processing/)
 - [sentences - A Go library for sentence segmentation](https://github.com/neurosnap/sentences)
 - [gse - An Efficient Go NLP library for text segmentation](https://github.com/go-ego/gse)
 - [segment - A Go library for performing Unicode Text Segmentation](https://github.com/blevesearch/segment?utm_campaign=awesomego)
 - [prose - A library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. Not maintained.](https://github.com/jdkato/prose)
 - [tokenizer - Tokenizer which can be used with pre-trained models. Inspired by huggingface/tokenizers](github.com/sugarme/tokenizer)
 - [huggingface/tokenizers - Fast State-of-the-Art Tokenizers optimized for Research and Production. Doesn't currently support Go.](https://github.com/huggingface/tokenizers)