package seek

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobLocation struct {
	Title string `json:"title"`
	Count int    `json:"count"`
}

type Company struct {
	Company string `json:"company"`
	Count   int    `json:"count"`
}

func (sc seekController) AnalyzeJobs(ctx *gin.Context) {
	locationTitleMap := make(map[string][]JobLocation)
	companyNameMap := make(map[string][]Company)
	locationTitle, err := sc.jobRepo.JobTitle(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	companyName, err := sc.jobRepo.CompanyTitle(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	for _, cn := range companyName {
		companyNameMap[cn.Location] = append(companyNameMap[cn.Location], Company{
			Company: cn.Company,
			Count:   cn.Total,
		})
	}

	for _, lt := range locationTitle {
		locationTitleMap[lt.Location] = append(locationTitleMap[lt.Location], JobLocation{
			Title: lt.Title,
			Count: lt.Total,
		})
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"location": locationTitleMap,
		"company":  companyNameMap,
	})
}
