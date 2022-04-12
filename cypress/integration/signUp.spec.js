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

  it('should register a user with valid data', () => {
    cy.intercept('POST', '/users/register').as('register');
    cy.getByDataCy('sign-up')
      .click();
    cy.getByDataCy('name')
      .type(user.username);
    cy.getByDataCy('email')
      .type(user.email);
    cy.getByDataCy('password')
      .type(user.password +'{enter}');

    cy.getByDataCy('username-home')
      .should('contain', user.username);
    cy.getByDataCy('register-alert')
      .should('contain', 'Welcome to RealWorld!');
  });

  it('should validate required fields', () => {
    cy.visit('/users/register');
    cy.getByDataCy('submit-register')
      .click();
    cy.contains('div', 'Name can not be blank.')
      .should('be.visible');
    cy.contains('div', 'Email can not be blank.')
      .should('be.visible');
    cy.contains('div', 'Password can not be blank.')
      .should('be.visible');
    cy.url()
      .should('include', '/users/register');
  });

  it('should redirect users to sign in page by "Have an account" button', () => {
    cy.visit('/users/register');
    cy.getByDataCy('have-acc-btn')
      .click();
    cy.url()
      .should('include', '/auth/login');
    cy.contains('h1', 'Sign in')
      .should('exist');
  });
});
