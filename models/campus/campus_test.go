package campusModel

import "testing"

func TestCampus(t *testing.T) {
	campuses, _ := GetAllCampuses()
	t.Log(campuses)
}
