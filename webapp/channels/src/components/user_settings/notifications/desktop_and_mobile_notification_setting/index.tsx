// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {Fragment, useCallback, useEffect, useMemo, useRef, memo} from 'react';
import type {ChangeEvent, ReactNode} from 'react';
import {FormattedMessage} from 'react-intl';
import ReactSelect from 'react-select';
import type {ValueType, OptionsType} from 'react-select';

import type {UserNotifyProps} from '@mattermost/types/users';

import SettingItemMax from 'components/setting_item_max';
import SettingItemMin from 'components/setting_item_min';
import type SettingItemMinComponent from 'components/setting_item_min';

import Constants, {NotificationLevels, UserSettingsNotificationSections} from 'utils/constants';

import type {Props as UserSettingsNotificationsProps} from '../user_settings_notifications';

export type SelectOption = {
    label: ReactNode;
    value: string;
};

export type Props = {
    active: boolean;
    updateSection: (section: string) => void;
    onSubmit: () => void;
    onCancel: () => void;
    saving: boolean;
    error: string;
    setParentState: (key: string, value: string | boolean) => void;
    areAllSectionsInactive: boolean;
    isCollapsedThreadsEnabled: boolean;
    desktopActivity: UserNotifyProps['desktop'];
    sendPushNotifications: UserSettingsNotificationsProps['sendPushNotifications'];
    pushActivity: UserNotifyProps['push'];
    pushStatus: UserNotifyProps['push_status'];
    desktopThreads: UserNotifyProps['desktop_threads'];
    pushThreads: UserNotifyProps['push_threads'];
    desktopAndMobileSettingsDifferent: boolean;
};

function DesktopAndMobileNotificationSettings({
    active,
    updateSection,
    onSubmit,
    onCancel,
    saving,
    error,
    setParentState,
    areAllSectionsInactive,
    isCollapsedThreadsEnabled,
    desktopActivity,
    sendPushNotifications,
    pushActivity,
    pushStatus,
    desktopThreads,
    pushThreads,
    desktopAndMobileSettingsDifferent,
}: Props) {
    const editButtonRef = useRef<SettingItemMinComponent>(null);
    const previousActiveRef = useRef(active);

    // Focus back on the edit button, after this section was closed after it was opened
    useEffect(() => {
        if (previousActiveRef.current && !active && areAllSectionsInactive) {
            editButtonRef.current?.focus();
        }

        previousActiveRef.current = active;
    }, [active, areAllSectionsInactive]);

    const handleChangeForSendDesktopNotificationsRadio = useCallback((event: ChangeEvent<HTMLInputElement>) => {
        const value = event.target.value;
        setParentState('desktopActivity', value);
    }, [setParentState]);

    const handleChangeForDesktopThreadsNotificationCheckbox = useCallback((event: ChangeEvent<HTMLInputElement>) => {
        const value = event.target.checked ? NotificationLevels.ALL : NotificationLevels.MENTION;
        setParentState('desktopThreads', value);
    }, [setParentState]);

    const handleChangeForDifferentMobileNotificationsCheckbox = useCallback((event: ChangeEvent<HTMLInputElement>) => {
        const value = event.target.checked;
        setParentState('desktopAndMobileSettingsDifferent', value);
    }, [setParentState]);

    const handleChangeForSendMobileNotificationForSelect = useCallback((selectedOption: ValueType<SelectOption>) => {
        if (selectedOption && 'value' in selectedOption) {
            setParentState('pushActivity', selectedOption.value);
        }
    }, [setParentState]);

    const handleChangeForMobileThreadsNotificationCheckbox = useCallback((event: ChangeEvent<HTMLInputElement>) => {
        const value = event.target.checked ? NotificationLevels.ALL : NotificationLevels.MENTION;
        setParentState('pushThreads', value);
    }, [setParentState]);

    const handleChangeForSendMobileNotificationsWhenSelect = useCallback((selectedOption: ValueType<SelectOption>) => {
        if (selectedOption && 'value' in selectedOption) {
            setParentState('pushStatus', selectedOption.value);
        }
    }, [setParentState]);

    const maximizedSettingsInputs = useMemo(() => {
        const maximizedSettingInputs = [];

        const sendDesktopNotificationsSection = (
            <Fragment key='sendDesktopNotificationsSection'>
                <fieldset>
                    <legend className='form-legend'>
                        <FormattedMessage
                            id='user.settings.notifications.desktopAndMobile.sendDesktopNotificationFor'
                            defaultMessage='Send desktop notifications for:'
                        />
                    </legend>
                    {optionsOfSendNotifications.map((optionOfSendNotifications) => (
                        <div
                            key={optionOfSendNotifications.value}
                            className='radio'
                        >
                            <label>
                                <input
                                    type='radio'
                                    checked={desktopActivity === optionOfSendNotifications.value}
                                    value={optionOfSendNotifications.value}
                                    onChange={handleChangeForSendDesktopNotificationsRadio}
                                />
                                {optionOfSendNotifications.label}
                            </label>
                        </div>
                    ))}
                </fieldset>
            </Fragment>
        );
        maximizedSettingInputs.push(sendDesktopNotificationsSection);

        if (shouldShowDesktopThreadNotificationCheckbox(isCollapsedThreadsEnabled, desktopActivity)) {
            const isChecked = desktopThreads === NotificationLevels.ALL;
            const desktopThreadNotificationSection = (
                <Fragment key='desktopThreadNotificationSection'>
                    <br/>
                    <div className='checkbox single-checkbox'>
                        <label>
                            <input
                                type='checkbox'
                                checked={isChecked}
                                onChange={handleChangeForDesktopThreadsNotificationCheckbox}
                            />
                            <FormattedMessage
                                id='user.settings.notifications.desktopAndMobile.notifyForDesktopthreads'
                                defaultMessage={'Notify me about replies to threads I\'m following'}
                            />
                        </label>
                    </div>
                </Fragment>
            );
            maximizedSettingInputs.push(desktopThreadNotificationSection);
        }

        if (sendPushNotifications) {
            const differentMobileNotificationsSection = (
                <Fragment key='differentMobileNotificationsSection'>
                    <hr/>
                    <div className='checkbox single-checkbox'>
                        <label>
                            <input
                                type='checkbox'
                                data-testid='differentMobileNotificationsCheckbox'
                                checked={desktopAndMobileSettingsDifferent}
                                onChange={handleChangeForDifferentMobileNotificationsCheckbox}
                            />
                            <FormattedMessage
                                id='user.settings.notifications.desktopAndMobile.differentMobileNotificationsTitle'
                                defaultMessage='Use different settings for <strong>mobile push notifications</strong>'
                                values={{
                                    strong: (content: string) => <strong>{content}</strong>,
                                }}
                            />
                        </label>
                    </div>
                </Fragment>
            );
            maximizedSettingInputs.push(differentMobileNotificationsSection);
        }

        if (shouldShowSendMobileNotificationsForSelect(sendPushNotifications, desktopAndMobileSettingsDifferent)) {
            const mobileNotificationsSection = (
                <React.Fragment key='mobileNotificationsSection'>
                    <br/>
                    <label
                        id='sendMobileNotificationsLabel'
                        htmlFor='sendMobileNotificationsSelectInput'
                        className='singleSelectLabel'
                    >
                        <FormattedMessage
                            id='user.settings.notifications.desktopAndMobile.sendMobileNotificationsFor'
                            defaultMessage='Send mobile notifications for:'
                        />
                    </label>
                    <ReactSelect
                        inputId='sendMobileNotificationsSelectInput'
                        aria-labelledby='sendMobileNotificationsLabel'
                        className='react-select singleSelect'
                        classNamePrefix='react-select'
                        options={optionsOfSendNotifications}
                        clearable={false}
                        isClearable={false}
                        isSearchable={false}
                        components={{IndicatorSeparator: NoIndicatorSeparatorComponent}}
                        value={getValueOfSendMobileNotificationForSelect(pushActivity)}
                        onChange={handleChangeForSendMobileNotificationForSelect}
                    />
                </React.Fragment>
            );
            maximizedSettingInputs.push(mobileNotificationsSection);
        }

        if (shouldShowMobileThreadNotificationCheckbox(sendPushNotifications, isCollapsedThreadsEnabled, desktopAndMobileSettingsDifferent, desktopActivity, pushActivity)) {
            const isChecked = pushThreads === NotificationLevels.ALL;
            const threadNotificationSection = (
                <Fragment key='threadNotificationSection'>
                    <br/>
                    <div className='checkbox single-checkbox'>
                        <label>
                            <input
                                type='checkbox'
                                checked={isChecked}
                                onChange={handleChangeForMobileThreadsNotificationCheckbox}
                            />
                            <FormattedMessage
                                id='user.settings.notifications.desktopAndMobile.notifyForMobilethreads'
                                defaultMessage={'Notify me on mobile about replies to threads I\'m following'}
                            />
                        </label>
                    </div>
                </Fragment>
            );
            maximizedSettingInputs.push(threadNotificationSection);
        }

        if (shouldShowSendMobileNotificationsWhenSelect(sendPushNotifications, desktopActivity, pushActivity, desktopAndMobileSettingsDifferent)) {
            const sendMobileNotificationsWhenSection = (
                <React.Fragment key='sendMobileNotificationsWhenSection'>
                    <hr/>
                    <label
                        id='pushMobileNotificationsLabel'
                        htmlFor='pushMobileNotificationSelectInput'
                        className='singleSelectLabel'
                    >
                        <FormattedMessage
                            id='user.settings.notifications.desktopAndMobile.pushNotification'
                            defaultMessage='Send mobile notifications when I am:'
                        />
                    </label>
                    <ReactSelect
                        inputId='pushMobileNotificationSelectInput'
                        aria-labelledby='pushMobileNotificationsLabel'
                        className='react-select singleSelect'
                        classNamePrefix='react-select'
                        options={optionsOfSendMobileNotificationsWhenSelect}
                        clearable={false}
                        isClearable={false}
                        isSearchable={false}
                        components={{IndicatorSeparator: NoIndicatorSeparatorComponent}}
                        value={getValueOfSendMobileNotificationWhenSelect(pushStatus)}
                        onChange={handleChangeForSendMobileNotificationsWhenSelect}
                    />
                </React.Fragment>
            );
            maximizedSettingInputs.push(sendMobileNotificationsWhenSection);
        }

        return maximizedSettingInputs;
    },
    [
        desktopActivity,
        handleChangeForSendDesktopNotificationsRadio,
        isCollapsedThreadsEnabled,
        desktopThreads,
        handleChangeForDesktopThreadsNotificationCheckbox,
        sendPushNotifications,
        desktopAndMobileSettingsDifferent,
        handleChangeForDifferentMobileNotificationsCheckbox,
        pushActivity,
        handleChangeForSendMobileNotificationForSelect,
        pushThreads,
        handleChangeForMobileThreadsNotificationCheckbox,
        pushStatus,
        handleChangeForSendMobileNotificationsWhenSelect,
    ]);

    function handleChangeForMaxSection(section: string) {
        updateSection(section);
    }

    function handleChangeForMinSection(section: string) {
        updateSection(section);
        onCancel();
    }

    if (active) {
        return (
            <SettingItemMax
                title={
                    <FormattedMessage
                        id={'user.settings.notifications.desktopAndMobile.title'}
                        defaultMessage='Desktop and mobile notifications'
                    />
                }
                inputs={maximizedSettingsInputs}
                submit={onSubmit}
                saving={saving}
                serverError={error}
                updateSection={handleChangeForMaxSection}
            />
        );
    }

    return (
        <SettingItemMin
            ref={editButtonRef}
            title={
                <FormattedMessage
                    id='user.settings.notifications.desktopAndMobile.title'
                    defaultMessage='Desktop and mobile notifications'
                />
            }
            describe={getCollapsedText(desktopActivity, pushActivity)}
            section={UserSettingsNotificationSections.DESKTOP_AND_MOBILE}
            updateSection={handleChangeForMinSection}
        />
    );
}

function NoIndicatorSeparatorComponent() {
    return null;
}

const optionsOfSendNotifications = [
    {
        label: (
            <FormattedMessage
                id='user.settings.notifications.desktopAndMobile.allNewMessages'
                defaultMessage='All new messages'
            />
        ),
        value: NotificationLevels.ALL,
    },
    {
        label: (
            <FormattedMessage
                id='user.settings.notifications.desktopAndMobile.onlyMentions'
                defaultMessage='Mentions, direct messages, and group messages'
            />
        ),
        value: NotificationLevels.MENTION,
    },
    {
        label: (
            <FormattedMessage
                id='user.settings.notifications.desktopAndMobile.nothing'
                defaultMessage='Nothing'
            />
        ),
        value: NotificationLevels.NONE,
    },
];

export function shouldShowDesktopThreadNotificationCheckbox(isCollapsedThreadsEnabled: boolean, desktopActivity: UserNotifyProps['desktop']) {
    if (!isCollapsedThreadsEnabled) {
        return false;
    }

    if (desktopActivity === NotificationLevels.ALL || desktopActivity === NotificationLevels.NONE) {
        return false;
    }

    return true;
}

export function shouldShowMobileThreadNotificationCheckbox(sendPushNotifications: UserSettingsNotificationsProps['sendPushNotifications'], isCollapsedThreadsEnabled: boolean, desktopAndMobileSettingsDifferent: boolean, desktopActivity: UserNotifyProps['desktop'], pushActivity: UserNotifyProps['push']) {
    if (!sendPushNotifications) {
        return false;
    }

    if (!isCollapsedThreadsEnabled) {
        return false;
    }

    if (!desktopAndMobileSettingsDifferent) {
        return false;
    }

    if (pushActivity === NotificationLevels.ALL || pushActivity === NotificationLevels.NONE) {
        return false;
    }

    return true;
}

function shouldShowSendMobileNotificationsForSelect(sendPushNotifications: UserSettingsNotificationsProps['sendPushNotifications'], desktopAndMobileSettingsDifferent: boolean) {
    if (!sendPushNotifications) {
        return false;
    }

    if (desktopAndMobileSettingsDifferent) {
        return true;
    }

    return false;
}

export function getValueOfSendMobileNotificationForSelect(pushActivity: UserNotifyProps['push']): ValueType<SelectOption> {
    if (!pushActivity) {
        return optionsOfSendNotifications[1];
    }

    const option = optionsOfSendNotifications.find((option) => option.value === pushActivity);
    if (!option) {
        return optionsOfSendNotifications[1];
    }

    return option;
}

export function shouldShowSendMobileNotificationsWhenSelect(sendPushNotifications: UserSettingsNotificationsProps['sendPushNotifications'], desktopActivity: UserNotifyProps['desktop'], pushActivity: UserNotifyProps['push'], desktopAndMobileSettingsDifferent: boolean): boolean {
    if (!sendPushNotifications) {
        return false;
    }

    if (!desktopActivity || !pushActivity) {
        return true;
    }

    let shouldShow: boolean;
    if (desktopActivity === NotificationLevels.ALL || desktopActivity === NotificationLevels.MENTION) {
        //  Here we explicitly pass the state of desktopAndMobileSettingsDifferent instead of deriving as
        //  we need to show the select for mobile notifications when the desktop and mobile settings are different
        if (desktopAndMobileSettingsDifferent === true) {
            if (pushActivity === NotificationLevels.NONE) {
                shouldShow = false;
            } else {
                shouldShow = true;
            }
        } else {
            shouldShow = true;
        }
    } else if (desktopActivity === NotificationLevels.NONE) {
        if (desktopAndMobileSettingsDifferent === true) {
            if (pushActivity === NotificationLevels.NONE) {
                shouldShow = false;
            } else {
                shouldShow = true;
            }
        } else {
            shouldShow = false;
        }
    } else {
        shouldShow = true;
    }

    return shouldShow;
}

const optionsOfSendMobileNotificationsWhenSelect: OptionsType<SelectOption> = [
    {
        label: (
            <FormattedMessage
                id='user.settings.notifications.desktopAndMobile.online'
                defaultMessage='Online, away or offline'
            />
        ),
        value: Constants.UserStatuses.ONLINE,
    },
    {
        label: (
            <FormattedMessage
                id='user.settings.notifications.desktopAndMobile.away'
                defaultMessage='Away or offline'
            />
        ),
        value: Constants.UserStatuses.AWAY,
    },
    {
        label: (
            <FormattedMessage
                id='user.settings.notifications.desktopAndMobile.offline'
                defaultMessage='Offline'
            />
        ),
        value: Constants.UserStatuses.OFFLINE,
    },
];

export function getValueOfSendMobileNotificationWhenSelect(pushStatus?: UserNotifyProps['push_status']): ValueType<SelectOption> {
    if (!pushStatus) {
        return optionsOfSendMobileNotificationsWhenSelect[2];
    }

    const option = optionsOfSendMobileNotificationsWhenSelect.find((option) => option.value === pushStatus);
    if (!option) {
        return optionsOfSendMobileNotificationsWhenSelect[2];
    }

    return option;
}

function getCollapsedText(desktopActivity: UserNotifyProps['desktop'], pushActivity: UserNotifyProps['push']) {
    let collapsedText: ReactNode = null;
    if (desktopActivity === NotificationLevels.ALL) {
        if (pushActivity === NotificationLevels.ALL) {
            collapsedText = (
                <FormattedMessage
                    id='user.settings.notifications.desktopAndMobile.allForDesktopAndMobile'
                    defaultMessage='All new messages'
                />
            );
        } else if (pushActivity === NotificationLevels.MENTION) {
            collapsedText = (
                <FormattedMessage
                    id='user.settings.notifications.desktopAndMobile.allDesktopButMobileMentions'
                    defaultMessage='All new messages on desktop; Mentions, direct messages, and group messages on mobile'
                />
            );
        } else if (pushActivity === NotificationLevels.NONE) {
            collapsedText = (
                <FormattedMessage
                    id='user.settings.notifications.desktopAndMobile.allDesktopButMobileNone'
                    defaultMessage='All new messages on desktop; Never on mobile'
                />
            );
        }
    } else if (desktopActivity === NotificationLevels.MENTION) {
        if (pushActivity === NotificationLevels.ALL) {
            collapsedText = (
                <FormattedMessage
                    id='user.settings.notifications.desktopAndMobile.mentionsDesktopButMobileAll'
                    defaultMessage='Mentions, direct messages, and group messages on desktop; All new messages on mobile'
                />
            );
        } else if (pushActivity === NotificationLevels.MENTION) {
            collapsedText = (
                <FormattedMessage
                    id='user.settings.notifications.desktopAndMobile.mentionsForDesktopAndMobile'
                    defaultMessage='Mentions, direct messages, and group messages'
                />
            );
        } else if (pushActivity === NotificationLevels.NONE) {
            collapsedText = (
                <FormattedMessage
                    id='user.settings.notifications.desktopAndMobile.mentionsForDesktopButMobileNone'
                    defaultMessage='Mentions, direct messages, and group messages on desktop; Never on mobile'
                />
            );
        }
    } else if (desktopActivity === NotificationLevels.NONE) {
        if (pushActivity === NotificationLevels.ALL) {
            collapsedText = (
                <FormattedMessage
                    id='user.settings.notifications.desktopAndMobile.noneDesktopButMobileAll'
                    defaultMessage='Never on desktop; All new messages on mobile'
                />
            );
        } else if (pushActivity === NotificationLevels.MENTION) {
            collapsedText = (
                <FormattedMessage
                    id='user.settings.notifications.desktopAndMobile.noneDesktopButMobileMentions'
                    defaultMessage='Never on desktop, Mentions, direct messages, and group messages on mobile'
                />
            );
        } else if (pushActivity === NotificationLevels.NONE) {
            collapsedText = (
                <FormattedMessage
                    id='user.settings.notifications.desktopAndMobile.noneForDesktopAndMobile'
                    defaultMessage='Never'
                />
            );
        }
    }

    return collapsedText;
}

export default memo(DesktopAndMobileNotificationSettings);
