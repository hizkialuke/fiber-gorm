package repo

import (
	"fiber-gorm/database"
	"fiber-gorm/models"
	"fiber-gorm/utils"
	"fmt"
)

func GetReportA(params *models.ReportRequest, userID int64) (*models.Pagination, error) {
	var (
		allDatas []*models.ReportATransformer
		datas    []*models.ReportATransformer
	)

	params.Sort = "date ASC"
	dateFormat := "%Y-%m-%d"

	rawSql := `SELECT DATE_FORMAT(t.created_at,'%s') as date, m.merchant_name, IFNULL(SUM(t.bill_total), 0) as omzet
	FROM transactions t, merchants m
	WHERE t.merchant_id = m.id
	AND t.merchant_id IN (SELECT id FROM merchants WHERE user_id IN (SELECT id FROM users WHERE id = %d))
	GROUP BY DATE_FORMAT(created_at,'%s')
	ORDER BY %s`
	filledRaw := fmt.Sprintf(rawSql, dateFormat, userID, dateFormat, params.Sort)

	paginateSql := rawSql + ` LIMIT %d OFFSET %d`
	filledPaginate := fmt.Sprintf(paginateSql, dateFormat, userID, dateFormat, params.Sort, params.GetLimit(), params.GetOffset())

	if params.StartDate != nil && params.EndDate != nil {
		params.Sort = "calendar.date ASC"
		rawSql = `SELECT DATE_FORMAT(calendar.date, '%s') as date, m.merchant_name, IFNULL(ttb.omzet, 0) as omzet
	FROM (SELECT * FROM merchants WHERE user_id IN (SELECT id FROM users WHERE id = %d)) m,
	(
    select curdate() - INTERVAL (a.a + (10 * b.a) + (100 * c.a) + (1000 * d.a) ) DAY as date
    FROM (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as a
    cross join (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as b
    cross join (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as c
    cross join (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as d
	) calendar
	LEFT JOIN (
		SELECT SUM(t.bill_total) as omzet, DATE_FORMAT(t.created_at, '%s') as transaction_date
		FROM transactions t
		WHERE merchant_id IN (SELECT id FROM merchants WHERE user_id IN (SELECT id FROM users WHERE id = %d))
		GROUP BY DATE_FORMAT(created_at, '%s')
	) as ttb
	ON calendar.date = ttb.transaction_date
	where calendar.date between '%s' and '%s'
	ORDER BY %s`
		filledRaw = fmt.Sprintf(rawSql, dateFormat, userID, dateFormat, userID, dateFormat, *params.StartDate, *params.EndDate, params.Sort)

		paginateSql = rawSql + ` LIMIT %d OFFSET %d`
		filledPaginate = fmt.Sprintf(paginateSql, dateFormat, userID, dateFormat, userID, dateFormat, *params.StartDate, *params.EndDate, params.Sort, params.GetLimit(), params.GetOffset())
	}

	// count all data
	database.DBConn.Raw(filledRaw).Find(&allDatas)
	// paginate data
	database.DBConn.Raw(filledPaginate).Find(&datas)

	params.Rows = datas
	params.Pagination = utils.CustomPaginate(int64(len(allDatas)), params.Pagination)

	return params.Pagination, nil
}

func GetReportB(params *models.ReportRequest, userID int64) (*models.Pagination, error) {
	var (
		allDatas []*models.ReportBTransformer
		datas    []*models.ReportBTransformer
	)

	params.Sort = "date ASC"
	dateFormat := "%Y-%m-%d"

	rawSql := `SELECT DATE_FORMAT(t.created_at, '%s') as date, m.merchant_name, o.outlet_name, SUM(t.bill_total) as omzet
	FROM transactions t, merchants m, outlets o
	WHERE t.merchant_id = m.id
	AND t.outlet_id = o.id
	AND o.merchant_id IN (SELECT id FROM merchants WHERE user_id IN (SELECT id FROM users WHERE id = %d))
	GROUP BY DATE_FORMAT(t.created_at, '%s'), outlet_id
	ORDER BY %s`
	filledRaw := fmt.Sprintf(rawSql, dateFormat, userID, dateFormat, params.Sort)

	paginateSql := rawSql + ` LIMIT %d OFFSET %d`
	filledPaginate := fmt.Sprintf(paginateSql, dateFormat, userID, dateFormat, params.Sort, params.GetLimit(), params.GetOffset())

	if params.StartDate != nil && params.EndDate != nil {
		params.Sort = "calendar.date ASC, o.id"

		rawSql = `SELECT DATE_FORMAT(calendar.date, '%s') as date, m.merchant_name, o.outlet_name, IFNULL(ttb.omzet, 0) as omzet
		FROM outlets o, merchants m,
		(
			select curdate() - INTERVAL (a.a + (10 * b.a) + (100 * c.a) + (1000 * d.a) ) DAY as date
			from (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as a
			cross join (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as b
	 		cross join (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as c
			cross join (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as d
	 	) calendar
		LEFT JOIN (
			SELECT DATE_FORMAT(t.created_at, '%s') as transaction_date, m.merchant_name, o.outlet_name, SUM(t.bill_total) as omzet
			FROM transactions t, merchants m, outlets o
			WHERE t.merchant_id = m.id
			AND t.outlet_id = o.id
			AND o.merchant_id IN (SELECT id FROM merchants WHERE user_id IN (SELECT id FROM users WHERE id = %d))
			GROUP BY DATE_FORMAT(t.created_at, '%s'), outlet_id
		) as ttb
		ON calendar.date = ttb.transaction_date
		WHERE o.merchant_id = m.id
		AND o.merchant_id IN (SELECT id FROM merchants WHERE user_id IN (SELECT id FROM users WHERE id = %d))
		AND calendar.date between '%s' and '%s'
		GROUP BY outlet_name, DATE_FORMAT(date, '%s')
		ORDER BY %s`
		filledRaw = fmt.Sprintf(rawSql, dateFormat, dateFormat, userID, dateFormat, userID, *params.StartDate, *params.EndDate, dateFormat, params.Sort)

		paginateSql = rawSql + ` LIMIT %d OFFSET %d`
		filledPaginate = fmt.Sprintf(paginateSql, dateFormat, dateFormat, userID, dateFormat, userID, *params.StartDate, *params.EndDate, dateFormat, params.Sort, params.GetLimit(), params.GetOffset())
	}

	// count all data
	database.DBConn.Raw(filledRaw).Find(&allDatas)
	// paginate data
	database.DBConn.Raw(filledPaginate).Find(&datas)

	params.Rows = datas
	params.Pagination = utils.CustomPaginate(int64(len(allDatas)), params.Pagination)

	return params.Pagination, nil
}
