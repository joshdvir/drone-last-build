package main

import (
	"fmt"
	"os"

	"github.com/drone/drone-go/drone"
	"gopkg.in/urfave/cli.v2"
)

var (
	build     string
	buildDate string
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "drone_server",
			Aliases: []string{"d"},
			Value:   "",
			Usage:   "Dorne server URL like http://example.com/api",
			EnvVars: []string{"DRONE_SERVER"},
		},
		&cli.StringFlag{
			Name:    "drone_token",
			Aliases: []string{"t"},
			Value:   "",
			Usage:   "Drone.io Token",
			EnvVars: []string{"DRONE_TOKEN"},
		},
		&cli.StringFlag{
			Name:    "team_name",
			Aliases: []string{"n"},
			Value:   "",
			Usage:   "Github/Bitbucket Team/Username name like joshdvir",
		}, &cli.StringFlag{
			Name:    "repo",
			Aliases: []string{"r"},
			Value:   "",
			Usage:   "Repository name joshdvir/drone-last-build",
		},
		&cli.StringFlag{
			Name:    "branch",
			Aliases: []string{"b"},
			Usage:   "Repository branch",
		},
	}

	app.Action = func(c *cli.Context) error {
		drone_server := c.String("drone_server")
		drone_token := c.String("drone_token")
		team_name := c.String("team_name")
		repo := c.String("repo")
		branch := c.String("branch")

		if drone_server == "" || drone_token == "" {
			panic("No DRONE_TOKEN or DRONE_SERVER available")
		}

		client := drone.NewClientToken(drone_server, drone_token)
		user, err := client.Self()
		if user.Email == "" {
			fmt.Println(err)
		}

		last_build, err := client.BuildLast(team_name, repo, branch)
		if last_build.Status == "success" {
			fmt.Println(last_build.Number)
			return nil
		}

		build_list, err := client.BuildList(team_name, repo)
		for i := 0; i < len(build_list); i++ {
			if build_list[i].Branch == branch && build_list[i].Status == "success" {
				fmt.Println(build_list[i].Number)
				return nil
			}
		}
		fmt.Println(build_list[0].Branch)

		return nil
	}

	app.Run(os.Args)
}
