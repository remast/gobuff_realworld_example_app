/// <reference types="cypress" />
// ***********************************************************
// This example plugins/index.js can be used to load plugins
//
// You can change the location of this file or turn off loading
// the plugins file with the 'pluginsFile' configuration option.
//
// You can read more here:
// https://on.cypress.io/plugins-guide
// ***********************************************************

// This function is called when a project is opened or re-opened (e.g. due to
// the project's config changing)

/**
 * @type {Cypress.PluginConfig}
 */
// eslint-disable-next-line no-unused-vars

const faker = require("faker");
const { clear } = require("../../server/db");

module.exports = (on, config) => {
  on("task", {
    newUser() {
        return {
          username: faker.name.firstName() + `${Math.round(Math.random(1000) * 1000)}`,
          email: 'test'+`${Math.round(Math.random(1000) * 1000)}`+'@mail.com',
          password: '12345Qwert!',
        };
    },
    newArticle() {
      article = {
        title: faker.lorem.word(),
        description: faker.lorem.words(),
        body: faker.lorem.words(),
        tag: faker.lorem.word()
      };
      return article;
    },
    'db:clear'() {
      clear();

      return null;
    }
  });
};
