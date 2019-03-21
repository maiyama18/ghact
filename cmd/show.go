package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var period string

type Activity struct {
	Date  string
	Count int
}

func (a *Activity) String() string {
	date := strings.Split(a.Date, "-")
	month := date[1]
	day := date[2]

	return fmt.Sprintf("%s/%s: %d contributions", month, day, a.Count)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show number of activities",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("show requires github username")
		}
		return cobra.OnlyValidArgs(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if period != "week" && period != "month" && period != "year" {
			exit(1, "period option should be one of week/month/year")
		}

		user := args[0]
		html, err := fetchActivitiesHTML(user)
		if err != nil {
			exit(1, err.Error())
		}

		activities := extractActivities(html, period)
		for _, activity := range activities {
			fmt.Println(activity)
		}
	},
}

func init() {
	showCmd.Flags().StringVarP(&period, "period", "p", "week", "period to show contributions (week/month/year)")
}

func fetchActivitiesHTML(user string) (string, error) {
	url := fmt.Sprintf("https://github.com/users/%s/contributions", user)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func extractActivities(html, period string) []*Activity {
	activities := make([]*Activity, 0)

	for _, line := range strings.Split(html, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "<rect") {
			dateRegex := regexp.MustCompile(`data-date="([\d\-]*)"`)
			countRegex := regexp.MustCompile(`data-count="(\d*)"`)

			dateGroup := dateRegex.FindStringSubmatch(line)
			countGroup := countRegex.FindStringSubmatch(line)
			if len(dateGroup) == 0 || len(countGroup) == 0 {
				continue
			}

			date := dateGroup[1]
			count, err := strconv.Atoi(countGroup[1])
			if err != nil {
				continue
			}

			activity := &Activity{date, count}

			activities = append(activities, activity)
		}
	}

	switch period {
	case "week":
		activities = activities[len(activities)-7:]
	case "month":
		activities = activities[len(activities)-30:]
	}

	return activities
}
