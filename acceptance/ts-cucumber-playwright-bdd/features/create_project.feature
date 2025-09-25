Feature: Create project

  Users can create projects, only visible to themselves

  Scenario: Create one project
    Given Sue has signed up
    When Sue creates a project
    Then Sue should see the project

  Scenario: Try to see someone else's project
    Given Sue has signed up
    And Bob has signed up
    When Sue creates a project
    Then Bob should not see any projects
