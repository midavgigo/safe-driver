package dbentry

import (
	"app/utils"
	"strconv"
)

func SetDriverStatus(
	dbman DBManager,
	id int,
	is_available bool,
	latitude float64,
	longitude float64,
) error {
	res, err := dbman.Exec("UPDATE Drivers Set IsAvailable = $1, Latitude = $2, Longitude = $3 WHERE Id = $4",
		is_available,
		longitude,
		latitude,
		id,
	)
	if err != nil {
		return utils.ReasonableError{
			Reason:  err,
			Message: "Error while updating driver status",
		}
	}
	n, err := res.RowsAffected()
	if err != nil {
		return utils.ReasonableError{
			Reason:  err,
			Message: "Error in getting number of affected rows",
		}
	}
	if n < 1 {
		return utils.ReasonableError{
			Reason:  err,
			Message: "Status of driver " + strconv.Itoa(id) + " not changed. Check id",
		}
	}
	return nil
}
