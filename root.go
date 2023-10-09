package main

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/mhmodo10/go-i18n-ai/internal/data"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

// This file contains the logic for the CLI tool
func Run() {
	app := &cli.App{
		Name:  "go-i18n-ai",
		Usage: "Generates i18n configuration files for different languages using openAI",
		Action: func(cCtx *cli.Context) error {
			return TranslateAndSave(cCtx)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "model",
				Aliases:     []string{"m"},
				Value:       "gpt-3.5-turbo",
				Usage:       "The openAI model to use for translation (has to be supported for chat completions)",
				Destination: &Model,
			},
			&cli.StringFlag{
				Name:        "key",
				Aliases:     []string{"k"},
				Usage:       "The openAI api key to be used to send requests",
				Destination: &ApiKey,
				Required: true,
			},
			&cli.StringFlag{
				Name:     "source",
				Aliases:  []string{"src", "s"},
				Usage:    "Source directory `/some/directory/`",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "base",
				Aliases:  []string{"b"},
				Usage:    "Language code of base language `en` (base configuration file will be source <source>/<base>.json)",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     "targets",
				Aliases:  []string{"t"},
				Usage:    "Array of target locales to translate. (repeat flag to pass multiple values)",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "out",
				Aliases:     []string{"o"},
				Usage:       "Directory path to write the results to",
				Value:       "./lang",
				DefaultText: "./lang",
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

//TODO: add concurrency
//TODO: cleanup
func TranslateAndSave(cCtx *cli.Context) error {
	s := spinner.New(spinner.CharSets[39], 500*time.Millisecond)
	y := color.New(color.FgYellow)
	g := color.New(color.FgGreen)
	s.Start()
	defer s.Stop()

	//parse flags
	base := cCtx.String("base")
	src := cCtx.String("source")
	// out := cCtx.String("out")
	targets := cCtx.StringSlice("targets")
	languageCode, err := language.Parse(base)
	en := display.English.Languages()
	if err != nil {
		return err
	}

	s.Prefix = y.Sprintf("Reading base file for %s", en.Name(languageCode))
	// read base
	baseDir := path.Join(src, base+".json")
	baseData, err := ReadJson(baseDir)
	if err != nil {
		return err
	}
	baseSpec := data.LanguageSpec{
		Language:     en.Name(languageCode),
		LanguageCode: languageCode.String(),
		Data:         baseData,
	}

	// read targets
	for _, target := range targets {
		targetLanguageCode, err := language.Parse(target)
		if err != nil {
			return err
		}
		s.Prefix = y.Sprintf("Reading target file for %s", en.Name(targetLanguageCode))
		targetDir := path.Join(src, target+".json")
		targetData, err := ReadJson(targetDir)

		if err != nil {
			return err
		}
		// create language specs for targets and base
		targetSpec := data.LanguageSpec{
			Language:     en.Name(targetLanguageCode),
			LanguageCode: targetLanguageCode.String(),
			Data:         targetData,
		}

		if targetSpec.Data == nil {
			targetSpec.Data = make(map[string]interface{})
		}

		s.Prefix = y.Sprintf("Finding difference for %s", en.Name(targetLanguageCode))
		// find diffs for each target spec
		targetDiff := Diff(baseSpec, targetSpec)

		s.Prefix = y.Sprintf("Translating text from %s to %s", en.Name(languageCode), en.Name(targetLanguageCode))
		// translate diff
		translatedDiff, err := Translate(baseSpec, targetDiff)
		if err != nil {
			return err
		}

		s.Prefix = y.Sprintf("Saving results for %s", en.Name(targetLanguageCode))
		for k, v := range translatedDiff.Data {
			targetSpec.Data[k] = v
		}

		err = WriteResultToJson(targetDir, targetSpec.Data)
		if err != nil {
			return err
		}
		s.FinalMSG = g.Sprintf("Saved translation to %s\n", targetDir)
	}
	// merge translation into parent spec
	// write results to target dir
	return nil
}
