package utils

const (
	LESS_THAN_ZERO int32 = iota + 1
	EQUAL_ZERO
	EQUAL_ONE
	GREATER_THAN_ZERO
	GREATER_THAN_ONE
	GREATER_THAN_TWO
	GREATER_THAN_THREE
	GREATER_THAN_FOUR
	GREATER_THAN_TEN
	ZERO_TO_TEN
	ZERO_TO_TWENTY
	TEN_TO_TWENTY
	TWENTY_TO_FIFTY
	ZERO_TO_FIFTY
	ZERO_TO_HUNDRED
	GREATER_THAN_FIFTY
	GREATER_THAN_TWENTY
	GREATER_THAN_HUNDRED
)

func MapQuantityFromTo(condition int64, res map[string]interface{}, key string) map[string]interface{} {
	switch int32(condition) {
	case EQUAL_ZERO:
		res["to_"+key] = int64(0)
	case EQUAL_ONE:
		res["to_"+key] = int64(1)
	case GREATER_THAN_ZERO:
		res["from_"+key] = int64(1)
	case GREATER_THAN_ONE:
		res["from_"+key] = int64(2)
	case GREATER_THAN_TWO:
		res["from_"+key] = int64(3)
	case GREATER_THAN_THREE:
		res["from_"+key] = int64(4)
	case GREATER_THAN_FOUR:
		res["from_"+key] = int64(5)
	case GREATER_THAN_TEN:
		res["from_"+key] = int64(11)
	case GREATER_THAN_TWENTY:
		res["from_"+key] = int64(21)
	case GREATER_THAN_FIFTY:
		res["from_"+key] = int64(51)
	case GREATER_THAN_HUNDRED:
		res["from_"+key] = int64(101)
	case ZERO_TO_TEN:
		res["to_"+key] = int64(10)
	case ZERO_TO_TWENTY:
		res["to_"+key] = int64(20)
	case ZERO_TO_FIFTY:
		res["to_"+key] = int64(50)
	case ZERO_TO_HUNDRED:
		res["to_"+key] = int64(100)
	case TEN_TO_TWENTY:
		res["from_"+key] = int64(10)
		res["to_"+key] = int64(20)
	case TWENTY_TO_FIFTY:
		res["from_"+key] = int64(20)
		res["to_"+key] = int64(50)
	}
	return res
}
