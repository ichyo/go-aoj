package aoj

import (
	"encoding/xml"
	"net/url"
	"time"
)

type User struct {
	ID                 string    `xml:"id"`
	Name               string    `xml:"name"`
	Affiliation        string    `xml:"affiliation"`
	RegisterDate       time.Time `xml:"-"`
	RegisterDateUnix   uint64    `xml:"registerdate"`
	LastSubmitDate     time.Time `xml:"-"`
	LastSubmitDateUnix uint64    `xml:"lastsubmitdate"`

	Status     UserStatus `xml:"status"`
	SolvedList []Solved   `xml:"solved_list>problem"`
}

type UserStatus struct {
	Submission   uint32 `xml:"submission"`
	Solved       uint32 `xml:"solved"`
	Accepted     uint32 `xml:"accepted"`
	WrongAnswer  uint32 `xml:"wronganswer"`
	TimeLimit    uint32 `xml:"timelimit"`
	MemoryLimit  uint32 `xml:"memorylimit"`
	OutputLimit  uint32 `xml:"outputlimit"`
	RuntimeError uint32 `xml:"runtimeerror"`
	CompileError uint32 `xml:"compileerror"`
}

type Solved struct {
	UserID             string    `xml:"-"` // TODO: fix it for parsing problem.xml
	ProblemID          string    `xml:"id"`
	SubmissionDate     time.Time `xml:"-"`
	SubmissionDateUnix uint64    `xml:"submissiondate"`
	Language           string    `xml:"language"`
	CPUTime            uint32    `xml:"cputime"`
	Memory             uint32    `xml:"memory"`
	CodeSize           uint32    `xml:"code_size"`
}

const (
	user_api = "http://judge.u-aizu.ac.jp/onlinejudge/webservice/user"
)

func GetUser(id string) (User, error) {
	values := make(url.Values)
	values.Add("id", id)
	xml_data, err := APIRequest(user_api, values)
	if err != nil {
		return User{}, err
	}
	return ParseUserXML(xml_data)
}

func ParseUserXML(xml_data []byte) (User, error) {
	var user User
	if err := xml.Unmarshal(xml_data, &user); err != nil {
		return User{}, err
	}
	user.RegisterDate = UnixToTime(user.RegisterDateUnix)
	user.LastSubmitDate = UnixToTime(user.LastSubmitDateUnix)
	for i, _ := range user.SolvedList {
		user.SolvedList[i].SubmissionDate =
			UnixToTime(user.SolvedList[i].SubmissionDateUnix)
		user.SolvedList[i].UserID = user.ID
	}
	return user, nil
}
