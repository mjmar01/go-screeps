package resources

import "errors"

type ErrorCode int

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
	NotOwner      error = errors.New("ERR_NOT_OWNER")
	NoPath        error = errors.New("ERR_NO_PATH")
	NameExists    error = errors.New("ERR_NAME_EXISTS")
	Busy          error = errors.New("ERR_BUSY")
	NotFound      error = errors.New("ERR_NOT_FOUND")
	NotEnough     error = errors.New("ERR_NOT_ENOUGH_ENERGY|ERR_NOT_ENOUGH_RESOURCES|ERR_NOT_ENOUGH_EXTENSIONS")
	InvalidTarget error = errors.New("ERR_INVALID_TARGET")
	Full          error = errors.New("ERR_FULL")
	NotInRange    error = errors.New("ERR_NOT_IN_RANGE")
	InvalidArgs   error = errors.New("ERR_INVALID_ARGS")
	Tired         error = errors.New("ERR_TIRED")
	NoBodypart    error = errors.New("ERR_NO_BODYPART")
	RclNotEnough  error = errors.New("ERR_RCL_NOT_ENOUGH")
	GclNotEnough  error = errors.New("ERR_GCL_NOT_ENOUGH")
	Unknown       error = errors.New("ERR_UNKNOWN")
)

func returnErr(code int) error {
	switch ErrorCode(code) {
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
