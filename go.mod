module github.com/jim380/kyoto-demo

go 1.18

require (
	github.com/kyoto-framework/kyoto v0.3.1-0.20220501105523-87c78018253e
	go.uber.org/zap v1.21.0
)

require (
	github.com/kyoto-framework/scheduler v0.0.0-20220419050632-472087dddf82 // indirect
	github.com/kyoto-framework/zen v0.0.0-20220418065731-a871e34f2f62 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/exp v0.0.0-20220328175248-053ad81199eb // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace github.com/kyoto-framework/uikit/twui => ./uikit/twui
