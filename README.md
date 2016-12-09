#Freetime

A CLI tool to help you find when you are available by querying your Google Calendar.

##To install Freetime

0) [Download and install Go, Git, and make sure your Google account has Google Calendar enabled](https://developers.google.com/google-apps/calendar/quickstart/go#prerequisites)

  * After installing Go, [make sure your `GOPATH` environmental variable is set](https://golang.org/doc/install)

1) Download this repo

    go get -d github.com/wholien/freetime

2) [Turn on Your Google Calendar API and download `client_secret.json` into the freetime directory](https://developers.google.com/google-apps/calendar/quickstart/go#step_1_turn_on_the_api_name)

3) Run the program and enter your queries

    //from inside the `freetime` directory
    go run *.go

##Usage Examples

After installing, if you see a new prompt `Freetime`, then Freetime is working!

Now, time for you to enter constraint statement for a query. Constraint statements are of the form:

    [duration] [hour range] [dates]

For example, `1hour 1pm-5pm 11/10` is a constraint statement that, when used to query, will return all non-busy 1 hour-long time slots from 1pm to 5pm on 11/10
