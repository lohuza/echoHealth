package echoHealth

type HealthChecker interface {
	CheckHealth() (string, bool)
}
