package config

type Config struct {
	MultiTenant       bool             `mapstructure:"multi_tenant"`
	LimitActiveSeries int              `mapstructure:"limit_active_series"`
	PrometheusTarget  prometheusTarget `mapstructure:"prometheus_target"`
}

type prometheusTarget struct {
	Url      string `mapstructure:"url"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
