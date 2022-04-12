/// <reference types="cypress" />

describe('Sign up page', () => {
  let user;

  before(() => {
    cy.task('db:clear');
    cy.task('newUser').then(newUser => {
      user = newUser;
    });
    cy.visit('/');
  });

  it('', () => {
    
  });
});
