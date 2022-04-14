/// <reference types="cypress" />

describe('User of the Counduit', () => {
  let user;
  let article;

  beforeEach(() => {
    cy.task('db:clear');

    cy.task('newUser').then(newUser => {
      user = newUser;
    });
    cy.task('newArticle').then(newArticle => {
      article = newArticle;
    });
    cy.visit('/');
  });

  it.only('should be able to create an article', () => {
    cy.register(user.username, user.email, user.password);
    cy.getByDataCy('new-article')
      .click();
    cy.getByDataCy('article-title')
      .type(article.title);
    cy.getByDataCy('article-description')
      .type(article.description);
    cy.getByDataCy('article-body')
      .type(article.body);
    cy.getByDataCy('article-tag')
      .type(article.tag);
    cy.getByDataCy('submit-article')
      .click();
    
    cy.getByDataCy('alert')
      .should('contain', 'Article created');
    cy.url()
      .should('include', article.title);
    cy.get('h1')
      .should('contain', article.title);
    cy.getByDataCy('body')
      .should('contain', article.body);
    cy.getByDataCy('tags')
      .should('contain', article.tag);
  });

  it('should be able to edit an article', () => {
    
  });

  it('should be able to delete an article', () => {
    
  });
});
