package utils

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func CreateProblemsTable(db *sql.DB) (sql.Result, error) {
	createProblemsTable := `CREATE TABLE IF NOT EXISTS problems (
	id INT PRIMARY KEY,
	title TEXT NOT NULL,
	titleSlug TEXT NOT NULL,
	difficulty TEXT NOT NULL,
	isPremium BOOLEAN NOT NULL,
	topic TEXT NOT NULL
	);`
	return db.Exec(createProblemsTable)
}

type Tag struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Question struct {
	Difficulty string `json:"difficulty"`
	Id         int    `json:"frontendQuestionId,string"`
	IsPremium  bool   `json:"paidOnly"`
	Title      string `json:"title"`
	TitleSlug  string `json:"titleSlug"`
	Tags       []Tag  `json:"topicTags"`
}

type Questions struct {
	Questions []Question `json:"questions"`
}

type problemSetQuestionList struct {
	QuestionList Questions `json:"problemsetQuestionList"`
}

type data struct {
	Data problemSetQuestionList `json:"data"`
}

func (q *Questions) writeJson(w io.Writer) {

	_, err := w.Write([]byte("[\n"))
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatal("unable to write")
	}

	for _, qu := range q.Questions {
		// TODO: add tags
		out := fmt.Sprintf(`{"id":"%d","title":"%s","titleSlug":"%s","difficulty":"%s","isPremium":"%t"},`,
			qu.Id, qu.Title, qu.TitleSlug, qu.Difficulty, qu.IsPremium,
		)
		out += "\n"

		_, err := w.Write([]byte(out))
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}
	_, err = w.Write([]byte("]\n"))
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func (q *Questions) writeCsv(w io.Writer) {

	_, err := w.Write([]byte("id,title,titleSlug,difficulty,isPremium\n"))
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	for _, qu := range q.Questions {
		// TODO: add tags
		out := fmt.Sprintf(`%d,"%s","%s","%s",%t`,
			qu.Id, qu.Title, qu.TitleSlug, qu.Difficulty, qu.IsPremium,
		)
		out += "\n"

		_, err := w.Write([]byte(out))
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}

}

func GetData(numQuestions int, w io.Writer, writeJson bool) {

	variables := fmt.Sprintf(`{"categorySlug": "", "skip": 0, "limit": %d, "filters": {}}`, numQuestions)
	jsonData := map[string]string{
		"query": `
			query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
			  problemsetQuestionList: questionList(
				categorySlug: $categorySlug
				limit: $limit
				skip: $skip
				filters: $filters
			  ) {
				questions: data {
				  difficulty
				  frontendQuestionId: questionId
				  paidOnly: isPaidOnly
				  status
				  title
				  titleSlug
				  topicTags {
					name
					id
					slug
				  }
				}
			  }
			}
	       `,
		"variables": variables,
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Printf("unable to marshal %s", err)
		os.Exit(1)
	}

	request, err := http.NewRequest("POST", "https://leetcode.com/graphql", bytes.NewBuffer(jsonValue))

	request.Header.Add("referrer", "https://leetcode.com/uwi/")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("cookie", "csrftoken=072x9rHWNg1kkNXwiJ3BH2RNd8sKxZ5jbByD8upwAX2sq9sLHxyLko43D4yq4UE7;  __cfduid=d200def86257d1bc29ba00c6c8b3ad20c1610345109;csrftoken=0HAcMpAuL6Fte9Af18y98z2T3gvV84PHTy673AOkl3q8mGebGIOHHgLwaUzBo11x")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()

	// io.Copy(w, response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("unable to read response %s\n", err)
	}

	var d data
	err = json.Unmarshal(body, &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatal("Unable to unmarshal")
	}

	if writeJson {
		d.Data.QuestionList.writeJson(w)
	} else {
		d.Data.QuestionList.writeCsv(w)
	}

}
