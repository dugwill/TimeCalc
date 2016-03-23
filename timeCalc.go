package timeCalc

import "time"

// the following statements set up time constants

const ThirtyDays time.Duration = 720 * time.Hour // 720 hours = 30 days
const OneDay time.Duration = 24 * time.Hour      // 24 hours = 1 day
const OneMinute time.Duration = 1 * time.Minute  // 1 minute
const OneHour time.Duration = 1 * time.Hour      // 1 hour
const OneYear time.Duration = 365 * OneDay       // 1 Day
