// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// GlobalLeaderboardCount represents a row from '[custom global_leaderboard_count]'.
type GlobalLeaderboardCount struct {
	Count int64 // count
}

// GetGlobalLeaderboardCounts runs a custom query, returning results as GlobalLeaderboardCount.
func GetGlobalLeaderboardCounts(db XODB) ([]*GlobalLeaderboardCount, error) {
	var err error

	// sql query
	const sqlstr = `select count("user") ` +
		`from global_rankings`

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*GlobalLeaderboardCount{}
	for q.Next() {
		glc := GlobalLeaderboardCount{}

		// scan
		err = q.Scan(&glc.Count)
		if err != nil {
			return nil, err
		}

		res = append(res, &glc)
	}

	return res, nil
}
