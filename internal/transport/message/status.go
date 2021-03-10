package message

type MessageStatus uint8

const (
	Warn                   MessageStatus = iota // 0000
	Successful                                  // 0001
	FailedConnection                            // 0010
	Error                                       // 0011
	UnauthorizedConnection                      // 0100
	NotFound                                    // 0101
	InvalidCommand                              // 0110
	InvalidQuery                                // 0111
	ConnectionTimeout                           // 1000
	_
	TooManyRequest // 1001
)

func (s MessageStatus) String() string {
	return []string{
		"WARNING",
		"SUCCESSFUL",
		"FAILED_CONNECTION",
		"ERROR",
		"UNAUTHORIZED_CONNECTION",
		"NOT_FOUND",
		"INVALID_COMMAND",
		"INVALID_QUERY",
		"CONNECTION_TIMEOUT",
		"",
		"TOO_MANY_REQUEST"}[s]
}
