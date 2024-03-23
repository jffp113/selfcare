package operation

import "fmt"

const login = "https://%s/login"
const signIn = "https://%s/users/signin"
const timesheet = "https://%s/coworker/timesheet"
const saveTimesheet = "https://%s/coworker/savetimesheethour"

func loginUrl(host string) string {
	return fmt.Sprintf(login, host)
}

func signInUrl(host string) string {
	return fmt.Sprintf(signIn, host)
}

func timesheetURL(host string) string {
	return fmt.Sprintf(timesheet, host)
}

func saveTimesheetUrl(host string) string {
	return fmt.Sprintf(saveTimesheet, host)
}
