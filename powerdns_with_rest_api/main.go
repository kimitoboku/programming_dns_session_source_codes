package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"net/url"
)

type LookupResult struct {
	Qtype   string `json:"qtype"`
	Qname   string `json:"qname"`
	Content string `json:"content"`
	Ttl     int    `json:"ttl"`
}

type Response struct {
	Result []LookupResult `json:"result"`
}

func handleLookup(c echo.Context) error {
	qname := c.Param("qname")
	qtype := c.Param("qtype")
	qname, _ = url.QueryUnescape(qname)

	switch qtype {
	case "SOA":
		result := Response{
			Result: []LookupResult{
				{
					Qname:   qname,
					Qtype:   qtype,
					Content: "example.com. hostmaster.example.com. 1 1800 3600 7200 5",
					Ttl:     60,
				},
			},
		}
		return c.JSON(http.StatusOK, result)
	case "A", "ANY":
		result := Response{
			Result: []LookupResult{
				{
					Qname:   qname,
					Qtype:   "A",
					Content: "10.1.1.1",
					Ttl:     60,
				},
			},
		}
		return c.JSON(http.StatusOK, result)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "Thes type is not supported")

	}

}

type GetDomainMetaResponse struct {
	Result []string `json:"result"`
}

func handleGetDomainMeta(c echo.Context) error {
	result := GetDomainMetaResponse{
		Result: []string{"0"},
	}

	return c.JSON(http.StatusOK, result)
}

type GetAllDomainMetaResult struct {
	Presigned []string `json:"PRESIGNED"`
}

type GetAllDomainMetaResponse struct {
	Result GetAllDomainMetaResult `json:"result"`
}

func handleGetAllDomainMeta(c echo.Context) error {
	result := GetAllDomainMetaResponse{
		Result: GetAllDomainMetaResult{
			Presigned: []string{"0"},
		},
	}

	return c.JSON(http.StatusOK, result)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/dnsapi/lookup/:qname/:qtype", handleLookup)
	e.GET("/dnsapi/getDomainMetadata/:name/:kind", handleGetDomainMeta)
	e.GET("/dnsapi/getAllDomainMetadata/:name", handleGetAllDomainMeta)

	e.Logger.Fatal(e.Start(":1323"))
}
