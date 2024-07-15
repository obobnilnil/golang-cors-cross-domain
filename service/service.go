// package service

// import (
// 	"bytes"
// 	"cyberreason_cross_domain/model"
// 	"cyberreason_cross_domain/utility"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"image/color"
// 	"io/ioutil"
// 	"math"
// 	"net/http"
// 	"net/url"
// 	"strings"
// 	"sync"
// 	"time"

// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/vg"
// )

// type ServicePort interface {
// 	LoginServices(login model.LoginRequest) (map[string]string, error)
// 	WidgetsServices(widgets map[string]interface{}, cookies map[string]string) ([]byte, error)
// 	GroupsServices(cookies map[string]string) ([]byte, error)
// 	GraphMalopsResolutionTrackingServices(req []model.MalopResolutionTracking) (string, error)
// }

// type serviceAdapter struct {
// 	cookies map[string]string // only one single cookie share all the rest
// 	mu      sync.Mutex
// }

// func NewServiceAdapter() ServicePort {
// 	return &serviceAdapter{
// 		cookies: make(map[string]string),
// 	}
// }

// func (s *serviceAdapter) LoginServices(login model.LoginRequest) (map[string]string, error) {
// 	if !login.Validate {
// 		return nil, errors.New("boolean value is false")
// 	}

// 	formData := url.Values{
// 		"username": {"your_username"},
// 		"password": {"your_password"},
// 	}

// 	req, err := http.NewRequest("POST", "https://sisth.cybereason.net/login.html", strings.NewReader(formData.Encode()))
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	client := &http.Client{
// 		CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 			return http.ErrUseLastResponse
// 		},
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// fmt.Println("Response Body:", string(body))

// 	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusFound {
// 		return nil, fmt.Errorf("login failed with status code: %d and body: %s", resp.StatusCode, string(body))
// 	}

// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	cookieMap := make(map[string]string)
// 	for _, cookie := range resp.Cookies() {
// 		cookieMap[cookie.Name] = cookie.Value
// 		s.cookies[cookie.Name] = cookie.Value // เก็บ cookie ใน session storage
// 	}

// 	if _, found := cookieMap["JSESSIONID"]; !found {
// 		return nil, errors.New("JSESSIONID cookie not found")
// 	}

// 	return cookieMap, nil
// }

// func (s *serviceAdapter) WidgetsServices(widgets map[string]interface{}, cookies map[string]string) ([]byte, error) {
// 	url := "https://sisth.cybereason.net/rest/dynamic/v1/hs-discovery/api/discovery/v1/widgets"

// 	jsonData, err := json.Marshal(widgets)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Accept", "application/json")

// 	fmt.Println("Cookies used for Widgets API:")

// 	// เพิ่ม Cookie ใน Header จาก session storage
// 	s.mu.Lock()
// 	for name, value := range s.cookies {
// 		req.AddCookie(&http.Cookie{Name: name, Value: value})
// 		req.Header.Add("Cookie", fmt.Sprintf("%s=%s", name, value))
// 		fmt.Printf("%s=%s\n", name, value)
// 	}
// 	s.mu.Unlock()

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	// fmt.Println("Response Headers:", resp.Header)

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// fmt.Println("Response Body:", string(body))

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("API call failed with status code: %d and body: %s", resp.StatusCode, string(body))
// 	}

// 	return body, nil
// }

// func (s *serviceAdapter) GroupsServices(cookies map[string]string) ([]byte, error) {
// 	url := "https://sisth.cybereason.net/rest/groups/permitted"

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Accept", "application/json")

// 	fmt.Println("Cookies used for Groups API:")
// 	// เพิ่ม Cookie ใน Header จาก session storage
// 	s.mu.Lock()
// 	for name, value := range s.cookies {
// 		req.AddCookie(&http.Cookie{Name: name, Value: value})
// 		req.Header.Add("Cookie", fmt.Sprintf("%s=%s", name, value))
// 		fmt.Printf("%s=%s\n", name, value)
// 	}
// 	s.mu.Unlock()

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	// fmt.Println("Response Headers:", resp.Header)

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// fmt.Println("Response Body:", string(body))

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("API call failed with status code: %d and body: %s", resp.StatusCode, string(body))
// 	}

// 	return body, nil
// }

package service

import (
	"bytes"
	"cyberreason_cross_domain/model"
	"cyberreason_cross_domain/utility"
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type ServicePort interface {
	LoginServices(login model.Login) (map[string]string, error)
	WidgetsServices(widgets map[string]interface{}, userID string) ([]byte, error)
	GroupsServices(userID string) ([]byte, error)
	GraphMalopsResolutionTrackingServices(req []model.MalopResolutionTracking) (string, error)
}

type serviceAdapter struct {
	cookies map[string]map[string]string // แผนที่ผู้ใช้เก็บ cookie
	mu      sync.Mutex
}

func NewServiceAdapter() ServicePort {
	return &serviceAdapter{
		cookies: make(map[string]map[string]string),
	}
}

func (s *serviceAdapter) LoginServices(login model.Login) (map[string]string, error) {
	// if !login.Validate {
	// 	return nil, errors.New("boolean value is false")
	// }

	formData := url.Values{
		"username": {login.Username},
		"password": {login.Password},
	}

	req, err := http.NewRequest("POST", "https://sisth.cybereason.net/login.html", strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println("Response Body:", string(body))

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusFound {
		return nil, fmt.Errorf("login failed with status code: %d and body: %s", resp.StatusCode, string(body))
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	cookieMap := make(map[string]string)
	for _, cookie := range resp.Cookies() {
		cookieMap[cookie.Name] = cookie.Value
	}

	if _, found := cookieMap["JSESSIONID"]; !found {
		return nil, errors.New("JSESSIONID cookie not found")
	}

	// เก็บ cookie โดยใช้ username เป็นคีย์
	s.cookies[login.Username] = cookieMap

	// พิมพ์ค่า cookies ที่เก็บใน session storage
	fmt.Println("Cookies stored in session for user:", login.Username)
	for name, value := range cookieMap {
		fmt.Printf("%s: %s\n", name, value)
	}

	return cookieMap, nil
}

func (s *serviceAdapter) WidgetsServices(widgets map[string]interface{}, userID string) ([]byte, error) {
	url := "https://sisth.cybereason.net/rest/dynamic/v1/hs-discovery/api/discovery/v1/widgets"

	jsonData, err := json.Marshal(widgets)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// เพิ่ม Cookie ใน Header จาก session storage
	s.mu.Lock()
	fmt.Println("Cookies used for Widgets API by user:", userID)
	for name, value := range s.cookies[userID] {
		req.AddCookie(&http.Cookie{Name: name, Value: value})
		req.Header.Add("Cookie", fmt.Sprintf("%s=%s", name, value))
		fmt.Printf("%s=%s\n", name, value)
	}
	s.mu.Unlock()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// fmt.Println("Response Headers:", resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println("Response Body:", string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API call failed with status code: %d and body: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func (s *serviceAdapter) GroupsServices(userID string) ([]byte, error) {
	url := "https://sisth.cybereason.net/rest/groups/permitted"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	// เพิ่ม Cookie ใน Header จาก session storage
	s.mu.Lock()
	fmt.Println("Cookies used for Groups API by user:", userID)
	for name, value := range s.cookies[userID] {
		req.AddCookie(&http.Cookie{Name: name, Value: value})
		req.Header.Add("Cookie", fmt.Sprintf("%s=%s", name, value))
		fmt.Printf("%s=%s\n", name, value)
	}
	s.mu.Unlock()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// fmt.Println("Response Headers:", resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println("Response Body:", string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API call failed with status code: %d and body: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func (s *serviceAdapter) GraphMalopsResolutionTrackingServices(req []model.MalopResolutionTracking) (string, error) {
	p := plot.New()
	p.Title.Text = "MalOp Resolution Tracking"
	p.X.Label.Text = "Date"
	p.Y.Label.Text = "Count"

	var minY, maxY float64 = math.Inf(1), math.Inf(-1)
	totalPoints := make(plotter.XYs, len(req))
	closedPoints := make(plotter.XYs, len(req))

	for i, d := range req {
		t := time.Unix(0, d.Timestamp*int64(time.Millisecond))
		totalPoints[i].X = float64(t.Unix())
		totalPoints[i].Y = float64(d.TotalMalopsCount)
		closedPoints[i].X = float64(t.Unix())
		closedPoints[i].Y = float64(d.ClosedMalopsCount)

		minY = math.Min(minY, totalPoints[i].Y)
		maxY = math.Max(maxY, totalPoints[i].Y)
	}

	totalLine, _ := plotter.NewLine(totalPoints)
	totalLine.LineStyle.Width = vg.Points(2)
	totalLine.LineStyle.Color = color.RGBA{R: 153, G: 102, B: 255, A: 255} // Purple

	closedLine, _ := plotter.NewLine(closedPoints)
	closedLine.LineStyle.Width = vg.Points(2)
	closedLine.LineStyle.Color = color.RGBA{R: 255, G: 153, B: 0, A: 255} // Orange

	p.Add(totalLine, closedLine)

	// Custom ticks for the X-axis
	p.X.Tick.Marker = plot.TickerFunc(func(min, max float64) []plot.Tick {
		ticks := make([]plot.Tick, 0, len(req))
		for _, d := range req {
			t := time.Unix(0, d.Timestamp*int64(time.Millisecond))
			tick := plot.Tick{
				Value: float64(t.Unix()),
				Label: t.Format("02/01/2006"), // DD/MM/YYYY format
			}
			ticks = append(ticks, tick)
		}
		return ticks
	})

	// Dynamic tick marks for Y-axis
	p.Y.Tick.Marker = plot.TickerFunc(func(min, max float64) []plot.Tick {
		ticks := []plot.Tick{}
		step := (max - min) / 5 // Adjust the number of ticks
		for value := min; value <= max; value += step {
			ticks = append(ticks, plot.Tick{
				Value: value,
				Label: utility.FormatK(value),
			})
		}
		return ticks
	})

	filename := "malop_resolution_tracking.svg"
	if err := p.Save(14*vg.Inch, 5*vg.Inch, filename); err != nil {
		return "", err
	}

	return filename, nil
}
