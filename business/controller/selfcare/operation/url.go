package operation

import "fmt"

const base = "https://%s"
const login = base + "/login"
const signIn = base + "/users/signin"
const timesheet = base + "/coworker/timesheet"
const saveTimesheet = base + "/coworker/savetimesheethour"
const getClients = base + "/coworker/getClients?timesheet_user=%s"

func baseUrl(host string) string {
	return fmt.Sprintf(base, host)
}

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

func getClientsUrl(host, userId string) string {
	return fmt.Sprintf(getClients, host, userId)
}
