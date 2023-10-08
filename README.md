# go-i18n-ai

**go-i18n-ai is a cli tool to automatically translate your base i18n json file into selected languages using openAI api**

## Prerequisites

- golang >= v1.20.0
- openai api key

## Installation

```bash
go install github.com/mhmodo10/go-i18n-ai
```

## Help

```bash
go-i18n-ai --help
```

## Usage

```bash
go-i18n-ai --base <code of base language> --src <path to base language directory> --t <target language code> -k <your openai api key>
```

## Example

```bash
go-i18n-ai --base en --src ./lang/ --t nl -k someOpenAiApiKey
```

## Flags

| Flag    | Description                                                                | Aliases | Required |
| ------- | -------------------------------------------------------------------------- | ------- | -------- |
| base    | the language code of the base language                                     | b       | Yes      |
| source  | The source directory where the base language file is located               | src, s  | Yes      |
| targets | The language code of the target language(s) repeat to pass multiple values | t       | Yes      |
| key     | The openAI API key to be used for translations                             | k       | Yes      |
| model   | The openAI API model to be used for translations                           | m       | No       |
| out     | Where the translation files should be saved                                | o       | No       |
