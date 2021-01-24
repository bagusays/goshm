package cmd

import (
	"errors"
	"goshm/caller"
	"goshm/executors"
	"goshm/exporter"
	"goshm/models"
	"goshm/utils"
	"strings"

	"github.com/spf13/cobra"
)

func fetchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "fetching daily idx stock",
		RunE:  fetchCmdHandler,
	}
	cmd.PersistentFlags().String("code", "", "kode emiten")
	cmd.PersistentFlags().String("date_from", "", "date from")
	cmd.PersistentFlags().String("date_to", "", "date to")
	cmd.PersistentFlags().String("output", string(exporter.OutputCSV), "date to")
	return cmd
}

func fetchCmdHandler(cmd *cobra.Command, args []string) error {
	code := cmd.Flag("code").Value.String()
	if code == "" {
		return errors.New("missing code arguments")
	}
	codes := strings.Split(code, ",")

	dateFrom := cmd.Flag("date_from").Value.String()
	if dateFrom == "" {
		return errors.New("missing date_from arguments")
	}
	dateFromTime, err := utils.FromArgToUnixDateFormatter(dateFrom)
	if err != nil {
		return err
	}

	dateTo := cmd.Flag("date_to").Value.String()
	if dateTo == "" {
		return errors.New("missing date_to arguments")
	}
	dateToTime, err := utils.FromArgToUnixDateFormatter(dateTo)
	if err != nil {
		return err
	}

	outputReq := cmd.Flag("output").Value.String()

	writer, ok := exporter.Type[models.Output(outputReq)]
	if !ok {
		return errors.New("generator not exist. Please use either csv or json")
	}

	shm := executors.New()
	ipot := caller.Ipot()
	for _, d := range codes {
		err := shm.Fetch(ipot, models.FetchArgs{
			Code:     d,
			DateFrom: dateFromTime,
			DateTo:   dateToTime,
		}, writer)
		if err != nil {
			return err
		}
	}

	println("Done!")
	return nil
}
