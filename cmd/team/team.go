package team

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"github.com/youssame/assistant-cli/internal"
	"log"
	"os"
	"text/tabwriter"
)

var Cmd = &cobra.Command{
	Use:     "team",
	Version: "0.1.0",
	Short:   "Manage my reports",
}

func state() {
	res, err := internal.ElasticsearchClient.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	internal.PrintMessage(res)
	defer res.Body.Close()
}
func addReport(doc map[string]string) {
	// Serialize the document to JSON
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(doc)

	// Index the document
	res, err := internal.ElasticsearchClient.Index("reports", &buf)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}
	defer res.Body.Close()
}
func eraseAllReports() {
	query := `{
		"query": {
			"match_all": {}
		}
	}`

	// Convert query to a bytes buffer
	var buf bytes.Buffer
	buf.WriteString(query)

	// Use the Delete By Query API
	res, err := internal.ElasticsearchClient.DeleteByQuery(
		[]string{"reports"}, // Replace with your index name
		&buf,
		internal.ElasticsearchClient.DeleteByQuery.WithContext(context.Background()),
	)
	if err != nil {
		log.Fatalf("Error executing DeleteByQuery: %s", err)
	}
	defer res.Body.Close()
}
func searchInReports(q string) {
	query := `{
  "query": {
    "query_string": {
      "query": "*` + q + `*"
    }
  },
  "size": 10,
  "from": 0,
  "sort": []
}`

	// Convert query to a bytes buffer
	var buf bytes.Buffer
	buf.WriteString(query)

	// Use the Delete By Query API
	res, err := internal.ElasticsearchClient.Search(
		internal.ElasticsearchClient.Search.WithIndex("reports"),
		internal.ElasticsearchClient.Search.WithBody(&buf),
	)
	if err != nil {
		log.Fatalf("Error executing DeleteByQuery: %s", err)
	}
	printReportsFromRes(res)
	defer res.Body.Close()
}

func printReportsFromRes(res *esapi.Response) {
	// Parse the search results
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response: %s", err)
	}

	// Extract hits
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})

	// Print results in a table format
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Name\tJob\tOrg\tEmail\tDotted Line Manager\tLOB\tTeam\tVP\t")

	for _, hit := range hits {
		hitMap := hit.(map[string]interface{})
		source := hitMap["_source"].(map[string]interface{}) // Extract _source

		// Extract and print each property
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n",
			source["name"],
			source["job"],
			source["org"],
			source["email"],
			source["dotted_line_manager"],
			source["lob"],
			source["team"],
			source["vp"],
		)
	}

	w.Flush()
}
func sync() {

	f, err := excelize.OpenFile("/Users/youssefameachaq/OneDrive - Oracle Corporation/Wim's org MADC.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("reports")
	if err != nil {
		log.Fatalf("Error reading from the excel file: %s", err)
		return
	}
	if err != nil {
		return
	}
	eraseAllReports()
	var firstLine = false
	for _, row := range rows {
		if !firstLine {
			firstLine = true
			continue
		}
		doc := map[string]string{
			"name":                row[0],
			"job":                 row[1],
			"org":                 row[2],
			"email":               row[3],
			"dotted_line_manager": row[4],
			"lob":                 row[5],
			"team":                row[6],
			"vp":                  row[7],
		}
		addReport(doc)
	}
	internal.SuccessAlert()
}

func init() {
	health := &cobra.Command{
		Use:     "health",
		Version: "0.1.0",
		Short:   "Check the elasticsearch health",
		Run: func(cmd *cobra.Command, args []string) {
			state()
		},
	}
	syncCmd := &cobra.Command{
		Use:     "sync",
		Version: "0.1.0",
		Short:   "Sync the team from the excel to the elasticsearch index",
		Run: func(cmd *cobra.Command, args []string) {
			sync()
		},
	}
	searchCmd := &cobra.Command{
		Use:     "search",
		Version: "0.1.0",
		Short:   "Sync the team from the excel to the elasticsearch index",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatalln("Empty argument")
			}
			searchInReports(args[0])
		},
	}
	Cmd.AddCommand(health, syncCmd, searchCmd)
}
