# dashboard-app

## Architecture Run-Down

### Main PHP Codebase
This will be the code which sends pages, retrieves data, and stores sessions
It will reach out to the microservices that are needed at a given time

### GO Microservices
These will handle all the modular things like sign-ins and using the ebay api

## Run the codebases

To run the codebases, treat each directory like a separate server. CD into go, and run main.go on port 3000. Then, go to the php directory, and run this on port 8000.

To load the webpages, use :8000. Every microservice will run under 3000 so don't create any conflicts with the port numbers