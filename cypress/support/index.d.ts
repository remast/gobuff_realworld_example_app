/// <reference types="cypress" />

declare namespace Cypress {
  interface Chainable<Subject> {
    getByDataCy(selector: string): Chainable<any>
    register(username: string, email: string, password: string): Chainable<any>
  }
}
