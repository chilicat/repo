package main

type Config struct {
	// Repository base url
	BaseUrl string
	// Destination folder.
	Dest string 
	// File name template
	Template string

	DryRun bool

	WorkersCount int
}