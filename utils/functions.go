package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os/user"
)

//AllMetaData holds the stuct of all the arguments
type AllMetaData struct {
	UserName  string
	HomeDir   string
	LocalHost string
	IP        string
	Timestamp string
}

//GetUserData return username and userdir
func GetUserData() *user.User {
	user, err := user.Current()
	if err != nil {
		panic(err)

	}

	return user

}

// GetIPAdress Get preferred outbound ip of this machine
func GetIPAdress() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

//OutputJSONall  outputs allMetaData struct in JSON format
func OutputJSONall(allData *AllMetaData) {
	data, _ := json.Marshal(allData)
	err := ioutil.WriteFile("./all.json", data, 0644)

	if err != nil {
		panic(err)
	}
}

//OutputJSONip outputs ip address  in JSON format
func OutputJSONip(ip *string) {
	data, _ := json.Marshal(ip)
	err := ioutil.WriteFile("./ip.json", data, 0644)

	if err != nil {
		panic(err)
	}
}

//OutputJSONts ouputs timestamp in JSON format
func OutputJSONts(ts *string) {
	data, _ := json.Marshal(ts)
	err := ioutil.WriteFile("./timestamp.json", data, 0644)

	if err != nil {
		panic(err)
	}
}

//OutputJSONuserName ouputs userName in JSON format
func OutputJSONuserName(userName *string) {
	data, _ := json.Marshal(userName)
	err := ioutil.WriteFile("./userName.json", data, 0644)

	if err != nil {
		panic(err)
	}
}

//OutputJSONhomeDir ouputs userName in JSON format
func OutputJSONhomeDir(homeDir *string) {
	data, _ := json.Marshal(homeDir)
	err := ioutil.WriteFile("./homeDir.json", data, 0644)

	if err != nil {
		panic(err)
	}
}
