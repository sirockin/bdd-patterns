Feature: Sign up

  New accounts need to confirm their email to activate
  their account.

  Scenario: Successful sign-up
    Given Tanya has created an account
    When Tanya activates her account
    Then Tanya should be authenticated

  Scenario: Try to sign in without activating account
    Given Bob has created an account
    When Bob tries to sign in
    Then Bob should not be authenticated
    And Bob should see an error telling him to activate the account
