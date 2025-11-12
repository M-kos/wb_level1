package calc

import (
	"math/big"
)

func Add[T any](x, y T) any {
	switch localX := any(x).(type) {
	case int:
		localY := any(y).(int)
		return localX + localY
	case int64:
		localY := any(y).(int64)
		return localX + localY
	case float64:
		localY := any(y).(float64)
		return localX + localY
	case *big.Int:
		localY := any(y).(*big.Int)
		return big.NewInt(0).Add(localX, localY)
	case *big.Float:
		localY := any(y).(*big.Float)
		return big.NewFloat(0).Add(localX, localY)
	default:
		return "type not supported"
	}
}

func Sub[T any](x, y T) any {
	switch localX := any(x).(type) {
	case int:
		localY := any(y).(int)
		return localX - localY
	case int64:
		localY := any(y).(int64)
		return localX + localY
	case float64:
		localY := any(y).(float64)
		return localX + localY
	case *big.Int:
		localY := any(y).(*big.Int)
		return big.NewInt(0).Add(localX, localY)
	case *big.Float:
		localY := any(y).(*big.Float)
		return big.NewFloat(0).Add(localX, localY)
	default:
		return "type not supported"
	}
}

func Multiple[T any](x, y T) any {
	switch localX := any(x).(type) {
	case int:
		localY := any(y).(int)
		return localX * localY
	case int64:
		localY := any(y).(int64)
		return localX * localY
	case float64:
		localY := any(y).(float64)
		return localX * localY
	case *big.Int:
		localY := any(y).(*big.Int)
		return big.NewInt(0).Mul(localX, localY)
	case *big.Float:
		localY := any(y).(*big.Float)
		return big.NewFloat(0).Mul(localX, localY)
	default:
		return "type not supported"
	}
}

func Div[T any](x, y T) any {
	switch localX := any(x).(type) {
	case int:
		localY := any(y).(int)
		if localY == 0 {
			return "division by zero"
		}
		return localX / localY
	case int64:
		localY := any(y).(int64)
		if localY == 0 {
			return "division by zero"
		}
		return localX / localY
	case float64:
		localY := any(y).(float64)
		if localY == 0 {
			return "division by zero"
		}
		return localX / localY
	case *big.Int:
		localY := any(y).(*big.Int)
		if localY.Cmp(big.NewInt(0)) == 0 {
			return "division by zero"
		}
		return big.NewInt(0).Div(localX, localY)
	case *big.Float:
		localY := any(y).(*big.Float)
		if localY.Cmp(big.NewFloat(0)) == 0 {
			return "division by zero"
		}
		return big.NewFloat(0).Quo(localX, localY)
	default:
		return "type not supported"
	}
}
