package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/HenriqueVuolo/daily-tools-cli/models"
	"github.com/HenriqueVuolo/daily-tools-cli/utils"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var (
	city        string
	countryCode string
	airportCode string
	lat         string
	long        string
	ip          string
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Get current weather information",
	Long: `Get the current weather information for a specific location.

Example:
	weather --city London`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		res := fetchWeather()
		weather := parseWeather(res)
		showWeather(weather)
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	weatherCmd.Flags().StringVarP(&city, "city", "c", "", "City name")
	weatherCmd.Flags().StringVarP(&countryCode, "countryCode", "k", "", "Country name")
	weatherCmd.Flags().StringVarP(&airportCode, "airportCode", "a", "", "Airport Code")
	weatherCmd.Flags().StringVarP(&lat, "lat", "", "", "Latitude")
	weatherCmd.Flags().StringVarP(&long, "lon", "", "", "Longitude")
	weatherCmd.Flags().StringVarP(&ip, "ip", "i", "", "IP Address")
}

func formatQuery(params []string) string {
	filteredParams := utils.RemoveEmptyStrings(params)
	return strings.Join(filteredParams, ",")
}

func fetchWeather() []byte {
	baseURL := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	days := os.Getenv("DAYS")
	params := url.Values{}
	params.Add("key", apiKey)
	params.Add("days", days)
	encodedParams := params.Encode()
	q := formatQuery([]string{city, countryCode, airportCode, lat, long, ip})

	requestUrl := fmt.Sprintf("%s?%s&q=%s", baseURL, encodedParams, q)

	res, err := http.Get(requestUrl)
	if err != nil {
		log.Fatal("Error on request", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal("Weather API not available. Status:", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	return data
}

func parseWeather(data []byte) models.Weather {
	var weather models.Weather
	if err := json.Unmarshal(data, &weather); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}
	return weather
}

func showWeather(weather models.Weather) {
	location := weather.Location
	currentWeather := weather.Current
	lastUpdated, err := time.Parse("2006-01-02 15:04", currentWeather.LastUpdated)

	if err != nil {
		log.Fatal("Error parsing date")
	}

	locationFmt := color.New(color.FgHiRed, color.Bold, color.BgHiWhite)
	currentFmt := color.New(color.FgMagenta, color.Bold, color.BgHiWhite)
	tableHeaderFmt := color.New(color.FgHiYellow, color.Bold)

	locationFmt.Printf("   %s, %s   ", location.Name, location.Country)
	fmt.Print("\n\n")
	currentFmt.Printf(" Current (%s): Condition: %s | Temp: %.1f°C - Feels like %.1f°C ", lastUpdated.Format("15:04"), currentWeather.Condition.Text, currentWeather.TempC, currentWeather.FeelslikeC)
	fmt.Print("\n\n")
	forecastTable := table.New("Date", "Condition", "Max. Temp. (ºC)", "Min. Temp. (ºC)", "Chance of Rain (%)")
	forecastTable.WithHeaderFormatter(tableHeaderFmt.SprintfFunc())

	for _, forecast := range weather.Forecast.Forecastday {
		forecastDay := forecast.Day

		date, err := time.Parse("2006-01-02", forecast.Date)
		if err != nil {
			log.Fatal("Error parsing date")
		}

		forecastTable.AddRow(date.Format("02/01"), forecastDay.Condition.Text, forecastDay.MaxTempC, forecastDay.MinTempC, forecastDay.DailyChanceOfRain)
	}

	forecastTable.Print()
}
