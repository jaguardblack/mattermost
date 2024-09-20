// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import classNames from 'classnames';
import React from 'react';

import './action.scss';
import WithTooltip from 'components/with_tooltip';

type Props = {
    icon: string;
    id: string;
    name: string;
    onClick: (e: React.MouseEvent<HTMLButtonElement>) => void;
    tooltipText: React.ReactElement | string;
    disabled?: boolean;
};

function Action({
    name,
    icon,
    onClick,
    id,
    tooltipText,
    disabled,
}: Props) {
    return (
        <div className='DraftAction'>
            <WithTooltip
                id={`drafts_action_tooltip_${id}`}
                placement='top'
                title={tooltipText}
            >
                <button
                    className={classNames(
                        'DraftAction__button',
                        {'DraftAction__button--delete': name === 'delete'},
                    )}
                    id={`draft_${icon}_${id}`}
                    onClick={onClick}
                    disabled={disabled}
                >
                    <i
                        className={classNames(
                            'icon',
                            icon,
                            {disabled},
                        )}
                    />
                </button>
            </WithTooltip>
        </div>
    );
}

export default Action;
