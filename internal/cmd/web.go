// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"net/http"
	"strings"

	"github.com/flamego/flamego"
	"github.com/urfave/cli/v2"
	log "unknwon.dev/clog/v2"

	"github.com/flamego/template"

	"github.com/wuhan005/icppp/internal/conf"
	"github.com/wuhan005/icppp/static"
	"github.com/wuhan005/icppp/templates"
)

var Web = &cli.Command{
	Name:  "web",
	Usage: "Start web server",
	Description: `icppp web server is the only thing you need to run,
and it takes care of all the other things for you`,
	Action: runWeb,
	Flags: []cli.Flag{
		intFlag("port, p", 9315, "Temporary port number to prevent conflict"),
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func runWeb(c *cli.Context) error {
	err := conf.Init(c.String("config"))
	if err != nil {
		log.Fatal("Failed to load config: %v", err)
	}

	f := flamego.Classic()

	// Embed static files.
	f.Use(flamego.Static(flamego.StaticOptions{
		FileSystem: http.FS(static.FS),
		Prefix:     "static",
	}))
	
	// Embed template files.
	templatesFS, err := template.EmbedFS(templates.FS, ".", []string{".html"})
	if err != nil {
		log.Fatal("Failed to embed templates file system: %v", err)
	}
	f.Use(template.Templater(template.Options{
		FileSystem: templatesFS,
	}))

	f.Get("/", func(ctx flamego.Context, t template.Template, data template.Data) {
		host := strings.Split(ctx.Request().Host, ":")[0]

		data["title"] = host
		data["icp"] = conf.GetICPByURL(host)
		t.HTML(http.StatusOK, "index")
	})

	f.NotFound(func(ctx flamego.Context, t template.Template, data template.Data) {
		ctx.ResponseWriter().Header().Set("Server", "Apache-Coyote/1.1")

		data["path"] = ctx.Request().URL
		t.HTML(http.StatusNotFound, "404")
	})

	f.Run("0.0.0.0", c.Int("port"))
	return nil
}
