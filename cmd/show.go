package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Activity struct {
	Date string
	Count int
}

var showCmd = &cobra.Command{
	Use: "show",
	Short: "Show number of activities",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("username should be provided as an arg")
			return
		}

		user := args[0]
		html, err := fetchActivitiesHTML(user)
		if err != nil {
			fmt.Println(err)
		}

		activities := extractActivities(html)
		for _, activity := range activities {
			fmt.Println(activity)
		}
	},
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

func extractActivities(html string) []*Activity {
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

	return activities
}
