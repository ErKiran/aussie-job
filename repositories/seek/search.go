package seek

import (
	"context"
	"fmt"
)

type JobLocation struct {
	Title    string `json:"title"`
	Location string `json:"location"`
	Total    int    `json:"total"`
}

type Company struct {
	Company  string `json:"company"`
	Location string `json:"location"`
	Total    int    `json:"total"`
}

func (sk seekRepo) JobTitle(ctx context.Context) ([]JobLocation, error) {
	var jobs []JobLocation
	if err := sk.db.WithContext(ctx).Raw(`select title, location, count(*) as total from jobs 
	group by title, location
	having count(*) >3
	order by total desc`).Scan(&jobs).Error; err != nil {
		fmt.Println("unable to get the jobs", err)
		return nil, err
	}
	return jobs, nil
}

func (sk seekRepo) CompanyTitle(ctx context.Context) ([]Company, error) {
	var jobs []Company
	if err := sk.db.WithContext(ctx).Raw(`select company, location, count(*) as total from jobs 
	group by company, location
	having count(*) >3
	order by total desc`).Scan(&jobs).Error; err != nil {
		fmt.Println("unable to get the jobs", err)
		return nil, err
	}
	return jobs, nil
}
