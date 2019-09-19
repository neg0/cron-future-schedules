# Human Readable Cron Parser
__Requirements:__
 * Golang 1.12.6+

The scheduler configuration looks like this:
```
30 1 /bin/run_me_daily
45 * /bin/run_me_hourly
* * /bin/run_me_every_minute
* 19 /bin/run_me_sixty_times
```

> The first field is the minutes past the hour, the second field is the hour of the day and the third is the command to run. For both cases * means that it should run for all values of that field. In the above example run_me_daily has been set to run at 1:30am every day and run_me_hourly at 45 minutes past the hour every hour. The fields are whitespace separated and each entry is on a separate line.

## Acceptance Criteria
We would like you to write a command line program 
that outputs the soonest time at which each of the commands will run 
and whether this will be today or tomorrow.

The config input will be via STDIN, and the output should be via STDOUT. 
The 'current time' will be the single command line argument to the program in the format HH:MM.

For example given the above examples as input and the simulated 'current time' 
command-line argument 16:10 the output should be:
```
1:30 tomorrow - /bin/run_me_daily
16:45 today - /bin/run_me_hourly
16:10 today - /bin/run_me_every_minute
19:00 today - /bin/run_me_sixty_times
```

# Build
I created the most simple model for cron jobs, each cron string will be converted to `Job` object and `Job` object 
holds the algorithm to show next to run commands. Architecture is Rich Domain Model that's why `processNoneRecurringJob` 
and `NextSchedule` is responsible for creation of next scheduled commands.

If I have taken Anemic Domain Model I should have created an abstract services to be injected or directly being used 
for `NextSchedule` or `ProcessNoneRecurringJob`. More information on why decision this way; please see Martin Fowler 
blog post regarding anti-pattern Anemic Model: [AnemicDomainModel](https://martinfowler.com/bliki/AnemicDomainModel.html).


## CLI App
Please use `make` command to build the binary

    ~$: make build

application gets compiled inside a newly created folder called `bin`, and you could run the app with an argument `current-time` 
to list the next scheduled cron jobs.

    ~$: ./bin/cron_parser --current-time=16:10

> By default if `current-time` is not passed in as an argument, application would assume the current time as default