Feature: Server is up and running
  In order to serve requests to the server
  As a client
  I need to have a server up and running

  Scenario: Server returns a successful response the from liveness probe URI
    Given an HTTP "GET" request with the URI "http://localhost/liveness"
    Then the server must reply with a status code 200
    And the server must add a trace ID
    And the server must reply with a body:
      """
      ["Server is live"]
      """

  Scenario: Server returns a successful response the from liveness probe URI reusing a trace ID
    Given an HTTP "GET" request with the URI "http://localhost/liveness" and a trace ID "a0os3jkcy"
    Then the server must reply with a status code 200
    And the server must reply with the trace ID "a0os3jkcy"
    And the server must reply with a body:
      """
      ["Server is live"]
      """

  Scenario: Server returns a successful response the from readiness probe URI
    Given an HTTP "GET" request with the URI "http://localhost/readiness"
    Then the server must reply with a status code 200
    And the server must add a trace ID
    And the server must reply with a body:
      """
      ["Server is ready"]
      """

  Scenario: Server returns a 405 status code from the root URI with an invalid request method
    Given an HTTP "POST" request with the URI "http://localhost/liveness"
    Then the server must reply with a status code 405
    And the server must add a trace ID
    And the server must reply with a body:
      """
      """

  Scenario: Server returns a 405 status code from a non existing routes
    Given an HTTP "GET" request with the URI "http://localhost/undefined"
    Then the server must reply with a status code 404
    And the server must add a trace ID
    And the server must reply with a body:
      """
      """