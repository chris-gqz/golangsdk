package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/geminidb/v3/instances"
	fake "github.com/huaweicloud/golangsdk/openstack/networking/v2/common"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v3/{project_id}/instances", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{ 
  "name": "test-cassandra-01", 
  "datastore": { 
    "type": "GeminiDB-Cassandra", 
    "version": "3.11", 
    "storage_engine": "rocksDB" 
  }, 
  "region": "aaa", 
  "availability_zone": "bbb", 
  "vpc_id": "674e9b42-cd8d-4d25-a2e6-5abcc565b961", 
  "subnet_id": "f1df08c5-71d1-406a-aff0-de435a51007b", 
  "security_group_id": "7aa51dbf-5b63-40db-9724-dad3c4828b58", 
  "password": "Test@123", 
  "mode": "Cluster", 
  "flavor": [ 
    { 
      "num": 3, 
      "size": 500,
      "storage": "ULTRAHIGH",
      "spec_code": "nosql.cassandra.4xlarge.4" 
    } 
  ], 
  "backup_strategy": { 
    "start_time": "08:15-09:15", 
    "keep_days": "8" 
  },
  "enterprise_project_id": "0" 
}     `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, `
{ 
  "id": "39b6a1a278844ac48119d86512e0000bin06", 
  "name": "test-cassandra-01", 
  "datastore": { 
    "type": "GeminiDB-Cassandra", 
    "version": "3.11", 
    "storage_engine": "rocksDB" 
  }, 
  "created": "2019-10-28 14:10:54",
  "status": "creating",
  "region": "aaa", 
  "availability_zone": "bbb", 
  "vpc_id": "674e9b42-cd8d-4d25-a2e6-5abcc565b961", 
  "subnet_id": "f1df08c5-71d1-406a-aff0-de435a51007b", 
  "security_group_id": "7aa51dbf-5b63-40db-9724-dad3c4828b58", 
  "mode": "Cluster", 
  "flavor": [ 
    { 
      "num": 3, 
      "size": 500,
      "storage": "ULTRAHIGH",
      "spec_code": "nosql.cassandra.4xlarge.4" 
    } 
  ], 
  "backup_strategy": { 
    "start_time": "08:15-09:15", 
    "keep_days": "8" 
  } ,
  "job_id": "c010abd0-48cf-4fa8-8cbc-090f093eaa2f",
  "enterprise_project_id": "0" 
}
    `)
	})


	instances.CreateURl

	options := instances.CreateGeminiDBOpts{
		Name: "test-cassandra-01",
		Datastore: &instances.Datastore{
			Type:          "GeminiDB-Cassandra",
			Version:       "3.11",
			StorageEngine: "rocksDB",
		},
		Region:           "aaa",
		AvailabilityZone: "bbb",
		VpcId:            "674e9b42-cd8d-4d25-a2e6-5abcc565b961",
		SubnetId:         "f1df08c5-71d1-406a-aff0-de435a51007b",
		SecurityGroupId:  "7aa51dbf-5b63-40db-9724-dad3c4828b58",
		Password:         "Test@123",
		Mode:             "Cluster",
		Flavor: []instances.Flavor{
			{
				Num:      3,
				Size:     500,
				Storage:  "ULTRAHIGH",
				SpecCode: "nosql.cassandra.4xlarge.4",
			},
		},
		BackupStrategy:&instances.BackupStrategy{
		StartTime: "08:15-09:15",
		KeepDays: 8,
	    },
		EnterpriseProjectId: "0",
	}

	actual, err := instances.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	expected := instances.CreateGeminiDB{
		Id: "39b6a1a278844ac48119d86512e0000bin06",
		Name: "test-cassandra-01",
		Datastore: instances.Datastore{
				Type: "GeminiDB-Cassandra",
				Version: "3.11",
				StorageEngine: "rocksDB",
			},
			Created: "2019-10-28 14:10:54",
			Status: "creating",
			Region: "aaa",
			AvailabilityZone: "bbb",
			VpcId: "674e9b42-cd8d-4d25-a2e6-5abcc565b961",
			SubnetId: "f1df08c5-71d1-406a-aff0-de435a51007b",
			SecurityGroupId: "7aa51dbf-5b63-40db-9724-dad3c4828b58",
			Mode: "Cluster",
			Flavor: []instances.Flavor{
				{
				Num: 3,
				Size: 500,
				Storage: "ULTRAHIGH",
				SpecCode: "nosql.cassandra.4xlarge.4",
				},
			},
			BackupStrategy: instances.BackupStrategy{
			StartTime: "08:15-09:15",
			KeepDays: 8,
			} ,
			JobId: "c010abd0-48cf-4fa8-8cbc-090f093eaa2f",
			EnterpriseProjectId: "0",
		}

	th.AssertDeepEquals(t, expected, *actual)
}
