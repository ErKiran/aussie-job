package seek

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Description struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Logo struct {
	Strategies Strategies `json:"strategies"`
}

type Assets struct {
	Logo Logo `json:"logo"`
}

type Branding struct {
	ID     string `json:"id"`
	Assets Assets `json:"assets"`
}

type Strategies struct {
	Jdplogo  string `json:"jdpLogo"`
	Serplogo string `json:"serpLogo"`
}

type JobLocation struct {
	Label        string `json:"label"`
	Countrycode  string `json:"countryCode"`
	Seohierarchy []struct {
		Contextualname string `json:"contextualName"`
	} `json:"seoHierarchy"`
}

type Solmetadata struct {
	Searchrequesttoken string `json:"searchRequestToken"`
	Token              string `json:"token"`
	Jobid              string `json:"jobId"`
	Section            string `json:"section"`
	Sectionrank        int    `json:"sectionRank"`
	Jobadtype          string `json:"jobAdType"`
	Tags               struct {
		MordorFlights string `json:"mordor__flights"`
		MordorS       string `json:"mordor__s"`
	} `json:"tags"`
}

type SolmetadataInfo struct {
	Requesttoken     string   `json:"requestToken"`
	Token            string   `json:"token"`
	Keywords         string   `json:"keywords"`
	Sortmode         string   `json:"sortMode"`
	Locations        []string `json:"locations"`
	Locationdistance int      `json:"locationDistance"`
	Pagesize         int      `json:"pageSize"`
	Pagenumber       int      `json:"pageNumber"`
	Totaljobcount    int      `json:"totalJobCount"`
	Tags             struct {
		MordorSearchmarket    string `json:"mordor:searchMarket"`
		MordorResultCountRst  string `json:"mordor:result_count_rst"`
		MordorResultCountVec  string `json:"mordor:result_count_vec"`
		MordorRt              string `json:"mordor:rt"`
		MordorCountVec        string `json:"mordor:count_vec"`
		MordorFlights         string `json:"mordor__flights"`
		MordorCountRst        string `json:"mordor:count_rst"`
		MordorCountIr         string `json:"mordor:count_ir"`
		MordorResultCountIr   string `json:"mordor:result_count_ir"`
		ChaliceSearchAPISolid string `json:"chalice-search-api:solId"`
	} `json:"tags"`
}

type Location struct {
	Areadescription         string `json:"areaDescription"`
	Areaid                  int    `json:"areaId"`
	Description             string `json:"description"`
	Locationdescription     string `json:"locationDescription"`
	Locationid              int    `json:"locationId"`
	Matched                 bool   `json:"matched"`
	Statedescription        string `json:"stateDescription"`
	Suburbparentdescription string `json:"suburbParentDescription"`
	Type                    string `json:"type"`
	Whereid                 int    `json:"whereId"`
	Descriptions            struct {
		En struct {
			Contextualname string `json:"contextualName"`
		} `json:"en"`
		ID struct {
			Contextualname string `json:"contextualName"`
		} `json:"id"`
		Th struct {
			Contextualname string `json:"contextualName"`
		} `json:"th"`
	} `json:"descriptions"`
}

type Searchparams struct {
	Page               string `json:"page"`
	Seekselectallpages string `json:"seekselectallpages"`
	Keywords           string `json:"keywords"`
	Locale             string `json:"locale"`
	Where              string `json:"where"`
	Solid              string `json:"solid"`
}

type Info struct {
	Timetaken  int    `json:"timeTaken"`
	Source     string `json:"source"`
	Experiment string `json:"experiment"`
}

type Sortmode struct {
	Isactive bool   `json:"isActive"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

type Paginationparameters struct {
	Seekselectallpages bool `json:"seekSelectAllPages"`
	Hadpremiumlistings bool `json:"hadPremiumListings"`
}

type Tags struct {
	Type  string `json:"type"`
	Label string `json:"label"`
}

type Joracrosslink struct {
	Cancrosslink bool `json:"canCrossLink"`
}

type Displaystyle struct {
	Search string `json:"search"`
}

type Searchinsights struct {
	Unmatchedkeywords []string `json:"unmatchedKeywords"`
}

type JobInfo struct {
	Advertiser                     Description    `json:"advertiser"`
	Area                           string         `json:"area,omitempty"`
	Areaid                         int            `json:"areaId,omitempty"`
	Areawherevalue                 string         `json:"areaWhereValue,omitempty"`
	Automaticinclusion             bool           `json:"automaticInclusion"`
	Branding                       Branding       `json:"branding,omitempty"`
	Bulletpoints                   []string       `json:"bulletPoints"`
	Classification                 Description    `json:"classification"`
	Companyname                    string         `json:"companyName,omitempty"`
	Companyprofilestructureddataid int            `json:"companyProfileStructuredDataId"`
	Displaystyle                   Displaystyle   `json:"displayStyle,omitempty"`
	Displaytype                    string         `json:"displayType"`
	Listingdatedisplay             string         `json:"listingDateDisplay"`
	Location                       string         `json:"location"`
	Locationid                     int            `json:"locationId"`
	Locationwherevalue             string         `json:"locationWhereValue"`
	ID                             int            `json:"id"`
	Ispremium                      bool           `json:"isPremium"`
	Isstandout                     bool           `json:"isStandOut"`
	Joblocation                    JobLocation    `json:"jobLocation"`
	Listingdate                    time.Time      `json:"listingDate"`
	Logo                           Description    `json:"logo"`
	Roleid                         string         `json:"roleId"`
	Salary                         string         `json:"salary"`
	Searchinsights                 Searchinsights `json:"searchInsights,omitempty"`
	Solmetadata                    Solmetadata    `json:"solMetadata"`
	Subclassification              Description    `json:"subClassification"`
	Suburb                         string         `json:"suburb,omitempty"`
	Suburbid                       int            `json:"suburbId,omitempty"`
	Suburbwherevalue               string         `json:"suburbWhereValue,omitempty"`
	Teaser                         string         `json:"teaser"`
	Title                          string         `json:"title"`
	Tracking                       string         `json:"tracking"`
	Worktype                       string         `json:"workType"`
	Isprivateadvertiser            bool           `json:"isPrivateAdvertiser"`
	Tags                           []Tags         `json:"tags,omitempty"`
}

type SearchedInfo struct {
	JobInfo              []JobInfo            `json:"data"`
	Title                string               `json:"title"`
	Totalcount           int                  `json:"totalCount"`
	Totalpages           int                  `json:"totalPages"`
	Paginationparameters Paginationparameters `json:"paginationParameters"`
	Info                 Info                 `json:"info"`
	Userqueryid          string               `json:"userQueryId"`
	Sortmode             []Sortmode           `json:"sortMode"`
	Solmetadata          SolmetadataInfo      `json:"solMetadata"`
	Location             Location             `json:"location"`
	Joracrosslink        Joracrosslink        `json:"joraCrossLink"`
	Searchparams         Searchparams         `json:"searchParams"`
}

type SummarizedData struct {
	JobID          string    `json:"id" gorm:"column:id;primary_key"`
	CompanyID      string    `json:"company_id" gorm:"column:company_id"`
	Title          string    `json:"title" gorm:"column:title"`
	Location       string    `json:"location"  gorm:"column:location"`
	Company        string    `json:"company"  gorm:"column:company"`
	ListingDate    time.Time `json:"listing_date"  gorm:"column:listing_date"`
	Salary         string    `json:"salary"  gorm:"column:salary"`
	Role           string    `json:"role" gorm:"column:role"`
	JobDescription string    `json:"job_description"  gorm:"column:job_description"`
	WorkType       string    `json:"work_type"  gorm:"column:work_type"`
	URL            string    `json:"url"  gorm:"column:url"`
	ExtraInfo      string    `json:"extra_info" gorm:"column:extra_info"`
}

func (sk *SeekAPI) SearchJobs(ctx context.Context, keyword string) ([]SummarizedData, error) {
	url := sk.SearchSlug(1, keyword)
	var jobsData []SummarizedData

	resultJobs := make(chan *SearchedInfo, 1)
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		sk.ProcessResult(ctx, url, resultJobs)
	}()

	go func() {
		wg.Wait()
		close(resultJobs)
	}()

	for {
		res, ok := <-resultJobs
		if !ok {
			fmt.Println("channel is closed")
			break // Exit the loop when the channel is closed
		}

		if res == nil {
			fmt.Println("res is nil", res)
			// Handle nil case if necessary
			continue
		}

		var pageWg sync.WaitGroup
		pageResultJobs := make(chan *SearchedInfo, 1)
		var pagesProcessed sync.WaitGroup

		for additionalPages := 2; additionalPages <= res.Totalpages; additionalPages++ {
			url := sk.SearchSlug(additionalPages, keyword)
			pageWg.Add(1)
			pagesProcessed.Add(1)
			go func(url string, wg *sync.WaitGroup, pagesProcessed *sync.WaitGroup) {
				defer wg.Done()
				defer pagesProcessed.Done()
				sk.ProcessResult(ctx, url, pageResultJobs)
			}(url, &pageWg, &pagesProcessed)
		}

		go func() {
			pageWg.Wait()
			close(pageResultJobs)
		}()

		go func() {
			pagesProcessed.Wait()
		}()

		for additionalJobs := range pageResultJobs {
			if additionalJobs != nil {
				res.JobInfo = append(res.JobInfo, additionalJobs.JobInfo...)
			}
		}

		for _, job := range res.JobInfo {
			jobsData = append(jobsData, SummarizedData{
				Location:       job.Location,
				JobID:          strconv.Itoa(job.ID),
				CompanyID:      job.Advertiser.ID,
				Company:        job.Advertiser.Description,
				ListingDate:    job.Listingdate,
				Salary:         job.Salary,
				Role:           job.Roleid,
				Title:          job.Title,
				JobDescription: job.Teaser,
				ExtraInfo: func() string {
					var info string
					if len(job.Bulletpoints) != 0 {
						return info
					}

					for _, points := range job.Bulletpoints {
						info += points
					}
					return info
				}(),
				WorkType: job.Worktype,
				URL:      fmt.Sprintf("https://www.seek.com.au/job/%d", job.ID),
			})
		}
	}

	return jobsData, nil
}

func (sk *SeekAPI) ProcessResult(ctx context.Context, url string, resultJobs chan<- *SearchedInfo) {
	req, err := sk.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error in request", err)
		resultJobs <- nil
		return
	}

	fmt.Println("req", req.URL.String())

	var res *SearchedInfo
	if _, err := sk.client.Do(ctx, req, &res); err != nil {
		fmt.Println("error in DO", err)
		resultJobs <- nil
		return
	}
	resultJobs <- res
}

func (sk *SeekAPI) SearchSlug(page int, keyword string) string {
	query := map[string]string{
		"page":               strconv.Itoa(page),
		"where":              "All+Australia",
		"seekSelectAllPages": "true",
		"locale":             "en-AU",
		"keywords":           keyword,
		"siteKey":            "AU-Main",
	}

	var slug string
	for k, v := range query {
		slug += fmt.Sprintf("%s=%s&", k, v)
	}
	return fmt.Sprintf("?%s", strings.TrimSuffix(slug, "&"))
}
