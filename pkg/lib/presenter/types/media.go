package types

type Source string

const (
	URL  Source = "URL"
	PATH Source = "PATH"
)

type Kind string

const (
	VIDEO Kind = "video"
	PHOTO Kind = "photo"
	AUDIO Kind = "audio"
)

type Media struct {
	Source     string
	SourceType Source
	Kind       Kind
}
