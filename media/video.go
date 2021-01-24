package media

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/dhowden/tag"
)

// Video represents metadata for a single video.
type Video struct {
	ID          string
	Title       string
	Album       string
	Description string
	Thumb       []byte
	ThumbType   string
	Modified    string
	Size        int64
	Path        string
	Timestamp   time.Time

	Views int64
}

// ParseVideo parses a video file's metadata and returns a Video.
func ParseVideo(p *Path, name string) (*Video, error) {
	pth := path.Join(p.Path, name)
	f, err := os.Open(pth)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	size := info.Size()
	timestamp := info.ModTime()
	modified := timestamp.Format("2006-01-02 03:04 PM")
	// ID is name without extension
	idx := strings.LastIndex(name, ".")
	if idx == -1 {
		idx = len(name)
	}
	id := name[:idx]
	if len(p.Prefix) > 0 {
		// if there's a prefix prepend it to the ID
		id = path.Join(p.Prefix, name[:idx])
	}
	m, err := tag.ReadFrom(f)
	if err != nil {
		return nil, err
	}
	title := m.Title()
	// Default title is filename
	if title == "" {
		title = strings.TrimSuffix(name, filepath.Ext(name))
	}
	v := &Video{
		ID:          id,
		Title:       title,
		Album:       m.Album(),
		Description: m.Comment(),
		Modified:    modified,
		Size:        size,
		Path:        pth,
		Timestamp:   timestamp,
	}
	// Add thumbnail (if exists)
	pic := m.Picture()
	if pic != nil {
		v.Thumb = pic.Data
		v.ThumbType = pic.MIMEType
	} else {
		thumbFn := fmt.Sprintf("%s.jpg", strings.TrimSuffix(pth, filepath.Ext(pth)))
		if FileExists(thumbFn) {
			data, err := ioutil.ReadFile(fmt.Sprintf("%s.jpg", strings.TrimSuffix(pth, filepath.Ext(pth))))
			if err != nil {
				return nil, err
			}
			v.Thumb = data
			v.ThumbType = "image/jpeg"
		}
	}

	return v, nil
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
