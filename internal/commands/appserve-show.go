package commands

import (
	"errors"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	_apiClient "github.com/openinfradev/tks-api/pkg/api-client"
	"github.com/openinfradev/tks-api/pkg/domain"
	"github.com/openinfradev/tks-client/internal/helper"
	"github.com/spf13/cobra"
	"strconv"
)

func NewAppServeShowCmd(globalOpts *GlobalOptions) *cobra.Command {
	var organizationId string
	var command = &cobra.Command{
		Use:   "show",
		Short: "show app info deployed by AppServing service",
		Long: `show app info deployed by AppServing service.
	
	Example:
	tks appserve show <APP_UUID> --organization-id <ORG_ID>`,
		RunE: func(cmd *cobra.Command, args []string) error {
			appId := args[0]
			if len(appId) < 1 {
				return errors.New("app UUID is mandatory. Run 'tks appserve show --help'")
			}

			if organizationId == "" {
				return errors.New("--organization-id is mandatory param")
			}

			apiClient, err := _apiClient.New(globalOpts.ServerAddr, globalOpts.AuthToken)
			helper.CheckError(err)

			url := fmt.Sprintf("organizations/%v/app-serve-apps/%s", organizationId, appId)
			body, err := apiClient.Get(url)
			if err != nil {
				return err
			}

			type DataInterface struct {
				AppServeApp domain.AppServeApp `json:"appServeApp"`
			}
			var out = DataInterface{}
			helper.Transcode(body, &out)

			printAppServeShow(out.AppServeApp, true)

			return nil
		},
	}

	command.Flags().StringVar(&organizationId, "organization-id", "", "a organizationId")

	return command
}

func printAppServeShow(d domain.AppServeApp, long bool) {
	t := table.NewWriter()
	tTemp := table.Table{}
	tTemp.Render()
	t.Style().Options.DrawBorder = false
	t.Style().Options.SeparateColumns = false
	t.Style().Options.SeparateFooter = false
	t.Style().Options.SeparateHeader = false
	t.Style().Options.SeparateRows = false
	if long {
		t.AppendHeader(table.Row{"ID", "Version", "Status", "Strategy", "Revision",
			"Image URL", "Profile", "CREATED_AT", "UPDATED_AT"})
		for _, i := range d.AppServeAppTasks {
			tCreatedAt := helper.ParseTime(i.CreatedAt)
			var tUpdatedAt string
			if i.UpdatedAt != nil {
				tUpdatedAt = helper.ParseTime(*i.UpdatedAt)
			}
			t.AppendRow(table.Row{i.ID, i.Version, i.Status, i.Strategy, strconv.Itoa(int(i.HelmRevision)),
				i.ImageUrl, i.Profile, tCreatedAt, tUpdatedAt})
		}
	} else {
		fmt.Println("Not implemented yet.")
	}
	fmt.Println(t.Render())
}
