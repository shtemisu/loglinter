package loglinter

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/shtemisu/loglinter/analyzer"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglinter", New)
}

type Plugin struct{}

func New(settings any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.New()}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
