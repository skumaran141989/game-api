package analytics

type AnalyticsEngine interface {
	Calculate(input any) (any, error)
}
