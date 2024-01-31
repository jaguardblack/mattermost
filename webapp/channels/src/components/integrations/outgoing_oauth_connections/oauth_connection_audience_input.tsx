// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {useEffect} from 'react';
import {FormattedMessage} from 'react-intl';
import {useDispatch, useSelector} from 'react-redux';

import {OauthIcon, InformationOutlineIcon} from '@mattermost/compass-icons/components';
import type {OutgoingOAuthConnection} from '@mattermost/types/integrations';

import {getOutgoingOAuthConnections as fetchOutgoingOAuthConnections} from 'mattermost-redux/actions/integrations';
import {getConfig} from 'mattermost-redux/selectors/entities/general';
import {getOutgoingOAuthConnections} from 'mattermost-redux/selectors/entities/integrations';

type Props = {
    value: string;
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
    placeholder: string;
}

const matchConnectionToSubmitUrl = (submitUrl: string, connections: OutgoingOAuthConnection[]): OutgoingOAuthConnection | null => {
    for (const conn of connections) {
        for (const audienceUrl of conn.audiences) {
            if (audienceUrl.endsWith('/*')) {
                const withoutSuffix = audienceUrl.substring(0, audienceUrl.length - 2);
                if (submitUrl === withoutSuffix || submitUrl.startsWith(withoutSuffix + '/')) {
                    return conn;
                }
            } else if (audienceUrl === submitUrl) {
                return conn;
            }
        }
    }

    return null;
};

const OAuthConnectionAudienceInput = (props: Props) => {
    const oauthConnections = useSelector(getOutgoingOAuthConnections);
    const oauthConnectionsEnabled = useSelector(getConfig).EnableOutgoingOAuthConnections === 'true';

    const dispatch = useDispatch();
    useEffect(() => {
        if (oauthConnectionsEnabled) {
            dispatch(fetchOutgoingOAuthConnections());
        }
    }, [dispatch, oauthConnectionsEnabled]);

    const connections = Object.values(oauthConnections);

    const input = (
        <input
            autoComplete='off'
            id='url'
            maxLength={1024}
            className='form-control'
            value={props.value}
            onChange={props.onChange}
            placeholder={props.placeholder}
        />
    );

    if (!connections.length) {
        return input;
    }

    const matchedConnection = matchConnectionToSubmitUrl(props.value, connections);
    let oauthMessage: React.ReactNode;

    if (matchedConnection) {
        oauthMessage = (
            <>
                <OauthIcon
                    size={20}
                />
                <strong>
                    <FormattedMessage
                        id='add_outgoing_oauth_connection.connected'
                        defaultMessage='Connected to "{connectionName}"'
                        values={{
                            connectionName: matchedConnection.name,
                        }}
                    />
                </strong>
            </>
        );
    } else {
        oauthMessage = (
            <>
                <InformationOutlineIcon
                    size={20}
                />
                <strong>
                    <FormattedMessage
                        id='add_outgoing_oauth_connection.not_connected'
                        defaultMessage='Not linked to an OAuth connection'
                    />
                </strong>
            </>
        );
    }

    return (
        <>
            {input}
            <div className='outgoing-oauth-audience-match-message'>
                {oauthMessage}
            </div>
        </>
    );
};

export default OAuthConnectionAudienceInput;
