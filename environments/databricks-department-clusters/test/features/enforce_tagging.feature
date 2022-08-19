Feature: Enfore tagging
  Scenario: Ensure all resources have custom_tags
    Given I have resource that supports custom_tags defined
    Then it must contain custom_tags
    And its value must not be null