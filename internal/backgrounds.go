package internal

import (
	"os"
)

// ReplaceBackgroundsFromImagePath replaces all background images in a beatmap with a new image.
func ReplaceBackgroundsFromImagePath(beatmap *BeatmapFolder, imagePath string) error {
	image, err := os.ReadFile(imagePath)
	if err != nil {
		return err
	}

	return ReplaceBackgrounds(beatmap, image)
}

// ReplaceBackgrounds replaces all background images in a beatmap with a new image.
func ReplaceBackgrounds(beatmap *BeatmapFolder, image []byte) error {
	for _, file := range beatmap.ImageFiles {
		destination := beatmap.FolderPath + "/" + file

		if !hasBackup(file, beatmap) {
			if err := createBackup(file, beatmap); err != nil {
				return err
			}
		}

		if err := writeImage(image, destination); err != nil {
			return err
		}
	}
	return nil
}

// RestoreBackground restores all background images in a beatmap to their original state.
func RestoreBackground(beatmap *BeatmapFolder) error {
	for _, file := range beatmap.ImageFiles {
		imageBackupPath := beatmap.FolderPath + "/" + file + ".imagebackup"
		destination := beatmap.FolderPath + "/" + file

		if !hasBackup(file, beatmap) {
			continue
		}

		data, err := readImage(imageBackupPath)
		if err != nil {
			return err
		}

		if err := writeImage(data, destination); err != nil {
			return err
		}

		if err := removeImage(imageBackupPath); err != nil {
			return err
		}
	}
	return nil
}

func createBackup(image string, beatmap *BeatmapFolder) error {
	imageBackupPath := beatmap.FolderPath + "/" + image + ".imagebackup"
	imagePath := beatmap.FolderPath + "/" + image

	imageData, err := readImage(imagePath)
	if err != nil {
		return err
	}

	err = writeImage(imageData, imageBackupPath)
	if err != nil {
		return err
	}

	return nil
}

func readImage(image string) ([]byte, error) {
	data, err := os.ReadFile(image)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func writeImage(imageData []byte, dst string) error {
	err := os.WriteFile(dst, imageData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func removeImage(image string) error {
	err := os.Remove(image)
	if err != nil {
		return err
	}

	return nil
}

func hasBackup(image string, beatmap *BeatmapFolder) bool {
	existingMiscFiles := make(map[string]bool)
	imageBackupName := image + ".imagebackup"

	for _, file := range beatmap.MiscFiles {
		existingMiscFiles[file] = true
	}

	if _, ok := existingMiscFiles[imageBackupName]; ok {
		return true
	}

	return false
}
