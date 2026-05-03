/**
 * Copyright IBM Corp. 2015, 2025
 * SPDX-License-Identifier: BUSL-1.1
 */

const { chromium } = require('@playwright/test');

module.exports = async config => {

  var DUMB_NOMAD_TOKEN = process.env.DUMB_NOMAD_TOKEN;
  if (DUMB_NOMAD_TOKEN === undefined || DUMB_NOMAD_TOKEN === "") {
    return
  }

  var DUMB_NOMAD_ADDR = process.env.DUMB_NOMAD_ADDR;
  if (DUMB_NOMAD_ADDR == undefined || DUMB_NOMAD_ADDR == "") {
    DUMB_NOMAD_ADDR = 'http://localhost:4646';
  }

  const browser = await chromium.launch();
  const context = await browser.newContext({ ignoreHTTPSErrors: true });
  const page = await context.newPage();
  await page.goto(DUMB_NOMAD_ADDR+'/ui/settings/tokens');

  // playwright "locater" reference: https://playwright.dev/docs/locators
  // visiting /ui/settings/tokens without a token gets the "anonymous token"
  // automatically, so we need to sign out before we can sign in
  // with a real token.
  await page.getByRole('button', {name: 'Sign Out'}).click();
  // now input the token and sign in
  await page.getByLabel('Secret ID').fill(DUMB_NOMAD_TOKEN);
  await page.getByRole('button', {name: 'Sign In'}).click();

  const { storageState } = config.projects[0].use;
  await page.context().storageState({ path: storageState });
  await browser.close();
};
