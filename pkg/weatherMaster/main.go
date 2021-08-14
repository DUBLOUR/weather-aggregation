package weatherMaster

import (
	"fmt"
)

type Weather struct {
	Temp float32
	Hum  int
	Wind float32
}

type ISource interface {
	InCity(string) (Weather, error)
}

type ILogger interface {
	Info(v ...interface{})
	Warning(v ...interface{})
}

type IMetricHandler interface {
	Inc(string) error
	Get(string) (int, error)
}

type Master struct {
	sources []ISource
	log ILogger
	metricHandler IMetricHandler
}

func (m *Master) ClearSource() {
	m.sources = nil
}

func (m *Master) AppendSource(newSource ISource) {
	m.sources = append(m.sources, newSource)
}

func (m *Master) InitSource(mainSource ISource, additionalSources ...ISource) {
	m.ClearSource()
	m.AppendSource(mainSource)
	for _, getter := range additionalSources {
		m.AppendSource(getter)
	}
}

func (m *Master) SetLogger(logger ILogger) {
	m.log = logger
}

func (m *Master) SetMetricHandler(handler IMetricHandler) {
	m.metricHandler = handler
}

func (m *Master) WeatherInCity(city string) (Weather, error) {
	if len(m.sources) == 0 {
		return Weather{}, fmt.Errorf("API ources not initialized")
	}
	if m.log == nil {
		return Weather{}, fmt.Errorf("logger not initialized")
	}
	if m.metricHandler == nil {
		return Weather{}, fmt.Errorf("metricHandler not initialized")
	}

	m.log.Info("Request: ", city)
	_ = m.metricHandler.Inc(city)
	cnt, _ := m.metricHandler.Get(city)
	m.log.Info("\tIs it", cnt, "-th request for", city)

	for _, s := range m.sources {
		w, err := s.InCity(city)
		if err == nil {
			m.log.Info("\tSuccessful")
			return w, nil
		}
	}
	m.log.Warning("\tCan't find working API")
	return Weather{}, fmt.Errorf("have not working API")
}
