@transfer
Feature: Send transfer transaction
  As an external developer
  I want to process a multiple transfer transaction using transaction scheduler API

  Background:
    Given I have the following tenants
      | alias   | tenantID        |
      | tenant1 | {{random.uuid}} |
    And I have created the following accounts
      | alias    | ID              | API-KEY            | Tenant               |
      | account1 | {{random.uuid}} | {{global.api-key}} | {{tenant1.tenantID}} |
      | account2 | {{random.uuid}} | {{global.api-key}} | {{tenant1.tenantID}} |
    Then I track the following envelopes
      | ID                  |
      | faucet-{{account1}} |
      | faucet-{{account2}} |
    Given I set the headers
      | Key         | Value                |
      | X-API-KEY   | {{global.api-key}}   |
      | X-TENANT-ID | {{tenant1.tenantID}} |
    When I send "POST" request to "{{global.api}}/transactions/transfer" with json:
      """
      {
        "chain": "{{chain.besu0.Name}}",
        "params": {
          "from": "{{global.nodes.besu[0].fundedPublicKeys[0]}}",
          "to": "{{account1}}",
          "value": "0x16345785D8A0000"
        },
        "labels": {
          "scenario.id": "{{scenarioID}}",
          "id": "faucet-{{account1}}"
        }
      }
      """
    Then the response code should be 202
    When I send "POST" request to "{{global.api}}/transactions/transfer" with json:
      """
      {
        "chain": "{{chain.geth0.Name}}",
        "params": {
          "from": "{{global.nodes.geth[0].fundedPublicKeys[0]}}",
          "to": "{{account2}}",
          "value": "0x16345785D8A0000"
        },
        "labels": {
          "scenario.id": "{{scenarioID}}",
          "id": "faucet-{{account2}}"
        }
      }
      """
    Then the response code should be 202
    Then Envelopes should be in topic "tx.decoded"

  @geth
  Scenario: Send transfer transaction
    Given I register the following alias
      | alias           | value              |
      | to1             | {{random.account}} |
      | to2             | {{random.account}} |
      | transferTxOneID | {{random.uuid}}    |
      | transferTxTwoID | {{random.uuid}}    |
    Then I track the following envelopes
      | ID                  |
      | {{transferTxOneID}} |
      | {{transferTxTwoID}} |
    Given I set the headers
      | Key         | Value                |
      | X-API-KEY   | {{global.api-key}}   |
      | X-TENANT-ID | {{tenant1.tenantID}} |
    When I send "POST" request to "{{global.api}}/transactions/transfer" with json:
      """
      {
        "chain": "{{chain.besu0.Name}}",
        "params": {
          "from": "{{account1}}",
          "to": "{{to1}}",
          "value": "0x1DCD6500"
        },
        "labels": {
          "scenario.id": "{{scenarioID}}",
          "id": "{{transferTxOneID}}"
        }
      }
      """
    Then the response code should be 202
    Then I register the following response fields
      | alias      | path         |
      | jobOneUUID | jobs[0].uuid |
    When I send "POST" request to "{{global.api}}/transactions/transfer" with json:
      """
      {
        "chain": "{{chain.geth0.Name}}",
        "params": {
          "from": "{{account2}}",
          "to": "{{to2}}",
          "value": "0x400000000",
          "transactionType": "legacy"
        },
        "labels": {
          "scenario.id": "{{scenarioID}}",
          "id": "{{transferTxTwoID}}"
        }
      }
      """
    Then the response code should be 202
    Then I register the following response fields
      | alias      | path         |
      | jobTwoUUID | jobs[0].uuid |
    Then Envelopes should be in topic "tx.sender"
    And Envelopes should have the following fields
      | Nonce |
      | 0     |
      | 0     |
    Then Envelopes should be in topic "tx.decoded"
    And Envelopes should have the following fields
      | Receipt.Status |
      | 1              |
      | 1              |
    When I send "GET" request to "{{global.api}}/jobs/{{jobOneUUID}}"
    Then the response code should be 200
    And Response should have the following fields
      | status | logs[0].status | logs[1].status | logs[2].status | logs[3].status |
      | MINED  | CREATED        | STARTED        | PENDING        | MINED          |
    When I send "GET" request to "{{global.api}}/jobs/{{jobTwoUUID}}"
    Then the response code should be 200
    And Response should have the following fields
      | status | logs[0].status | logs[1].status | logs[2].status | logs[3].status |
      | MINED  | CREATED        | STARTED        | PENDING        | MINED          |
    When I send "POST" request to "{{global.api}}/proxy/chains/{{chain.besu0.UUID}}" with json:
      """
      {
        "jsonrpc": "2.0",
        "method": "eth_getBalance",
        "params": [
          "{{to1}}",
          "latest"
        ],
        "id": 1
      }
      """
    Then the response code should be 200
    And Response should have the following fields
      | result     |
      | 0x1dcd6500 |
    When I send "POST" request to "{{global.api}}/proxy/chains/{{chain.geth0.UUID}}" with json:
      """
      {
        "jsonrpc": "2.0",
        "method": "eth_getBalance",
        "params": [
          "{{to2}}",
          "latest"
        ],
        "id": 1
      }
      """
    Then the response code should be 200
    And Response should have the following fields
      | result     |
      | 0x17d78400 |

  Scenario: Fail to send transfer transaction with missing value
    Given I register the following alias
      | alias | value              |
      | to1   | {{random.account}} |
    Given I set the headers
      | Key         | Value                |
      | X-API-KEY   | {{global.api-key}}   |
      | X-TENANT-ID | {{tenant1.tenantID}} |
    When I send "POST" request to "{{global.api}}/transactions/transfer" with json:
      """
      {
        "chain": "{{chain.besu0.Name}}",
        "params": {
          "from": "{{account1}}",
          "to": "{{to1}}"
        },
        "labels": {
          "scenario.id": "{{scenarioID}}"
        }
      }
      """
    Then the response code should be 400
    And Response should have the following fields
      | code   | message |
      | 271104 | ~       |

  Scenario: Fail to send transfer transaction with missing To
    Given I register the following alias
      | alias | value              |
      | to1   | {{random.account}} |
    Given I set the headers
      | Key         | Value                |
      | X-API-KEY   | {{global.api-key}}   |
      | X-TENANT-ID | {{tenant1.tenantID}} |
    When I send "POST" request to "{{global.api}}/transactions/transfer" with json:
      """
      {
        "chain": "{{chain.besu0.Name}}",
        "params": {
          "from": "{{account1}}",
          "value": "0x17D78400"
        },
        "labels": {
          "scenario.id": "{{scenarioID}}"
        }
      }
      """
    Then the response code should be 400
    And Response should have the following fields
      | code   | message |
      | 271104 | ~       |

  Scenario: Fail to send transfer transaction with missing private-key
    Given I register the following alias
      | alias           | value              |
      | to1             | {{random.account}} |
      | transferTxOneID | {{random.uuid}}    |
      | account3        | {{random.account}} |
    Then I track the following envelopes
      | ID                  |
      | {{transferTxOneID}} |
    Given I set the headers
      | Key         | Value                |
      | X-API-KEY   | {{global.api-key}}   |
      | X-TENANT-ID | {{tenant1.tenantID}} |
    When I send "POST" request to "{{global.api}}/transactions/transfer" with json:
      """
      {
        "chain": "{{chain.besu0.Name}}",
        "params": {
          "from": "{{account3}}",
          "to": "{{to1}}",
          "value": "0x0"
        },
        "labels": {
          "scenario.id": "{{scenarioID}}",
          "id": "{{transferTxOneID}}"
        }
      }
      """
    Then the response code should be 422
