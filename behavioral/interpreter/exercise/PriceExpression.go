package interpreter

type PriceExpression interface {
	calculate(price int, customer Customer) int
}
