package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"io/ioutil"
	"os"
)

type apiR struct {
	Pujara				Provider 		`json:"provider"`
	SecurityGroup 		SecurityGroup 	`json:"SecurityGroup"`
	Vpc 				Vpc 			`json:"Vpc"`
	VSwitch 			VSwitch 		`json:"VSwitch"`
}

type Provider struct {
	Name string `json:"name"`
	Credential Credential `json:"credentials"`
}

type Credential struct {
	Access_key		string    `json:"access_key"`
	Secret_key     	string    `json:"secret_key"`
	Region 			string    `json:"region"`
}



type SecurityGroup struct {
	Type string `json:"type"`
	Property Property `json:"properties"`
}

type Property struct {
	SecurityGroupName SecurityGroupName `json:"securityGroupName"`
	VpcIdsac VpcIdsac `json:"vpcId"`
}

type SecurityGroupName struct {
	Ref string `json:"ref"`
}

type VpcIdsac struct {
	FnGetAtt       []string      `json:"fnGetAtt"`
}


type Vpc struct {
	Types string `json:"type"`
	Proper Proper `json:"properties"`
}

type Proper struct {
	CidrBlock CidrBlock `json:"cidrBlock"`
	VpcName VpcName `json:"vpcName"`
}

type CidrBlock  struct {
	Ref string `json:"ref"`
}

type VpcName  struct {
	Ref string `json:"ref"`
}

type VSwitch struct {
	Type123 string `json:"type123"`
	DependsOn string `json:"dependsOn"`
	Prop Prop `json:"properties"`
}

type Prop struct {
	Cidr	 		Cidr	 	`json:"cidrBlock"`
	VSwitchName 	VSwitchName `json:"vSwitchName"`
	ZoneId 			ZoneId 		`json:"zoneId"`
	VpcIdmis 		VpcIdmis 	`json:"vpcIdmis"`
}

type Cidr  struct {
	Ref string `json:"ref"`
}

type VSwitchName  struct {
	Ref string `json:"ref"`
}

type ZoneId  struct {
	Rsssssef string `json:"rsssssef"`
}

type VpcIdmis  struct {
	FnGetAtt       []string      `json:"fnGetAtt"`
}

func main() {
	Server()
}

func Server() {
	router := mux.NewRouter()

	router.HandleFunc("/parsejson", ParsejSON).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func ParsejSON(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(w, `error`)
		return
	}
	var cric apiR
	err = json.Unmarshal(b, &cric)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", cric)
	result := fmt.Sprintf("%v\n", cric)
	createtf(result) 

}

func createtf(temp string) {

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(temp)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

