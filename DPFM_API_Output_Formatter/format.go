package dpfm_api_output_formatter

import (
	"data-platform-api-site-type-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToSiteType(rows *sql.Rows) (*[]SiteType, error) {
	defer rows.Close()
	siteType := make([]SiteType, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SiteType{}

		err := rows.Scan(
			&pm.SiteType,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &siteType, nil
		}

		data := pm
		siteType = append(siteType, SiteType{
			SiteType:				data.SiteType,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &siteType, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.SiteType,
			&pm.Language,
			&pm.SiteTypeName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			SiteType:     			data.SiteType,
			Language:          		data.Language,
			SiteTypeName:			data.SiteTypeName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}
