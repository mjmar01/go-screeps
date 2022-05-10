package resources

import "errors"

type ErrorCode int
type ScreepsError error

const (
	OK                        ErrorCode = 0
	ERR_NOT_OWNER             ErrorCode = -1
	ERR_NO_PATH               ErrorCode = -2
	ERR_NAME_EXISTS           ErrorCode = -3
	ERR_BUSY                  ErrorCode = -4
	ERR_NOT_FOUND             ErrorCode = -5
	ERR_NOT_ENOUGH_ENERGY     ErrorCode = -6
	ERR_NOT_ENOUGH_RESOURCES  ErrorCode = -6
	ERR_INVALID_TARGET        ErrorCode = -7
	ERR_FULL                  ErrorCode = -8
	ERR_NOT_IN_RANGE          ErrorCode = -9
	ERR_INVALID_ARGS          ErrorCode = -10
	ERR_TIRED                 ErrorCode = -11
	ERR_NO_BODYPART           ErrorCode = -12
	ERR_NOT_ENOUGH_EXTENSIONS ErrorCode = -6
	ERR_RCL_NOT_ENOUGH        ErrorCode = -14
	ERR_GCL_NOT_ENOUGH        ErrorCode = -15
)

var (
	NotOwner            ScreepsError = errors.New("ERR_NOT_OWNER")
	NoPath              ScreepsError = errors.New("ERR_NO_PATH")
	NameExists          ScreepsError = errors.New("ERR_NAME_EXISTS")
	Busy                ScreepsError = errors.New("ERR_BUSY")
	NotFound            ScreepsError = errors.New("ERR_NOT_FOUND")
	NotEnough		    ScreepsError = errors.New("ERR_NOT_ENOUGH_ENERGY|ERR_NOT_ENOUGH_RESOURCES|ERR_NOT_ENOUGH_EXTENSIONS")
	InvalidTarget       ScreepsError = errors.New("ERR_INVALID_TARGET")
	Full                ScreepsError = errors.New("ERR_FULL")
	NotInRange          ScreepsError = errors.New("ERR_NOT_IN_RANGE")
	InvalidArgs         ScreepsError = errors.New("ERR_INVALID_ARGS")
	Tired               ScreepsError = errors.New("ERR_TIRED")
	NoBodypart          ScreepsError = errors.New("ERR_NO_BODYPART")
	RclNotEnough        ScreepsError = errors.New("ERR_RCL_NOT_ENOUGH")
	GclNotEnough        ScreepsError = errors.New("ERR_GCL_NOT_ENOUGH")
	Unknown				ScreepsError = errors.New("ERR_UNKNOWN")
)

func ReturnErr(code ErrorCode) ScreepsError {
	switch code {
	case OK:
		return nil
	case ERR_NOT_OWNER:
		return NotOwner
	case ERR_NO_PATH:
		return NoPath
	case ERR_NAME_EXISTS:
		return NameExists
	case ERR_BUSY:
		return Busy
	case ERR_NOT_FOUND:
		return NotFound
	case ERR_INVALID_TARGET:
		return InvalidTarget
	case ERR_FULL:
		return Full
	case ERR_NOT_IN_RANGE:
		return NotInRange
	case ERR_INVALID_ARGS:
		return InvalidArgs
	case ERR_TIRED:
		return Tired
	case ERR_NO_BODYPART:
		return NoBodypart
	case ERR_RCL_NOT_ENOUGH:
		return RclNotEnough
	case ERR_GCL_NOT_ENOUGH:
		return GclNotEnough
	case -6:
		return NotEnough
	default:
		return Unknown
	}
}
