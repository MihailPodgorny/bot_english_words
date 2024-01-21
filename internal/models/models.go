package models

type Word struct {
	ID               int          `json:"id"`
	WordID           int          `json:"wordId"`
	DifficultyLevel  any          `json:"difficultyLevel"`
	PartOfSpeechCode string       `json:"partOfSpeechCode"`
	Text             string       `json:"text"`
	SoundURL         string       `json:"soundUrl"`
	Transcription    string       `json:"transcription"`
	Properties       any          `json:"properties"`
	IsGold3000       bool         `json:"isGold3000"`
	Translation      *Translation `json:"translation"`
	Images           *Images      `json:"images"`
	Definition       *Definition  `json:"definition"`
	Prefix           any          `json:"prefix"`
	Mnemonics        []any        `json:"mnemonics"`
	Examples         []Example    `json:"examples"`
}

type Translation struct {
	Text string `json:"text"`
	Note any    `json:"note"`
}

type Images struct {
	URL string `json:"url"`
}

type Definition struct {
	Text     string `json:"text"`
	SoundURL string `json:"soundUrl"`
}

type Example struct {
	Text     string `json:"text"`
	SoundURL string `json:"soundUrl"`
}
