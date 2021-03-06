[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/ApiGateway)](https://goreportcard.com/report/github.com/Ulbora/ApiGateway)
[![](https://img.shields.io/docker/build/mariobehling/loklak.svg)](https://hub.docker.com/r/ulboralabs/apigateway/builds/)
[![sqale_rating](https://sonarcloud.io/api/project_badges/measure?project=apigateway&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=apigateway)
[![reliability_rating](https://sonarcloud.io/api/project_badges/measure?project=apigateway&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=apigateway)
[![coverage](https://sonarcloud.io/api/project_badges/measure?project=apigateway&metric=coverage)](https://sonarcloud.io/dashboard?id=apigateway)



API Gateway (runs inside user's private network)
==============

Get a **free** account on the **Free Tier** at: http://www.myapigateway.com

A lightweight API Gateway that runs inside a user's private network with a self service portal at: http://www.myapigateway.com

Copyright (C) 2018 Ulbora Labs LLC. (www.ulboralabs.com)
All rights reserved.

Copyright (C) 2018 Ken Williamson
All rights reserved.

# Run on
- Pivotal Cloud Foundry
- Docker Swarm
- Kubernetes
- Many others


User Admin Portal: https://github.com/Ulbora/ApiGatewayUserPortal.git

# Features
- Circuit Breaker
- Health Check
- Self Healing when breaker is open
- Gateway Analytics
- Blue/Green/Active Routes
- Gateway Error Loggin
- Admin Portal (written in Golang)


## Headers For Gateway Route Calls
- clientId: Your assigned client id
- apiKey: Your assigned API Key
- Any other headers required for your micro services

## Allowed HTTP Methods
- POST
- PUT
- PATCH
- GET
- DELETE
- OPTIONS


## Gateway Routes
### Local Non-Prod
- http://localhost:3011/np/routeID/routeName/yourRoute
- (example): http://localhost:3011/np/challenge/blue/rs/challenge/en_us?g=g&b=b
- Note: 
- routeID is: challenge
- routeName is: blue
- yourRoute is: rs/challenge/en_us?g=g&b=b which can be mappend in the user portal to something like https://www.youapi/rs/challenge/en_us?g=g&b=b

### Local Prod

- http://localhost:3011/routeID/yourRoute
- (example): http://localhost:3011/challenge/rs/challenge?name=sam&age=44
- Note: 
- routeID is: challenge
- yourRoute is: /rs/challenge?name=sam&age=44 which can be mappend in the user portal to something like https://www.youapi/rs/challenge?name=sam&age=44

### Active Production Route
The User Admin Portal allows you to make any route URL the active production route with the click of a switch.
Using Non-Prod routes allows you to test services before placing them in production.

