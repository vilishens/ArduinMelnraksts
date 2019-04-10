package a_runningpoints

type WPointData struct {
	Point    string
	Active   bool
	Frozen   bool
	Descr    string
	Type     int
	CfgRun   interface{}
	CfgSaved interface{}
	Index    interface{}
	State    int
}

type pointDescr struct {
	pType  int
	pDescr string
}
