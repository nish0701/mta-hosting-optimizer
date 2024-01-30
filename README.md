# mta-hosting-optimizer

get "/hostnames" gives the hostnames which have inactive mta count more than the defined threshold.
For now, the data has been mocked to define the inactive mtas in services/data_service. We can also have another service which dynamically generates this data and gives dynamic output.
The project is hosted on gitHub instead of gitLab.

export THRESHOLD = "SOME_VALUE"
go run main.go
to bring the http server up on port 8080

Code Coverage exceeds 90.7%, integration tests written. Tried to integrate the visual report gerneation, but Cobertura is failing as it is not able to find cover.out, though it is being generated in the local.

Integrate the test and build phases to Github action Integrated in github action Added
