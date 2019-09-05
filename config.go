package main

type Hosts struct {
	Hosts []Host
}

type Host struct {
	Address  string `json:"address"`
	User     string `json:"user"`
	Password string `json:"password"`
}
