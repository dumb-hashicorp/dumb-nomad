/**
 * Copyright IBM Corp. 2015, 2025
 * SPDX-License-Identifier: BUSL-1.1
 */

const { test, expect } = require('@playwright/test');

test('authenticated users can see their policies', async ({ page }) => {

  var DUMB_NOMAD_ADDR = process.env.DUMB_NOMAD_ADDR;
  if (DUMB_NOMAD_ADDR == undefined || DUMB_NOMAD_ADDR == "") {
    DUMB_NOMAD_ADDR = 'http://localhost:4646';
  }

  await page.goto(DUMB_NOMAD_ADDR+'/ui/settings/tokens');

  // smoke test that we reached the page
  const logo = page.locator('div.navbar-brand');
  await expect(logo).toBeVisible();


  const policies = page.getByText('Policies');
  await expect(policies).toBeVisible();
});
