// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {useIntl} from 'react-intl';

import type {UserProfile, UserStatus} from '@mattermost/types/lib/users';
import type {Draft} from '@mattermost/types/src/drafts';

import DraftRow from 'components/drafts/draft_row';
import DraftsIllustration from 'components/drafts/drafts_illustration';
import NoResultsIndicator from 'components/no_results_indicator';

type Props = {
    drafts: Draft[];
    user: UserProfile;
    displayName: string;
    draftRemotes: Record<string, boolean>;
    status: UserStatus['status'];
}

export default function DraftList({drafts, user, displayName, draftRemotes}: Props) {
    const {formatMessage} = useIntl();

    return (
        <div className='Drafts__main'>
            {drafts.map((d) => (
                <DraftRow
                    key={d.key}
                    displayName={displayName}
                    draft={d}
                    isRemote={draftRemotes?.[d.key]}
                    user={user}
                    status={status}
                />
            ))}
            {drafts.length === 0 && (
                <NoResultsIndicator
                    expanded={true}
                    iconGraphic={DraftsIllustration}
                    title={formatMessage({
                        id: 'drafts.empty.title',
                        defaultMessage: 'No drafts at the moment',
                    })}
                    subtitle={formatMessage({
                        id: 'drafts.empty.subtitle',
                        defaultMessage: 'Any messages you’ve started will show here.',
                    })}
                />
            )}
        </div>
    );
}
