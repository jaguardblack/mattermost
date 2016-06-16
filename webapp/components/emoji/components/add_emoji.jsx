// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

import React from 'react';

import * as AsyncClient from 'utils/async_client.jsx';

import BackstageHeader from 'components/backstage/components/backstage_header.jsx';
import {FormattedMessage} from 'react-intl';
import FormError from 'components/form_error.jsx';
import {browserHistory, Link} from 'react-router';
import SpinnerButton from 'components/spinner_button.jsx';

export default class AddEmoji extends React.Component {
    static get propTypes() {
        return {
            team: React.PropTypes.object.isRequired,
            user: React.PropTypes.object.isRequired
        };
    }

    constructor(props) {
        super(props);

        this.handleSubmit = this.handleSubmit.bind(this);

        this.updateName = this.updateName.bind(this);
        this.updateImage = this.updateImage.bind(this);

        this.state = {
            name: '',
            image: null,
            imageUrl: '',
            saving: false,
            error: null
        };
    }

    handleSubmit(e) {
        e.preventDefault();

        if (this.state.saving) {
            return;
        }

        this.setState({
            saving: true,
            error: null
        });

        const emoji = {
            creator_id: this.props.user.id,
            name: this.state.name.trim().toLowerCase()
        };

        if (!emoji.name) {
            this.setState({
                saving: false,
                error: (
                    <FormattedMessage
                        id='add_emoji.nameRequired'
                        defaultMessage='A name is required for the emoji'
                    />
                )
            });

            return;
        } else if (/[^a-z0-9_-]/.test(emoji.name)) {
            this.setState({
                saving: false,
                error: (
                    <FormattedMessage
                        id='add_emoji.nameInvalid'
                        defaultMessage="An emoji's name can only contain lowercase letters, numbers, and the symbols '-' and '_'."
                    />
                )
            });

            return;
        }

        if (!this.state.image) {
            this.setState({
                saving: false,
                error: (
                    <FormattedMessage
                        id='add_emoji.imageRequired'
                        defaultMessage='An image is required for the emoji'
                    />
                )
            });

            return;
        }

        AsyncClient.addEmoji(
            emoji,
            this.state.image,
            () => {
                browserHistory.push('/' + this.props.team.name + '/emoji');
            },
            (err) => {
                this.setState({
                    saving: false,
                    error: err.message
                });
            }
        );
    }

    updateName(e) {
        this.setState({
            name: e.target.value
        });
    }

    updateImage(e) {
        if (e.target.files.length === 0) {
            this.setState({
                image: null,
                imageUrl: ''
            });

            return;
        }

        const image = e.target.files[0];

        const reader = new FileReader();
        reader.onload = () => {
            this.setState({
                image,
                imageUrl: reader.result
            });
        };
        reader.readAsDataURL(image);
    }

    render() {
        let filename = null;
        if (this.state.image) {
            filename = (
                <span className='add-emoji__filename'>
                    {this.state.image.name}
                </span>
            );
        }

        let preview = null;
        if (this.state.imageUrl) {
            preview = (
                <div className='form-group'>
                    <label
                        className='control-label col-sm-4'
                        htmlFor='preview'
                    >
                        <FormattedMessage
                            id='add_emoji.preview'
                            defaultMessage='Preview'
                        />
                    </label>
                    <div className='col-md-5 col-sm-8 add-emoji__preview'>
                        <FormattedMessage
                            id='add_emoji.preview.sentence'
                            defaultMessage='This is a sentence with {image} in it.'
                            values={{
                                image: (
                                    <img
                                        className='emoticon'
                                        src={this.state.imageUrl}
                                    />
                                )
                            }}
                        />
                    </div>
                </div>
            );
        }

        return (
            <div className='backstage-content row'>
                <BackstageHeader>
                    <Link to={'/' + this.props.team.name + '/emoji'}>
                        <FormattedMessage
                            id='emoji_list.header'
                            defaultMessage='Custom Emoji'
                        />
                    </Link>
                    <FormattedMessage
                        id='add_emoji.header'
                        defaultMessage='Add'
                    />
                </BackstageHeader>
                <div className='backstage-form'>
                    <form className='form-horizontal'>
                        <div className='form-group'>
                            <label
                                className='control-label col-sm-4'
                                htmlFor='name'
                            >
                                <FormattedMessage
                                    id='add_emoji.name'
                                    defaultMessage='Name'
                                />
                            </label>
                            <div className='col-md-5 col-sm-8'>
                                <input
                                    id='name'
                                    type='text'
                                    maxLength='64'
                                    className='form-control'
                                    value={this.state.name}
                                    onChange={this.updateName}
                                />
                                <div className='form__help'>
                                    <FormattedMessage
                                        id='add_emoji.name.help'
                                        defaultMessage="Choose a name for your emoji made of up to 64 characters consisting of lowercase letters, numbers, and the symbols '-' and '_'."
                                    />
                                </div>
                            </div>
                        </div>
                        <div className='form-group'>
                            <label
                                className='control-label col-sm-4'
                                htmlFor='image'
                            >
                                <FormattedMessage
                                    id='add_emoji.image'
                                    defaultMessage='Image'
                                />
                            </label>
                            <div className='col-md-5 col-sm-8'>
                                <div>
                                    <div className='add-emoji__upload'>
                                        <button className='btn btn-primary'>
                                            <FormattedMessage
                                                id='add_emoji.image.button'
                                                defaultMessage='Select'
                                            />
                                        </button>
                                        <input
                                            type='file'
                                            accept='.jpg,.png,.gif'
                                            multiple={false}
                                            onChange={this.updateImage}
                                        />
                                    </div>
                                    {filename}
                                    <div className='form__help'>
                                        <FormattedMessage
                                            id='add_emoji.image.help'
                                            defaultMessage='Choose the image for your emoji. The image can be a gif, png, or jpeg file with a max size of 64 KB and dimensions up to 128 by 128 pixels.'
                                        />
                                    </div>
                                </div>
                            </div>
                        </div>
                        {preview}
                        <div className='backstage-form__footer'>
                            <FormError error={this.state.error}/>
                            <Link
                                className='btn btn-sm'
                                to={'/' + this.props.team.name + '/emoji'}
                            >
                                <FormattedMessage
                                    id='add_emoji.cancel'
                                    defaultMessage='Cancel'
                                />
                            </Link>
                            <SpinnerButton
                                className='btn btn-primary'
                                type='submit'
                                spinning={this.state.saving}
                                onClick={this.handleSubmit}
                            >
                                <FormattedMessage
                                    id='add_emoji.save'
                                    defaultMessage='Save'
                                />
                            </SpinnerButton>
                        </div>
                    </form>
                </div>
            </div>
        );
    }
}
