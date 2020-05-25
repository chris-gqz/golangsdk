package instances

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

type commonResult struct {
	golangsdk.Result
}

type CreateResult struct {
	commonResult
}

type CreateGeminiDB struct {
	Id                  string         `json:"id"`
	Name                string         `json:"name"`
	Datastore           Datastore      `json:"datastore"`
	Created             string         `json:"created"`
	Status              string         `json:"status"`
	Region              string         `json:"region"`
	AvailabilityZone    string         `json:"availability_zone"`
	VpcId               string         `json:"vpc_id"`
	SubnetId            string         `json:"subnet_id"`
	SecurityGroupId     string         `json:"security_group_id"`
	Mode                string         `json:"mode"`
	Flavor              []Flavor       `json:"flavor"`
	BackupStrategy      BackupStrategy `json:"backup_strategy"`
	EnterpriseProjectId string         `json:"enterprise_project_id"`
	JobId               string         `json:"job_id"`
}

func (r CreateResult) Extract() (*CreateGeminiDB, error) {
	var response CreateGeminiDB
	err := r.ExtractInto(&response)
	return &response, err
}

type DeleteInstanceGeminiDBResult struct {
	commonResult
}

type DeleteInstanceGeminiDBResponse struct {
	JobId string `json:"job_id"`
}

func (r DeleteInstanceGeminiDBResult) Extract() (*DeleteInstanceGeminiDBResponse, error) {
	var response DeleteInstanceGeminiDBResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ListGeminiDBResult struct {
	commonResult
}

type ListGeminiDBResponse struct {
	Instances  []GeminiDBInstanceResponse  `json:"instances"`
	TotalCount int                         `json:"total_count"`
}

type GeminiDBInstanceResponse struct {
	Id                  string            `json:"id"`
	Name                string            `json:"name"`
	Status              string            `json:"status"`
	Port                int               `json:"port"`
	Mode                string            `json:"mode"`
	Region              string            `json:"region"`
	DataStore           Datastore         `json:"datastore"`
	Engine              string            `json:"engine"`
	Created             string            `json:"created"`
	Updated             string            `json:"updated"`
	DbUserName          string            `json:"db_user_name"`
	VpcId               string            `json:"vpc_id"`
	SubnetId            string            `json:"subnet_id"`
	SecurityGroupId     string            `json:"security_group_id"`
	BackupStrategy      BackupStrategy    `json:"backup_strategy"`
	PayMode             string            `json:"pay_mode"`
	MaintenanceWindow   string            `json:"maintenance_window"`
	Groups              Groups            `json:"groups"`
	EnterpriseProjectId string            `json:"enterprise_project_id"`
	TimeZone            string            `json:"time_zone"`
	Actions             []string          `json:"actions"`
}

type Groups struct {
	Id               string         `json:"id"`
	Status           string         `json:"status"`
	Volume           Volume         `json:"volume"`
	nodes     		 []Nodes 		`json:"nodes"`
}

type Volume struct {
	Size	 string 	`json:"size"`
	used 	 string		`json:"used"`
}

type Nodes struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Status           string `json:"status"`
	PrivateIp        string `json:"private_ip"`
	SpecCode 		 string `json:"spec_code"`
	AvailabilityZone string `json:"availability_zone"`
}

type GeminiDBPage struct {
	pagination.SinglePageBase
}

func (r GeminiDBPage) IsEmpty() (bool, error) {
	data, err := ExtractGeminiDBInstances(r)
	if err != nil {
		return false, err
	}
	return len(data.Instances) == 0, err
}

// ExtractGeminiDBInstances is a function that takes a ListResult and returns the services' information.
func ExtractGeminiDBInstances(r pagination.Page) (ListGeminiDBResponse, error) {
	var s ListGeminiDBResponse
	err := (r.(GeminiDBPage)).ExtractInto(&s)
	return s, err
}
