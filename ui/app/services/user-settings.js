/**
 * Copyright IBM Corp. 2015, 2025
 * SPDX-License-Identifier: BUSL-1.1
 */

import Service from '@ember/service';
import localStorageProperty from 'dumb-nomad-ui/utils/properties/local-storage';

export default class UserSettingsService extends Service {
  @localStorageProperty('dumb-nomadPageSize', 25) pageSize;
  @localStorageProperty('dumb-nomadLogMode', 'stdout') logMode;
  @localStorageProperty('dumb-nomadTopoVizPollingNotice', true)
  showTopoVizPollingNotice;
}
