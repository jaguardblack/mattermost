// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {screen, fireEvent} from '@testing-library/react';
import React from 'react';

import type {AllowedIPRange} from '@mattermost/types/config';

import {renderWithIntl} from 'tests/react_testing_utils';

import EditSection from './';

describe('EditSection', () => {
    const ipFilters = [
        {
            CIDRBlock: '192.168.0.0/24',
            Description: 'Test Filter',
        },
    ] as AllowedIPRange[];
    const currentUsersIP = '192.168.0.1';
    const setShowAddModal = jest.fn();
    const setEditFilter = jest.fn();
    const handleConfirmDeleteFilter = jest.fn();
    const currentIPIsInRange = true;

    const baseProps = {
        ipFilters,
        currentUsersIP,
        setShowAddModal,
        setEditFilter,
        handleConfirmDeleteFilter,
        currentIPIsInRange,
    };

    test('renders the component', () => {
        renderWithIntl(
            <EditSection
                {...baseProps}
            />,
        );

        expect(screen.getByText('Allowed IP Addresses')).toBeInTheDocument();
        expect(screen.getByText('Create rules to allow access to the workspace for specified IP addresses only.')).toBeInTheDocument();
        expect(screen.getByText('If no rules are added, all IP addresses will be allowed.')).toBeInTheDocument();
        expect(screen.getByText('Add Filter')).toBeInTheDocument();
        expect(screen.getByText('Filter Name')).toBeInTheDocument();
        expect(screen.getByText('IP Address Range')).toBeInTheDocument();
        expect(screen.getByText('Test Filter')).toBeInTheDocument();
        expect(screen.getByText('192.168.0.0/24')).toBeInTheDocument();
    });

    test('clicking the Add Filter button calls setShowAddModal', () => {
        renderWithIntl(
            <EditSection
                {...baseProps}
            />,
        );

        fireEvent.click(screen.getByText('Add Filter'));

        expect(setShowAddModal).toHaveBeenCalledTimes(1);
        expect(setShowAddModal).toHaveBeenCalledWith(true);
    });

    test('clicking the Edit button calls setEditFilter', () => {
        renderWithIntl(
            <EditSection
                {...baseProps}
            />,
        );

        fireEvent.mouseEnter(screen.getByText('Test Filter'));
        fireEvent.click(screen.getByRole('button', {
            name: /Edit/i,
        }));

        expect(setEditFilter).toHaveBeenCalledTimes(1);
        expect(setEditFilter).toHaveBeenCalledWith(ipFilters[0]);
    });

    test('clicking the Delete button calls handleConfirmDeleteFilter', () => {
        renderWithIntl(
            <EditSection
                {...baseProps}
            />,
        );

        fireEvent.mouseEnter(screen.getByText('Test Filter'));
        fireEvent.click(screen.getByRole('button', {
            name: /Delete/i,
        }));

        expect(handleConfirmDeleteFilter).toHaveBeenCalledTimes(1);
        expect(handleConfirmDeleteFilter).toHaveBeenCalledWith(ipFilters[0]);
    });

    test('displays an error panel if current IP is not in range', () => {
        renderWithIntl(
            <EditSection
                {...baseProps}
                currentUsersIP='192.168.1.1'
                currentIPIsInRange={false}
            />,
        );

        expect(screen.getByText('Your IP address 192.168.1.1 is not included in your allowed IP address rules.')).toBeInTheDocument();
        expect(screen.getByText('Include your IP address in at least one of the rules below to continue.')).toBeInTheDocument();
        expect(screen.getByText('Add your IP address')).toBeInTheDocument();
    });

    test('displays a message if no filters are added', () => {
        renderWithIntl(
            <EditSection
                {...baseProps}
                ipFilters={[]}
            />,
        );

        expect(screen.getByText('No IP filtering rules added')).toBeInTheDocument();
        expect(screen.getByText('Add a filter')).toBeInTheDocument();
    });
});
