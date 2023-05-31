// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import PropTypes from 'prop-types';
import React from 'react';

type Props = {
    id?: string;
    show: boolean;
    header?: React.ReactNode;
    title?: React.ReactNode;
    subtitle?: React.ReactNode;
    children?: React.ReactNode;
    container?: boolean;
};

export default class SettingsGroup extends React.PureComponent<Props> {
    static propTypes = {
        id: PropTypes.string,
        show: PropTypes.bool.isRequired,
        header: PropTypes.node,
        title: PropTypes.node,
        subtitle: PropTypes.node,
        children: PropTypes.node,
        container: PropTypes.bool,
    };

    static defaultProps = {
        show: true,
        container: true,
    };

    render(): React.ReactNode {
        let wrapperClass = '';
        let contentClass = '';

        if (!this.props.show) {
            return null;
        }

        if (this.props.container) {
            wrapperClass = 'admin-console__wrapper';
            contentClass = 'admin-console__content';
        }

        let header = null;
        if (this.props.header) {
            header = <h4>{this.props.header}</h4>;
        }

        let title = null;
        if (!this.props.header && this.props.title) {
            title = <div className={'section-title'}>{this.props.title}</div>;
        }

        let subtitle = null;
        if (!this.props.header && this.props.subtitle) {
            subtitle = (
                <div className={'section-subtitle'}>{this.props.subtitle}</div>
            );
        }

        let sectionHeader = null;
        if (title || subtitle) {
            sectionHeader = (
                <div className={'section-header'}>
                    {title}
                    {subtitle}
                </div>
            );
        }

        return (
            <div className={wrapperClass}>
                <div className={contentClass}>
                    {header}
                    {sectionHeader}
                    {this.props.children}
                </div>
            </div>
        );
    }
}
