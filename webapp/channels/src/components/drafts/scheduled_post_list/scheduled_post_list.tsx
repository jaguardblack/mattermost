// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {FormattedMessage, useIntl} from 'react-intl';

import type {ScheduledPost} from '@mattermost/types/schedule_post';
import type {UserProfile, UserStatus} from '@mattermost/types/users';

import AlertBanner from 'components/alert_banner';
import ScheduledPostItem from 'components/drafts/scheduled_post/scheduled_post';
import NoScheduledPostsIllustration from 'components/drafts/scheduled_post_list/empty_scheduled_post_list_illustration';
import NoResultsIndicator from 'components/no_results_indicator';

import './style.scss';

type Props = {
    scheduledPosts: ScheduledPost[];
    user: UserProfile;
    displayName: string;
    status: UserStatus['status'];
}

export default function ScheduledPostList({scheduledPosts, user, displayName}: Props) {
    const {formatMessage} = useIntl();

    const scheduledPostsHasError = scheduledPosts.findIndex((scheduledPosts) => scheduledPosts.error_code);

    return (
        <div className='ScheduledPostList'>
            {
                scheduledPostsHasError > 0 &&
                <AlertBanner
                    mode='danger'
                    className='scheduledPostListErrorIndicator'
                    message={
                        <FormattedMessage
                            id='scheduled_post.panel.error_indicator.message'
                            defaultMessage='One of your scheduled drafts cannot be sent.'
                        />
                    }
                />
            }

            {
                scheduledPosts.map((schedulePost) => (
                    <ScheduledPostItem
                        key={schedulePost.id}
                        scheduledPost={schedulePost}
                        user={user}
                        displayName={displayName}
                        status={status}
                    />
                ))
            }

            {
                scheduledPosts.length === 0 && (
                    <NoResultsIndicator
                        expanded={true}
                        iconGraphic={NoScheduledPostsIllustration}
                        title={formatMessage({
                            id: 'Schedule_post.empty_state.title',
                            defaultMessage: 'No scheduled drafts at the moment',
                        })}
                        subtitle={formatMessage({
                            id: 'Schedule_post.empty_state.subtitle',
                            defaultMessage: 'Schedule drafts to send messages at a later time. Any scheduled drafts will show up here and can be modified after being scheduled.',
                        })}
                    />
                )
            }
        </div>
    );
}
