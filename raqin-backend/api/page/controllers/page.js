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
    const { id } = ctx.params;

    let entity;

    const [page] = await strapi.services.page.find({
      pageNumber: ctx.request.body.pageNumber,
      'book.id': ctx.request.body.book,
      'user.id': ctx.state.user.id
    });

    if (page) {
      return ctx.response.conflict('resource already exists');
    }

    if (ctx.is('multipart')) {
      const { data, files } = parseMultipartData(ctx);
      data.user = ctx.state.user.id;
      entity = await strapi.services.page.create(data, { files });
    } else {
      ctx.request.body.user = ctx.state.user.id;
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

    const [page] = await strapi.services.page.find({
      id: ctx.params.id,
      'user.id': ctx.state.user.id,
    });

    if (!page) {
      return ctx.unauthorized(`You can't update this entry`);
    }

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

    const [page] = await strapi.services.page.find({
      id: ctx.params.id,
      'user.id': ctx.state.user.id,
    });

    if (!page) {
      return ctx.unauthorized(`You can't delete this entry`);
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