package core

type Kind int

// xodm type
const (
	// bool
	Bool Kind = iota
	// bit
	Bit
	// byte
	Byte
	// num
	Int
	Float
	// time
	Interval
	Time
	Date
	TimeStamp
	// string
	String
	// group
	Array
	Map
	Struct
	// ip
	IP
	// uuid
	Uuid
	// custom
	Custom
	// any
	Any
)
