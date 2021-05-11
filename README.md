
BE Candidate Take Home Test
Have you shopped online? Let’s imagine that you need to build the checkout backend service
that will support different promotions with the given inventory.

### Documentation
git pull https://github.com/luterrinding/TestShopped.git


runing$ go build && ./TestShopped 

Unit Test : go test main.go main_test.go -v

Example Scenarios:
Scanned Items: MacBook Pro, Raspberry Pi B
Total: $5,399.99
Scanned Items: Google Home, Google Home, Google Home
Total: $99.98
Scanned Items: Alexa Speaker, Alexa Speaker, Alexa Speaker
Total: $295.65
Please write it in Golang or Node with a CI script that runs tests and produces a binary.


Result Test

=== RUN   TestScenario1

    main_test.go:154: Scenario1 : 5399.99

--- PASS: TestScenario1 (0.00s)

=== RUN   TestScenario2

    main_test.go:164: Scenario2 : 99.98

--- PASS: TestScenario2 (0.00s)

=== RUN   TestScenario3

    main_test.go:173: Scenario3 : 295.65

--- PASS: TestScenario3 (0.00s)

PASS

ok  	command-line-arguments	(cached)

