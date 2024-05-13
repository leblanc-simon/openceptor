package project

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
	"regexp"
	"strings"

	"openceptor.eu/connection"
)

type Project struct {
	Id       string
	Name     string
	Endpoint string
	MockingRules []MockingRule
}

type MockingRule struct {
	Id		string
	Method  string
	ConditionType string
	Path string
	ResponseStatus uint8
	ResponseHeaders MockingRuleHeaders
	ResponseBody string
}

type MockingRuleHeaders map[string]interface{}
func (a MockingRuleHeaders) Value() (driver.Value, error) {
    return json.Marshal(a)
}

func (a *MockingRuleHeaders) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(b, &a)
}

func (o *Project) Load(id string) {
	db := connection.GetDbInstance(nil)

	sqlStatement := `
	SELECT p.name, p.endpoint,
			mr.id, mr.method, mr.condition_type, mr.path, mr.response_status, mr.response_headers, mr.response_body
	FROM project p
		LEFT JOIN mocking_rule mr ON p.id = mr.project_id
	WHERE p.id = $1
	ORDER BY mr.condition_type ASC
	`

	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	
	isFirstIterate := true
	for rows.Next() {
		mr := new(MockingRule)
		var name string
		var endpoint string

		err = rows.Scan(&name, &endpoint, &mr.Id, &mr.Method, &mr.ConditionType, &mr.Path, &mr.ResponseStatus, &mr.ResponseHeaders, &mr.ResponseBody)
		if err != nil {
			panic(err)
		}

		if isFirstIterate {
			o.Name = name
			o.Endpoint = endpoint

			isFirstIterate = false
		}

		o.MockingRules = append(o.MockingRules, *mr)
	}

	o.Id = id
}

func (o *Project) GetMockingRule(method string, path string) *MockingRule {
	for _, mockingRule := range o.MockingRules {
		if method != mockingRule.Method {
			continue
		}

		if mockingRule.ConditionType == "match" && path == mockingRule.Path {
			return &mockingRule
		}
		
		if mockingRule.ConditionType == "start" && strings.HasPrefix(path, mockingRule.Path) {
			return &mockingRule
		}

		if mockingRule.ConditionType != "regex" {
			continue
		}

		match, err := regexp.MatchString(mockingRule.Path, path)
		if err != nil {
			log.Fatalln(err)
		}

		if match {
			return &mockingRule
		}
	}
	
	return nil
}
