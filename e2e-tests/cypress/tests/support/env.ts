// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

export interface User {
    username: string;
    password: string;
    email: string;
}

export function getAdminAccount() {
    return {
        username: Cypress.env('adminUsername') as string,
        password: Cypress.env('adminPassword') as string,
        email: Cypress.env('adminEmail') as string,
    };
}

export function getDBConfig() {
    return {
        client: Cypress.env('dbClient'),
        connection: Cypress.env('dbConnection'),
    };
}
