package yahoo

import (
	"fmt"
	"testing"
)

func TestWeatherForecast(t *testing.T) {
	var s Session
	if err := s.Init(); err != nil {
		t.Log(err)
		t.Fail()
	}
	wfs, err := s.GetWeatherForecast("London", 5, TemperatureUnit_Celcius)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	for _, wf := range wfs {
		fmt.Printf("%#v\n", fmt.Sprintf("%s %s %s - *%s °C* / *%s °C*",
			wf.Code,
			wf.Day,
			wf.Date,
			wf.High,
			wf.Low))
	}
}
