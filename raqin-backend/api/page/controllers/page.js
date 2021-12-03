'use strict';

/**
 * Read the documentation (https://strapi.io/documentation/developer-docs/latest/development/backend-customization.html#core-controllers)
 * to customize this controller
 */

const { parseMultipartData, sanitizeEntity } = require('strapi-utils');

module.exports = {
  /**
   * Create a record.
   *
   * @return {Object}
   */

  async create(ctx) {

    let entity;

    const userId = ctx.state.user.id;
    ctx.request.body.user = userId;
    const page = await strapi.services.page.findOne({
      pageNumber: ctx.request.body.pageNumber,
      'book.id': ctx.request.body.book,
      'user.id': userId
    });

    if (page) {
      return ctx.response.conflict('page already exists');
    }

    const book = await strapi.services.book.findOne({
      id: ctx.request.body.book,
    });
    if (ctx.request.body.pageNumber > book?.numberOfPages) {
      return ctx.response.conflict(`number of page can't exceed max number of book pages`);
    }

    if (ctx.is('multipart')) {
      const { data, files } = parseMultipartData(ctx);
      entity = await strapi.services.page.create(data, { files });
    } else {
      entity = await strapi.services.page.create(ctx.request.body);
    }
    return sanitizeEntity(entity, { model: strapi.models.page });
  },

  /**
  * Update a record.
  *
  * @return {Object}
  */

  async update(ctx) {
    const { id } = ctx.params;

    let entity;

    const userId = ctx.state.user.id;
    const page = await strapi.services.page.findOne({
      id: id,
      'user.id': userId,
    });

    if (!page) {
      return ctx.unauthorized(`You can't update this entry`);
    }
    if (page.stage == "rev" || page.stage == "done") {
      return ctx.locked(`page is already in revision state or has been approved`);
    }

    // delete attributes that shouldn't be updated
    delete ctx.request.body.pageNumber;
    delete ctx.request.body.book;
    delete ctx.request.body.user;

    if (ctx.is('multipart')) {
      const { data, files } = parseMultipartData(ctx);
      entity = await strapi.services.page.update({ id }, data, {
        files,
      });
    } else {
      entity = await strapi.services.page.update({ id }, ctx.request.body);
    }

    return sanitizeEntity(entity, { model: strapi.models.page });
  },

  /**
   * Update a record.
   *
   * @return {Object}
   */

  async delete(ctx) {
    const { id } = ctx.params;

    let entity;

    const page = await strapi.services.page.findOne({
      id: ctx.params.id,
      'user.id': ctx.state.user.id,
    });

    if (!page) {
      return ctx.unauthorized(`You can't delete this entry`);
    }
    if (page.stage == "rev" || page.stage == "done") {
      return ctx.locked(`page can't be deleted if in revision state or has been approved`);
    }

    if (ctx.is('multipart')) {
      const { data, files } = parseMultipartData(ctx);
      entity = await strapi.services.page.delete({ id }, data, {
        files,
      });
    } else {
      entity = await strapi.services.page.delete({ id }, ctx.request.body);
    }

    return sanitizeEntity(entity, { model: strapi.models.page });
  },

  /**
  * Count records.
  *
  * @return {Number}
  */

  async countDoneOfBook(ctx) {

    const { id } = ctx.params;

    const pages = await strapi.services.page.find({
      book: id,
      stage: 'done',
    });

    let distinctPages = [...new Set(pages.map(x => x.pageNumber))]
    return distinctPages.length

  },
};