package main

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func existFile(name string) bool {
	stat, err := os.Stat(name)
	return err == nil && stat.Mode().IsRegular()
}

func existDir(name string) bool {
	stat, err := os.Stat(name)
	return err == nil && stat.IsDir()
}

func performEach(name string) []string {
	return []string{name, strconv.FormatBool(existFile(name)), strconv.FormatBool(existDir(name))}
}

func render(data [][]string) int {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"File name", "IsFile", "IsDir"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
	return 0
}

func goMain(args []string) int {
	data := [][]string{}
	for _, arg := range args {
		data = append(data, performEach(arg))
	}
	return render(data)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
