// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import moment from 'moment';
import React, {useMemo} from 'react';
import {FormattedMessage, useIntl} from 'react-intl';

import OverlayTrigger from 'components/overlay_trigger';
import Tooltip from 'components/tooltip';

import Constants from 'utils/constants';

interface Props {
    date?: number;
}

export function ElapsedDurationCell(props: Props) {
    const todaysDate = moment().startOf('day').valueOf();

    const {formatMessage} = useIntl();

    const {elapsedDays, exactPassedInDate} = useMemo(() => {
        const todayMoment = moment();
        const passedInDateMoment = moment(props.date);

        // TODO: Use Timestamp component here
        const elapsedDays = todayMoment.diff(passedInDateMoment, 'days');
        const exactPassedInDate = passedInDateMoment.format(
            `MMMM DD, Y [${formatMessage({id: 'admin.console.list.table.exactTime.at', defaultMessage: 'at'})}] hh:mm:ss A`,
        );

        return {
            elapsedDays,
            exactPassedInDate,
        };
    }, [props.date, todaysDate]);

    if (!props.date) {
        return null;
    }

    let elapsedDaysText = null;
    if (elapsedDays < 1) {
        elapsedDaysText = (
            <FormattedMessage
                id='admin.system_users.list.memberSince.today'
                defaultMessage='Today'
            />
        );
    } else if (elapsedDays === 1) {
        elapsedDaysText = (
            <FormattedMessage
                id='admin.system_users.list.memberSince.yesterday'
                defaultMessage='Yesterday'
            />
        );
    } else {
        elapsedDaysText = (
            <FormattedMessage
                id='admin.system_users.list.memberSince.days'
                defaultMessage='{days} days'
                values={{days: elapsedDays}}
            />
        );
    }

    const sharedTooltip = (
        <Tooltip id='system-users-cell-tooltip'>
            {exactPassedInDate}
        </Tooltip>
    );

    return (
        <OverlayTrigger
            delayShow={Constants.OVERLAY_TIME_DELAY}
            placement='bottom'
            overlay={sharedTooltip}
        >
            <span>
                {elapsedDaysText}
            </span>
        </OverlayTrigger>
    );
}
