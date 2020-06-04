package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/nviso-be/thehive-sentinel-integration/thehive-sentinel-hooks/config"
	"github.com/nviso-be/thehive-sentinel-integration/thehive-sentinel-hooks/thehive"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var conf config.Conf

func callURL(u string, o *thehive.Capsule) error {
	jsonData, err := json.Marshal(o)
	if err != nil {
		log.Error().Msgf("json.Marshal(o): %v ", err)
		return err
	}

	log.Trace().Msgf("callURL json body: %v", string(jsonData))

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transCfg}

	response, err := client.Post(u, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error().Msgf("ioutil.ReadAll(response.Body): %v ", err)
		return err
	}

	if response.StatusCode != http.StatusOK {
		err := errors.New("http.Post: Request status code: " + strconv.Itoa(response.StatusCode) + string(data))
		return err
	}

	log.Info().Msgf("http.Post: Request status code:%d (%s) %v", response.StatusCode, http.StatusText(response.StatusCode), string(data))
	return nil
}

func parseRequestBody(b []byte) error {
	o := thehive.NewCapsule(&conf)
	if err := json.Unmarshal(b, &o); err != nil {
		log.Error().Msgf("json.Unmarshal %v", err)
		return err
	}

	log.Trace().Msgf("TheHive Webhook json body: %v", string(b))

	switch o.ObjectType {
	case "case":
		return handleCase(o)
	case "alert":
		return handleAlert(o)
	default:
		log.Debug().Msgf("request body contains unkown object type = %+v", o.ObjectType)
	}

	return nil
}

func handleCase(c *thehive.Capsule) error {
	switch c.Operation {
	case "Update":
		switch c.Details.Status {
		case "Resolved":
			if conf.ResolvedCaseURL == "" {
				log.Warn().Msg("ResolvedCaseURL not set in config.yml")
				return nil
			}
			if err := callURL(conf.ResolvedCaseURL, c); err != nil {
				log.Error().Err(err).Msg("callURL(conf.ResolvedCaseURL, &c)")
				return err
			}
		default:
			log.Debug().Msgf("Request body does not contain match in switch objecttype ='case': operation = 'Update'; Details.Status = %+v", c.Details.Status)
		}
	case "Creation":
		switch c.Details.Status {
		case "Open":
			if conf.NewCaseURL == "" {
				log.Warn().Msg("NewCaseURL not set in config.yml")
				return nil
			}
			if err := callURL(conf.NewCaseURL, c); err != nil {
				log.Error().Err(err).Msg("callURL(conf.NewCaseURL, &c)")
				return err
			}
		default:
			log.Debug().Msgf("debrequest body does not contain match in switch objecttype ='case': operation = 'Creation'; Details.Status = %+v", c.Details.Status)
		}
	default:
		log.Debug().Msgf("debug: request body does not contain match in switch objecttype ='case': operation = %+v", c.Operation)
	}
	return nil
}

func handleAlert(c *thehive.Capsule) error {
	switch c.Operation {
	case "Creation":
		switch c.Details.Status {
		case "New":
			if conf.NewAlertURL == "" {
				log.Warn().Msg("NewAlertURL not set in config.yml")
				return nil
			}
			if err := callURL(conf.NewAlertURL, c); err != nil {
				log.Error().Err(err).Msg("callURL(conf.NewAlertURL, c)")
				return err
			}
		default:
			err := fmt.Errorf("request body does not contain match in switch objecttype ='alert': operation = 'Creation'; Details.Status = %+v", c.Details.Status)
			log.Debug().Err(err).Send()
			return err
		}
	case "Update":
		switch c.Details.Status {
		case "Ignored":
			if conf.IgnoredAlertURL == "" {
				log.Warn().Msg("NewAlertURL not set in config.yml")
				return nil
			}
			if err := callURL(conf.IgnoredAlertURL, c); err != nil {
				log.Error().Err(err).Msg("callURL(conf.IgnoredAlertURL, c)")
				return err
			}
		case "Imported":
			if conf.ImportedAlertURL == "" {
				log.Warn().Msg("ImportedAlertURL not set in config.yml")
				return nil
			}
			if err := callURL(conf.ImportedAlertURL, c); err != nil {
				log.Error().Err(err).Msg("callURL(conf.ImportedAlertURL, c)")
				return err
			}
		default:
			log.Debug().Msgf("request body does not contain match in switch objecttype ='alert': operation = 'Update'; Details.Status = %+v", c.Details.Status)
		}
	default:
		log.Debug().Msgf("request body does not contain match in switch objecttype ='alert'; operation = %+v", c.Operation)
	}
	return nil
}

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message": "Goodday, this is thehivesentinelhooks REST API"}`))
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error": "Failed to read request body"}`))
			log.Error().Err(err).Msg("ioutil.ReadAll")
		} else {
			if err := parseRequestBody(body); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(`{"error": "POST method failed"}`))
				log.Error().Err(err).Msg("parseRequestBody(body)")
			} else {
				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte(`{"message": "POST method success"}`))
			}
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message": "Method not supported"}`))
		log.Error().Msg("http.request: Request Method not supported")
	}
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC822})

	configfile := flag.String("config", "./config.yml", "Configfile")
	flag.Parse()

	if err := config.GetConfig(&conf, configfile); err != nil {
		log.Fatal().Err(err).Msg("config.GetConfig(&conf, configfile)")
	}

	switch conf.LogLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Debug().Msgf("IgnoredAlertURL: %v", conf.IgnoredAlertURL)
	log.Debug().Msgf("ImportedAlertURL: %v", conf.ImportedAlertURL)
	log.Debug().Msgf("NewAlertURL: %v", conf.NewAlertURL)
	log.Debug().Msgf("NewCaseURL: %v", conf.NewCaseURL)
	log.Debug().Msgf("ResolvedCaseURL: %v", conf.ResolvedCaseURL)
	log.Debug().Msgf("Organization: %v", conf.Organization)

	http.HandleFunc("/", apiResponse)
	log.Info().Msg("Starting http listener")
	if err := http.ListenAndServe(":9002", nil); err != nil {
		log.Fatal().Err(err).Msg("Start http listener failed")
	}
}
