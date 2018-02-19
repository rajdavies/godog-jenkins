@basic
Feature: Basic
  As a user I can perform basic operations of Jenkins X

  Scenario: User will get a list of available commands when running jx
    When executing "jx"
    Then stdout should contain
         """
         Usage
         """
   Scenario: get shell code for using completion with jx
    When executing "jx completion"
    Then stdout should contain
         """
         _jx_
          """


   Scenario: change Kubernetes namespace for context using jx
    When executing "jx namespace"
    Then stdout should contain
          """
          Using namespace
          """

    Scenario: generate shell prompt using jx
    When executing "jx prompt"
    Then stdout should contain
          """
          k8s
          """

   Scenario: open the Jenkins console running on my cluster
       When executing "jx console" my web browser should open a page to the Jenkins console of Jenkins running in my cluster

