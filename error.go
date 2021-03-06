package infobip

var (
	// ErrForDestinationNonAlphanumeric ...
	ErrForDestinationNonAlphanumeric = Error{Err: "non-alphanumeric 'Destination' value must be between 3 and 14 numbers"}

	// ErrForFromNonAlphanumeric ...
	ErrForFromNonAlphanumeric = Error{Err: "non-alphanumeric 'From' value must be between 3 and 14 numbers"}

	// ErrForFromAlphanumeric ...
	ErrForFromAlphanumeric = Error{Err: "alphanumeric 'From' value must be between 3 and 13 characters"}

	// ErrForToNonAlphanumeric ...
	ErrForToNonAlphanumeric = Error{Err: "non-alphanumeric 'To' value must be between 3 and 14 numbers"}

	// ErrSMSStatusNotFound ...
	ErrSMSStatusNotFound = Error{Err: "SMS Status not found"}

	// ErrNoAuthentication ...
	ErrNoAuthentication = Error{Err: "Not auth format available"}
)

// Error for Infobip
type Error struct {
	Err string `json:"error,omitempty"`
}

// Error func to implements error interface
func (e Error) Error() string {
	return e.Err
}
