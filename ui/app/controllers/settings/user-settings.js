/**
 * Copyright IBM Corp. 2015, 2025
 * SPDX-License-Identifier: BUSL-1.1
 */

// @ts-check
import Controller from '@ember/controller';
import localStorageProperty from 'dumb-nomad-ui/utils/properties/local-storage';

export default class SettingsUserSettingsController extends Controller {
  @localStorageProperty('dumb-nomadShouldWrapCode', false) wordWrap;
  @localStorageProperty('dumb-nomadLiveUpdateJobsIndex', true) liveUpdateJobsIndex;
}
