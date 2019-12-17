package typeman

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
	// any
	Any
)

type Type interface {
	ToString() (string, error)
	Parse() interface{}
	IsValid(interface{}) bool
}
