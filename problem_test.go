package aoj

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestParseProblem(t *testing.T) {
	const test_data = "./test_data/problem.xml"
	f, err := os.Open(test_data)
	if err != nil {
		t.Fatalf("fail to read %s : %s\n", test_data, err)
	}
	xml_data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("fail to read %s : %s\n", test_data, err)
	}

	actual, err := ParseProblemXML(xml_data)
	if err != nil {
		t.Fatalf("fail to parse xml %s : %s\n", test_data, err)
	}

	expected := Problem{
		ID:          "2349",
		Name:        "World Trip",
		Available:   1,
		TimeLimit:   5,
		MemoryLimit: 262144,
		Status: ProblemStatus{
			Submission:   21,
			Accepted:     2,
			WrongAnswer:  12,
			TimeLimit:    5,
			MemoryLimit:  1,
			OutputLimit:  0,
			RuntimeError: 1,
		},
		SolvedList: []Solved{
			Solved{
				UserID:             "jag_staff",
				ProblemID:          "2349",
				SubmissionDate:     UnixToTime(1330964917292),
				SubmissionDateUnix: 1330964917292,
				Language:           "C++",
				CPUTime:            135,
				Memory:             1,
				CodeSize:           7480,
			},
			Solved{
				UserID:             "jag_staff",
				ProblemID:          "2349",
				SubmissionDate:     UnixToTime(1330964917292),
				SubmissionDateUnix: 1330964917292,
				Language:           "C++",
				CPUTime:            135,
				Memory:             1,
				CodeSize:           7480,
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpected:\t%+v\nacutual:\t%+v\n", expected, actual)
	}
}

func TestParseProblemList(t *testing.T) {
	const test_data = "./test_data/problem_list.xml"
	f, err := os.Open(test_data)
	if err != nil {
		t.Fatalf("fail to read %s : %s\n", test_data, err)
	}
	xml_data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("fail to read %s : %s\n", test_data, err)
	}

	actual, err := ParseProblemListXML(xml_data)
	if err != nil {
		t.Fatalf("fail to parse xml %s : %s\n", test_data, err)
	}

	expected := Volume{
		Problems: []Problem{
			Problem{
				ID:          "1000",
				Name:        "A + B Problem",
				TimeLimit:   1,
				MemoryLimit: 65536,
			},
			Problem{
				ID:          "1001",
				Name:        "Binary Tree Intersection And Union",
				TimeLimit:   1,
				MemoryLimit: 65536,
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpected:\t%+v\nacutual:\t%+v\n", expected, actual)
	}
}
