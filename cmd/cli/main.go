package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Lekuruu/osu-background-nuke/internal"
	"github.com/urfave/cli/v3"
)

func main() {
	replaceCmd := &cli.Command{
		Name:  "replace",
		Usage: "Replaces all osu! background images from a song folder with a new image",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "song-folder",
				Aliases: []string{"f"},
				Usage:   "Path to the song folder",
			},
			&cli.StringFlag{
				Name:    "image",
				Aliases: []string{"i"},
				Usage:   "Path to the new image",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			songFolder := cmd.String("song-folder")
			image := cmd.String("image")

			if songFolder == "" || image == "" {
				return cli.ShowCommandHelp(ctx, cmd, "osu! background nuke")
			}

			beatmaps, err := internal.ListBeatmaps(songFolder)
			if err != nil {
				return err
			}

			fmt.Printf("Replacing backgrounds for %d beatmaps...\n", len(beatmaps))

			for _, beatmap := range beatmaps {
				fmt.Printf("-> %s\n", beatmap.FolderPath)
				err := internal.ReplaceBackgroundsFromImagePath(beatmap, image)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	restoreCmd := &cli.Command{
		Name:  "restore",
		Usage: "Restores all osu! background images from a song folder",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "song-folder",
				Aliases: []string{"f"},
				Usage:   "Path to the song folder",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			songFolder := cmd.String("song-folder")

			if songFolder == "" {
				return cli.ShowCommandHelp(ctx, cmd, "osu! background nuke")
			}

			beatmaps, err := internal.ListBeatmaps(songFolder)
			if err != nil {
				return err
			}

			fmt.Printf("Restoring backgrounds for %d beatmaps...\n", len(beatmaps))

			for _, beatmap := range beatmaps {
				fmt.Printf("-> %s\n", beatmap.FolderPath)
				if err := internal.RestoreBackground(beatmap); err != nil {
					return err
				}
			}

			return nil
		},
	}

	app := &cli.Command{
		Name:     "osu-background-nuke",
		Usage:    "Manage your osu! backgrounds",
		Version:  "1.0.0",
		Commands: []*cli.Command{replaceCmd, restoreCmd},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
