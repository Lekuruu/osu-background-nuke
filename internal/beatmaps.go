package internal

import (
	"errors"
	"os"
	"strings"
)

type BeatmapFolder struct {
	FolderPath      string
	ImageFiles      []string
	VideoFiles      []string
	AudioFiles      []string
	BeatmapFiles    []string
	StoryboardFiles []string
	MiscFiles       []string
}

// ListBeatmaps returns a list of beatmaps in the given song folder.
func ListBeatmaps(songFolder string) ([]*BeatmapFolder, error) {
	files, err := os.ReadDir(songFolder)
	if err != nil {
		return nil, err
	}

	var beatmaps []*BeatmapFolder

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		beatmap, err := GetBeatmapFromFolder(songFolder + "/" + file.Name())
		if err != nil {
			continue
		}

		beatmaps = append(beatmaps, beatmap)
	}

	return beatmaps, nil
}

// GetBeatmapFromFolder returns a BeatmapFolder struct from the given folder path.
func GetBeatmapFromFolder(folderPath string) (*BeatmapFolder, error) {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	beatmap := &BeatmapFolder{FolderPath: folderPath}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := strings.ToLower(file.Name())

		if isImageFile(filename) {
			beatmap.ImageFiles = append(beatmap.ImageFiles, file.Name())
			continue
		}

		if isVideoFile(filename) {
			beatmap.VideoFiles = append(beatmap.VideoFiles, file.Name())
			continue
		}

		if isAudioFile(filename) {
			beatmap.AudioFiles = append(beatmap.AudioFiles, file.Name())
			continue
		}

		if isBeatmapFile(filename) {
			beatmap.BeatmapFiles = append(beatmap.BeatmapFiles, file.Name())
			continue
		}

		if isStoryboardFile(filename) {
			beatmap.StoryboardFiles = append(beatmap.StoryboardFiles, file.Name())
			continue
		}

		beatmap.MiscFiles = append(beatmap.MiscFiles, file.Name())
	}

	if len(beatmap.BeatmapFiles) == 0 {
		return nil, errors.New("beatmap folder does not contain a beatmap file")
	}

	return beatmap, nil
}

func isImageFile(file string) bool {
	fileExtensions := []string{
		".jpg", ".jpeg", ".png",
		".bmp", ".gif", ".webp",
	}

	for _, ext := range fileExtensions {
		if strings.HasSuffix(file, ext) {
			return true
		}
	}

	return false
}

func isVideoFile(file string) bool {
	fileExtensions := []string{
		".mp4", ".avi", ".mkv", ".m4v",
		".flv", ".mov", ".wmv",
	}

	for _, ext := range fileExtensions {
		if strings.HasSuffix(file, ext) {
			return true
		}
	}

	return false
}

func isAudioFile(file string) bool {
	fileExtensions := []string{
		".mp3", ".wav", ".ogg",
		".flac", ".m4a", ".wma",
	}

	for _, ext := range fileExtensions {
		if strings.HasSuffix(file, ext) {
			return true
		}
	}

	return false
}

func isBeatmapFile(file string) bool {
	if strings.HasSuffix(file, ".osu") {
		return true
	}
	return false
}

func isStoryboardFile(file string) bool {
	if strings.HasSuffix(file, ".osb") {
		return true
	}
	return false
}
