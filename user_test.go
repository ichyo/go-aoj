package aoj

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

const (
	test_data = "./test_data/user.xml"
)

func TestParseUser(t *testing.T) {
	f, err := os.Open(test_data)
	if err != nil {
		t.Errorf("fail to read %s : %s\n", test_data, err)
	}
	xml_data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Errorf("fail to read %s : %s\n", test_data, err)
	}

	actual, err := ParseUserXML(xml_data)
	if err != nil {
		t.Errorf("fail to parse xml %s : %s\n", test_data, err)
	}

	expected := User{
		ID:                 "ichyo",
		Name:               "ichyo",
		Affiliation:        "Kyoto",
		RegisterDateUnix:   1277708050162,
		RegisterDate:       UnixToTime(1277708050162),
		LastSubmitDateUnix: 1409842187603,
		LastSubmitDate:     UnixToTime(1409842187603),
		Status: UserStatus{
			Submission:   2789,
			Solved:       996,
			Accepted:     1321,
			WrongAnswer:  930,
			TimeLimit:    234,
			MemoryLimit:  66,
			OutputLimit:  2,
			RuntimeError: 154,
			CompileError: 81,
		},
		SolvedList: []Solved{
			Solved{
				UserID:             "ichyo",
				ProblemID:          "1086",
				SubmissionDateUnix: 1364428321290,
				SubmissionDate:     UnixToTime(1364428321290),
				Language:           "C++",
				CPUTime:            14,
				Memory:             1208,
				CodeSize:           2101,
			},
			Solved{
				UserID:             "ichyo",
				ProblemID:          "1087",
				SubmissionDateUnix: 1371475860312,
				SubmissionDate:     UnixToTime(1371475860312),
				Language:           "C++",
				CPUTime:            3,
				Memory:             1312,
				CodeSize:           2737,
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpected:\t%+v\nacutual:\t%+v\n", expected, actual)
	}

}
