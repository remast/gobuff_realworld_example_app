/// <reference types="cypress" />

describe('Sign in page', () => {
  let user;

  beforeEach(() => {
    cy.task('db:clear');

    cy.task('newUser').then(newUser => {
      user = newUser;
    });
    cy.visit('/');
  });

  it('should provide an ability to sign in with existing credentials', () => {
    cy.task('db:seed');

    cy.getByDataCy('sign-in')
      .click();
    cy.getByDataCy('email-sign-in')
      .type('test@mail.com');
    cy.getByDataCy('password-sign-in')
      .type('12345Qwert!');
    cy.getByDataCy('sign-in-submit')
      .click();
    
    cy.getByDataCy('username-home')
      .should('contain', 'test');
    cy.getByDataCy('register-alert')
      .should('contain', 'Welcome Back to Buffalo!');
  });

  it('should redirect to sign up page by click on the "Need an account?" link', () => {
    cy.visit('/auth/login');
    cy.getByDataCy('redirect-sign-up')
      .click();
    
    cy.url()
      .should('include', '/users/register');
    cy.contains('h1', 'Sign up')
      .should('exist');
  });

  it('should validate inputs for non existing credentials', () => {
    cy.visit('/auth/login');
    cy.getByDataCy('email-sign-in')
      .type(user.email);
    cy.getByDataCy('password-sign-in')
      .type(user.password);
    cy.getByDataCy('sign-in-submit')
      .click();

    cy.contains('div', 'invalid email/password')
      .should('exist');
  });
});
