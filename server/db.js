const { Sequelize } = require('sequelize');

const sequilize = new Sequelize('gobuff_realworld_example_app_development', 'postgres', 'postgres', {
  host: '127.0.0.1',
  dialect: 'postgres',
  port: 5432,
});

async function clear() {
  const t = await sequilize.transaction();

  try {
    await sequilize.query('DELETE FROM article_tags;')
    await sequilize.query('DELETE FROM articles;')
    await sequilize.query('DELETE FROM tags;')
    await sequilize.query('DELETE FROM users;')
    await sequilize.query('DELETE FROM comments;')

    await t.commit();

    console.log('DB was cleared');
  } catch (error) {
    await t.rollback();

    console.log(`Can't clear DB`);
  }
}

module.exports = { clear };
