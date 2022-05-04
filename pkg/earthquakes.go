package pkg

import (
	//"errors"
	"net/http"
	//"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
)

type Earthquake struct {
	Date      string
	Latitude  string
	Longitude string
	Magnitude string
	Depth     string
	Location  string
	IrisId    string
}

func EarthquakeTable(w http.ResponseWriter, r *http.Request) {

	e := &Earthquake{}
	var quakes []Earthquake

	response, err := http.Get("http://ds.iris.edu/seismon/eventlist/index.phtml")
	if err != nil {
		return //err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return //errors.New("non 200 status code received")
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return //err
	}

	doc.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
		if i >= 10 {
			return
		}

		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			switch j {
			case 0:
				e.Date = td.Text()
			case 1:
				e.Latitude = td.Text()
			case 2:
				e.Longitude = td.Text()
			case 3:
				e.Magnitude = td.Text()
			case 4:
				e.Depth = td.Text()
			case 5:
				e.Location = td.Text()
			case 6:
				e.IrisId = td.Text()
			}
		})

		quakes = append(quakes, *e)
	})

	// table := tablewriter.NewWriter(os.Stdout)
	table := tablewriter.NewWriter(w)
	header := []string{"Date", "Location", "Magnitude", "Longitude", "Latitude", "Depth", "IrisID"}
	table.SetHeader(header)
	table.SetCaption(true, "Last Ten Earthquakes")

	for _, quake := range quakes {
		s := []string{
			quake.Date,
			quake.Location,
			quake.Magnitude,
			quake.Longitude,
			quake.Latitude,
			quake.Depth,
			quake.IrisId,
		}
		table.Append(s)
	}

	table.Render()
	//return //nil
}
