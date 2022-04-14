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

async function seed() {
  const t = await sequilize.transaction();

  try {
    await sequilize.query(`INSERT INTO users (
      "created_at", 
      "email",
      "id",
      "name",
      "password_hash",
      "updated_at")
    VALUES (
      '2022-04-13 19:46:14.539064',
      'test@mail.com',
      '4ebd00a9-b527-4802-9aae-abf5bfed5f76',
      'test',
      '$2a$10$obnFbyhuea941I8KT43lIeoCBb3Ms/0Ltwj65lO/iTpMDfpPtThby',
      '2022-04-13 19:46:14.539064');`)

    await t.commit();

    console.log('User was added');
  } catch (error) {
    await t.rollback();

    console.log(`Can't seed DB`);
  }
}

module.exports = { clear, seed };
