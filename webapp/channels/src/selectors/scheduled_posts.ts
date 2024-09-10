// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import type {GlobalState} from 'types/store';

export function getScheduledPostsByTeam(state: GlobalState, teamId: string) {
    return state.views.scheduledPosts.byTeamId[teamId];
}

export function getScheduledPostsByTeamCount(state: GlobalState, teamId: string) {
    return state.views.scheduledPosts.byTeamId[teamId]?.length || 0;
}
