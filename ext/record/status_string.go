// generated by stringer -type=Status status.go; DO NOT EDIT

package record

import "fmt"

const _Status_name = "InProgressCanceledFailedFinishedHangup"

var _Status_index = [...]uint8{0, 10, 18, 24, 32, 38}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return fmt.Sprintf("Status(%d)", i)
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}