package aoj

import (
	"encoding/xml"
	"net/url"
	"strconv"
)

type Problem struct {
	ID          string        `xml:"id"`
	Name        string        `xml:"name"`
	Available   uint8         `xml:"available"`
	TimeLimit   uint32        `xml:"problemtimelimit"`
	MemoryLimit uint32        `xml:"problemmemorylimit"`
	Status      ProblemStatus `xml:"status"`
	SolvedList  []Solved      `xml:"solved_list>user"`
}

type ProblemStatus struct {
	Submission   uint32 `xml:"submission"`
	Accepted     uint32 `xml:"accepted"`
	WrongAnswer  uint32 `xml:"wronganswer"`
	TimeLimit    uint32 `xml:"timelimit"`
	MemoryLimit  uint32 `xml:"memorylimit"`
	OutputLimit  uint32 `xml:"outputlimit"`
	RuntimeError uint32 `xml:"runtimeerror"`
}

const (
	problem_api = "http://judge.u-aizu.ac.jp/onlinejudge/webservice/problem"
)

func GetProblem(id string, status bool) (Problem, error) {
	values := make(url.Values)
	values.Add("id", id)
	values.Add("status", strconv.FormatBool(status))
	xml_data, err := APIRequest(problem_api, values)
	if err != nil {
		return Problem{}, err
	}
	return ParseProblemXML(xml_data)
}

func ParseProblemXML(xml_data []byte) (Problem, error) {
	var p Problem
	if err := xml.Unmarshal(xml_data, &p); err != nil {
		return Problem{}, err
	}
	for i, _ := range p.SolvedList {
		// FIXME: It is because sharing Solved Type wich User Type
		p.SolvedList[i].UserID, p.SolvedList[i].ProblemID =
			p.SolvedList[i].ProblemID, p.ID
		p.SolvedList[i].SubmissionDate = UnixToTime(p.SolvedList[i].SubmissionDateUnix)
	}
	return p, nil
}
