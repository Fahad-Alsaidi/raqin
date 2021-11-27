'use strict';

const _ = require('lodash');

const emailRegExp = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

const formatError = error => [
    { messages: [{ id: error.id, message: error.message, field: error.field }] },
];

const { sanitizeEntity, parseMultipartData } = require('strapi-utils');

module.exports = {

    /**
    * Update a record.
    *
    * @return {Object}
    */

    async updateMe(ctx) {
        const advancedConfigs = await strapi
            .store({
                environment: '',
                type: 'plugin',
                name: 'users-permissions',
                key: 'advanced',
            })
            .get();

        const { id } = ctx.state.user;
        const {
            request: { body },
        } = ctx;
        const { email, username, password } = body;

        let entity;

        const user = await strapi.plugins['users-permissions'].services.user.fetch({
            id,
        });

        if (!user || user.id !== id) {
            return ctx.unauthorized(`You can't update this entry`);
        }

        if (_.has(body, 'email') && !email) {
            return ctx.badRequest('email.notNull');
        }

        if (_.has(body, 'username') && !username) {
            return ctx.badRequest('username.notNull');
        }

        if (_.has(body, 'password') && !password && user.provider === 'local') {
            return ctx.badRequest('password.notNull');
        }

        // Check if the provided email is valid or not.
        const isEmail = emailRegExp.test(email);

        if (_.has(body, 'email') && advancedConfigs.unique_email) {
            const userWithSameEmail = await strapi
                .query('user', 'users-permissions')
                .findOne({ email: email.toLowerCase() });

            if (userWithSameEmail && userWithSameEmail.id != id) {
                return ctx.badRequest(
                    null,
                    formatError({
                        id: 'Auth.form.error.email.taken',
                        message: 'Email already taken',
                        field: ['email'],
                    })
                );
            }
            if (!isEmail) {
                return ctx.badRequest(
                    null,
                    formatError({
                        id: 'Auth.form.error.email.format',
                        message: 'Please provide valid email address.',
                    })
                );
            }
            body.email = body.email.toLowerCase();
        }

        entity = await strapi.plugins['users-permissions'].services.user.edit({ id }, { username, email, password });

        return sanitizeEntity(entity, { model: strapi.plugins['users-permissions'].models.user });
    },

};
