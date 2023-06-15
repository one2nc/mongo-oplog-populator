package types

// TODO-DONE: give better name
type PersonnelInfo struct {
	FirstNames, LastNames, Subjects, StreetAddresses, Positions, Zips, PhoneNumbers []string
	Ages, Workhours                                                                 []int
	Salaries                                                                        []float64
}
