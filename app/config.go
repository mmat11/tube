package app

import (
	"encoding/json"
	"os"
)

// Config settings for main App.
type Config struct {
	Library     []*PathConfig      `json:"library"`
	Server      *ServerConfig      `json:"server"`
	Thumbnailer *ThumbnailerConfig `json:"thumbnailer"`
	Transcoder  *TranscoderConfig  `json:"transcoder"`
}

// PathConfig settings for media library path.
type PathConfig struct {
	Path   string `json:"path"`
	Prefix string `json:"prefix"`
}

// ServerConfig settings for App Server.
type ServerConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	StorePath string `json:"store_path"`
}

// ThumbnailerConfig settings for Transcoder
type ThumbnailerConfig struct {
	Timeout int `json:"timeout"`
}

// Sizes a map of ffmpeg -s option to suffix. e.g: hd720 -> #720p
type Sizes map[string]string

// TranscoderConfig settings for Transcoder
type TranscoderConfig struct {
	Timeout int   `json:"timeout"`
	Sizes   Sizes `json:"sizes"`
}

// DefaultConfig returns Config initialized with default values.
func DefaultConfig() *Config {
	return &Config{
		Library: []*PathConfig{
			&PathConfig{
				Path:   "videos",
				Prefix: "",
			},
		},
		Server: &ServerConfig{
			Host:      "0.0.0.0",
			Port:      8000,
			StorePath: "tube.db",
		},
		Thumbnailer: &ThumbnailerConfig{
			Timeout: 60,
		},
		Transcoder: &TranscoderConfig{
			Timeout: 300,
			Sizes:   Sizes(nil),
		},
	}
}

// ReadFile reads a JSON file into Config.
func (c *Config) ReadFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	d := json.NewDecoder(f)
	return d.Decode(c)
}
