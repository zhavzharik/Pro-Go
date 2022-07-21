package main

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
}

type ServiceWithFeatures struct {
	description    string
	durationMonths int
	monthlyFee     float64
	features []string
}
