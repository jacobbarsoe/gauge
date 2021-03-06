// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"github.com/getgauge/gauge/logger"
	"github.com/getgauge/gauge/projectInit"
	"github.com/getgauge/gauge/track"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:     "init [flags] <template>",
		Short:   "Initialize project structure in the current directory",
		Long:    `Initialize project structure in the current directory.`,
		Example: "  gauge init java",
		Run: func(cmd *cobra.Command, args []string) {
			setGlobalFlags()
			if templates {
				track.ListTemplates()
				projectInit.ListTemplates()
				return
			}
			if len(args) < 1 {
				logger.Fatalf("Error: Missing argument <template name>. To see all the templates, run 'gauge list-templates'.\n%s", cmd.UsageString())
			}
			track.ProjectInit(args[0])
			projectInit.InitializeProject(args[0])
		},
	}
	templates bool
)

func init() {
	GaugeCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&templates, "templates", "t", false, "Lists all available templates")
}
